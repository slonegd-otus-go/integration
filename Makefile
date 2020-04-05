all: 
	go build

up:
	docker-compose up -d

down:
	docker-compose down

client:
	./integration client --a=3 --b=4

example:
	docker-compose build; \
	docker-compose up -d; \
	./integration client --a=3 --b=4;\
	docker-compose down;