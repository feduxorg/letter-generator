build:
.PHONY: all
all: build

.PHONY: build
build:
	go build -o dist/lg github.com/fedux-org/letter-generator-go/cmd/lg

.PHONY: test
test:
ifeq ($(DEBUG), 1)
	  go test ./... -d
else
	  go test ./...
endif


.PHONY: clean
clean:
	rm -rf dist/
