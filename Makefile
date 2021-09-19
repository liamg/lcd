default: test

.PHONY: test
test:
	go test -race ./...

.PHONY: cyclo
cyclo:
	which gocyclo || go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	gocyclo -over 15 -ignore 'vendor/' .

.PHONY: vet
vet:
	go vet ./...


