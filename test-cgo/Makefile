.PHONY: install

linuxinstall:
	gcc -c -fPIC foo.c
	gcc -shared -fPIC foo.o -o libfoo.so
	sudo cp libfoo.so /usr/lib
	go build main.go && ./main

macosinstall:
	gcc -c -fPIC foo.c
	gcc -shared -fPIC foo.o -o libfoo.so
	go build main.go && ./main

remove:
	rm foo.o libfoo.so main
	sudo rm /usr/lib/libfoo.so
