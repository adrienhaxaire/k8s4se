

install kubectl, minikube, kustomize and helm

follow the [minikube tutorial](https://kubernetes.io/docs/tutorials/hello-minikube/)


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












install golang




kubectl get pod,svc -n kube-system

