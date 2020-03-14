
from invoke import task

@task
def depend(c):
    c.run('npm install -g less', echo=True)

@task
def build(c):
    c.run('lessc main.less main.css', echo=True)
