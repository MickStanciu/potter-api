.PHONY: buildProtos
buildProtos:
	@echo "building proto files..."
	@rm -rf gen
	@buf generate
