#!/usr/bin/env bash

SOURCE_DIR="kaitai"
SOURCES="${SOURCE_DIR}/*.ksy"

KAITAI_COMPILER="${KAITAI_COMPILER:-kaitai-struct-compiler}"

KAITAI_GOLANG_PACKAGE="generated"
KAITAI_GOLANG_OPTIONS="--target go --outdir ./ --go-package ${KAITAI_GOLANG_PACKAGE}"

for f in "$SOURCES"; do
    $KAITAI_COMPILER $KAITAI_GOLANG_OPTIONS $f
done

go mod tidy;
