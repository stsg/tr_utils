B=$(shell git rev-parse --abbrev-ref HEAD)
BRANCH=$(subst /,-,$(B))
GITREV=$(shell git describe --abbrev=7 --always --tags)
DATE=$(shell date +%Y%m%d-%H:%M:%S)
REV=$(GITREV)-$(BRANCH)-$(DATE)

info:
	- @echo "revision $(REV)"

build: info
	@ echo
	@ echo "Compiling Binary"
	@ echo
	go build -ldflags "-X main.revision=$(REV) -s -w" -o bin/trl trl/main.go

docker:
	docker build -t starky/gophkeeper:master .

clean:
	@ echo
	@ echo "Cleaning"
	@ echo
	rm bin/tr || true
	rm bin/clip

tidy:
	@ echo
	@ echo "Tidying"
	@ echo
	go mod tidy

run:
	go run tr/main.go

lint:
	@ echo
	@ echo "Linting"
	@ echo
	golangci-lint run

test:
	@ echo
	@ echo "Testing"
	@ echo
	go test ./...

.PHONY: *
