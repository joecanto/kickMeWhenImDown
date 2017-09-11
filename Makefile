
.PHONY: all
all: clean
	mkdir -p build/
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/ingress ./ingress

clean:
	rm -rf build

build:
	docker build -t ingress:latest -f docker/Dockerfile.scratch .

run:
	docker run -it ingress 
