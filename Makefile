
build: hooks  ## build, install, lint
	go build ./...
	go install ./cmd/cssformat ./cmd/csscut ./cmd/csscount
	gometalinter \
                 --vendor \
                 --deadline=60s \
                 --disable-all \
		 --enable=vet \
                 --enable=golint \
                 --enable=gofmt \
                 --enable=goimports \
                 --enable=gosimple \
                 --enable=staticcheck \
                 --enable=ineffassign \
                 --exclude=/usr/local/go/src/net/lookup_unix.go \
                 ./...
	go test .

# nothing special to do for travis-ci.org
ci: build

test:  ## run all tests
	go test .

clean:  ## clean up time
	rm -rf dist/ bin/
	go clean ./...
	git gc --aggressive

.PHONY: help ci console bench

# https://www.client9.com/automatically-install-git-hooks/
.git/hooks/pre-commit: scripts/pre-commit.sh
	cp -f scripts/pre-commit.sh .git/hooks/pre-commit
.git/hooks/commit-msg: scripts/commit-msg.sh
	cp -f scripts/commit-msg.sh .git/hooks/commit-msg
hooks: .git/hooks/pre-commit .git/hooks/commit-msg  ## install git precommit hooks


# https://www.client9.com/self-documenting-makefiles/
help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ {\
	printf "\033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
.DEFAULT_GOAL=help
.PHONY=help

