#!/bin/bash
chmod 755 rsshub

echo "start running..."

# shellcheck disable=SC2046
# shellcheck disable=SC2006
kill -9 `cat pidfile.txt`

rm pidfile.txt

nohup ./rsshub -port 8083 & echo $! > pidfile.txt

echo "end"

exit
