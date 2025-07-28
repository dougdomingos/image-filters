# Image Filters

An image processing application focused on performance, with multiple filters
(e.g., Binarization, Sobel, Gaussian Blur) implemented using only built-in
features of Go's standard library.

## Features

- CLI and REST API entrypoints
- Support for PNG and JPEG images
- Configurable performance parameters (e.g., worker count)
- Thread-safe filter implementations
- Composable filters, allowing multiple transformations at once
- Built-in benchmark suites

## Installation and Usage

### System requirements

- Go v1.17+
- Git and GNU Make (for usage through source code)

### Installing the package with Go CLI

To install the package using Go's package manager, you can use `go install`:

```bash
# Installs the CLI package
go install github.com/dougdomingos/image-filters/cmd/cli@latest

# Installs the API package
go install github.com/dougdomingos/image-filters/cmd/api@latest
```

Once you've installed the package, you can use one of the following commands:

```bash
# For CLI package, run:
image-filters-cli -img [path/to/image.png|jpg] -outDir [path/to/output_dir] -filters [filter1,filter2,...] [-listPipelines]

# For API package, run:
image-filters-api
```

### Cloning the repository

If you want to modify or build the application from source, you can clone the
repository with Git:

```bash
git clone https://github.com/dougdomingos/image-filters.git
```

Then, you can use `make` to run some commands:

```bash
# Run the CLI entrypoint
make run-cli IMG=path/to/image OUT_DIR=path/to/output_dir FILTERS=filter1,filter2,...

# Start the REST API server
make run-api

# List the avaliable filter
make list-filters

# Run a benchmark test of a list of filters with different worker configuration
make bench FILTERS=filter1,filter2,... IMG_SIZE=size_in_pixels

# Run a benchmark test while profiling CPU usage
make bench-cpuprof FILTERS=filter1,filter2,... IMG_SIZE=size_in_pixels

# Run a benchmark test while profiling memory usage
make bench-memprof FILTERS=filter1,filter2,... IMG_SIZE=size_in_pixels

# Build the CLI binary. The binary will be named "image-filters-cli"
make build-cli

# Build the API binary. The binary will be named "image-filters-api"
make build-api

# Remove old binaries
make clean

# Show help for each make command
make help
```

## Benchmarking and Profiling

The project includes a benchmark suite that accepts multiple filters at once.
It runs two tests: one with a single worker for each filter and the other with
K workers, where K is equal to the number of logical CPUs avaliable in the
system. The test images have an arbitrary size of NxN, where N is the
`IMG_SIZE` parameter passed to the `make bench` command.

Below follows the output of a benchmark test of Grayscale with 5000x5000 images:

```bash
$ make bench FILTERS=grayscale IMG_SIZE=5000
go test -bench=. -run=^$ -benchmem ./engines -args -filters grayscale -imageSize 5000
goos: linux
goarch: amd64
pkg: dougdomingos.com/image-filters/engines
cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
BenchmarkExecuteSerial-8              99          12944693 ns/op        100008021 B/op         2 allocs/op
BenchmarkExecuteConcurrent-8         123           8684297 ns/op        100008072 B/op         2 allocs/op
PASS
ok      dougdomingos.com/image-filters/engines  9.245s
```

You can pass as many filters as you want, simply by separating the names by
commas. The benchmark will consider the total execution time:

```bash
$ make bench FILTERS=sobel,vertical-flip IMG_SIZE=5000
go test -bench=. -run=^$ -benchmem ./engines -args -filters sobel,vertical-flip -imageSize 5000
goos: linux
goarch: amd64
pkg: dougdomingos.com/image-filters/engines
cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
BenchmarkExecuteSerial-8               1        5476185448 ns/op        300177648 B/op        26 allocs/op
BenchmarkExecuteConcurrent-8           1        1765656321 ns/op        300197584 B/op        96 allocs/op
PASS
ok      dougdomingos.com/image-filters/engines  7.261s
```

You can also generate CPU and memory profile reports for benchmarks, by
running the `bench-cpuprof` and `bench-memprof` Make targets. Profile files
will be stored in the `./profiles` directory:

```bash
$ make bench-cpuprof FILTERS=sobel,vertical-flip IMG_SIZE=500
go test -bench=. -run=^$ -benchmem ./engines \
        -args -filters sobel,vertical-flip \
        -imageSize 500 \
        -test.cpuprofile ../profiles/cpu.prof
goos: linux
goarch: amd64
pkg: dougdomingos.com/image-filters/engines
cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
BenchmarkExecuteSerial-8           11364            106708 ns/op         1007861 B/op          2 allocs/op
BenchmarkExecuteConcurrent-8       11851            102083 ns/op         1007854 B/op          2 allocs/op
PASS
ok      dougdomingos.com/image-filters/engines  3.724s
```

- Be careful with the size of the test image! Large images may present high
  memory consumption (e.g., a 5000x5000 image consumes ~100MB)

## Design choices

### In-place image processing

Filter implementations are built with an in-place heuristic to keep memory
consumption as low as possible. For filters that do require a copy of the
original image (e.g., Sobel, Gaussian Blur), an internal copy is created
and used as reference to compute values. After that, resultant pixels
are written back to the original image.

### Padded copies and convolutional kernels

Convolution-based filters compute the color values of each pixel based on its
neighbors, as defined by a convolution kernel (typically a square matrix).
For such scenarios, there are two main concerns that must be addressed to
ensure the correct processing of the image:

1. Source immutability: The source image must remain unchanged until all pixels
   have been computed, to prevent processing of finished pixels.

2. Edge handling: Pixels near the borders of the image lack a full set of
   neighbors and may otherwise attempt to access positions outside the image
   bounds.

To address both issues, the selected approach is to create a padded copy of the
image. A padding of K pixels is added on all sides, where each pixel is filled
with an opaque black value (i.e., `RGBA(0, 0, 0, 255)`). All pixel computations
are then performed by reading from the padded copy and writing back to the
original image.

The padding size K depends on the kernel size. If the kernel has dimensions
N × N and N is odd, then K = `floor(N / 2)`. For example, a 3×3 kernel has
K = 1, since the center is surrounded by neighbors one pixel away in all
directions. A 5×5 kernel would require a padding of K = 2, and so on.

For every position (x, y) in the original image, its corresponding position in
the padded copy is given by (x + k, y + k). This ensures that, even at the
edges, any kernel-relative neighbor access stays within the bounds of the copy.
Whether the padding pixels should influence the filter result is up to its
implementation.
