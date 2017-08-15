## Development Guide
This document is intended to be the canonical source of truth for things like supported toolchain versions for building Pharmer.
If you find a requirement that this doc does not capture, please submit an issue on github.

This document is intended to be relative to the branch in which it is found. It is guaranteed that requirements will change over time
for the development branch, but release branches of Pharmer should not change.

### Build Pharmer
Some of the Pharmer development helper scripts rely on a fairly up-to-date GNU tools environment, so most recent Linux distros should
work just fine out-of-the-box.

#### Setup GO
Pharmer is written in Google's GO programming language. Currently, Pharmer is developed and tested on **go 1.8.3**. If you haven't set up a GO
development environment, please follow [these instructions](https://golang.org/doc/code.html) to install GO.

#### Download Source

```console
$ go get github.com/appscode/pharmer
$ cd $(go env GOPATH)/src/github.com/appscode/pharmer
```

#### Install Dev tools
To install various dev tools for Pharmer, run the following command:

```console
# setting up dependencies for compiling pharmer...
$ ./hack/builddeps.sh
```

#### Build Binary
```
$ ./hack/make.py
$ pharmer version
```

#### Dependency management
Pharmer uses [Glide](https://github.com/Masterminds/glide) to manage dependencies. Dependencies are already checked in the `vendor` folder.
If you want to update/add dependencies, run:
```console
$ glide slow
```

#### Build Docker images
To build and push your custom Docker image, follow the steps below. To release a new version of Pharmer, please follow the [release guide](/docs/developer-guide/release.md).

```console
# Build Docker image
$ ./hack/docker/pharmer/setup.sh; ./hack/docker/pharmer/setup.sh push

# Add docker tag for your repository
$ docker tag appscode/pharmer:<tag> <image>:<tag>

# Push Image
$ docker push <image>:<tag>
```

#### Generate CLI Reference Docs
```console
$ ./hack/gendocs/make.sh
```