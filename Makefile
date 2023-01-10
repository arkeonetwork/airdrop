build:
	go build ./...

run-indexer: build
	go run cmd/indexer/main.go --env=./docker/dev/docker.env

generate-contracts:
	@./scripts/contractGen.sh

db-migrate:
	tern migrate -c db/tern.conf -m db