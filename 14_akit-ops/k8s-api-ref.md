# K8S API Reference

```bash
# It gives k8s api information
kubectl api-resources -o wide
# it gives how kubectl calls k8s api vi url
kubectl get pods -v 6
#https://192.168.49.2:8443/api/v1/namespaces/default/pods?limit=500


# Make Kubernetes API available on localhost:8080
# to bypass the auth step in subsequent queries:
kubectl proxy --port=8080 &

# List all known API paths
curl http://localhost:8080/
# List known versions of the `core` group
curl http://localhost:8080/api
# List known resources of the `core/v1` group
curl http://localhost:8080/api/v1
# Get a particular Pod resource
curl http://localhost:8080/api/v1/namespaces/default/pods/sleep-7c7db887d8-dkkcg

# List known groups (all but `core`)
curl http://localhost:8080/apis
# List known versions of the `apps` group
curl http://localhost:8080/apis/apps
# List known resources of the `apps/v1` group
curl http://localhost:8080/apis/apps/v1
# Get a particular Deployment resource
curl http://localhost:8080/apis/apps/v1/namespaces/default/deployments/sleep

# json response
kubectl get --raw /api/v1/namespaces/default/pods/httpd | python3 -m json.tool

```

### K8S - API

- k8s.io/api and k8s.io/apimachinery modules - these are the two main dependencies of the official Go client.
  The api module defines Go structs for the Kubernetes Objects, and the apimachinery module brings lower-level building blocks and common API functionality like serialization, type conversion, or error handling

```bash
# get k8s api endpoint url with port
kubectl cluster-info
kubectl config view

# Find k8s version
curl https://192.168.49.2:8443/version --insecure
curl https://192.168.49.2:8443/version --cacert ~/.minikube/ca.crt

# call api with authentication
# With certificates
# It throws forbidder error
curl --cacert ~/.minikube/ca.crt https://192.168.49.2:8443/apis/apps/v1/deployments
# with certs
curl --cert ~/.minikube/profiles/minikube/client.crt --key ~/.minikube/profiles/minikube/client.key \
--cacert ~/.minikube/ca.crt https://192.168.49.2:8443/apis/apps/v1/deployments

# With Token
# generate token
TOKEN=$(kubectl create token default)
curl --cacert ~/.minikube/ca.crt https://192.168.49.2:8443/apis/apps/v1/deployments \
--header "Authorization: Bearer $TOKEN"

# service account from kube-system namespace
TOKEN=$(kubectl -n kube-system create token default)
curl --cacert ~/.minikube/ca.crt https://192.168.49.2:8443/apis/apps/v1/deployments \
--header "Authorization: Bearer $TOKEN"

# Create deployment
curl --cert ~/.minikube/profiles/minikube/client.crt --key ~/.minikube/profiles/minikube/client.key \
--cacert ~/.minikube/ca.crt https://192.168.49.2:8443/apis/apps/v1/namespaces/default/deployments \
-X POST \
-H 'Content-Type: application/yaml' \
-d '---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: demo-app
  name: demo-app
spec:
  replicas: 2
  selector:
    matchLabels:
      app: demo-app
  template:
    metadata:
      labels:
        app: demo-app
    spec:
      containers:
      - image: httpd
        name: httpd
'

# Update deployment
curl --cert ~/.minikube/profiles/minikube/client.crt --key ~/.minikube/profiles/minikube/client.key \
--cacert ~/.minikube/ca.crt https://192.168.49.2:8443/apis/apps/v1/namespaces/default/deployments/demo-app \
-X PUT \
-H 'Content-Type: application/yaml' \
-d '---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: demo-app
  name: demo-app
spec:
  replicas: 4
  selector:
    matchLabels:
      app: demo-app
  template:
    metadata:
      labels:
        app: demo-app
    spec:
      containers:
      - image: httpd
        name: httpd
'

#Patch deployment
curl https://192.168.49.2:8443/apis/apps/v1/namespaces/default/deployments/demo-app \
--cacert ~/.minikube/ca.crt \
--cert ~/.minikube/profiles/minikube/client.crt \
--key ~/.minikube/profiles/minikube/client.key \
-X PATCH \
-H 'Content-Type: application/merge-patch+json' \
-d '{
  "spec": {
    "template": {
      "spec": {
        "containers": [
          {
            "name": "httpd",
            "image": "httpd"
          }
        ]
      }
    }
  }
}'

# update only replicas
curl https://192.168.49.2:8443/apis/apps/v1/namespaces/default/deployments/demo-app \
--cacert ~/.minikube/ca.crt \
--cert ~/.minikube/profiles/minikube/client.crt \
--key ~/.minikube/profiles/minikube/client.key \
-X PATCH \
-H 'Content-Type: application/merge-patch+json' \
-d '{
  "spec": {
    "replicas": 1
  }
}'

# Delete deployment
curl https://192.168.49.2:8443/apis/apps/v1/namespaces/default/deployments/demo-app \
--cacert ~/.minikube/ca.crt \
--cert ~/.minikube/profiles/minikube/client.crt \
--key ~/.minikube/profiles/minikube/client.key \
-X DELETE
```

### k8s.io/api, k8s.io/apimachinery

- While the k8s.io/api module focuses on the concrete higher-level types like Deployments, Secrets, or Pods, the k8s.io/apimachinery is a home for lower-level but more universal data structures.
