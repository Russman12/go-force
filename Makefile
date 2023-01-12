codegen:
	openapi-generator-cli generate -c codegencfg.json -i pkg/bulkv2/api/openapi.yaml --package-name bulkv2 -o pkg/bulkv2
	openapi-generator-cli generate -c codegencfg.json -i pkg/auth/api/openapi.yaml --package-name auth -o pkg/auth
	openapi-generator-cli generate -c codegencfg.json -i pkg/rest/api/openapi.yaml --package-name rest -o pkg/rest
	go fmt ./...
	go fmt ./...
codegen-dry:
	openapi-generator-cli generate -c codegencfg.json --dry-run
init-project:
	git config --local core.hooksPath .githooks/
