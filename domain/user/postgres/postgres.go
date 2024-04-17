package postgres

//this code is not auto generated
import (
	"context"

	"github.com/mamadeusia/AuthSrv/client/postgres"
	"github.com/mamadeusia/AuthSrv/entity"

	"go-micro.dev/v4/logger"
)

type PostgresRepository struct {
	Postgres *postgres.PostgresRepository
}

func NewRepository(postgres *postgres.PostgresRepository) *PostgresRepository {
	return &PostgresRepository{
		Postgres: postgres,
	}
}

// Create - is responsible for calling sqlc to create a person
func (pr *PostgresRepository) CreatePerson(ctx context.Context, person entity.Person) error {

	params := postgres.CreatePersonParams{
		TelegramID:       person.TelegramID,
		FirstName:        person.FirstName,
		LastName:         person.LastName,
		Language:         person.Language,
		TelegramLanguage: person.TelegramLanguage,
		MainPasswordHash: person.MainPasswordHash,
		FakePasswordHash: person.FakePasswordHash,
	}
	if _, err := pr.Postgres.Queries.CreatePerson(ctx, params); err != nil {
		logger.Info("DOMAIN::Create, has failed with error %v", err)
		return err
	}

	return nil

}

// GetByTelegramID - is responsible for calling sqlc
func (pr *PostgresRepository) GetByTelegramID(ctx context.Context, telegramID int64) (*entity.Person, error) {
	o, err := pr.Postgres.Queries.GetPersonByTelegramID(ctx, telegramID)
	if err != nil {
		logger.Info("DOMAIN::GetByTelegramID, has failed with error : %v", err)
		return nil, err
	}
	return &entity.Person{
		ID:               o.ID,
		TelegramID:       o.TelegramID,
		FirstName:        o.FirstName,
		LastName:         o.LastName,
		Language:         o.Language,
		TelegramLanguage: o.TelegramLanguage,
		MainPasswordHash: o.MainPasswordHash,
		FakePasswordHash: o.FakePasswordHash,
		CreatedAt:        o.CreatedAt,
	}, nil

}

// GetNearValidators - is responsible to call sqlc
func (pr *PostgresRepository) GetNearValidators(ctx context.Context, lat, lon float64, distance int64, limit, offset int32) ([]int64, error) {
	result, err := pr.Postgres.Queries.GetNearValidators(ctx, postgres.GetNearValidatorsParams{
		Radians:     lat,
		Radians_2:   lon,
		LocationLat: distance,
		Limit:       limit,
		Offset:      offset,
	})
	if err != nil {
		logger.Info("DOMAIN::GetNearValidators, has failed with error : %v", err)
		return nil, err
	}

	return result, nil
}

// SetAdmin - is responsible for calling sqlc
func (pr *PostgresRepository) SetAdmin(ctx context.Context, telegramID int64) error {
	if err := pr.Postgres.Queries.SetAdmin(ctx, telegramID); err != nil {
		logger.Info("DOMAIN::SetAdmin, has failed with error , %v", err)
		return err
	}
	return nil
}
