SERVICE=app

.PHONY: run

run:
	go build -o $(SERVICE) main.go
	clear		
	./$(SERVICE) run


.PHONY: count

gen:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

count:
	find . -name tests -prune -o -type f -name '*.go' | xargs wc -l


.DEFAULT_GOAL := run