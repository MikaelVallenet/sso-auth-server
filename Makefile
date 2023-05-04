GOPATH ?= $(HOME)/go
PROTOS_SRC = $(wildcard ./api/*.proto)


.PHONY: run
run:
	@echo "Running..."
	docker-compose up --build

.PHONY: stop
stop:
	@echo "Stopping..."
	docker-compose down

.PHONY: generate
generate:
	@echo "Generating..."
	protoc -I=. -I=$(GOPATH)/src --gogofaster_out=. api/models.proto
	goimports -w .