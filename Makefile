GIT_COMMIT := $(shell git describe --tags | tr -d v)
GIT_DIRTY := $(if $(shell git status --porcelain | grep -qE .),~$(shell git rev-parse --short HEAD))
VERSION := $(GIT_COMMIT)$(GIT_DIRTY)
ARCH ?= amd64
PACKAGE := winterhill-web

.PHONY: all
all: $(PACKAGE)

test:
	go vet ./...
	go test -race ./...

$(PACKAGE): app/winterhill-web/dist test
	CGO_ENABLED=1 go build -o $(PACKAGE) app/$(PACKAGE)/main.go

app/winterhill-web/dist:
	cd web && npm install && npm run build && mv dist ../app/winterhill-web

.PHONY: clean
clean:
	rm -rf app/winterhill-web/dist $(PACKAGE) build

