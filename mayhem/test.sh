#!/usr/bin/env bash
#
# mayhem/test.sh — behavioral KAT via the clang-linked fuzz binary (dyn-linked, sabotage-sensitive).
set -uo pipefail
[ -n "${SOURCE_DATE_EPOCH:-}" ] || unset SOURCE_DATE_EPOCH
cd "$SRC"

emit_ctrf() {
  local tool="$1" passed="$2" failed="$3" skipped="${4:-0}" pending="${5:-0}" other="${6:-0}"
  local tests=$(( passed + failed + skipped + pending + other ))
  cat > "${CTRF_REPORT:-$SRC/ctrf-report.json}" <<JSON
{
  "results": {
    "tool": { "name": "$tool" },
    "summary": {
      "tests": $tests,
      "passed": $passed,
      "failed": $failed,
      "pending": $pending,
      "skipped": $skipped,
      "other": $other
    }
  }
}
JSON
  printf 'CTRF {"results":{"tool":{"name":"%s"},"summary":{"tests":%d,"passed":%d,"failed":%d,"pending":%d,"skipped":%d,"other":%d}}}\n' \
    "$tool" "$tests" "$passed" "$failed" "$pending" "$skipped" "$other"
  [ "$failed" -eq 0 ]
}

KAT_SEED="$SRC/mayhem/seeds/kat"
FUZZ_BIN="/mayhem/fuzz_lipgloss_style"

if [ ! -f "$KAT_SEED" ]; then
  echo "missing KAT seed: $KAT_SEED" >&2
  emit_ctrf "lipgloss-kat" 0 1
  exit 1
fi
if [ ! -x "$FUZZ_BIN" ]; then
  echo "missing fuzz binary: $FUZZ_BIN" >&2
  emit_ctrf "lipgloss-kat" 0 1
  exit 1
fi

rm -f /tmp/lipgloss-kat.out
out="$("$FUZZ_BIN" -runs=1 -max_len=64 -malloc_limit_mb=256 "$KAT_SEED" 2>&1)" || true
kat_out=""
if [ -f /tmp/lipgloss-kat.out ]; then
  kat_out="$(cat /tmp/lipgloss-kat.out)"
fi
if ! grep -q 'KAT:5:1' <<< "$kat_out"; then
  echo "KAT failed — expected KAT:5:1 from lipgloss.Size(\"hello\")" >&2
  echo "fuzz: $out" >&2
  echo "kat: $kat_out" >&2
  emit_ctrf "lipgloss-kat" 0 1
  exit 1
fi

emit_ctrf "lipgloss-kat" 1 0
