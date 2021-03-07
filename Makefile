#Docker = docker-compose
#
#.PHONY : run
#run :
#	$(Docker) build
#	$(Docker) up -d
#
#.PHONY : stop
#stop :
#	$(Docker) stop
#	$(Docker) rm -v -f
#
#.PHONY: build
#build:
#	$(DOCKER) docker-compose build
#

.PHONY : zip
zip :
	GOOS=linux go build main.go
	zip function.zip main

.PHONY : env
env :
	cp .env.example .env