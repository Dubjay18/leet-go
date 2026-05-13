#!/usr/bin/env bash
set -euo pipefail

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
repo_root=$(cd "$script_dir/.." && pwd)
readme="$repo_root/README.md"

if [[ ! -f "$readme" ]]; then
  echo "README.md not found at repo root." >&2
  exit 1
fi

start_marker="<!-- AUTO-GENERATED:START -->"
end_marker="<!-- AUTO-GENERATED:END -->"

if ! grep -q "$start_marker" "$readme" || ! grep -q "$end_marker" "$readme"; then
  echo "Auto-generated markers not found in README.md." >&2
  exit 1
fi

mapfile -t folders < <(
  find "$repo_root" -maxdepth 1 -mindepth 1 -type d \
    ! -name ".git" \
    ! -name "scripts" \
    -print \
  | while IFS= read -r dir; do
      if [[ -f "$dir/main.go" ]]; then
        basename "$dir"
      fi
    done \
  | sort
)

{
  echo "$start_marker"
  for folder in "${folders[@]}"; do
    echo "- \`$folder/\`"
  done
  echo "$end_marker"
} > "$readme.tmp"

awk -v start="$start_marker" -v end="$end_marker" -v block="$(cat "$readme.tmp")" '
  $0 ~ start {print block; in_block=1; next}
  $0 ~ end {in_block=0; next}
  !in_block {print}
' "$readme" > "$readme.new"

mv "$readme.new" "$readme"
rm -f "$readme.tmp"

echo "README.md updated."
