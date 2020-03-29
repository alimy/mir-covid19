#!/bin/sh

create_volume_subfolder() {
    # Create VOLUME subfolder
    for f in /data/covid19/data /data/covid19/conf /data/covid19/log; do
        if ! test -d $f; then
            mkdir -p $f
        fi
    done
}

setids() {
    PUID=${PUID:-1000}
    PGID=${PGID:-1000}
    groupmod -o -g "$PGID" covid19
    usermod -o -u "$PUID" covid19
}

setids
create_volume_subfolder

# Exec CMD or S6 by default if nothing present
if [ $# -gt 0 ];then
    exec "$@"
else
    exec /bin/s6-svscan /app/covid19/docker/s6/
fi
