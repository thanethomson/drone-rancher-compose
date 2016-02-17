# Rancher Compose Plugin for Drone

## Overview
[Drone](http://readme.drone.io/) is a great application for continuous
integration. I wasn't too happy with how limited the
[Rancher plugin](http://addons.drone.io/rancher/) for Drone was, so I quickly
slapped together a Go-based plugin to execute
[Rancher Compose](https://github.com/rancher/rancher-compose).

## Binary
Build the binary using `make`:

```bash
make clean build
```

## Docker
First edit the `Makefile` to set up the `IMAGE_TAG` and `IMAGE` parameters, then
run the following:

```bash
make docker
```

## Parameters
In the `vargs` parameter, set yourself a `commands` list that supplies
parameters to the `rancher-compose` executable. See the
[Rancher Compose docs](http://docs.rancher.com/rancher/rancher-compose/)
for information on how to use this.

## Example
As per the [Drone docs](http://readme.drone.io/devs/plugins/) for custom
plugins, you will first need to configure your Drone instance with the
`PLUGIN_FILTER` environment variable to allow your custom plugin.

Then, you can execute something similar to the following. This will be executed
within the same folder as your repository, so it should automatically pick
up your `docker-compose.yml` and `rancher-compose.yml` files if you have
any in the root of your repo.

```bash
docker run -i registry:5001/labs/drone-rancher-compose <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com",
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "commands": [
          "--access-key <ACCESS_KEY> --secret-key <SECRET_KEY> -p <STACK_NAME> create",
          "--access-key <ACCESS_KEY> --secret-key <SECRET_KEY> -p <STACK_NAME> up --upgrade -d"
        ]
    }
}
EOF
```
