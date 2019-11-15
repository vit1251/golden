
from invoke import task

@task
def depend_dev(c):
    c.run('npm install -g sass')

@task
def depend(c):
    c.run('/usr/local/go/bin/go get -u')

@task(default=True)
def build(c):
    c.run('/usr/local/go/bin/go build -o __main__')

@task
def debug(c):
    c.run('./__main__ 2> debug.txt')

@task
def install(c):
    c.run('install -m 0644 ./contrib/golden.service /etc/systemd/system/golden.service')
    c.run('systemctl daemon-reload')

@task
def start(c):
    c.run('systemctl start golden.service')

@task
def stop(c):
    c.run('systemctl stop golden.service')

@task
def restart(c):
    c.run('systemctl restart golden.service')

@task
def status(c):
    c.run('systemctl status golden.service', pty=True)

@task
def watch(c):
    c.run('journalctl -f -u golden.service', pty=True)
