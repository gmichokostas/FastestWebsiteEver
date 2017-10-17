all: server/index.html server/cpkthttp server/gopkthttp

server/index.html: client/index.html
	./utils/deflate.py client/index.html server/index.html

server/cpkthttp: server/c/main.c
	gcc -O2 -o server/cpkthttp server/c/main.c

server/gopkthttp: server/go/main.go
	go build -o server/gopkthttp server/go/main.go

clean:
	rm server/index.html
	rm server/cpkthttp
	rm server/gopkthttp
