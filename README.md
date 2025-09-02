# Chaos Game — N-gon

Minimal Go + Ebiten implementation of the Chaos Game for generating fractal patterns on arbitrary N-gons. 
Configure N (vertices), r (step fraction), and points-per-frame to explore classic Sierpinski triangles and a wide range of fractal/chaotic patterns.

## Usage
```shell
git clone https://github.com/anarhehest/chaos-game.git
go get github.com/hajimehoshi/ebiten/v2
```
**Run with default parameters (Sierpinski triangle):**
```shell
make
```
**Run with custom parameters:**
```shell
make build
./build/chaos-game -n 6 -r 0.55 -ppf 10000
```

## Notes

* Typical r = 0.5 produces classical results for triangles; other N/r combinations produce varied fractals.
* Increase -ppf for faster rendering; higher values use more CPU.