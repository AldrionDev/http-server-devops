output "namespace" {
  description = "Kubernetes namespace used by the application"
  value       = kubernetes_manifest.namespace.manifest.metadata.name
}

output "deployment_name" {
  description = "Kubernetes deployment name"
  value       = kubernetes_manifest.deployment.manifest.metadata.name
}

output "service_name" {
  description = "Kubernetes service name"
  value       = kubernetes_manifest.service.manifest.metadata.name
}