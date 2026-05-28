resource "kubernetes_manifest" "namespace" {
  manifest = yamldecode(file("${path.module}/../k8s/namespace.yaml"))
}

resource "kubernetes_manifest" "deployment" {
  manifest = yamldecode(file("${path.module}/../k8s/deployment.yaml"))

  depends_on = [
    kubernetes_manifest.namespace
  ]
}

resource "kubernetes_manifest" "service" {
  manifest = yamldecode(file("${path.module}/../k8s/service.yaml"))

  depends_on = [
    kubernetes_manifest.deployment
  ]
}