provider "kubernetes" {
  config_path    = pathexpand("~/.kube/config")
  config_context = "http-server-devops"
}
