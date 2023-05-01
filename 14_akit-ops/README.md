[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fakilans%2Fgolang-mini-projects%2Ftree%2Fmain%2F14_akit-ops&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# Akit-Ops - simple gitops solution

- This application detect any changes in the provided git repository and create/update k8s resources
- Provide a repo url and pull interval as env variable
- Commit a valid k8s manifest files in the provided git repository

### Prerequisites

- Go
- K8s cluster
- Create a service account with Proper permissions and deploy this application - [sa-role-rb-akit-ops.yaml](https://github.com/akilans/golang-mini-projects/blob/main/14_akit-ops/sa-role-rb-akit-ops.yaml)
- Public github repo with simple k8s manifest files
- Understanding of k8s API architecture - [k8s-api-ref.md](https://github.com/akilans/golang-mini-projects/blob/main/14_akit-ops/k8s-api-ref.md)

### Flowchart

![Akit-Ops](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/images/akit-ops.png?raw=true)

### Demo

![Akit-Ops Demo](https://raw.githubusercontent.com/akilans/golang-mini-projects/main/demos/akit-ops.gif)

### Usage

```bash
# clone a repo
git clone https://github.com/akilans/golang-mini-projects.git

# go to the 14_akit-ops
cd 14_akit-ops

# create docker image for akit-ops and push it
docker image build -t akilan/akit-ops:1 .
docker image push akilan/akit-ops:1

# create a service account,deployment with proper roles for deployment
# update repo url and pull interval here
kubectl apply -f sa-role-rb-akit-ops.yaml

```

## Credits and references

1. [Ivan Velichko](https://iximiuz.com/en/posts/kubernetes-api-structure-and-terminology/)
2. [dx13.co.uk](https://dx13.co.uk/articles/2021/01/15/kubernetes-types-using-go/)
3. [Client-go Examples](https://github.com/kubernetes/client-go/tree/master/examples)
