# go-demo-service
Demo project for unit testing presentation @ GoJKT meetup

## This is a demo project to show examples of unit testing for GoJKT meetup

## Useful libraries:
    github.com/DATA-DOG/go-sqlmock
	github.com/golang/mock/gomock
	github.com/nsf/jsondiff
	github.com/stretchr/testify
    github.com/wolfcw/libfaketime

## Command cheatsheet
single package test in current directory:
>go test -cover . 

run test and generate coverprofile for all packages
>go test ./... -coverprofile cover.out 

calculate total coverage
>go tool cover -func cover.out 

create html to show line-by-line coverage
>go tool cover -html=cover.out
