// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package schema

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Mutation struct {
	ID              uuid.UUID
	UserID          pgtype.UUID
	Amount          pgtype.Numeric
	BalanceBefore   pgtype.Numeric
	BalanceAfter    pgtype.Numeric
	TransactionType string
	Status          pgtype.Text
	CreatedDate     pgtype.Timestamp
}

type Payment struct {
	PaymentID   uuid.UUID
	UserID      pgtype.UUID
	Amount      pgtype.Numeric
	Status      string
	Remarks     string
	CreatedDate pgtype.Timestamp
}

type TopUp struct {
	TopUpID     uuid.UUID
	UserID      pgtype.UUID
	AmountTopUp pgtype.Numeric
	Status      string
	CreatedDate pgtype.Timestamp
}

type Transfer struct {
	TransferID  uuid.UUID
	UserID      pgtype.UUID
	TargetUser  pgtype.UUID
	Amount      pgtype.Numeric
	Remarks     pgtype.Text
	CreatedDate pgtype.Timestamp
}

type User struct {
	UserID      uuid.UUID
	FirstName   string
	LastName    pgtype.Text
	PhoneNumber pgtype.Text
	Address     pgtype.Text
	Pin         string
	Balance     pgtype.Numeric
	CreatedDate pgtype.Timestamp
}
