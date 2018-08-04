get:
	go get -v ./...

test:
	go test ./...

build:
	cd serverless && $(MAKE) build

deploy-dev:
	cd serverless && $(MAKE) deploy-dev

deploy-prod:
	cd serverless && $(MAKE) deploy-prod
