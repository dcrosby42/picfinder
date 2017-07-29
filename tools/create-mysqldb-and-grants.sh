#!/bin/bash
pushd `dirname ${BASH_SOURCE[0]}` > /dev/null; HERE=`pwd`; popd > /dev/null
cd $HERE
    
mysql_cmd="sudo mysql"
database=picfinder
username=picfinder
password=picfinder
echo "create database $database" | $mysql_cmd
echo "grant all on ${database}.* to ${username} identified by '${password}'" | $mysql_cmd
echo "grant all on ${database}.* to ${username}@localhost identified by '${password}'" | $mysql_cmd
