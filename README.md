<p align="center">
    <a alt="GitHub Actions CI" href="https://github.com/otaviof/path-helper/actions">
        <img src="https://github.com/otaviof/path-helper/actions/workflows/test.yaml/badge.svg">
    </a>
    <a alt="go.pkg.dev project documentation" href="https://pkg.go.dev/mod/github.com/otaviof/path-helper">
        <img src="https://img.shields.io/badge/go.pkg.dev-docs-007d9c?logo=go&logoColor=white">
    </a>
    <a alt="goreportcard.com project report" href="https://goreportcard.com/report/github.com/otaviof/path-helper">
        <img src="https://goreportcard.com/badge/github.com/otaviof/path-helper">
    </a>
</p>

# `path-helper`

Command-line application to generate `PATH` and `MANPATH` environment variables based in a
configuration directory, the same approach used in Mac OS. By default, `PATH` is genarated based in
`/etc/paths.d`, while `MANPATH` is based on `/etc/manpaths.d`.

Inside those base directories the order of the files is kept, and each line in a file corresponds to
a directory in the file-system that should become part of `PATH` or `MANPATH`. Please consider the
local example of [`paths.d`](./test/paths.d).

In short, the usage of `path-helper` is:

```bash
eval `path-helper`
```

## Install

To install `path-helper` you can simply `go install`, as per:

```bash
go install github.com/otaviof/path-helper/cmd/path-helper@latest
```

Alternatively, you can run the following target in the project folder, `sudo` might be required for completion.

```bash
make install INSTALL_DIR="/usr/local/bin"
```

## Usage Examples

Please consider `--help` to see all possible options:

```bash
path-helper --help
```

Skipping duplicated entries, if any:

```bash
path-helper -s
```

Skipping non-existing directory:

```bash
path-helper -d
```

### Shell Configuration Example

Evaluate `path-helper` output in order to export `PATH` and `MANPATH` environment variables. The
following example checks if `path-helper` is present in default location, and later runs `eval`
against:

```bash
[ -x ${GOPATH}/bin/path-helper ] && \
    eval `${GOPATH}/bin/path-helper`
```

Running `path-helper` without `eval`, would return:

```
$ path-helper
PATH="..." ; MANPATH="..." ; export PATH MANPATH ;
```
