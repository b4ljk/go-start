-- Add user_name and gender columns to users
ALTER TABLE users
ADD COLUMN user_name VARCHAR(255) NOT NULL;
ALTER TABLE users
ADD COLUMN gender VARCHAR(25) NULL DEFAULT 'PREFFER_NOT_TO_SAY';