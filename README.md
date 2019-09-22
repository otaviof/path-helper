<p align="center">
    <a alt="GoReport" href="https://goreportcard.com/report/github.com/otaviof/path-helper">
        <img alt="GoReport" src="https://goreportcard.com/badge/github.com/otaviof/path-helper">
    </a>
    <a alt="Code Coverage" href="https://codecov.io/gh/otaviof/path-helper">
        <img alt="Code Coverage" src="https://codecov.io/gh/otaviof/path-helper/branch/master/graph/badge.svg">
    </a>
    <a href="https://godoc.org/github.com/otaviof/path-helper">
        <img alt="GoDoc Reference" src="https://godoc.org/github.com/otaviof/path-helper?status.svg">
    </a>
    <a alt="CI Status" href="https://travis-ci.com/otaviof/path-helper">
        <img alt="CI Status" src="https://travis-ci.com/otaviof/path-helper.svg?branch=master">
    </a>
</p>

# `path-helper`

Command-line application to generate `$PATH` based in a `paths.d` (`/etc/paths.d`) directory.

## Install

```bash
go get -u github.com/otaviof/path-helper
```

```bash
make instal
```

# Usage Examples

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