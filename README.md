# glvd-client

`glvd-client` is a command-line tool to query the [Garden Linux Vulnerability Database (GLVD)](https://security.gardenlinux.org/) API for CVEs affecting installed packages on Garden Linux nodes.

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

## Building

Build the client with:

```sh
go build -o glvd .
```

## License

MIT License. See [LICENSE](LICENSE) for details.
