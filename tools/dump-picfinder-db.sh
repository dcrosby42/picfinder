#!/bin/bash
pushd `dirname ${BASH_SOURCE[0]}` > /dev/null; HERE=`pwd`; popd > /dev/null
cd $HERE/..
mysqldump picfinder -u picfinder -ppicfinder --add-drop-table > dumped-picfinder-db.sql
