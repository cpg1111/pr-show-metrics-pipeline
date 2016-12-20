VERSION=0.1.0

all: build

build:
	glide install
	mkdir -p www/
	go-bindata -o www/assets.go assets/
	go build -o ./dist/pr-metrics main.go

test:
	go test ./...

install:
	ln -s `pwd`/dist/pr-metrics /opt/bin/pr-metrics

docker:
	docker build -t prmetrics:build .
	docker run --rm -v `pwd`:/go/src/github.com/cpg1111/pr-show-metrics-pipeline prmetrics
	docker build -t premetrics:${VERSION}

docker-run:
	docker run ${RUN_ARGS} -e DB_CONNECT_STRING=${DB_CONNECT_STRING} prmetrics
