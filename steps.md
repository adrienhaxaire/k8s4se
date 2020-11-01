

install kubectl, minikube, kustomize

follow the [minikube tutorial](https://kubernetes.io/docs/tutorials/hello-minikube/)

kubectl create deployment hello --image=k8s4se/hello


configure the registry: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
kubectl create secret generic dockerhub  --from-file=.dockerconfigjson=$HOME/.docker/config.json --type=kubernetes.io/dockerconfigjson
kubectl get secret dockerhub --output="jsonpath={.data.\.dockerconfigjson}" | base64 --decode

https://minikube.sigs.k8s.io/docs/handbook/registry/

kubectl get service $SERVICE --output='jsonpath="{.spec.ports[0].nodePort}"'

λ> k get service hello-service --output='jsonpath="{.spec.ports[0].nodePort}"'
"32055"λ> 
λ> kubectl get svc
NAME            TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello-service   LoadBalancer   10.110.239.63   <pending>     8080:32055/TCP   5m27s
kubernetes      ClusterIP      10.96.0.1       <none>        443/TCP          16h
λ> minikube service --url hello-service
http://172.17.0.2:32055
λ> curl http://172.17.0.2:32055/toto
Hello, you've requested: /toto


This `pending` is normal:

    λ> kubectl get svc
    NAME            TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
    hello-service   LoadBalancer   10.110.239.63   <pending>     8080:32055/TCP   5m27s
    kubernetes      ClusterIP      10.96.0.1       <none>        443/TCP          16h

as `minikube` doesn't have an integrated load balancer: https://stackoverflow.com/questions/44110876/kubernetes-service-external-ip-pending



## helm

https://helm.sh/docs/intro/quickstart/#initialize-a-helm-chart-repository

     λ> helm repo add stable https://charts.helm.sh/stable
     "stable" has been added to your repositories
     helm repo add stable https://charts.helm.sh/stable
     λ> "bitnami" has been added to your repositories

https://helm.sh/docs/intro/using_helm/#creating-your-own-charts


install golang

https://golang.org/doc/articles/wiki/


## terraform

https://www.hashicorp.com/blog/using-the-kubernetes-and-helm-providers-with-terraform-0-12

https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs

For elasticsearch
minikube addons enable default-storageclass
minikube addons enable storage-provisioner
minikube addons enable ingress

minikube start --cpus=6 --memory=16g

minikube config set cpus 6 
minikube config set memory 16384


helm install hello applications/hello/charts

https://artifacthub.io/packages/helm/helm-stable/fluentd
helm install fluentd stable/fluentd

helm repo add elastic https://helm.elastic.co

https://artifacthub.io/packages/helm/elastic/kibana
helm install kibana elastic/kibana --version 7.9.3 --set service.type=LoadBalancer

https://artifacthub.io/packages/helm/elastic/elasticsearch
helm install elasticsearch elastic/elasticsearch --version 7.9.3
helm install elasticsearch elastic/elasticsearch --version 7.9.3 --values elasticsearch-values.yaml 
kubectl expose deploy elasticsearch --port 9200


minikube ssh

Watch out! you need to call `minikube addons enable ingress` **after** having created the ingress resource!
