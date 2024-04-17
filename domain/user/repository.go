package user

import (
	"context"
	"errors"

	"github.com/mamadeusia/AuthSrv/entity"
)

var (
	ErrOrderNotFound = errors.New("the order not found")
)

type Repository interface {
	CreatePerson(ctx context.Context, person entity.Person) error
	GetByTelegramID(ctx context.Context, telegramID int64) (*entity.Person, error)
	GetNearValidators(ctx context.Context, lat, lon float64, distance int64, limit, offset int32) ([]int64, error)
	SetAdmin(ctx context.Context, telegramID int64) error
}
