#!/usr/bin/env bash
#
# Generates protos for Teleport and Teleport API.
set -eu

main() {
  cd "$(dirname "$0")"  # ./build-assets/
  cd ../                # teleport root

  # Clean gen/proto directories before regenerating them. Legacy protos are
  # generated all over the directory tree, so they won't get cleaned up
  # automatically if the proto is deleted.
  rm -fr api/gen/proto gen/proto

  # Generate Gogo protos. Generated protos are written to
  # gogogen/github.com/gravitational/teleport/..., so we copy them to the
  # correct relative path. This is in lieu of the module= option, which would do
  # this for us (and which is what we use for the non-gogo protogen).
  rm -fr gogogen
  trap 'rm -fr gogogen' EXIT # don't leave files behind
  buf generate --template=buf-gogo.gen.yaml
  cp -r gogogen/github.com/gravitational/teleport/. .
  # error out if there's anything outside of github.com/gravitational/teleport
  rm -fr gogogen/github.com/gravitational/teleport
  rmdir gogogen/github.com/gravitational gogogen/github.com gogogen

  # Generate protoc-gen-go protos (preferred).
  # Add your protos to the list if you can.
  buf generate --template=buf-go.gen.yaml \
    --path=api/proto/teleport/devicetrust/ \
    --path=api/proto/teleport/integration/ \
    --path=api/proto/teleport/kube/ \
    --path=api/proto/teleport/loginrule/ \
    --path=api/proto/teleport/okta/ \
    --path=api/proto/teleport/plugins/ \
    --path=api/proto/teleport/samlidp/ \
    --path=api/proto/teleport/transport/ \
    --path=api/proto/teleport/trust/ \
    --path=proto/teleport/lib/multiplexer/ \
    --path=proto/teleport/lib/teleterm/

  # Generate connect-go protos.
  buf generate --template=buf-connect-go.gen.yaml \
    --path=proto/prehog/

  # Generate JS protos.
	buf generate --template=buf-js.gen.yaml \
    --path=proto/prehog/ \
    --path=proto/teleport/lib/teleterm/
}

main "$@"
