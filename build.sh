#!/usr/bin/env sh

cd assets/excalidraw
npm install
npm run build
rm -rf ../dist/*
cp -rf build/* ../dist/
rm -f ../dist/*.map ../dist/static/css/*.map ../dist/static/js/*.map
cd -
