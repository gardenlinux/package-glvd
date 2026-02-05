# package-glvd

`package-glvd` is a command-line tool for querying the [Garden Linux Vulnerability Database (GLVD)](https://security.gardenlinux.org/) API. It helps you identify CVEs affecting installed packages on Garden Linux nodes or container images.

You can install `package-glvd` in Garden Linux images by enabling the `glvd` feature at build time, or at runtime via `apt` (the package name is `glvd`).

While primarily designed for use on Garden Linux nodes, `package-glvd` also supports a developer mode for local development and testing.

## Features

- Check for CVEs affecting installed packages.
- Query CVEs for a custom list of packages ("what-if" analysis).
- Print an executive summary of potential security issues.
- Supports both human-readable and JSON output.

## Usage

```sh
glvd [command] [args]
```

### Commands

- `check`  
  Query CVEs for all installed source packages.

- `what-if <pkg1> <pkg2> ...`  
  Query CVEs for a custom list of source packages.

- `executive-summary`  
  Print a summary of the number of potential security issues.

### Options

- Set `GLVD_CLIENT_JSON_OUTPUT=true` to get JSON output.
- Set `GLVD_CLIENT_DEV_MODE=true` to run using test data from `test-data/`.

### Examples

Check for CVEs affecting installed packages:

```sh
glvd check
```

Check for CVEs affecting specific packages:

```sh
glvd what-if vim bash coreutils
```

Print an executive summary:

```sh
glvd executive-summary
```

Get JSON output:

```sh
GLVD_CLIENT_JSON_OUTPUT=true glvd check
```

## Development

To run locally with test data:

```sh
GLVD_CLIENT_DEV_MODE=true go run .
```

### Configuring the API Base URL

By default, `package-glvd` uses `https://security.gardenlinux.org` as the API endpoint.  
You can override this by setting the `GLVD_API_BASE_URL` environment variable:

```sh
export GLVD_API_BASE_URL="http://localhost:8080"
glvd check
```

## Building

This project provides a `Makefile` for common development tasks.

### Format the code

```sh
make fmt
```

### Build the binary for your current platform

```sh
make build
```

The output will be a binary named `glvd`.

### Build Linux binaries for amd64 and arm64

```sh
make build-linux
```

This will produce:
- `glvd-linux-amd64`
- `glvd-linux-arm64`

### Clean build artifacts

```sh
make clean
```

## Release a new version of the client

To release a new version of the client:

1. Edit the `debian/changelog` file and add a new version entry, following the format of previous entries.
2. Commit your changes and push them to the `main` branch.

This will trigger the pipeline to build and publish the new version automatically.

For reference, see [this example commit](https://github.com/gardenlinux/package-glvd/commit/10209351ca301cdb091ed9fc40dff9a59e7345e3).

## License

MIT License. See [LICENSE](LICENSE) for details.
