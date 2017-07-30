#!/bin/bash
pushd `dirname ${BASH_SOURCE[0]}` > /dev/null; HERE=`pwd`; popd > /dev/null

cfg=$HERE/../local-config.yaml
env=$1
if [ -z "$env" ]; then
  env=dev
fi

set -e
`picfinder --config $cfg --env $env db shellvars`

echo "mysql --host $DBHOST --port $DBPORT -u $DBUSER -p$DBPASS $DBNAME"
mysql --host $DBHOST --port $DBPORT -u $DBUSER -p$DBPASS $DBNAME
