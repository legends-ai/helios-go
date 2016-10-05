all: clean build

clean:
	rm -f helios

cleanall:
	rm -rf helios vendor

build: genproto
	go build .

syncbuild: syncmodule genproto
	go build .

genproto:
	./proto/gen_go.sh

syncmodule:
	cd proto && git pull origin master

init:
	git submodule update --init

install: init genproto
	glide install

docker-build:
	docker build -t helios .

docker-push:
	docker tag helios:latest 096202052535.dkr.ecr.us-west-2.amazonaws.com/helios:latest
	docker push 096202052535.dkr.ecr.us-west-2.amazonaws.com/helios:latest

docker-build-dev:
	docker build -t helios:dev .

docker-push-dev:
	docker tag helios:dev 096202052535.dkr.ecr.us-west-2.amazonaws.com/helios:dev
	docker push 096202052535.dkr.ecr.us-west-2.amazonaws.com/helios:dev
