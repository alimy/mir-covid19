#!/bin/sh
set -x
set -e

# Create user for covid19
addgroup -S covid19
adduser -G covid19 -H -D -g 'covid19 User' covid19 -h /data/covid19 -s /bin/bash && usermod -p '*' covid19 && passwd -u covid19
echo "export COVID19_CUSTOM=${COVID19_CUSTOM}" >> /etc/profile

# Final cleaning
rm /app/covid19/docker/finalize.sh
rm /app/covid19/docker/nsswitch.conf
