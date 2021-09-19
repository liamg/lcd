default: test

.PHONY: test deploy

test:
	go test -race ./...

deploy: 
	./scripts/deploy.sh

