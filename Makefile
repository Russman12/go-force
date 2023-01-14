all: clean auth bulkv2 rest fmt

fmt:
	go fmt ./... >/dev/null
	go fmt ./... >/dev/null
clean:
	find pkg -type f -not \( -name .openapi-generator-ignore -o -name openapi.yaml \) -delete
auth:
	openapi-generator-cli generate -c codegencfg.json -i pkg/auth/api/openapi.yaml --package-name auth -o pkg/auth >/dev/null
bulkv2:
	openapi-generator-cli generate -c codegencfg.json -i pkg/bulkv2/api/openapi.yaml --package-name bulkv2 -o pkg/bulkv2 >/dev/null
rest:
	openapi-generator-cli generate -c codegencfg.json -i pkg/rest/api/openapi.yaml --package-name rest -o pkg/rest >/dev/null
init:
	git config --local core.hooksPath .githooks/
