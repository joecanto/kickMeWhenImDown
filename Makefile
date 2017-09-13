
build-ingress:
	mkdir -p build/
	GOOS=darwin CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/ingress ./Ingress

build-aggregator:
	mkdir -p build/
	GOOS=darwin CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/aggregator ./Aggregator

clean:
	rm -rf build

.PHONY: all
all: clean build-aggregator build-ingress
	docker build -t aggregator:latest -f docker/Dockerfile.aggregator .
	docker build -t ingress:latest -f docker/Dockerfile.ingress .
