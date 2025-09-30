.PHONY: all build run clean

all: build run

build:
	go get github.com/hajimehoshi/ebiten/v2
	go build -o build/chaos-game
	@echo "Built: build/chaos-game"

run:
	build/chaos-game

clean:
	rm -rf build