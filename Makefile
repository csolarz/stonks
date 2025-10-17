test:
	go test -race -covermode=atomic -v ./... -coverprofile=coverage.out

cover: test
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

lint:
	golangci-lint run;