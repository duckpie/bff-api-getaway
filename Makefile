ENV=local

.PHONY: run
run:
	sudo docker-compose -f docker-compose.$(ENV).yml build
	sudo docker-compose -f docker-compose.$(ENV).yml up


.PHONY: build
build:
	sudo docker-compose -f docker-compose.$(ENV).yml build


.PHONY: test
test:
	sudo docker-compose -f docker-compose.test.yml build
	sudo docker-compose -f docker-compose.test.yml up --remove-orphans


.PHONY: gen
gen:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate


.PHONY: count
count:
	find . -name tests -prune -o -type f -name '*.go' | xargs wc -l


.DEFAULT_GOAL := run