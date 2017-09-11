
.PHONY: all
all:
	mkdir -p build/
	GOOS=darwin CGO_ENABLED=0 go build -a -installsuffix cgo \
		-o build/ingress ./ingress