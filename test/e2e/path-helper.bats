#!/usr/bin/env bats

BASE_DIR="./test/paths.d"
PATH_HELPER="./build/path-helper -b $BASE_DIR"

@test "without-duplicates-and-witout-not-founds" {
    run $PATH_HELPER
    [ "$status" = 0 ]
    [ "$output" = "export PATH=\"\"" ]
}

@test "with-duplicates" {
    run $PATH_HELPER -s=false
    [ "$status" = 0 ]
    [ "$output" = "export PATH=\"\"" ]
}

@test "with-not-founds" {
    run $PATH_HELPER -d=false
    [ "$status" = 0 ]
    [ "$output" = "export PATH=\"/a/a/a:/b/b/b:/c/c/c:/d/d/d\"" ]
}

@test "with-duplicates-and-with-not-founds" {
    run $PATH_HELPER -s=false -d=false
    [ "$status" = 0 ]
    [ "$output" = "export PATH=\"/a/a/a:/b/b/b:/c/c/c:/d/d/d:/d/d/d\"" ]
}
