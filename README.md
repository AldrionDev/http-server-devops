# HTTP Server DevOps Project

A simple HTTP server demonstrating a full DevOps workflow:
containerization, CI/CD, Kubernetes deployment, and Terraform infrastructure.

## Stack

- **Language:** Go
- **Container:** Docker
- **CI/CD:** GitHub Actions
- **Registry:** Docker Hub
- **Orchestration:** Kubernetes
- **IaC:** Terraform

## Planned Project Structure

```text
http-server-devops/
├── .github/
│   └── workflows/       # CI/CD pipeline
├── k8s/                 # Kubernetes manifests
├── terraform/           # Terraform configuration
├── Dockerfile
├── go.mod
├── main.go
├── main_test.go
└── README.md
```

## Tasks

- [x] Task 1: HTTP Server
- [ ] Task 2: Dockerize
- [ ] Task 3: CI/CD Pipeline
- [ ] Task 4: Kubernetes Deployment
- [ ] Task 5: Terraform Configuration

## Installation & Usage

Run the server locally:

```bash
go run .
```

The server will start on the default port `8080`. You can access it in your browser or using `curl`:

```bash
curl http://localhost:8080/
```

Configure port

```bash
PORT=9090 go run .
```

```bash
curl http://localhost:9090/
```

Run tests:

```bash
go test ./...
```
