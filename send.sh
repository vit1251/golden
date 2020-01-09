#!/bin/sh

touch /var/spool/ftn/outb/139f0018.dlo
mkdir -p /var/spool/ftn/outb/139f0018.sep
mv ./compose.pkt /var/spool/ftn/outb/139f0018.sep/compose.pkt
echo "^/var/spool/ftn/outb/139f0018.sep/compose.pkt" > /var/spool/ftn/outb/139f0018.dlo
