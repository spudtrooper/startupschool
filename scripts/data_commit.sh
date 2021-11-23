#!/bin/sh

pushd ../startupschooldata
git pull
popd

./scripts/report.sh

pushd ../startupschooldata
./scripts/commit.sh
popd

