#!/bin/sh
set -x
set -e

# Create user for Gin
addgroup -S echomusic
adduser -G echomusic -H -D -g 'echomusic User' echomusic -h /data/echomusic -s /bin/bash && usermod -p '*' echomusic && passwd -u echomusic
echo "export ECHOMUSIC_CUSTOM=${ECHOMUSIC_CUSTOM}" >> /etc/profile

# Final cleaning
rm /app/echomusic/docker/finalize.sh
rm /app/echomusic/docker/nsswitch.conf
