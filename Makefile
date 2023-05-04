GOPATH ?= $(HOME)/go
PROTOS_SRC = $(wildcard ./protos/*.proto)


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
	@for proto_file in $(PROTOS_SRC); do \
		protoc -I=. -I=$(GOPATH)/src --gogofaster_out=. $$proto_file; \
	done
	goimports -w .