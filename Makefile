codegen:
	openapi-generator-cli generate -c codegencfg.json
	go fmt
codegen-dry:
	openapi-generator-cli generate -c codegencfg.json --dry-run