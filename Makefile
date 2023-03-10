.PHONY: buildProtoFiles
buildProtoFiles:
	@echo "building proto files using protoc..."
	@rm -rf potterapi/internal/gen
	@mkdir -p potterapi/internal/gen
	@protoc --proto_path=./proto \
		--go_out=potterapi/internal/gen --go_opt=paths=source_relative \
		--go-grpc_out=potterapi/internal/gen --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out potterapi/internal/gen --grpc-gateway_opt paths=source_relative \
		$(shell find ./proto -name '*.proto')
	@echo "... done"

.PHONY: run
run:
	@echo "starting server..."
	@go run potterapi/cmd/main.go

.PHONY: httpExample
httpExample:
	@echo "getting the Good Guys"
	@curl localhost:8080/v1/goodguys | json_pp

.PHONY: httpExample2
httpExample2:
	@echo "getting the Bad Guys"
	@curl localhost:8080/v1/badguys | json_pp

.PHONY: grpcExample
grpcExample:
	@echo "getting the Good Guys"
	@grpcurl -plaintext localhost:9000 potter.HogwartsService/GetTheGoodGuys

