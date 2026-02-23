# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Go CLI that composes `PATH` and `MANPATH` from files in config directories (like macOS `/etc/paths.d`), with env var expansion and duplicate/missing-dir filtering. Output is `eval`-able shell export syntax.

## Commands

```bash
make build          # → bin/path-helper
make test           # unit + e2e
make test-unit      # go test -race -cover
make test-e2e       # bats-core (needs git submodules)
make install        # → /usr/local/bin
```

Single test: `PATH_HELPER_TEST_DIR="/test" go test -v -run TestPathHelper/with-not-founds ./pkg/path-helper/...`

## Critical: Test Environment

Tests require `PATH_HELPER_TEST_DIR="/test"` (Makefile exports this). Without it, env-expansion assertions fail silently.

E2E tests need bats-core submodule: `git submodule update --init --recursive`

## Architecture

~200 lines, zero runtime dependencies (`testify` is test-only). Package name is `pathhelper` (not `path-helper`).

**Flow:** `main.go` → `flag.Parse` → `Config` → `PathHelper.RenderExpression()` → stdout

`RenderExpression` calls `collect(baseDir)` twice (PATH, MANPATH). Each `collect`: globs files in dir → reads lines → skips `#` comments → `os.ExpandEnv` → filters by `SkipNotFound`/`SkipDuplicates` → colon-joins → formats `PATH="..." ; MANPATH="..." ; export PATH MANPATH ;`

Test fixtures in `test/paths.d/` (numbered files, processed in sort order).