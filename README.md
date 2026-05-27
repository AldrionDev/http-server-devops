# HTTP Server DevOps Project

A simple HTTP server demonstrating a full DevOps workflow:
containerization, CI/CD, Kubernetes deployment, and Terraform infrastructure.

## Stack

- **Language:** Go
- **Container:** Docker
- **CI/CD:** GitHub Actions
- **Registry:** GitHub Container Registry
- **Orchestration:** Kubernetes
- **IaC:** Terraform

## Planned Project Structure

```text
http-server-devops/
├── .github/
│   └── workflows/
│       └── docker-publish.yml
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
- [x] Task 2: Dockerize
- [x] Task 3: CI/CD Pipeline
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

## Docker Usage

### Build image

```bash
docker build -t http-server-devops:local .
```

### Run container

```bash
docker run --rm -p 8080:8080 http-server-devops:local
```

### Test container

```bash
curl http://localhost:8080/
```

### Run with custom port

```bash
docker run --rm -p 9090:9090 -e PORT=9090 http-server-devops:local
```

```bash
curl http://localhost:9090/
```

## CI/CD

This project uses GitHub Actions to run tests, build the Docker image, and publish it to GitHub Container Registry.

The workflow is defined in:

```text
.github/workflows/docker-publish.yml
```

On pull requests to `main`, the pipeline:

- runs Go tests
- builds the Docker image without publishing it

On pushes to `main`, the pipeline:

- runs Go tests
- builds the Docker image
- publishes the image to GitHub Container Registry

Published image:

```text
ghcr.io/aldriondev/http-server-devops
```

Published tags:

- `latest`
- ghcr.io/aldriondev/http-server-devops:latest

The workflow uses the built-in GITHUB_TOKEN with following permissions:

```yaml
permissions:
  contents: read
  packages: write
```

No external Docker registry credentials are required.
