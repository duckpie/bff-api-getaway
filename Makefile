ENV=local
DOCKER_PATH=docker
CHERRYV=v0.1.0

.PHONY: run
run:
	sudo docker-compose -f $(DOCKER_PATH)/docker-compose.$(ENV).yml build
	sudo docker-compose -f $(DOCKER_PATH)/docker-compose.$(ENV).yml up


.PHONY: build
build:
	sudo docker-compose -f $(DOCKER_PATH)/docker-compose.$(ENV).yml build


.PHONY: test
test:
	sudo docker-compose -f $(DOCKER_PATH)/docker-compose.test.yml build
	sudo docker-compose -f $(DOCKER_PATH)/docker-compose.test.yml up \
		--remove-orphans \
		--abort-on-container-exit \
		--exit-code-from api_getaway_test


.PHONY: gen
gen:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate


.PHONY: upcherry
upcherry:
	go get -u github.com/duckpie/cherry@v$(CHERRYV)


.PHONY: count
count:
	find . -name tests -prune -o -type f -name '*.go' | xargs wc -l


.DEFAULT_GOAL := run