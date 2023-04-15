.PHONY: serve
serve: network
	cd devzone \
	&& docker-compose up

.PHONY: build
build:
	go build -o ./build/wave ./cmd/wave

.PHONY: run
run:
	./build/wave -config ./.config.toml
