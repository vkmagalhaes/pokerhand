GOVERSION=1.7

run:
	go run cli/main.go

build:
	go build -v -o dist/pokerhand cli/main.go

clean:
	rm -rf dist

test:
	go test -cover -race $$(go list ./... | grep -v /vendor/)

bench:
	go test -cover -benchmem -race -bench=. $$(go list ./... | grep -v /vendor/)

run_docker:
	docker run --rm -v "$(PWD)":/go/src/github.com/vkmagalhaes/pokerhand -w /go/src/github.com/vkmagalhaes/pokerhand -ti golang:$(GOVERSION) bash -c "make run"

test_docker:
	docker run --rm -v "$(PWD)":/go/src/github.com/vkmagalhaes/pokerhand -w /go/src/github.com/vkmagalhaes/pokerhand golang:$(GOVERSION) bash -c "make test"

bench_docker:
	docker run --rm -v "$(PWD)":/go/src/github.com/vkmagalhaes/pokerhand -w /go/src/github.com/vkmagalhaes/pokerhand golang:$(GOVERSION) bash -c "make bench"

build_bin:
	docker run --rm -v "$(PWD)":/go/src/github.com/vkmagalhaes/pokerhand -w /go/src/github.com/vkmagalhaes/pokerhand golang:$(GOVERSION) bash -c "make build"
