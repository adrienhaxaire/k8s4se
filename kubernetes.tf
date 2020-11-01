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

resource "kubernetes_namespace" "k8s4se" {
  metadata {
    name = "k8s4se"
  }
}

