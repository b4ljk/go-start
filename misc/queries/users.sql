-- name: CreateUser :one
INSERT INTO users (
        email,
        password,
        user_name,
        gender
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
-- name: ListAllUsers :many
SELECT *
FROM users
ORDER BY created_at DESC;
-- name: UpdateUser :one
UPDATE users
SET email = COALESCE(sqlc.narg('email'), email),
    user_name = COALESCE(sqlc.narg('user_name'), user_name),
    gender = COALESCE(sqlc.narg('gender'), gender),
    updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
-- name: CreatePost :one
INSERT INTO posts (title, content, user_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetPostByID :one
SELECT p.*,
    u.user_name,
    u.email
FROM posts p
    JOIN users u ON u.id = p.user_id
WHERE p.id = $1
LIMIT 1;
-- name: ListPosts :many
SELECT p.*,
    u.user_name,
    u.email
FROM posts p
    JOIN users u ON u.id = p.user_id
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;
-- name: ListUserPosts :many
SELECT p.*,
    u.user_name,
    u.email
FROM posts p
    JOIN users u ON u.id = p.user_id
WHERE p.user_id = $1
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3;
-- name: UpdatePost :one
UPDATE posts
SET title = COALESCE(sqlc.narg('title'), title),
    content = COALESCE(sqlc.narg('content'), content),
    updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
    AND user_id = sqlc.arg('user_id')
RETURNING *;
-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
    AND user_id = $2;
-- name: SearchPosts :many
SELECT p.*,
    u.user_name,
    u.email
FROM posts p
    JOIN users u ON u.id = p.user_id
WHERE p.title ILIKE '%' || $1 || '%'
    OR p.content ILIKE '%' || $1 || '%'
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3;