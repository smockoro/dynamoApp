all: build

.PHONY: build
build:
	GOOS=linux go build -o bin/create table/create/main.go
	GOOS=linux go build -o bin/delete table/delete/main.go
	GOOS=linux go build -o bin/import data/import/main.go
	GOOS=windows go build -o bat/create.exe table/create/main.go
	GOOS=windows go build -o bat/delete.exe table/delete/main.go
	GOOS=windows go build -o bat/import.exe data/import/main.go
