GOPATH ?= $(HOME)/go
PROTOS_SRC = $(wildcard ./api/*.proto)


.PHONY: run
run:
	@echo "Running..."
	docker-compose up

.PHONY: stop
stop:
	@echo "Stopping..."
	docker-compose down

.PHONY: generate
generate:
	@echo "Generating..."
	protoc -I=. -I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --gogofaster_out=. api/models.proto