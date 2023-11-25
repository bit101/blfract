# blfract

Static and animated fractals. 

## Basics
Each render consists of:

1. An iterator which combines all the other components.
2. A fractal algorithm. e.g. Mandelbrot.
3. A warping algorithm, which warps the complex plane in some way during the render.
4. A coloring algorithm, which determins what color to give each pixel based on the results of the fractal algorithm.

## Dependencies
- https://github.com/bit101/blcairo
- https://github.com/bit101/bitlib

As configured, this project is using a `go.work` file that expects the above two dependencies to be checked out in the same parent directory as this project.

Alternately, you can remove the `go.work` file and `go get ...` the dependencies instead.

## Todo
- More fractal algorithms
- More warpers
- More colorizers
- Allow configs to be saved and loaded
- Lots of optimizations
- Improved docs
- Tests
