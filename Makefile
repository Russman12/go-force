all: clean bulkv2 rest tooling

fmt:
	go fmt ./... >/dev/null
	go fmt ./... >/dev/null
clean:
	find pkg -type f -not \( -name .openapi-generator-ignore -o -name openapi.yaml \) -delete
bulkv2:
	openapi-generator-cli generate -c codegencfg.json -i pkg/bulkv2/api/openapi.yaml --package-name bulkv2 -o pkg/bulkv2 >/dev/null
	go fmt ./pkg/bulkv2/... >/dev/null
	go fmt ./pkg/bulkv2/... >/dev/null
rest:
	openapi-generator-cli generate -c codegencfg.json -i pkg/rest/api/openapi.yaml --package-name rest -o pkg/rest >/dev/null
	go fmt ./pkg/rest/... >/dev/null
	go fmt ./pkg/rest/... >/dev/null
tooling:
	openapi-generator-cli generate -c codegencfg.json -i pkg/tooling/api/openapi.yaml --package-name tooling -o pkg/tooling >/dev/null
	go fmt ./pkg/rest/... >/dev/null
	go fmt ./pkg/rest/... >/dev/null
init:
	git config --local core.hooksPath .githooks/
watcher:
	scripts/watcher.sh
