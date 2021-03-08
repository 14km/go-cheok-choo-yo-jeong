.PHONY : zip
zip :
	GOOS=linux go build main.go
	zip function.zip main

.PHONY : env
env :
	cp .env.example .env