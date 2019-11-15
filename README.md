# golden

Golden is Web-base version Fido message reader 

## Prepare environment

You SHOULD have Python 3 package in Debian system:

    # apt-get install python3 python3-pip
    # python3 -m pip install invoke

You SHOULD have modern Golang installation:

    # cd /opt
    # wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
    # tar xf go1.13.4.linux-amd64.tar.gz 
    # rm -rf /usr/local/go
    # mv /opt/go /usr/local

## Installation instrustion

You SHOULD clone source code:

    # mkdir -p /srv/golden-1.0.0
    # git clone https://github.com/vit1251/golden /srv/golden-1.0.0

You SHOULD make application by next command:

    # cd /srv/golden-1.0.0
    # inv build

## User instrustion

You SHOULD have HPT configuration in ```/etc/hpt/config```.

You MAY use Golden in debug mode:

    # cd /srv/golden-1.0.0
    # inv debug

and you MAY watching about application processing in console now ;)

You MAY open browser and navigate address http://127.0.0.1:8080/ and
watch you Fido squish messages.

