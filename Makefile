build:
	go build -o go-waitgroup github.com/pieterclaerhout/go-waitgroup/cmd

run: build
	./go-waitgroup

test:
	go test -coverage ./...