
GOOS := linux
GOARCH := amd64
CGO_ENABLED := 1

TARGET := golden-${GOOS}-${GOARCH}

.DEFAULT_GOAL := all

.PHONY: all
all:
	make clean
	make build

.PHONY: clean
clean:
	rm -f ${TARGET}
	rm -f mod.sum	

.PHONY: check
check:
	go get -v ./...
	#go generate
	go test ./...

.PHONY: build
build:
	go get -u ./...
	go get -v ./...
	#go generate
	go build -o ${TARGET} ./cmd/golden
