all: build

build: main

main: main.go
	go build -o main main.go

install: build
	go install 

clean:
	rm -f main
