makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

process: purge
	go generate ./...

purge:
	rm -rf file_extention/*.gen.go
	rm -rf file_extention/*.enum.go
	rm -rf mime_type/*.gen.go	
	rm -rf mime_type/*.enum.go

test: process
	go test ./...

publish:
	git add . || true
	git commit -m "$(tag)"
	git push origin master
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)
