
from invoke import task
from datetime import datetime
from shutil import copyfile
from platform import system as platform_system

@task
def clean(c):
    c.run('rm -rf ./package')
    c.run('rm -rf ./dist')

@task
def depend(c):
    c.run('go get -v -u', echo=True)

@task
def package(c, version="1.2.11"):
    c.run('install -m 0755 -d ./package')
    c.run('install -m 0755 -d ./package/DEBIAN')
    c.run('install -m 0755 -d ./package/usr/local/bin')
    c.run('install -m 0755 -d ./dist')
    c.run('cp ./DEBIAN/control ./package/DEBIAN/control')
    c.run('cp ./golden ./package/usr/local/bin/golden')
    c.run('dpkg-deb -v --build ./package golden-{version}.deb'.format(version=version))
    c.run('cp ./golden-{version}.deb ./dist/golden-{version}.deb'.format(version=version))

@task
def check(c):
    c.run('go test ./...', echo=True)

@task
def prepare(c):
    #c.run('npm install -g less', echo=True)
    c.run('cd contrib && lessc main.less main.css', echo=True)
    copyfile("contrib/main.css", "static/css/main.css")

@task
def build_w32(c):
    platform_system_name = platform_system()
    if platform_system_name == "Windows":
        env = {
            "GOOS": "windows",
            "GOARCH": "386",
            "CGO_ENABLED": "1",
        }
        c.run('go build -o golden-windows-386.exe .', echo=True, env=env)

@task
def build_w64(c):
    env = {
        "GOOS": "windows",
        "GOARCH": "amd64",
        "CGO_ENABLED": "1",
        "CC": "x86_64-w64-mingw32-gcc",
        "CXX": "x86_64-w64-mingw32-g++",
    }
    c.run('go build -o golden-windows-amd64.exe .', echo=True, env=env)

@task
def build_linux(c):
    env = {
        "GOOS": "linux",
        "GOARCH": "amd64",
        "CGO_ENABLED": "1"
    }
    c.run('go build -o golden-linux-amd64 .', echo=True, env=env)

@task
def build_darwin(c):
    platform_system_name = platform_system()
    if platform_system_name == "Darwin":
        env = {
            "GOOS": "darwin",
            "GOARCH": "amd64",
            "CGO_ENABLED": "1"
        }
        c.run('go build -o golden-darwin-amd64 .', echo=True, env=env)

@task(default=True)
def build(c):
    build_w32(c)
    build_w64(c)
    build_linux(c)
    build_darwin(c)

@task
def package(c):
    """ Create 
    """
    now = datetime.now()
    stamp = now.strftime("%Y%m%d%H%M%S")
    package_name = "GoldenPoint-windows-{stamp}.zip".format(stamp=stamp)
    #
    platform_system_name = platform_system()
    if platform_system_name == "Darwin":
        env = {
            "PATH": "C:\\Program Files\\7-Zip",
        }
        c.run('7z a -tzip {package_name} golden-windows-386.exe'.format(package_name=package_name), echo=True, env=env, warn=True)
        c.run('7z a -tzip {package_name} golden-linux-amd64'.format(package_name=package_name), echo=True, env=env, warn=True)
        c.run('7z a -tzip {package_name} golden-windows-amd64.exe'.format(package_name=package_name), echo=True, env=env, warn=True)
        c.run('7z a -tzip {package_name} README.md'.format(package_name=package_name), echo=True, env=env, warn=True)
        c.run('7z a -tzip {package_name} LICENSE'.format(package_name=package_name), echo=True, env=env, warn=True)
        c.run('7z a -tzip {package_name} static'.format(package_name=package_name), echo=True, env=env, warn=True)
    elif platform_system_name == "Linux":
        c.run('zip {package_name} golden-windows-386.exe'.format(package_name=package_name), echo=True, warn=True)
        c.run('zip {package_name} golden-linux-amd64'.format(package_name=package_name), echo=True, warn=True)
        c.run('zip {package_name} golden-windows-amd64.exe'.format(package_name=package_name), echo=True, warn=True)
        c.run('zip {package_name} README.md'.format(package_name=package_name), echo=True, warn=True)
        c.run('zip {package_name} LICENSE'.format(package_name=package_name), echo=True, warn=True)
        c.run('zip -r {package_name} static'.format(package_name=package_name), echo=True, warn=True)
    else:
        raise RuntimeError('Unknown system {platform_system_name}'.format(platform_system_name=platform_system_name))

@task
def debug(c):
    c.run('golden.exe', echo=True)
