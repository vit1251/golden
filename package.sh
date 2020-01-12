#!/bin/sh

set -x

# Step 1. Clean
rm -rf ./package

# Step 2. Create directories
install -m 0755 -d ./package
install -m 0755 -d ./package/DEBIAN
install -m 0755 -d ./package/usr/local/bin

# Step 3. Make Debian package description
cp ./DEBIAN/control ./package/DEBIAN/control

# Step 4. Copy require binary
cp ./golden ./package/usr/local/bin/golden

# Step 5. Create Debian package
dpkg-deb -v --build ./package golden-1.2.3.deb
