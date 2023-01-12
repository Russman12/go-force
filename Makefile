codegen:
	openapi-generator-cli generate -c codegencfg.json
	go fmt
	go fmt
codegen-dry:
	openapi-generator-cli generate -c codegencfg.json --dry-run
init-project:
	git config --local core.hooksPath .githooks/
