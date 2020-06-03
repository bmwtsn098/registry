#!/bin/bash

# This should point to the .proto files distributed with protoc.
export PROTO_PROTOS="$HOME/local/include"

# This is a local directory containing .proto files used by many APIs.
export ANNOTATIONS="../third_party/api-common-protos"

mkdir -p lib/generated

echo "Generating dart support code."
protoc --proto_path=../proto --proto_path=${ANNOTATIONS} \
        ${PROTO_PROTOS}/google/protobuf/timestamp.proto \
        ${PROTO_PROTOS}/google/protobuf/field_mask.proto \
        ${PROTO_PROTOS}/google/protobuf/empty.proto \
	../proto/registry_models.proto \
	../proto/registry_service.proto \
	--dart_out=grpc:lib/generated
