
build: hooks  ## build, install, lint, test
	./scripts/build.sh

test: build

cov:
	go test -coverprofile=coverage.out 
	go tool cover -html=coverage.out

clean:  ## clean up time
	rm -rf dist/ bin/
	go clean ./...
	git gc --aggressive

.PHONY: help ci bench

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

