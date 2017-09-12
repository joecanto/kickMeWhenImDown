
ingress:
	mkdir -p build/
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/ingress ./Ingress

aggregator:
	mkdir -p build/
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/aggregator ./Aggregator

clean:
	rm -rf build

.PHONY: all
all: clean aggregator ingress
    docker build -t aggregator:latest -f docker/Dockerfile.aggregator .
	docker build -t ingress:latest -f docker/Dockerfile.ingress .
