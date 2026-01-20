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
