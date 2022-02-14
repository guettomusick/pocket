.SILENT:

.PHONY: help
		# build build_and_watch \
		# client \
		# compose_and_watch compose_and_watch_d compose_and_watch_single \
		# neo_d

help:
	printf "Available targets\n\n"
	awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "%-30s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)


## Build consensus/main.go
build:
	go build -v cmd/consensus/main.go

## Continous build of consensus/main.go as files changel.
build_and_watch:
	/bin/sh ${PWD}/scripts/watch_build.sh

## Run a client daemon which is only used for debugging.
client_start:
	docker-compose -f deployments/docker-compose.yaml up -d client

## Connect to the client daemon running in a background docker container.
client_connect:
	docker exec -it client /bin/bash -c "go run cmd/client/*.go"

## Attached docker compose of all the services except for neo4j w/ hot reload.
compose_and_watch:
	docker-compose -f deployments/docker-compose.yaml up --force-recreate node1.consensus node2.consensus node3.consensus node4.consensus

## Detached docker compose of all the services except for neo4j  w/ hot reload.
compose_and_watch_d:
	docker-compose -f deployments/docker-compose.yaml up -d --force-recreate --scale neo4j=0 --scale client=0

## Detached deployment of the neo4j container.
neo_d:
	docker-compose -f deployments/docker-compose.yaml up -d neo4j

## Use `mockgen` to generate mocks used for testing.
generate_mocks:
	mockgen --source=pkg/shared/modules/pocket_module.go -destination=pkg/shared/modules/mocks/pocket_module_mock.go

	mockgen --source=pkg/shared/modules/persistence_module.go -destination=pkg/shared/modules/mocks/persistence_module_mock.go -aux_files=github.com/pocket-network/pkg/shared/modules=./pkg/shared/modules/pocket_module.go
	mockgen --source=pkg/shared/modules/utility_module.go -destination=pkg/shared/modules/mocks/utility_module_mock.go -aux_files=github.com/pocket-network/pkg/shared/modules=./pkg/shared/modules/pocket_module.go
	mockgen --source=pkg/shared/modules/p2p_module.go -destination=pkg/shared/modules/mocks/p2p_module_mock.go -aux_files=github.com/pocket-network/pkg/shared/modules=./pkg/shared/modules/pocket_module.go
	mockgen --source=pkg/p2p/p2p_types/network.go -destination=pkg/p2p/p2p_types/mocks/network_mock.go
	mockgen --source=pkg/shared/modules/consensus_module.go -destination=pkg/shared/modules/mocks/consensus_module_mock.go -aux_files=github.com/pocket-network/pkg/shared/modules=./pkg/shared/modules/pocket_module.go

# TODO: Mocks

# mockgen --source=pkg/pocket/node.go -destination=pkg/pocket/mocks/node_mock.go

## Use `protoc` to generate consensus .go files from .proto files.
generate_protos:
	$(eval types_dir := "pkg/types/")
	protoc -I=${types_dir}/protos --go_out=${types_dir} ${types_dir}/protos/*.proto

## V1 Integration - Use `protoc` to generate consensus .go files from .proto files.
v1_generate_protos:
	protoc -I=./shared/protos --go_out=./shared shared/protos/*.proto

# Good stack overflow page for organizing tests: https://stackoverflow.com/questions/25965584/separating-unit-tests-and-integration-tests-in-go
# Setting cout=1 for tests to avoid caching.

## Run all the unit tests
test_all: # generate_mocks
	go test ./...

## Run unit tests for consensus
test_hotstuff: # generate_mocks
	go test -count=1 -v -run Hotstuff ${PWD}/pkg/pocket/consensus_tests

## Run unit tests for the pacemaker
test_pacemaker: # generate_mocks
	go test -count=1 --tags=flaky -v -run Pacemaker ${PWD}/pkg/pocket/consensus_tests

## Run unit tests for crypto related code paths
test_crypto:
	go test -count=1 -v -run Crypto ${PWD}/pkg/pocket/consensus_tests

## Run the leader election unit tests
test_leader_election:
	go test --count=1 -v ./pkg/consensus/leader_election

## Run the sortition unit tests
test_sortition:
	go test --count=1 -v ./pkg/consensus/leader_election/sortition

## Benchmarking the sortition unit tests
benchmark_sortition:
	go test --count=1 -v ./pkg/consensus/leader_election/sortition -bench=.