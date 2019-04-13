PROJECT_NAME := pexels
PKG := "github.com/martinomburajr/pexels"
PKG_LIST := ($shell go list {$PKG}/...)

.PHONY: build test run

run: ##builds and runs the program and cleans the build file after
	@go build -i -v -o pexelss && ./pexelss

build: ##builds and runs the program
	@go build -i -v -o pexels

test: ##tests the available test files
	@go test -short ${PKG_LIST}

