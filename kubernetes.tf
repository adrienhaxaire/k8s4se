terraform {
  required_providers {
    kubernetes = {
      source = "hashicorp/kubernetes"
      version = "~> 1.13.3"
    }
  }
}

provider "kubernetes" {
  config_context_cluster = "minikube"
}

resource "kubernetes_deployment" "hello" {
  metadata {
    name = "hello"
    labels = {
      App = "hello"
    }
  }

  spec {
    replicas = 2
    selector {
      match_labels = {
        App = "hello"
      }
    }
    template {
      metadata {
        labels = {
          App = "hello"
        }
      }
      spec {
        container {
          image = "k8s4se/hello:latest"
          name  = "hello"

          port {
            container_port = 8080
          }

          resources {
            limits {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "hello" {
  metadata {
    name = "hello"
  }
  spec {
    selector = {
      App = kubernetes_deployment.hello.spec.0.template.0.metadata[0].labels.App
    }
    port {
      port        = 8080
      target_port = 8080
    }

    type = "LoadBalancer"
  }
}

output "lb_ip" {
  value = kubernetes_service.hello.load_balancer_ingress[0].ip
}

