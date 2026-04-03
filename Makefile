.PHONY: test test-exercise run fmt

test:
	go test ./...

test-exercise:
	@test -n "$(EX)" || (echo "Uso: make test-exercise EX=01_if_for_switch_functions/01_numeri_pari" && exit 1)
	go test ./exercises/$(EX)

test-exercise-v:
	@test -n "$(EX)" || (echo "Uso: make test-exercise-v EX=01_if_for_switch_functions/01_numeri_pari" && exit 1)
	go test -v ./exercises/$(EX)

run:
	@test -n "$(EX)" || (echo "Uso: make run EX=01_if_for_switch_functions/01_numeri_pari" && exit 1)
	go run ./exercises/$(EX)

fmt:
	gofmt -w .
