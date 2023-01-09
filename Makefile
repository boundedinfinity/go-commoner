makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

generate:
	go generate ./...

test: generate
	go test ./...

commit:
	git add . || true
	git commit -m "$(m)" || true
	git push origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

publish: test
	@if ack replace go.mod ;then echo 'Remove the "replace" line from the go.mod file'; exit 1; fi
	make commit m=$(m)
	make tag tag=$(m)