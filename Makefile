REPLACE=`semver tag`

# Build
# Usage:
#   make build RELEASE=major
#   make build RELEASE=minor
#   make build RELEASE=patch

build:
	@echo "Building release: `semver tag`"
	@go build  -ldflags "-X main.version `semver tag`" oss.go

tag:
	@semver inc $(RELEASE)
	@echo "New release: `semver tag`"

# Tag git with last release
git_tag:
	@git tag -a $(semver tag) -m "tagging $(semver tag)"
