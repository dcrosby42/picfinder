#!/bin/bash
pushd `dirname ${BASH_SOURCE[0]}` > /dev/null; HERE=`pwd`; popd > /dev/null
cd $HERE

cmd=$HERE/picfinder-db.sh

sql="select count(*) from file_info"
echo $sql
echo $sql | $cmd

sql="select count(*) number,sum(size) bytes, kind,type from file_info group by type order by count(*) desc;" 
echo $sql
echo $sql | $cmd

sql="select count(*) number,sum(size) bytes, kind from file_info group by kind order by count(*) desc;"
echo $sql
echo $sql | $cmd
