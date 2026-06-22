s9-database:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/s9-database

clean:
	rm s9-database

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	s9-database \
	clean \
	test \
	lint