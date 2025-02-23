#!/bin/sh

cd ../client
find dist -print
npm run build
cp ./dist/main.css ../internal/site2/public/main.css
cp ./dist/main.js ../internal/site2/public/main.js

