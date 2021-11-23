#!/bin/sh

set -e

pushd ../startupschooldata
git pull
popd

./scripts/report.sh --data=../startupschooldata/data

pushd ../startupschooldata
./scripts/commit.sh
popd

