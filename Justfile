m   := "updates"

list:
    just --list

push:
    git add . || true
    git commit -m "{{ m }}" || true
    git push origin master

commit:
	git add . || true
	git commit -m "{{ m }}" || true

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

publish:
	just commit m={{ m }}
	just tag tag={{ m }}


test:
	go test ./...

bootstrap:
	go install golang.org/x/perf/cmd/benchstat@latest
	go mod tidy

bench:
	go test -bench idiomatic/stringer/... -count 5 | tee new.txt
