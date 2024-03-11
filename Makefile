REPLACE=`semver tag`
.PHONY: build

test:
	@go test ./...

build:
	@echo "\e[34m\033[1m-> Building go binaries for supported platforms\033[0m\e[39m\n"
	rm -Rf dist || true
	go install github.com/goreleaser/goreleaser@latest
	GOPATH=$(HOME)/go PATH=$$PATH:$(HOME)/go/bin goreleaser build --snapshot
	@echo "\e[34m\033[1mDONE \033[0m\e[39m\n"


tag:
	@semver inc $(RELEASE)
	@echo "New release: `semver tag`"

git_tag:
	@git tag -a $(semver tag) -m "tagging $(semver tag)"
