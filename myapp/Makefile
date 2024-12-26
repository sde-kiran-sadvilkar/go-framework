BINARY_NAME=sopeApp.exe

## build: builds all binaries
build:
	@go mod vendor
	@go build -o tmp/${BINARY_NAME} .
	@echo Sope built!

run: build
	@echo Staring Sope...
	@.\tmp\${BINARY_NAME} &
	@echo Sope started!

clean:
	@echo Cleaning...
	@go clean
	@del .\tmp\${BINARY_NAME}
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run
	
stop:
	@echo "Starting the front end..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped Sope

restart: stop start

docker_up:
	docker-compose up -d

docker_down:
	docker-compose down
