#!/bin/sh

set -e

go run report.go --data=../startupschooldata/data "$@"