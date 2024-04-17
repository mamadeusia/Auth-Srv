-- name: CreatePerson :one
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
) RETURNING *;

-- name: GetPersonByTelegramID :one
SELECT * FROM person
WHERE telegram_id = $1
LIMIT 1;

-- name: GetNearValidators :many
SELECT telegram_id
FROM person
WHERE is_admin = 0
HAVING ( 6371 * acos( cos( radians($1) ) * cos( radians( location_lat ) ) * cos( radians( location_lon ) - radians($2) ) + sin( radians($1) ) * sin( radians( location_lat ) ) ) ) < $3
ORDER BY created_at DESC
LIMIT $4 OFFSET $5;

-- name: SetAdmin :exec
UPDATE person SET is_admin = 1
WHERE telegram_id = $1;
