
.PHONY: build
build:
	go build -o blockus cmd/*.go

.PHONY: build
run: build
	./blockus
