
### Running the application 
``` 
cd go
go run main.go
curl localhost:8181/app
curl localhost:8181/metrics
``` 
### Running the tests 
``` 
cd go
go run main.go
go test
``` 
### Running the application using docker
```
docker build -t hostname-app:v0.1 .
docker run -p 8181:8181 hostname-app:v0.1
```

### Deploying kube cluster using kind 
``` 
kind create cluster  --config kube/kind.yml

kind get kubeconfig > ~/.kube/config  # Warning!!! This will rewrite your kube config file
kubectl get nodes

# Load ingress into the new kind cluster
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

docker build -t hostname-app:v0.1 .
# load the image into each node, so we don't have to fetch from any registry 
kind load  docker-image hostname-app:v0.1

kubectl apply -f kube/app.yaml

# calling the api
curl localhost/app
curl localhost/metrics
``` 

### Delete the cluster
```
kind delete cluster
``` 