
build-ingress:
	mkdir -p build/
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/ingress ./ingress

build-aggregator:
	mkdir -p build/
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/insult-engine ./insult-aggregator-grpc

clean:
	rm -rf build

all: clean build-ingress build-aggregator
	docker build -t ingress:latest -f docker/Dockerfile.ingress .
	docker build -t aggregator:latest -f docker/Dockerfile.aggregator .