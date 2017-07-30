#!/bin/bash
pushd `dirname ${BASH_SOURCE[0]}` > /dev/null; HERE=`pwd`; popd > /dev/null

cd $HERE/..
log=server.log
go install && picfinder server | tee -a $log
