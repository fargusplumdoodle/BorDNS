#!/bin/bash

echo "Building version $1"

if [ $# -ne 1 ];
then
	echo "must supply version"
	echo "./scripts/build_version.sh <version>"
	exit -1
fi


docker build -t bordns:$1 -f docker/Dockerfile .
