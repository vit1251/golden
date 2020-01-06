#!/bin/sh

touch /var/spool/ftn/outb/139f0018.dlo
mkdir -p /var/spool/ftn/outb/139f0018.sep
cp ./compose.pkt /var/spool/ftn/outb/139f0018.sep/compose.pkt
echo "^/var/spool/ftn/outb/139f0018.sep/compose.pkt" > /var/spool/ftn/outb/139f0018.dlo
