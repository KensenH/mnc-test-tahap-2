// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package schema

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createMutation = `-- name: CreateMutation :one
INSERT INTO mutations (user_id, amount, balance_before, balance_after, transaction_type, status)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, amount, balance_before, balance_after, transaction_type, status, created_date
`

type CreateMutationParams struct {
	UserID          pgtype.UUID
	Amount          pgtype.Numeric
	BalanceBefore   pgtype.Numeric
	BalanceAfter    pgtype.Numeric
	TransactionType string
	Status          pgtype.Text
}

func (q *Queries) CreateMutation(ctx context.Context, arg CreateMutationParams) (Mutation, error) {
	row := q.db.QueryRow(ctx, createMutation,
		arg.UserID,
		arg.Amount,
		arg.BalanceBefore,
		arg.BalanceAfter,
		arg.TransactionType,
		arg.Status,
	)
	var i Mutation
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.BalanceBefore,
		&i.BalanceAfter,
		&i.TransactionType,
		&i.Status,
		&i.CreatedDate,
	)
	return i, err
}

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (amount, user_id, remarks) VALUES ($1, $2, $3) RETURNING payment_id, user_id, amount, status, remarks, created_date
`

type CreatePaymentParams struct {
	Amount  pgtype.Numeric
	UserID  pgtype.UUID
	Remarks string
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, createPayment, arg.Amount, arg.UserID, arg.Remarks)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.UserID,
		&i.Amount,
		&i.Status,
		&i.Remarks,
		&i.CreatedDate,
	)
	return i, err
}

const createTopUp = `-- name: CreateTopUp :one
INSERT INTO top_ups (amount_top_up, user_id) VALUES ($1, $2) RETURNING top_up_id, user_id, amount_top_up, status, created_date
`

type CreateTopUpParams struct {
	AmountTopUp pgtype.Numeric
	UserID      pgtype.UUID
}

func (q *Queries) CreateTopUp(ctx context.Context, arg CreateTopUpParams) (TopUp, error) {
	row := q.db.QueryRow(ctx, createTopUp, arg.AmountTopUp, arg.UserID)
	var i TopUp
	err := row.Scan(
		&i.TopUpID,
		&i.UserID,
		&i.AmountTopUp,
		&i.Status,
		&i.CreatedDate,
	)
	return i, err
}

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (target_user, user_id, amount, remarks) VALUES ($1, $2, $3, $4) RETURNING transfer_id, user_id, target_user, amount, remarks, created_date
`

type CreateTransferParams struct {
	TargetUser pgtype.UUID
	UserID     pgtype.UUID
	Amount     pgtype.Numeric
	Remarks    pgtype.Text
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, createTransfer,
		arg.TargetUser,
		arg.UserID,
		arg.Amount,
		arg.Remarks,
	)
	var i Transfer
	err := row.Scan(
		&i.TransferID,
		&i.UserID,
		&i.TargetUser,
		&i.Amount,
		&i.Remarks,
		&i.CreatedDate,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, phone_number, address, pin
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING user_id, first_name, last_name, phone_number, address, pin, balance, created_date
`

type CreateUserParams struct {
	FirstName   string
	LastName    pgtype.Text
	PhoneNumber pgtype.Text
	Address     pgtype.Text
	Pin         string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.Address,
		arg.Pin,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Address,
		&i.Pin,
		&i.Balance,
		&i.CreatedDate,
	)
	return i, err
}

const creditBalance = `-- name: CreditBalance :one
UPDATE users SET balance = balance + $1 WHERE user_id = $2 RETURNING user_id, first_name, last_name, phone_number, address, pin, balance, created_date
`

type CreditBalanceParams struct {
	Balance pgtype.Numeric
	UserID  uuid.UUID
}

func (q *Queries) CreditBalance(ctx context.Context, arg CreditBalanceParams) (User, error) {
	row := q.db.QueryRow(ctx, creditBalance, arg.Balance, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Address,
		&i.Pin,
		&i.Balance,
		&i.CreatedDate,
	)
	return i, err
}

const debitBalance = `-- name: DebitBalance :one
UPDATE users SET balance = balance - $1 WHERE user_id = $2 RETURNING user_id, first_name, last_name, phone_number, address, pin, balance, created_date
`

type DebitBalanceParams struct {
	Balance pgtype.Numeric
	UserID  uuid.UUID
}

func (q *Queries) DebitBalance(ctx context.Context, arg DebitBalanceParams) (User, error) {
	row := q.db.QueryRow(ctx, debitBalance, arg.Balance, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Address,
		&i.Pin,
		&i.Balance,
		&i.CreatedDate,
	)
	return i, err
}

const getUserByPhoneNumber = `-- name: GetUserByPhoneNumber :one
SELECT user_id, first_name, last_name, phone_number, address, pin, balance, created_date FROM users WHERE phone_number = $1
`

func (q *Queries) GetUserByPhoneNumber(ctx context.Context, phoneNumber pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByPhoneNumber, phoneNumber)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Address,
		&i.Pin,
		&i.Balance,
		&i.CreatedDate,
	)
	return i, err
}
