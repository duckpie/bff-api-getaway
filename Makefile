SERVICE=app

.PHONY: run

run:
	sudo docker-compose -f docker-compose.local.yml up

.PHONY: count

gen:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

count:
	find . -name tests -prune -o -type f -name '*.go' | xargs wc -l


.DEFAULT_GOAL := run