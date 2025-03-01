BIN=build/bin/app

.PHONY: tailwindcss sqlc templ

build: build/tailwind build/templ go.sum go.mod main.go
	GOOS=linux GOARCH=amd64 go build -o $(BIN)

build/tailwind: tailwindcss
	tailwindcss -i web/static/css/input.css -o web/static/css/output.css

build/templ: templ
	templ generate

# build/sqlc: sqlc $(wildcard *.sql) sqlc.yaml
# 	sqlc generate

clean:
	rm $(BIN)

cleanall:
	rm -rf $(shell dirname $(BIN))/*
