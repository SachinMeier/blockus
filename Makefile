
.PHONY: build
build:
	go build -o blockus cmd/blockus/*.go

.PHONY: build
run: build
	./blockus
