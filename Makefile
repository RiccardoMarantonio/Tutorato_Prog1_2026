.PHONY: test test-topic run fmt

test:
	go test ./...

test-topic:
	@test -n "$(TOPIC)" || (echo "Uso: make test-topic TOPIC=01_if_for" && exit 1)
	go test ./exercises/$(TOPIC)/...

run:
	@test -n "$(EX)" || (echo "Uso: make run EX=exercises/01_if_for/easy" && exit 1)
	go run ./$(EX)

fmt:
	gofmt -w .
