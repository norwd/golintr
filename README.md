# golintr

GitHub Action to run [golint]

## Usage

### Inputs

`path`

Path where your Go files are.
This path will be used by golint command to check code of this files.
Defaults to the repository root including all sub-directories (`./...`).

*FYI, the `/...` suffix is used to includes all sub-directories, remove it if you only want to check files of the given directory.*

### Outputs

`errors`

The output of [golint] if the command fail.

For Example:

```
internal/cmd/root.go:17:1: exported function Execute should have comment or be unexported
```

### Basic Setup

```yaml
---

name: "Lint Go Files"

on:
  push:

jobs:
  golintr:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v3

      - name: Run Golint
        uses: norwd/golintr@v2
        with:
          path: ./src/...
````

[golint]: https://github.com/golang/lint
