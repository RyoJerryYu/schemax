#!/bin/bash

# go install the tools defined in tools.go

set -e
set -o errexit
set -o pipefail

cd $(dirname "$0")

# Read the tools.go file and extract import paths
IMPORT_PATHS=$(awk '/import \(/,/\)/' ./tools.go | sed -n 's/.*_ "\([^"]*\)".*/\1/p')

# Convert the import paths to an array
ALL_TOOLS=()
while IFS= read -r line; do
    ALL_TOOLS+=("$line")
done <<<"$IMPORT_PATHS"

# ALL_TOOLS will be like this:
# ALL_TOOLS=(
#     "github.com/bufbuild/buf/cmd/buf"
#     "github.com/envoyproxy/protoc-gen-validate"
#     "github.com/gogf/gf/cmd/gf/v2"
# )

function tools::install() {
    local __tool=$1
    echo "Installing $__tool"
    go install "$__tool"
}

for tool in "${ALL_TOOLS[@]}"; do
    tools::install "$tool"
done

echo ""
echo "install tools success"

# you can install the completion for your cli, example when using oh-my-zsh:
# buf completion zsh > $ZSH/completions/_buf
# then restart your shell
