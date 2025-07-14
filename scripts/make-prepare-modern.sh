#!/bin/sh

set -x

cd ../client

find dist -print

./make-build.sh

cp ./dist/main.css ../internal/site2/public/main.css
cp ./dist/main.js ../internal/site2/public/main.js

