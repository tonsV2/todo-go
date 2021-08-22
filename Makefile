tag ?= latest
clean-cmd = docker compose down --remove-orphans --volumes

keys:
	openssl genpkey -algorithm RSA -out ./rsa_private.pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in ./rsa_private.pem -pubout -out ./rsa_public.pem

build-image:
	IMAGE_TAG=$(tag) docker compose build prod

push-image:
	IMAGE_TAG=$(tag) docker compose push prod

di:
	wire gen ./pgk/di

build-dev:
	docker compose build dev

dev:
	docker compose up redis database dev

build-test:
	docker compose build test

test: clean
	docker compose run --no-deps test
	$(clean-cmd)

clean:
	$(clean-cmd)

.PHONY: keys build-image push-image di build-dev launch-dev build-test test clean
