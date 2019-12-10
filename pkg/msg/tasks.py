
from invoke import task

@task(default=True)
def check(c):
    c.run('/usr/local/go/bin/go test .')
