REPLACE=`semver tag`

init:
	@export GOPATH=~/go:`pwd`

install:
	go get github.com/mitchellh/gox
	go get github.com/dotcypress/phonetics
	go get github.com/stretchr/testify

test:
	@go test github.com/halleck45/...

build: test
	@echo "Building release: `semver tag`"
	@gox -build-toolchain -ldflags "-X main.version `semver tag`" oss.go

publish:
	@./publish.sh


tag:
	@semver inc $(RELEASE)
	@echo "New release: `semver tag`"

git_tag:
	@git tag -a $(semver tag) -m "tagging $(semver tag)"
