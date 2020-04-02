
from invoke import task
from datetime import datetime

@task
def clean(c):
    c.run('rm -rf ./package')
    c.run('rm -rf ./dist')

@task
def depend(c):
    c.run('go get -v -u', echo=True)

@task
def package(c):
    c.run('install -m 0755 -d ./package')
    c.run('install -m 0755 -d ./package/DEBIAN')
    c.run('install -m 0755 -d ./package/usr/local/bin')
    c.run('install -m 0755 -d ./dist')
    c.run('cp ./DEBIAN/control ./package/DEBIAN/control')
    c.run('cp ./golden ./package/usr/local/bin/golden')
    c.run('dpkg-deb -v --build ./package golden-1.2.3.deb')
    c.run('cp ./golden-1.2.3.deb ./dist/golden-1.2.3.deb')

@task
def check(c):
    c.run('go test ./...', echo=True)

@task(default=True)
def build(c):
    c.run('go build -o golden-windows-386.exe .', echo=True, env={"GOOS": "windows", "GOARCH": "386", "CGO_ENABLED": "1"})
    c.run('go build -o golden-windows-amd64.exe .', echo=True, env={"GOOS": "windows", "GOARCH": "amd64", "CGO_ENABLED": "1"})

def build2(c):
    c.run('go build -o golden-linux-i386-1.2.11-amd64.exe .', echo=True, env={"GOOS": "linux", "GOARCH": "amd64", "CGO_ENABLED": "1"})
    c.run('go build -o golden-darwin-i386-1.2.11-amd64.exe .', echo=True, env={"GOOS": "darwin", "GOARCH": "amd64", "CGO_ENABLED": "1"})

@task
def package(c):
    """ Create 
    """
    now = datetime.now()
    stamp = now.strftime("%Y%m%d%H%M%S")
    package_name = "GoldenPoint-windows-{stamp}.zip".format(stamp=stamp)
    c.run('7z a -tzip {package_name} golden-windows-386.exe'.format(package_name=package_name), echo=True, env={"PATH": "C:\\Program Files\\7-Zip"})
    c.run('7z a -tzip {package_name} golden-windows-amd64.exe'.format(package_name=package_name), echo=True, env={"PATH": "C:\\Program Files\\7-Zip"})
    c.run('7z a -tzip {package_name} README.md'.format(package_name=package_name), echo=True, env={"PATH": "C:\\Program Files\\7-Zip"})
    c.run('7z a -tzip {package_name} LICENSE'.format(package_name=package_name), echo=True, env={"PATH": "C:\\Program Files\\7-Zip"})

@task
def debug(c):
    c.run('golden.exe >golden_service.log 2>golden_service_err.log', echo=True)
