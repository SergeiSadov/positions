PROGRAM=positions


all: help
help:
	@echo "build	    : Build app binary"
	@echo "run         : Run program."
	@echo "clean       : Remove auto-generated files."
	@echo "test        : Run tests."

build:
		go build -o ${PROGRAM} cmd/${PROGRAM}/main.go

clean:
		rm ${PROGRAM}

run:
		go run cmd/${PROGRAM}/main.go

test:
	go test ./tests -covermode=atomic -coverpkg=./... -coverprofile cover.out
