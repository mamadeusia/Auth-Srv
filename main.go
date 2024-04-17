package main

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/mamadeusia/AuthSrv/client/postgres"
	"github.com/mamadeusia/AuthSrv/config"
	authHandler "github.com/mamadeusia/AuthSrv/handler/authSrv"
	pb "github.com/mamadeusia/AuthSrv/proto"
	AuthService "github.com/mamadeusia/AuthSrv/service/authSrv"

	pgRepo "github.com/mamadeusia/AuthSrv/domain/user/postgres"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "authsrv"
	version = "latest"
)

func main() {

	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	//create db connection
	ctxDB := context.Background()
	postgresClientInstance, err := postgres.NewPostgres(ctxDB, config.PostgresURL())
	if err != nil {
		logger.Fatal(err)
	}
	defer postgresClientInstance.DB.Close()

	//create aggregator repository
	repo := pgRepo.NewRepository(postgresClientInstance)

	//create auth service
	authSrv := AuthService.NewAuthService(repo)

	//create auth hanler

	authHandler := authHandler.NewAuthHandler(authSrv, validator.New())

	// Register handler
	if err := pb.RegisterAuthSrvHandler(srv.Server(), authHandler); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

// we can add role
// what can we do when we have a validator that can be an admin or a requester?
// should we count on the time

// we know when the person is validator , we are selecting him and we can add them

// admins also are people that we are aware of

// but admins can be different in different cases

// a better solution would be having a request id in our table too

// but how can be possible?? and the process is too complicated

// let's consider it as simple as possible

//how much is the possibility of an admin being a requester???
//how much is the possibility of an validator being a requester?? -> by now this is the most impostant bottleneck
// and i am considering that admins and validators never can be switched
//i think we can seperate them by time and then when we are in need we can query the request service
//if we keep it here it seems to be a duplicate
