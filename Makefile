build:
	go build

dev-dependencies:
	./scripts/dev-dependencies.sh

list-packages:
	go list ./...

test:
	./scripts/test.sh

watch:
	gin run main.go

db.create:
	soda create -a

db.migrate:
	soda migrate up

.PHONY: install watch test list-packages dev-dependencies
