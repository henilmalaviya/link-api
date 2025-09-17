.PHONY: docs build all

all:
	$(MAKE) build

docs:
	go run github.com/swaggo/swag/cmd/swag@latest init

build:
	$(MAKE) docs
	go build -o ./tmp/main .

