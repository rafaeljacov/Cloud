BIN=build/bin/app

.PHONY: build tailwindcss sqlc templ zip

build: build/tailwind build/templ go.sum go.mod main.go zip
	GOOS=linux GOARCH=amd64 go build -o $(BIN)
	@mkdir -p app/web
	@cp -r build app
	@cp -r web/static app/web
	zip -r build/app.zip app
	@rm -rf app

build/tailwind: tailwindcss
	tailwindcss -i web/static/css/input.css -o web/static/css/output.css

build/templ: templ
	templ generate

# build/sqlc: sqlc sqlc.yaml
# 	sqlc generate

clean:
	@rm -rf build
	@echo 'Build successfully cleaned.'
