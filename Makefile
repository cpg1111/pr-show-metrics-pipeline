VERSION=0.1.0

all: build

build:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/influxdata/influxdb/client/v2/...
	mkdir -p www/
	go-bindata -pkg www -o www/assets.go assets/...
#	glide install
	go build -o ./dist/pr-metrics ./srv/main.go

test:
	go test ./...

install:
	ln -s `pwd`/dist/pr-metrics /opt/bin/pr-metrics

docker:
	docker build -t prmetrics:build -f Dockerfile_build .
	docker run --rm -v `pwd`:/go/src/github.com/cpg1111/pr-show-metrics-pipeline prmetrics:build
	docker build -t prmetrics:${VERSION} .

docker-run:
	docker run ${RUN_ARGS} -e HOST=${SRV_HOST} -e PROTO=${PROTO} -e PORT=${SRV_PORT} -e DB_CONNECT_STRING=${DB_CONNECT_STRING} prmetrics:${VERSION}
