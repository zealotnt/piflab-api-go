#!/bin/bash

ENTRY_HEADER='[ENTRY]'
usage() {
	echo
	echo "Syntax: api-go-entry.sh <db-type>"
	echo "	<db-type>	= postgres | mysql"
}

case $# in
1)	dbTyped=$1
	;;
*)	usage
	exit 1
	;;
esac

case $dbTyped in
"mysql") 	dbPort=3306
			;;
"postgres")	dbPort=5432
			;;
*) 	echo "$dbTyped not supported"
	exit 2
	;;
esac

echo "$ENTRY_HEADER Updating Go packages"
godep restore
gin -p $PORT run
