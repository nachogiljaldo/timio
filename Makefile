all: deps test
.PHONY: all deps test run

deps:
	echo "Downloading dependencies..."
	go get -u github.com/kardianos/govendor
	go get -u github.com/gorilla/mux
	go get -u github.com/pborman/uuid
	go get -u gopkg.in/yaml.v2

test:
	echo "Running tests..."
	govendor test -v -p=1 +local

run:
	echo "Starting..."
	go run main/main.go