<p align="center">
    <a alt="GoReport" href="https://goreportcard.com/report/github.com/otaviof/path-helper">
        <img alt="GoReport" src="https://goreportcard.com/badge/github.com/otaviof/path-helper">
    </a>
    <a alt="Code Coverage" href="https://codecov.io/gh/otaviof/path-helper">
        <img alt="Code Coverage" src="https://codecov.io/gh/otaviof/path-helper/branch/master/graph/badge.svg">
    </a>
    <a href="https://godoc.org/github.com/otaviof/path-helper/pkg/path-helper">
        <img alt="GoDoc Reference" src="https://godoc.org/github.com/otaviof/path-helper/pkg/path-helper?status.svg">
    </a>
    <a alt="CI Status" href="https://travis-ci.com/otaviof/path-helper">
        <img alt="CI Status" src="https://travis-ci.com/otaviof/path-helper.svg?branch=master">
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

To install `path-helper` you can simply `go get`, for instance:

```bash
go get -u github.com/otaviof/path-helper/cmd/path-helper
```

Alternatively, you can:

```bash
make install
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
