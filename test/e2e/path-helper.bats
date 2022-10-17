#!/usr/bin/env bats

BASE_DIR="./test/paths.d"
PATH_HELPER="./_output/path-helper -p $BASE_DIR -m $BASE_DIR"

PATHS="/a/a/a:/b/b/b:/c/c/c:/d/d/d"
PATHS_DUP="/a/a/a:/b/b/b:/c/c/c:/d/d/d:/d/d/d"

EXPR="PATH=\"$PATHS\" ; MANPATH=\"$PATHS\" ; export PATH MANPATH ;"
EXPR_DUP="PATH=\"$PATHS_DUP\" ; MANPATH=\"$PATHS_DUP\" ; export PATH MANPATH ;"

@test "without-duplicates-and-witout-not-founds" {
    run $PATH_HELPER
    [ "$status" = 0 ]
    [ "$output" = "PATH=\"\" ; MANPATH=\"\" ; export PATH MANPATH ;" ]
}

@test "with-duplicates" {
    run $PATH_HELPER -s=false
    [ "$status" = 0 ]
    [ "$output" = "PATH=\"\" ; MANPATH=\"\" ; export PATH MANPATH ;" ]
}

@test "with-not-founds" {
    run $PATH_HELPER -d=false
    [ "$status" = 0 ]
    [ "$output" = "$EXPR" ]
}

@test "with-duplicates-and-with-not-founds" {
    run $PATH_HELPER -s=false -d=false
    [ "$status" = 0 ]
    [ "$output" = "$EXPR_DUP" ]
}
