#!/bin/bash -e

echo "mode: atomic" > piflab-store-api-go.coverprofile

packages=(
    "handlers"
)

for package in ${packages[@]};
do
	path=./$package/$package.coverprofile
    cat $path | grep -v "mode: atomic" >> piflab-store-api-go.coverprofile
	rm $path
done

if [ -n "$COVERALLS_TOKEN" ]
then
    goveralls -coverprofile=piflab-store-api-go.coverprofile -service circleci -repotoken $COVERALLS_TOKEN
    rm ./piflab-store-api-go.coverprofile
fi	
