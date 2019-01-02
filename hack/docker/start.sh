#!/bin/sh

create_volume_subfolder() {
    # Create VOLUME subfolder
    for f in /data/echomusic/data /data/echomusic/conf /data/echomusic/log; do
        if ! test -d $f; then
            mkdir -p $f
        fi
    done
}

setids() {
    PUID=${PUID:-1000}
    PGID=${PGID:-1000}
    groupmod -o -g "$PGID" echomusic
    usermod -o -u "$PUID" echomusic
}

setids
create_volume_subfolder

# Exec CMD or S6 by default if nothing present
if [ $# -gt 0 ];then
    exec "$@"
else
    exec /bin/s6-svscan /app/echomusic/docker/s6/
fi
