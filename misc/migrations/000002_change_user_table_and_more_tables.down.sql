-- drop updated at column from users
ALTER TABLE users DROP COLUMN updated_at;
-- drop posts table
DROP TABLE posts;