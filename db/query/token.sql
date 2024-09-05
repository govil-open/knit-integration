-- name: CreateToken :one
INSERT INTO tokens (
  audience,
  token,
  expires_in,
  scope,
  authorities
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING jti;