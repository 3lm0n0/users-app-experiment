##################################################
# Usage:
##################################################
# make          # compile all binary
# make hello    # prints hello
# make init     # creates module
# make setup    # sets up the microservice
# make build    # builds the microservice
# make run      # runs the microservice
# make test     # run tests
# make clean    # remove ALL binaries and objects

MICROSERVICE_NAME=user

.PHONY:= hello init setup build run test clean
.DEFAULT_GOAL:= setup build run

hello:
	echo "Hello"

init:
	@echo "=> Go module ${MICROSERVICE_NAME} initializing"
	@go mod init '${MICROSERVICE_NAME}'
	@echo "=> Go module initialized"

setup:
	@echo "=> Stetting microservice"
	@export GOSUMDB=off
	@go mod tidy
	@go mod download
	@echo "=> Setup completed"

build:
	@echo "=> Building microservice"
	@go build -o ./bin/${MICROSERVICE_NAME}
	@echo "=> Building completed"
	
run:
	./bin/${MICROSERVICE_NAME}

test:
	go test -v ./...

clean:
	@echo "Cleaning up all binaries, objects and sum ..."
	@go clean
	@rm -rvf *.o ./bin/${MICROSERVICE_NAME} go.sum
	@echo "Cleaning completed"