

docker build -t adjust:v0.1 .

kind create cluster  --config kind.yml  
kind get kubeconfig > ~/.kube/config  

kubectl get nodes

kind load  docker-image adjust:v0.1


kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

kubectl apply -f deployment.yaml



kind delete cluster
