#
# Arguments: all arguments are passed to svu
# Tasks
# - verify tests are running
# - verify goreleaser with "--snapshot --skip-publish --rm-dist"
# - replace cmd/version via sed
# 
FULL_TAG=$(svu $@)
NO_PREFIX_TAG=$(svu $@ --strip-prefix)
echo $FULL_TAG
echo $NO_PREFIX_TAG

# updating the version in cmd/version.go
go test ./... && \
goreleaser --snapshot --skip-publish --rm-dist && \
cp cmd/version.go.template cmd/version.go && \
sed -i -E "s/(version\W+=\W*)\"(.*?)\"/\1\"$NO_PREFIX_TAG\"/" cmd/version.go && \
git add cmd/version.go && git commit -m "chore: Release $FULL_TAG"
