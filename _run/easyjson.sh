#!/bin/bash

root_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dir="$(dirname "$root_dir")"

#############################################################################

echo "Обход структуры Time-Rules"
easyjson -all -omit_empty "$root_dir/struct.go"
easyjson -all -omit_empty "$root_dir/def.go"

