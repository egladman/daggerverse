# daggerverse

A monorepo of [Dagger](https://dagger.io/) modules.

## Create a New Module

The following snippet demonstrates how to create a new module named `hello-world`

```sh
mkdir hello-world
cd hello-world
dagger init --sdk=go --source=.
```

## Release a New Version of an Existing Module

The following snippet demonstrates how to cut a release for an existing module named `hello-world`

```
make VERSION=0.1.0 hello-world
```

This creates/pushes a git tag with name `hello-world/v0.1.0`
