#!/bin/bash

chsum1=""

echo "listening for openapi.yaml changes..."

while [[ true ]]
do
    chsum2=`find pkg -type f -wholename 'pkg/*/api/openapi.yaml' -exec md5sum {} \;`
    if [[ $chsum1 != $chsum2 ]]; then
        if [ -n "$chsum1" ]; then
            while IFS= read -r line; do
                if ! [[ $chsum1 == *"$line"* ]]; then
                    IFS='/'
                    echo $line

                    #Read the split words into an array based on comma delimiter
                    read -a strarr <<< "$line"

                    pkgName="${strarr[1]}"
                    echo "regenerating $pkgName code"

                    sed -i '$a api/openapi.yaml' "pkg/$pkgName/.openapi-generator-ignore"

                    openapi-generator-cli generate -c codegencfg.json -i pkg/$pkgName/api/openapi.yaml --package-name $pkgName -o pkg/$pkgName --minimal-update >/dev/null && go fmt ./pkg/$pkgName/... >/dev/null && go fmt ./pkg/$pkgName/... >/dev/null

                    sed -i '$d' "pkg/$pkgName/.openapi-generator-ignore"

                    echo "done"
                fi
            done <<< "$chsum2"
        fi
        chsum1=$chsum2
    fi
    sleep 2
done
