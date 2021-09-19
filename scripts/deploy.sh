#!/usr/bin/env bash

set -e

rsync -av --exclude .git '../lcd1602/' root@192.168.1.237:/tmp/blah
ssh root@192.168.1.237 -C 'sh -c "cp -r /tmp/blah/ /root/ && cd /root/blah/ && (killall go; killall pi || true) && echo Starting... && /usr/local/go/bin/go run ./_examples/pi"'

