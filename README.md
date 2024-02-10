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

Command-line application to generate `PATH` and `MANPATH` environment variables based on a configuration directory, the same approach used in MacOS.

By default, `PATH` is genarated based in `/etc/paths.d` directory, while `MANPATH` is based on `/etc/manpaths.d`, you can configure these locations.

Files created on these directories are listed alphabetically, each line in a file corresponds to a another directory that should become part of `PATH` and `MANPATH`.

Configuration files may as well contain environment variables which will be expanded during execution.

Please consider the local example of [`paths.d`](./test/paths.d).

In short, the usage of `path-helper` is:

```bash
eval `path-helper`
```

## Install

The most convenient way would be:

```sh
curl -sL https://raw.githubusercontent.com/otaviof/path-helper/main/hack/install-lastest-release.sh | sh
```

This process is [automated by this script](./hack/install-lastest-release.sh), you should consider running on your repository clone, i.e.:

```sh
hack/install-lastest-release.sh
```

To install `path-helper` consider the [release page][releaseURL] and download pre-compiled binaries for your platform. Once the tarball is downloaded, you need to extract and install on the desired location, like the example snippet below

```sh
cd /tmp
tar -zxvpf path-helper-ostype-arch.tar.gz path-helper
install -m 0755 path-helper /usr/local/bin/path-helper
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

Evaluate `path-helper` output in order to export `PATH` and `MANPATH` environment variables. The following example checks if `path-helper` is present in default location, and later runs `eval` against:

```sh
declare -r PATH_HELPER_BIN="/usr/local/bin/path-helper"

[[ -x "${PATH_HELPER_BIN}" ]] &&
    eval "$(${PATH_HELPER_BIN})"
```

Running `path-helper` without `eval`, would print out the Shell script snippet it generateds. For instance:

```
$ path-helper
PATH="..." ; MANPATH="..." ; export PATH MANPATH ;
```

[releaseURL]: https://github.com/otaviof/path-helper/releases/tag/0.1.1
