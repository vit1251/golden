
from invoke import task

@task
def clean(c):
    c.run('rm -rf ./package')
    c.run('rm -rf ./dist')

@task
def depend(c):
    c.run('go get -v', echo=True)

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
    c.run('go build -o golden .', echo=True)

@task
def debug(c):
    c.run('./golden >golden_service.log 2>golden_service_err.log', echo=True)
