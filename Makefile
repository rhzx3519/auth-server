SHELL=/bin/bash

.PHONY: build clean dist
docker=
build: clean
	sh build.sh $(docker)

clean:
	rm -fr ./build
	rm -fr ./bin
	rm -fr ./dist
	mkdir -p ./build
	docker compose rm -f

dist:
	rm -fr dist/
	mkdir -p dist
	cp -R ./scripts dist/
	cp Makefile dist/
	cp .env dist/
	cp compose.yaml dist/

run:
	./bin/auth-server

run-docker:
	docker compose up --force-recreate
	#docker compose rm -f

