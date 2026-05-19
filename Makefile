.PHONY: compose build clean

compose:
	docker compose up -d --build --no-deps gainsmith 

build:
	docker build -t gainsmith-test .

clean:
	docker compose kill
	docker compose rm -f
	-docker run --rm -v $(PWD):/app busybox chown -R $(shell id -u):$(shell id -g) /app/pb_data
	rm -rf pb_data
	-docker image rm gainsmith-gainsmith
