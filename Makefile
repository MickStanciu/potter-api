.PHONY: buildProtos
buildProtoFiles:
	@echo "building proto files using protoc..."
	@rm -rf potterapi/internal/gen
	@mkdir -p potterapi/internal/gen
	@protoc -I=./proto \
		--go_out=potterapi/internal/gen --go_opt=paths=source_relative \
		--go-grpc_out=potterapi/internal/gen --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out potterapi/internal/gen --grpc-gateway_opt paths=source_relative \
		$(shell find ./proto -name '*.proto')
	@echo "... done"
