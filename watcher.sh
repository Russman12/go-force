#!/bin/bash

chsum1=""

echo "listening for openapi.yaml changes..."

while [[ true ]]
do
    chsum2=`find pkg -type f -wholename 'pkg/*/api/openapi.yaml' -exec md5sum {} \;`
    if [[ $chsum1 != $chsum2 ]] ; then
        if [ -n "$chsum1" ]; then
            IFS='/'

            #Read the split words into an array based on comma delimiter
            read -a strarr <<< "$chsum1"

            pkgName="${strarr[1]}"
            echo "regenerating $pkgName code"
            printf "\napi/openapi.yaml" >> "pkg/$pkgName/.openapi-generator-ignore"
            sed -i '$a api/openapi.yaml' "pkg/$pkgName/.openapi-generator-ignore"

            openapi-generator-cli generate -c codegencfg.json -i pkg/$pkgName/api/openapi.yaml --package-name $pkgName -o pkg/$pkgName --minimal-update >/dev/null && go fmt ./pkg/$pkgName/... >/dev/null && go fmt ./pkg/$pkgName/... >/dev/null

            sed -i '$d' "pkg/$pkgName/.openapi-generator-ignore"

            echo "done"
        fi
        chsum1=$chsum2
    fi
    sleep 2
done
