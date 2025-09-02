.PHONY: all build run clean

all: build run

build:
	go build -o build/chaos-game
	@echo "Built: build/chaos-game"

run:
	build/chaos-game

clean:
	rm -rf build