makefile_dir	:= $(abspath $(shell pwd))
m				:= "updates"

.PHONY: list bootstrap init build clean

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	find . -name '*.enum.go' -type f -delete

generate:
	go generate ./...

test: generate
	go test -v ./...

commit: test
	git add . || true
	git commit -m "$(m)" || true

push: commit
	git push origin master

pull:
	git fetch
	git pull origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

tag-list:
	git tag --list | sort -V

publish: generate test
	@if ack replace go.mod ;then echo 'Remove the "replace" line from the go.mod file'; exit 1; fi
	make commit m=$(m)
	make tag tag=$(m)
