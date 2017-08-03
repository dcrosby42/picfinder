#!/bin/bash
pushd `dirname ${BASH_SOURCE[0]}` > /dev/null; HERE=`pwd`; popd > /dev/null
cd $HERE

cfg=$HERE/../local-config.yaml
env=$1
if [ -z "$env" ]; then
  env=dev
fi

set -e
`picfinder --config $cfg --env $env db shellvars`

    
mysql_cmd="sudo mysql"
database=picfinder
username=picfinder
password=picfinder
set -x
echo "create database $DBNAME" | $mysql_cmd
echo "grant all on ${DBNAME}.* to ${DBUSER} identified by '${DBPASS}'" | $mysql_cmd
echo "grant all on ${DBNAME}.* to ${DBUSER}@localhost identified by '${DBPASS}'" | $mysql_cmd
