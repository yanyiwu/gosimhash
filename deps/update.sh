#!/bin/bash
set -x
REPO_NAME=$1
VERSION=$2
echo $REPO_NAME $VERSION

rm -rf ./$REPO_NAME
git clone --recurse-submodules --branch $VERSION --single-branch git@github.com:yanyiwu/$REPO_NAME.git
cd $REPO_NAME
rm -rf .git
cd ..
git add $REPO_NAME
