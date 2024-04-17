package handler

import (
	"context"
	"errors"

	"go-micro.dev/v4/logger"

	"github.com/go-playground/validator"
	"github.com/mamadeusia/AuthSrv/entity"
	pb "github.com/mamadeusia/AuthSrv/proto"
	service "github.com/mamadeusia/AuthSrv/service/authSrv"
)

// handler
type AuthHandler struct {
	AuthService *service.AuthSrv
	Validate    *validator.Validate
}

func NewAuthHandler(srv *service.AuthSrv, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		AuthService: srv,
		Validate:    validate,
	}
}

// SetPerson - handle the comming req and do a simple validation then route to service
func (e *AuthHandler) CreatePerson(ctx context.Context, req *pb.CreatePersonRequest, rsp *pb.CreatePersonResponse) error {
	if req.Person.TelegramID == 0 || req.Person.Language == "" ||
		req.Person.MainPasswordHash == "" || req.Person.FakePasswordHash == "" || req.Person.TelegramLanguage == "" {
		logger.Error("HANDLER::CreatePerson, has failed with error :Invalid input parameter : %v", req.Person)
		rsp.Msg = "Failed"
		return errors.New("invalid input parameters")
	}
	if err := e.AuthService.CreatePerson(ctx, entity.Person{
		TelegramID:       req.Person.TelegramID,
		FirstName:        req.Person.FirstName,
		LastName:         req.Person.LastName,
		Language:         req.Person.Language,
		TelegramLanguage: req.Person.TelegramLanguage,
		MainPasswordHash: req.Person.MainPasswordHash,
		FakePasswordHash: req.Person.FakePasswordHash,
	}); err != nil {
		logger.Error("HANDLER::CreatePerson, has failed with error : %v", err)
		rsp.Msg = "Failed"
		return err
	}
	rsp.Msg = "Success"

	return nil
}

// GetPersonByTelegramID - handle the comming req and do a simple validation then route to service
func (e *AuthHandler) GetPersonByTelegramID(ctx context.Context, req *pb.GetPersonByTelegramIDRequest, rsp *pb.GetPersonByTelegramIDResponse) error {
	logger.Infof("HANDLER::Received AuthSrv.GetPersonByTelegramID request: %v", req)
	if req.TelegramID == 0 || req.PasswordHash == "" {
		logger.Error("HANDLER:: GetPersonByTelegramID, has failed with error Invalid input parameter : %v", req)
		return errors.New("invalid input parameters")
	}
	person, err := e.AuthService.GetPersonByTelegramID(ctx, req.TelegramID, req.PasswordHash)
	if err != nil {
		logger.Error("HANDLER::GetPersonByTelegramID, has failed with error  : %v", err)
		return err
	}
	rsp.Person = &pb.Person{
		TelegramID:       person.ID,
		FirstName:        person.FirstName,
		LastName:         person.LastName,
		Language:         person.Language,
		TelegramLanguage: person.TelegramLanguage,
		LocationLat:      person.LocationLat,
		LocationLon:      person.LocationLon,
	}
	return nil
}

// GetNearValidators implements AuthSrv.AuthSrvHandler
func (a *AuthHandler) GetNearValidators(ctx context.Context, req *pb.GetNearValidatorsRequest, res *pb.GetNearValidatorsResponse) error {
	//we have to add additional limitation for country boundry
	if req.LocationLat == 0 || req.LocationLon == 0 || req.Distance == 0 {
		logger.Info("HANDLER::GetNearValidators, has failed with error , invalid input type : %v", req)
		return errors.New("invalid input type")
	}

	if req.Limit == 0 {
		req.Limit = 1000
	}
	result, err := a.AuthService.GetNearValidators(ctx, req.LocationLat, req.LocationLon, req.Distance, req.Limit, req.Offset)
	if err != nil {
		logger.Info("HANDLER::GetNearValidators, has failed with error : %v", err)
		return nil
	}
	res.Validators = result
	return nil
}

// SetAdmin implements AuthSrv.AuthSrvHandler
func (a *AuthHandler) SetAdmin(ctx context.Context, req *pb.SetAdminRequest, res *pb.SetAdminResponse) error {
	if req.TelegramID == 0 {
		logger.Info("HANDLER::SetAdmin, has failed with error , input type is not valid : %v", req)
		return errors.New("input type is not valid")
	}

	if err := a.AuthService.SetAdmin(ctx, req.TelegramID); err != nil {
		logger.Info("HANDLER::SetAdmin, has failed with error : %v", err)
		return err
	}
	return nil
}

// CheckPersonExistByTelegramID - check this person has profile or not
func (e *AuthHandler) CheckPersonExistByTelegramID(ctx context.Context, req *pb.CheckPersonExistByTelegramIDRequest, rsp *pb.CheckPersonExistByTelegramIDResponse) error {
	logger.Infof("HANDLER::Received AuthSrv.CheckPersonExistByTelegramID request: %v", req)

	request := entity.CheckPersonExistByTelegramIDRequest{
		TelegramID: req.TelegramID,
	}

	err := e.Validate.Struct(request)
	if err != nil {
		logger.Error("HANDLER:: CheckPersonExistByTelegramID, has failed with error Invalid input parameter : %v", req)
		return errors.New("invalid input parameters")
	}

	result, err := e.AuthService.CheckPersonExistByTelegramID(ctx, request)
	if err != nil {
		logger.Error("HANDLER::GetPersonByTelegramID, has failed with error  : %v", err)
		return err
	}

	rsp.Result = result
	return nil
}
