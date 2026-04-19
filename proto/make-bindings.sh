#!/usr/bin/env bash

set -e

PROTOS=(
    services.proto
    sysconfig.proto
    health.proto
    api.proto
    notification.proto
    reflect.proto
    tests.proto
    web.proto
    system.proto
)

# Use the protoc image to run protoc.sh and generate Go + Rust bindings.
for p in "${PROTOS[@]}"; do
    docker run --user "$(id -u):$(id -g)" -e PROTO="$p" --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
done

# --- Go bindings ---
rm -rf ../go/types
mkdir -p ../go/types
mv ./types/* ../go/types/.
rm -rf ./types

mkdir -p ../go/testtypes
mv ./testtypes/*.pb.go ../go/testtypes/.
rm -rf ./testtypes

pushd ../go > /dev/null
find . -name "*.go" -type f -exec sed -i 's|"./types/l8services"|"github.com/saichler/l8types/go/types/l8services"|g' {} +
find . -name "*.go" -type f -exec sed -i 's|"./types/l8api"|"github.com/saichler/l8types/go/types/l8api"|g' {} +
popd > /dev/null

# --- Rust bindings ---
RUST_PB_DIR=../rust/crates/l8types/src/pb
mkdir -p "$RUST_PB_DIR"
# Clear existing generated bindings so deletions in proto/ propagate.
find "$RUST_PB_DIR" -maxdepth 1 -name '*.rs' -delete
# Move newly generated per-proto bindings. The mod.rs emitted by protoc-gen-rs
# only references the last proto compiled, so we drop it and regenerate below.
for f in ./*.rs; do
    [[ -e "$f" ]] || continue
    base="$(basename "$f")"
    if [[ "$base" == "mod.rs" ]]; then
        rm -f "$f"
        continue
    fi
    mv "$f" "$RUST_PB_DIR/"
done

# Regenerate pb/mod.rs listing every generated module.
{
    echo "// @generated"
    echo ""
    for f in "$RUST_PB_DIR"/*.rs; do
        name="$(basename "$f" .rs)"
        [[ "$name" == "mod" ]] && continue
        echo "pub mod $name;"
    done
} > "$RUST_PB_DIR/mod.rs"
