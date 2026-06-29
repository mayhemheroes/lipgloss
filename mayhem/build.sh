#!/usr/bin/env bash
#
# mayhem/build.sh — build lipgloss go-fuzz harnesses as sanitized libFuzzer binaries.
set -euo pipefail

[ -n "${SOURCE_DATE_EPOCH:-}" ] || unset SOURCE_DATE_EPOCH

: "${CC:=clang}" ; : "${CXX:=clang++}" ; : "${LIB_FUZZING_ENGINE:=-fsanitize=fuzzer}"
: "${SANITIZER_FLAGS=-fsanitize=address}"
: "${GO_DEBUG_FLAGS:=-g -gdwarf-3}"
: "${MAYHEM_JOBS:=$(nproc)}"
export CC CXX LIB_FUZZING_ENGINE SANITIZER_FLAGS GO_DEBUG_FLAGS MAYHEM_JOBS
export CGO_CFLAGS="${CGO_CFLAGS:+$CGO_CFLAGS }$GO_DEBUG_FLAGS"
export CGO_CXXFLAGS="${CGO_CXXFLAGS:+$CGO_CXXFLAGS }$GO_DEBUG_FLAGS"

export GOFLAGS="${GOFLAGS:--mod=mod}"
export GOPROXY="${GOPROXY:-file://$(go env GOMODCACHE)/cache/download,https://proxy.golang.org,direct}"

cd "$SRC"
go version

go get github.com/dvyukov/go-fuzz/go-fuzz-dep
go get github.com/AdaLogics/go-fuzz-headers

build_target() {
  local harness_dir="$1"
  local target="$2"
  echo "=== building $target (go-fuzz-build -libfuzzer) ==="
  (
    cd "$SRC/$harness_dir"
    go-fuzz-build -libfuzzer -o "$SRC/mayhem-build/$target.a"
  )
  $CXX $GO_DEBUG_FLAGS $SANITIZER_FLAGS $LIB_FUZZING_ENGINE \
    "$SRC/mayhem-build/$target.a" -o "/mayhem/$target"
  echo "built /mayhem/$target"
}

mkdir -p "$SRC/mayhem-build"
build_target mayhem/fuzz_lipgloss_position fuzz_lipgloss_position
build_target mayhem/fuzz_lipgloss_style fuzz_lipgloss_style

echo "build.sh complete"
