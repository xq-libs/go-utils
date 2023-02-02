VERSION = `cat VERSION`

release:
	git fetch --tags
	git tag -a v${VERSION} -m "Release version ${VERSION}"
	git push --tags