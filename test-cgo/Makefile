.PHONY: install
install:
	gcc -c foo.c
	gcc -shared -fPIC foo.o -o libfoo.so
	go build main.go && ./main
