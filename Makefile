
APP=mf-analyst
APP_VERSION:=$(shell cat .version)
APP_COMMIT:=$(shell git rev-parse HEAD)
APP_EXECUTABLE="./out/$(APP)"
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE) -ldflags "-X main.version=$(APP_VERSION) -X main.commit=$(APP_COMMIT)" cmd/*.go

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

lint:
	@for package in $(ALL_PACKAGES); do\
		echo	"==> Linting $$package";\
		golint $$package | { grep -vwE "exported (var|function|method|type|const) \S should have comment" || true;	}\
		done

build: compile fmt vet lint

clean:
	rm -rf out/

test:
	go test -p 1 -tags integration -cover $(ALL_PACKAGES)

test-cover:
	mkdir -p out/
	go test -covermode=count -coverprofile=coverage-all.out	./...
	@go tool cover -html=coverage-all.out -o out/coverage.html
	@go tool cover -func=coverage-all.out

ci:clean build test

build-run-server: build
	out/mf-analyst start