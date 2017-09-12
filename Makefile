
.PHONY: all
build-ingress: clean
	mkdir -p build/
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/Ingress ./Ingress

clean:
	rm -rf build

all: build-ingress
	docker build -t ingress:latest -f docker/Dockerfile.scratch .

run:
	docker run -it ingress
