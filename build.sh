#!/usr/bin/env sh

cd assets/excalidraw
yarn --ignore-optional
NODE_ENV=production yarn build:app
rm -rf ../dist/*
cp -rf build/* ../dist/
rm -f ../dist/*.map ../dist/static/css/*.map ../dist/static/js/*.map
cd -
