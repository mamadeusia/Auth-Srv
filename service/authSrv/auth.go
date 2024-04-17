package service

import (
	"context"
	"errors"

	"github.com/mamadeusia/AuthSrv/domain/user/postgres"
	"github.com/mamadeusia/AuthSrv/entity"

	"go-micro.dev/v4/logger"
)

// AuthSrv
type AuthSrv struct {
	UserStore *postgres.PostgresRepository
}

func NewAuthService(store *postgres.PostgresRepository) *AuthSrv {
	return &AuthSrv{
		UserStore: store,
	}
}

// SetPerson - call the aggregator to store the data
func (e *AuthSrv) CreatePerson(ctx context.Context, req entity.Person) error {
	logger.Infof("SERVICE::Received AuthSrv.SetPerson request: %v", req)
	if err := e.UserStore.CreatePerson(ctx, req); err != nil {
		logger.Error("SERVICE::CreatePerson, has failed with error : %v", err)
		return err
	}
	return nil
}

// GetPersonByTelegramID - is responsible for business logic around fetching person data and routing to repository
func (e *AuthSrv) GetPersonByTelegramID(ctx context.Context, telegram_id int64, password_hash string) (*entity.Person, error) {
	logger.Infof("SERVICE::Received AuthSrv.GetPersonByTelegramID request: %v", telegram_id)
	person, err := e.UserStore.GetByTelegramID(ctx, telegram_id)
	if err != nil {
		logger.Error("SERVICE::GetPersonByTelegramID, has failed with error : %V", err)
		return nil, err
	}
	if person.FakePasswordHash == password_hash || person.MainPasswordHash == password_hash {
		return person, nil
	}
	return nil, errors.New("password is not correct")
}

// GetNearValidators - is responsible for business logic arounf finding near validators and routing to repository
func (e *AuthSrv) GetNearValidators(ctx context.Context, lat, lon float64, distance int64, limit, offset int32) ([]int64, error) {
	result, err := e.UserStore.GetNearValidators(ctx, lat, lon, distance, limit, offset)
	if err != nil {
		logger.Info("SERVICE::GetNearValidators, has failed with error : %v", err)
		return nil, err
	}

	return result, nil
}

func (e *AuthSrv) SetAdmin(ctx context.Context, telegramID int64) error {
	if err := e.UserStore.SetAdmin(ctx, telegramID); err != nil {
		logger.Info("SERVICE::SetAdmin, has failed with error : %v", err)
		return err
	}
	return nil
}

// CheckPersonExistByTelegramID - is responsible for business logic around fetching person data and routing to repository
func (e *AuthSrv) CheckPersonExistByTelegramID(ctx context.Context, request entity.CheckPersonExistByTelegramIDRequest) (bool, error) {
	logger.Infof("SERVICE::Received AuthSrv.CheckPersonExistByTelegramID request: %v", request.TelegramID)
	person, err := e.UserStore.GetByTelegramID(ctx, request.TelegramID)
	if err != nil {
		return false, nil
	}
	if person != nil {
		return true, nil
	}

	return false, nil
}
