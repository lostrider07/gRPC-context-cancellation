## gRPC Context Cancellation

### Prerequisites
- Go
- Docker
- kubectl
- Kubernetes cluster (e.g., Minikube or Docker Desktop Kubernetes)

---
### Build & Run Locally with Docker
Step 1: Build Docker images
From the project root:
```
docker build -t grpc-server -f server/Dockerfile .
docker build -t grpc-client -f client/Dockerfile .
```
Step 2: Run the server
```
docker run --rm -p 50051:50051 grpc-server
```
You should see:
```
gRPC server running on :50051
```
Step 3: Run the client (in another terminal)
```
docker run --rm --network host grpc-client
```
Note: The client connects to localhost:50051 by default.
You should see the client output, either the response or a cancellation error after 2 seconds.

---
### Deploy on Kubernetes
Step 1: Push images to your container registry (Optional)
Tag and push both images:
```
docker tag grpc-server <your-registry>/grpc-server:latest
docker push <your-registry>/grpc-server:latest
docker tag grpc-client <your-registry>/grpc-client:latest
docker push <your-registry>/grpc-client:latest
```
Step 2: Update Kubernetes manifests (k8s/ folder)
Update image references in:

`grpc-server-deployment.yaml`
`grpc-client-job.yaml`

to use your registry images.

Step 3: Start a minikube
```
minikube start
```

Step 4: Deploy the server
```
kubectl apply -f k8s/grpc-server-deployment.yaml
kubectl apply -f k8s/grpc-server-service.yaml
```
Check the server pod status:
```
kubectl get pods
```
Step 5: Run the client Job
```
kubectl apply -f k8s/grpc-client-job.yaml
kubectl logs job/grpc-client
```
---
### Notes
- The client connects to the gRPC server inside Kubernetes using the service DNS name grpc-server:50051.
- The server listens on all interfaces (:50051) for incoming connections.
- The client sets a 2-second timeout and expects the server to respond in 5 seconds, demonstrating context cancellation.