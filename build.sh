#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build -o homef-gin

rsync -azP public/ root@homefbase:/home/apps/homef/public/
rsync -azP views/ root@homefbase:/home/apps/homef/views/

ssh -l root homefbase "systemctl stop homef.service; systemctl status homef.service;"
scp homef-gin root@homefbase:/home/apps/homef/

ssh -l root homefbase "systemctl start homef.service; systemctl status homef.service;"
