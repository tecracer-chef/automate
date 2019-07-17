#!/bin/bash

set -e

echo "Install latest dep and other tools"
if [[ -n "$HABITAT" ]]; then
  curl -s https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
else
  hab pkg install -b core/dep/0.5.0 core/ruby core/jq-static core/shellcheck
fi

echo "Checking Go Dependencies"
dep version
dep ensure -v
git diff --exit-code --ignore-submodules=all # fail if anything's been changed

echo "Checking automate-deployment binds.txt"
(
  cd ./components/automate-deployment
  make check-bindings
)

yml2json() {
    ruby -ryaml -rjson -e 'puts JSON.pretty_generate(YAML.load(ARGF))' "$1"
}

echo "Checking if Golang license fallbacks/exceptions are needed"
license_scout=$(yml2json .license_scout.yml)
for d in $(jq -ner --argjson data "$license_scout" '$data | (.fallbacks, .exceptions) | (.golang // [])[].name'); do
    if [[ ! -d "vendor/$d" ]]; then
        echo "dependency \"$d\" not required anymore"
        exit 1
    fi
done

echo "Checking for up-to-date bldr config"
go run ./tools/bldr-config-gen
if ! git diff --exit-code --ignore-submodules=all; then
    echo "The bldr config appears to be out of date!"
    echo "To fix this, run:"
    echo ""
    echo "   hab studio run \"source .studiorc && generate_bldr_config\""
    echo ""
    echo "Inspect and commit the resulting changes if they look reasonable"
    exit 1
fi

echo "Shellchecking!"
shellcheck -s bash -ax \
  .expeditor/*.sh \
  .expeditor/**/*.sh \
  .buildkite/hooks/* \
  scripts/*.sh \
  integration/tests/*.sh

echo "Checking for possible credentials in the source code"
go run ./tools/credscan
