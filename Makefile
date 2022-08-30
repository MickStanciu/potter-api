.PHONY: buildProtos
buildProtos:
	@echo "building proto files using buf..."
	@rm -rf gen
	@buf generate

buildProtos2:
	@echo "building proto files using protoc..."
	@rm -rf gen2
	@mkdir gen2
	@protoc -I=./proto --go_out=gen2 $(shell find ./proto -name '*.proto')
