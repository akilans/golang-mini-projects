[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2F13-k8s-client-go&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Kubernetes Client-go

This program lists pods and deployments in the kubernetes cluster.
Also creates/updates a deployment with k8s manifest yaml file
[dep.yaml](https://github.com/akilans/golang-mini-projects/blob/main/13-k8s-client-go/dep.yaml)

## Prerequisites

- Go
- Running k8s cluster with kubeconfig file (v1.25.3)
- client-go package (v0.25.3)
- k8s API structure and basic understanding
- [k8s-api-ref.md](https://github.com/akilans/golang-mini-projects/blob/main/13-k8s-client-go/k8s-api-ref.md)

### Demo

![List k8s pods, Deployment and create/update deployment](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/k8s-client-go.gif)

```bash

# clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

# go to the 13-k8s-client-go
cd 13-k8s-client-go

# build
go build

# run

./k8s-client-go
# Sample output
: '
Testing client go...
Pod - test-dep-ffd6fccb7-4cc7g
Deployment - test-dep
test-dep is found in default namespace. So Updating....
Updated deployment successfully...
'
```

## Credits and references

1. [Ivan Velichko](https://iximiuz.com/en/posts/kubernetes-api-structure-and-terminology/)
2. [dx13.co.uk](https://dx13.co.uk/articles/2021/01/15/kubernetes-types-using-go/)
3. [Client-go Examples](https://github.com/kubernetes/client-go/tree/master/examples)
