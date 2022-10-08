BINARY_NAME=youtube-cli
DIST=bin

build:
	GOOS=darwin GOARCH=amd64 go build -o ${DIST}/${BINARY_NAME}-mac main.go
	GOOS=darwin GOARCH=arm64 go build -o ${DIST}/${BINARY_NAME}-mac-silicon main.go
	GOOS=linux GOARCH=amd64 go build -o ${DIST}/${BINARY_NAME}-linux main.go
	GOOS=windows GOARCH=amd64 go build -o ${DIST}/${BINARY_NAME}-windows main.go

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

lint:
	golangci-lint run

clean:
	go clean
	rm ${DIST}/${BINARY_NAME}*