DROP INDEX IF EXISTS "accounts_owner_currency_idx";
DROP INDEX IF EXISTS "users_email_idx";
DROP INDEX IF EXISTS "users_username_idx";
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";
DROP TABLE "users";
