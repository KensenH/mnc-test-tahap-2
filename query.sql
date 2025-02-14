-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, phone_number, address, pin
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1;

-- name: CreateTopUp :one
INSERT INTO top_ups (amount_top_up, user_id) VALUES ($1, $2) RETURNING *;

-- name: CreatePayment :one
INSERT INTO payments (amount, user_id, remarks) VALUES ($1, $2, $3) RETURNING *;

-- name: CreateTransfer :one
INSERT INTO transfers (target_user, user_id, amount, remarks) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreditBalance :one
UPDATE users SET balance = balance + $1 WHERE user_id = $2 RETURNING *;

-- name: DebitBalance :one
UPDATE users SET balance = balance - $1 WHERE user_id = $2 RETURNING *;

-- name: CreateMutation :one
INSERT INTO mutations (user_id, amount, balance_before, balance_after, transaction_type, status)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;