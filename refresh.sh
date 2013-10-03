#!/bin/sh

## Refresh golangjp/go documents on godoc.org

## Usage: ./refresh.sh

## args: path
refresh_godoc_org() {
    local path=$1

    echo "refresh $path"
    curl -d "path=$path" http://godoc.org/-/refresh
}

## 10分以内に更新されたものだけを更新する
updateList=`git log --name-only --pretty=format:"" --since="10 minutes ago"|grep /|sort|uniq`

for fname in $updateList
do 
    target=`dirname $fname`
    refresh_godoc_org "github.com/gophersjp/go/$target"
done

## トップページを更新する
refresh_godoc_org github.com/gophersjp/go
 

## おしまい。

