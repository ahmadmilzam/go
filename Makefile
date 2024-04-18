APP_EXECUTABLE="bin/app"

build:
	mkdir -p bin/
	go build -o $(APP_EXECUTABLE) cmd/app/main.go

run: build
	./bin/app start

migrate-create:
	go run cmd/app/main.go migrate create --filename $(FILENAME)

migrate-up:
	go run cmd/app/main.go migrate up

migrate-down:
	go run cmd/app/main.go migrate down

test:
	go test ./...

copy-config:
	cp ./config/config.yaml.example ./config/config.yaml