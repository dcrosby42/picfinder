#!/bin/bash
pushd `dirname ${BASH_SOURCE[0]}` > /dev/null; HERE=`pwd`; popd > /dev/null
cd $HERE

protoc -I . ./picfinder.proto --go_out=plugins=grpc:.
