#!/usr/bin/env bats

declare -r BIN="${BIN:-}"

BASE_DIR="${PWD}/test/paths.d"
PATH_HELPER="${BIN} -p ${BASE_DIR} -m ${BASE_DIR}"

PATHS="/a/a/a:/b/b/b:/c/c/c:/d/d/d:${PATH_HELPER_TEST_DIR}/bin"
PATHS_DUP="/a/a/a:/b/b/b:/c/c/c:/d/d/d:/d/d/d:${PATH_HELPER_TEST_DIR}/bin"

EXPR="PATH=\"${PATHS}\" ; MANPATH=\"${PATHS}\" ; export PATH MANPATH ;"
EXPR_DUP="PATH=\"${PATHS_DUP}\" ; MANPATH=\"${PATHS_DUP}\" ; export PATH MANPATH ;"

@test "without-duplicates-and-witout-not-founds" {
    echo ${BIN}
    [[ -f "${BIN}" && -n "${PATH_HELPER}" && -d "${BASE_DIR}" ]]

    run ${PATH_HELPER}

    [ "$status" = 0 ]
    [ "$output" = "PATH=\"\" ; MANPATH=\"\" ; export PATH MANPATH ;" ]
}

@test "with-duplicates" {
    [[ -f "${BIN}" && -n "${PATH_HELPER}" && -d "${BASE_DIR}" ]]

    run ${PATH_HELPER} -s=false

    [ "$status" = 0 ]
    [ "$output" = "PATH=\"\" ; MANPATH=\"\" ; export PATH MANPATH ;" ]
}

@test "with-not-founds" {
    [[ -f "${BIN}" && -n "${PATH_HELPER}" && -d "${BASE_DIR}" ]]

    run ${PATH_HELPER} -d=false

    [ "$status" = 0 ]
    [ "$output" = "$EXPR" ]
}

@test "with-duplicates-and-with-not-founds" {
    [[ -f "${BIN}" && -n "${PATH_HELPER}" && -d "${BASE_DIR}" ]]

    run ${PATH_HELPER} -s=false -d=false

    [ "$status" = 0 ]
    [ "$output" = "$EXPR_DUP" ]
}
