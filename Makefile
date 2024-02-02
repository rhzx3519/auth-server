SHELL=/bin/bash

.PHONY: build clean dist clean-images
docker=
arch=
os=
build: clean
	sh build.sh $(docker) $(arch) $(os)

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

docker=
run:
ifdef docker
	docker compose up --force-recreate
else
	./bin/auth-server
endif

push:
	docker push rhzx3519/auth-server:latest
