codegen:
	openapi-generator-cli generate -c codegencfg.json
codegen-dry:
	openapi-generator-cli generate -c codegencfg.json --dry-run