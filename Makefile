protobuf-stock-generate:
	protoc \
		-I api/proto \
		--go_out=pkg/broker \
		--go_opt=paths=source_relative \
		stockaggs.proto

protobuf-event-generate:
	protoc \
		-I api/proto \
		--go_out=pkg/kafka \
		--go_opt=paths=source_relative \
		event.proto

build-kafka:
	docker compose -f ./build/kafka/docker-compose.yml up --build -d
clean-kafka:
	docker compose -f ./build/kafka/docker-compose.yml down

sqlc-generate:
	sqlc generate -f api/sql/sqlc.yml
sqlc-check:
	sqlc compile -f api/sql/sqlc.yml