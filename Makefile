SHELL=/bin/bash

.PHONY: build clean dist clean-images
docker=
build: clean
	sh build.sh $(docker)

clean:
	rm -fr ./build
	rm -fr ./bin
	rm -fr ./dist
	mkdir -p ./build
	docker compose rm -f

clean-images:
	docker rmi -f $(docker images | awk 'NR>1{if($2=="<none>")print$3}')

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

