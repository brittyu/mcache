.PHONY: install

install:
	gcc -c -fPIC foo.c
	gcc -shared -fPIC foo.o -o libfoo.so
	go build main.go && ./main

remove:
	rm foo.o libfoo.so main
