#!/usr/bin/env bash

mkdir dist
go build -o startup
mv startup dist
cp -R public  dist/public
cp -R views dist/views
