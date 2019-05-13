#Makefile

all:
	go run cmd/db/main.go
	go run cmd/server/main.go

setup:
	go run cmd/db/main.go init
	go run cmd/db/main.go
	go install cmd/server/main.go

db-migrate:
	go run cmd/db/main.go

start:
	go run cmd/server/main.go

build:
	go build cmd/server/main.go

test:
	go test -v ./...

upgrade:
	go get -u
