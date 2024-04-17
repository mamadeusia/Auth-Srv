# AuthSrv

Authentication service is auth service and all micro services that need auth they need to connect to AuthSrv.


## Go-micro


Install Go-micro from the following repository:

```sh
    https://github.com/go-micro
```


## RUN

```
cp sample.env .env
```


postgres viewer 

```
http://127.0.0.1:10000/?pgsql=db-auth&username=postgres&db=postgres&ns=public
```

export env
```
export $(xargs < .env)
go-micro run
```


# Migration Database


for installation migration command

```
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
```


```
$ go get -u -d github.com/golang-migrate/migrate/cmd/migrate
$ cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
$ git checkout $TAG  # e.g. v4.1.0
$ # Go 1.15 and below
$ go build -tags 'postgres' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
$ # Go 1.16+
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$TAG
```


```
migrate -path ./sql/migration -database postgresql://postgres:postgres@0.0.0.0:3002/?sslmode=disable up
```