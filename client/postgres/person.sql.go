// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: person.sql

package postgres

import (
	"context"
)

const createPerson = `-- name: CreatePerson :one
INSERT INTO person (
  telegram_id,
  first_name,
  last_name, 
  language,
  telegram_language,
  main_password_hash,
  fake_password_hash,
  location_lat,
  location_lon
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9
) RETURNING id, telegram_id, first_name, last_name, language, telegram_language, main_password_hash, fake_password_hash, location_lat, location_lon, is_admin, created_at, updated_at
`

type CreatePersonParams struct {
	TelegramID       int64  `db:"telegram_id" json:"telegram_id"`
	FirstName        string `db:"first_name" json:"first_name"`
	LastName         string `db:"last_name" json:"last_name"`
	Language         string `db:"language" json:"language"`
	TelegramLanguage string `db:"telegram_language" json:"telegram_language"`
	MainPasswordHash string `db:"main_password_hash" json:"main_password_hash"`
	FakePasswordHash string `db:"fake_password_hash" json:"fake_password_hash"`
	LocationLat      int64  `db:"location_lat" json:"location_lat"`
	LocationLon      int64  `db:"location_lon" json:"location_lon"`
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) (Person, error) {
	row := q.db.QueryRow(ctx, createPerson,
		arg.TelegramID,
		arg.FirstName,
		arg.LastName,
		arg.Language,
		arg.TelegramLanguage,
		arg.MainPasswordHash,
		arg.FakePasswordHash,
		arg.LocationLat,
		arg.LocationLon,
	)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.TelegramID,
		&i.FirstName,
		&i.LastName,
		&i.Language,
		&i.TelegramLanguage,
		&i.MainPasswordHash,
		&i.FakePasswordHash,
		&i.LocationLat,
		&i.LocationLon,
		&i.IsAdmin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getNearValidators = `-- name: GetNearValidators :many
SELECT telegram_id
FROM person
WHERE is_admin = 0
HAVING ( 6371 * acos( cos( radians($1) ) * cos( radians( location_lat ) ) * cos( radians( location_lon ) - radians($2) ) + sin( radians($1) ) * sin( radians( location_lat ) ) ) ) < $3
ORDER BY created_at DESC
LIMIT $4 OFFSET $5
`

type GetNearValidatorsParams struct {
	Radians     float64 `db:"radians" json:"radians"`
	Radians_2   float64 `db:"radians_2" json:"radians_2"`
	LocationLat int64   `db:"location_lat" json:"location_lat"`
	Limit       int32   `db:"limit" json:"limit"`
	Offset      int32   `db:"offset" json:"offset"`
}

func (q *Queries) GetNearValidators(ctx context.Context, arg GetNearValidatorsParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, getNearValidators,
		arg.Radians,
		arg.Radians_2,
		arg.LocationLat,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var telegram_id int64
		if err := rows.Scan(&telegram_id); err != nil {
			return nil, err
		}
		items = append(items, telegram_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPersonByTelegramID = `-- name: GetPersonByTelegramID :one
SELECT id, telegram_id, first_name, last_name, language, telegram_language, main_password_hash, fake_password_hash, location_lat, location_lon, is_admin, created_at, updated_at FROM person
WHERE telegram_id = $1
LIMIT 1
`

func (q *Queries) GetPersonByTelegramID(ctx context.Context, telegramID int64) (Person, error) {
	row := q.db.QueryRow(ctx, getPersonByTelegramID, telegramID)
	var i Person
	err := row.Scan(
		&i.ID,
		&i.TelegramID,
		&i.FirstName,
		&i.LastName,
		&i.Language,
		&i.TelegramLanguage,
		&i.MainPasswordHash,
		&i.FakePasswordHash,
		&i.LocationLat,
		&i.LocationLon,
		&i.IsAdmin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const setAdmin = `-- name: SetAdmin :exec
UPDATE person SET is_admin = 1
WHERE telegram_id = $1
`

func (q *Queries) SetAdmin(ctx context.Context, telegramID int64) error {
	_, err := q.db.Exec(ctx, setAdmin, telegramID)
	return err
}