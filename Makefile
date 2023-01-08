build:
	go build ./...

run-data-gen: build
	go run cmd/datagen/main.go --env=.env

generate-contracts:
	@./scripts/contractGen.sh

db-migrate:
	tern migrate -c db/tern.conf -m db