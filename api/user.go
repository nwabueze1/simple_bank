package api

import (
	"fmt"
	"net/http"
	"time"

	db "fidelis.com/simple_bank/db/sqlc"
	"fidelis.com/simple_bank/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"fullname" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}


type createUserResponse struct{
	UserName string `json:"username"`
	FullName string `json:"fullname"`
	Email string	`json:"email"`
	PasswordChangedAt time.Time`json:"password_changed_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context){

	var req createUserRequest

	if err:= ctx.ShouldBindJSON(&req); err!=nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hash user password
	hashed, err := util.HashPassword(req.Password)

	if err != nil {
		errMessage := fmt.Errorf("hash password error: %v",err )
		ctx.JSON(http.StatusBadRequest, errorResponse(errMessage))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		HashedPassword: hashed,
		FullName: req.FullName,
		Email: req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		pqErr, ok := err.(*pq.Error)

		if ok{
			switch pqErr.Code.Name(){
			case "unique_violation":
				errorMessage := fmt.Errorf("A user with this username already exists")
				ctx.JSON(http.StatusForbidden, errorResponse(errorMessage))
				return;
			} 
		}

	
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := createUserResponse{
		UserName: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt: user.CreatedAt,
	}
	
	ctx.JSON(http.StatusOK, resp)
}