
from invoke import task

@task
def depend(c):
    c.run('/usr/local/go/bin/go get -v', echo=True, pty=True)


@task
def install(c):
    c.run('install -m 0644 ./contrib/golden.service /etc/systemd/system/golden.service', echo=True, pty=True)
    c.run('systemctl daemon-reload', echo=True, pty=True)

@task
def start(c):
    c.run('systemctl start golden.service', echo=True, pty=True)

@task
def stop(c):
    c.run('systemctl stop golden.service', echo=True, pty=True)

@task
def restart(c):
    c.run('systemctl restart golden.service', echo=True, pty=True)

@task
def status(c):
    c.run('systemctl status golden.service', echo=True, pty=True)

@task
def watch(c):
    c.run('journalctl -f -u golden.service', echo=True, pty=True)

@task(default=True, pre=[ depend ])
def build(c):
    c.run('/usr/local/go/bin/go build -o reader .', echo=True, pty=True)

@task
def check(c):
    c.run('/usr/local/go/bin/go test', pty=True, echo=True)
