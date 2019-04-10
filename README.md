# cluster-api-provider-generic

An experimental provider which uses webhooks as an extension mechanism for
allocating machine infrastructure while sharing software provisioning.

# Developer

## Build artifacts

```
export IMG=<CONTAINER_REPOSITORY_URL>
make docker-build
make docker-push
```

## Deploy controller

```
kind create cluster
export KUBECONFIG="$(kind get kubeconfig-path --name="1")"
make deploy
```

