# YAML to ENV Converter

This Golang application converts a YAML configuration file into an environment file with keys in uppercase and separated by underscores.

## Features

- Convert YAML files to environment files
- Supports nested YAML structures
- Allows for optional prefix to the output environment keys

## Prerequisites

- Golang 1.20 or higher, Or
- Docker

## Usage

0. Clone the repository or download the source files.
1. Build the application using 
    ```bash
    go build -o yaml2env
    ```
2. Run the application using 
    ```bash
    ./yaml2env [--prefix prefix] [--format format] input.yaml
    ```

The output environment file will be created with the same name as the input file but with a `.env` extension (e.g., `example.env`).

### Example
```bash
go run main.go --prefix MY_PREFIX --format yaml example.yaml
```
or
```bash
./yaml2env --prefix MY_PREFIX --format yaml example.yaml
```

## Docker

A `Dockerfile` is provided for those who wish to run the application in a Docker container.

### Build

```bash
docker build -t yaml2env .
```

### Run
```bash
docker run --rm -v "$(pwd)":/data -w /data yaml2env --prefix MY_PREFIX --format yaml example.yaml
```

This command mounts the current directory to `/data` in the container and sets the working directory to `/data`. The input YAML file should be located in the current directory on your host machine. The output file will be written to the same directory. The `--prefix` option can be customized as needed.