// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_payment_methods.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createUserPaymentMethod = `-- name: CreateUserPaymentMethod :execresult
INSERT INTO user_payment_methods (
    user_id, payment_type, provider, is_default,
    status, card_info, bank_info, wallet_info
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateUserPaymentMethodParams struct {
	UserID      int64                         `json:"user_id"`
	PaymentType UserPaymentMethodsPaymentType `json:"payment_type"`
	Provider    string                        `json:"provider"`
	IsDefault   sql.NullBool                  `json:"is_default"`
	Status      NullUserPaymentMethodsStatus  `json:"status"`
	CardInfo    json.RawMessage               `json:"card_info"`
	BankInfo    json.RawMessage               `json:"bank_info"`
	WalletInfo  json.RawMessage               `json:"wallet_info"`
}

func (q *Queries) CreateUserPaymentMethod(ctx context.Context, arg CreateUserPaymentMethodParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUserPaymentMethod,
		arg.UserID,
		arg.PaymentType,
		arg.Provider,
		arg.IsDefault,
		arg.Status,
		arg.CardInfo,
		arg.BankInfo,
		arg.WalletInfo,
	)
}

const deleteUserPaymentMethod = `-- name: DeleteUserPaymentMethod :exec
UPDATE user_payment_methods
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) DeleteUserPaymentMethod(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserPaymentMethod, id)
	return err
}

const getDefaultUserPaymentMethod = `-- name: GetDefaultUserPaymentMethod :one
SELECT id, user_id, uuid, payment_type, provider, is_default, status, card_info, bank_info, wallet_info, created_at, updated_at, deleted_at FROM user_payment_methods
WHERE user_id = ? AND is_default = true AND deleted_at IS NULL
`

func (q *Queries) GetDefaultUserPaymentMethod(ctx context.Context, userID int64) (UserPaymentMethod, error) {
	row := q.db.QueryRowContext(ctx, getDefaultUserPaymentMethod, userID)
	var i UserPaymentMethod
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.PaymentType,
		&i.Provider,
		&i.IsDefault,
		&i.Status,
		&i.CardInfo,
		&i.BankInfo,
		&i.WalletInfo,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserPaymentMethodByID = `-- name: GetUserPaymentMethodByID :one
SELECT id, user_id, uuid, payment_type, provider, is_default, status, card_info, bank_info, wallet_info, created_at, updated_at, deleted_at FROM user_payment_methods
WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetUserPaymentMethodByID(ctx context.Context, id int64) (UserPaymentMethod, error) {
	row := q.db.QueryRowContext(ctx, getUserPaymentMethodByID, id)
	var i UserPaymentMethod
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Uuid,
		&i.PaymentType,
		&i.Provider,
		&i.IsDefault,
		&i.Status,
		&i.CardInfo,
		&i.BankInfo,
		&i.WalletInfo,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listUserPaymentMethods = `-- name: ListUserPaymentMethods :many
SELECT id, user_id, uuid, payment_type, provider, is_default, status, card_info, bank_info, wallet_info, created_at, updated_at, deleted_at FROM user_payment_methods
WHERE user_id = ? AND deleted_at IS NULL
ORDER BY is_default DESC, id DESC
`

func (q *Queries) ListUserPaymentMethods(ctx context.Context, userID int64) ([]UserPaymentMethod, error) {
	rows, err := q.db.QueryContext(ctx, listUserPaymentMethods, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserPaymentMethod{}
	for rows.Next() {
		var i UserPaymentMethod
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Uuid,
			&i.PaymentType,
			&i.Provider,
			&i.IsDefault,
			&i.Status,
			&i.CardInfo,
			&i.BankInfo,
			&i.WalletInfo,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserPaymentMethod = `-- name: UpdateUserPaymentMethod :exec
UPDATE user_payment_methods
SET 
    payment_type = COALESCE(?, payment_type),
    provider = COALESCE(?, provider),
    is_default = COALESCE(?, is_default),
    status = COALESCE(?, status),
    card_info = COALESCE(?, card_info),
    bank_info = COALESCE(?, bank_info),
    wallet_info = COALESCE(?, wallet_info)
WHERE id = ? AND deleted_at IS NULL
`

type UpdateUserPaymentMethodParams struct {
	PaymentType UserPaymentMethodsPaymentType `json:"payment_type"`
	Provider    string                        `json:"provider"`
	IsDefault   sql.NullBool                  `json:"is_default"`
	Status      NullUserPaymentMethodsStatus  `json:"status"`
	CardInfo    json.RawMessage               `json:"card_info"`
	BankInfo    json.RawMessage               `json:"bank_info"`
	WalletInfo  json.RawMessage               `json:"wallet_info"`
	ID          int64                         `json:"id"`
}

func (q *Queries) UpdateUserPaymentMethod(ctx context.Context, arg UpdateUserPaymentMethodParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPaymentMethod,
		arg.PaymentType,
		arg.Provider,
		arg.IsDefault,
		arg.Status,
		arg.CardInfo,
		arg.BankInfo,
		arg.WalletInfo,
		arg.ID,
	)
	return err
}
