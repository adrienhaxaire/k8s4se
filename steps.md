

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


## helm
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


## elasticsearch
https://www.elastic.co/blog/getting-started-with-elastic-cloud-on-kubernetes-deployment

    λ> kubectl apply -f https://download.elastic.co/downloads/eck/0.9.0/all-in-one.yam
    λ> kubectl apply -f https://raw.githubusercontent.com/adamquan/ECK/master/apm_es_kibana.yaml

λ> mkdir -p volumes/elasticsearch/data

https://lernentec.com/post/running-simple-elasticsearch-kibana-minikube/

    λ> kubectl create -f base/elasticsearch.yaml 
    statefulset.apps/elasticsearch created
    service/elasticsearch created



mkdir volumes
minikube start --mount --mount-string=$PWD/volumes:/volumes

docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.9.3
https://hub.docker.com/_/elasticsearch

λ> minikube start --mount --mount-string=$PWD/volumes/elasticsearch:/volumes/elasticsearch

minikube start --mount --mount-string=$PWD/volumes:/volumes --driver=virtualbox

Check DNS: https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/

λ> k port-forward services/elasticsearch 9200 -n monitoring

https://dev.to/amitsaha/how-to-set-up-log-forwarding-in-a-kubernetes-cluster-using-fluent-bit-3bgk

λ> k -n monitoring exec --stdin --tty fluentd-xxxx -- /bin/sh


Okay, EFK is too complicated and resource intensive, not worth my free time
Going "Native Kubernetes"

Use cAdvisor instead, officially supported by Google: https://github.com/google/cadvisor/tree/master/deploy/kubernetes

Rancher tutorials to setup dashboard and cAdvisor:
https://rancher.com/blog/2019/native-kubernetes-monitoring-tools-part-1
https://rancher.com/blog/2019/native-kubernetes-monitoring-tools-part-2

cAdvisor UI deprecated, Heapster too, this stinks. Looking into the metrics server: https://github.com/kubernetes-sigs/metrics-server

λ> kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

Something to read during the winter, by the fire: https://github.com/kubernetes/community/tree/master/contributors/design-proposals

minikube addons enable metrics-server
minikube addons enable efk

    λ> kubectl top pod -n kube-system
    NAME                               CPU(cores)   MEMORY(bytes)   
    coredns-f9fd979d6-tjj46            7m           9Mi             
    elasticsearch-logging-tp2zh        476m         1431Mi          
    etcd-minikube                      34m          26Mi            
    fluentd-es-6k82v                   17m          91Mi            
    kibana-logging-qbkl6               32m          618Mi           
    kube-apiserver-minikube            99m          262Mi           
    kube-controller-manager-minikube   37m          41Mi            
    kube-proxy-7kl6f                   0m           14Mi            
    kube-scheduler-minikube            5m           14Mi            
    metrics-server-d9b576748-r974c     1m           10Mi            
    storage-provisioner                2m           8Mi      

ES + kibana take 2Gi!!! fluentd still takes 91Mi

RBAC: TODO

