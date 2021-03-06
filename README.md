A super simple hello world server written in Go.


# What for?
<li> Testing your K8s setup (bare metal or otherwise)
<li> Testing connectivity, firewall or routing rules in a cluster
<li> Testing your CI/CD pipeline
<li> Just about anything you might need a standalone test web server for 🍕
  
# Usage
  
### Local Testing

Step 1. Clone the repo
  
    git clone https://github.com/pugad/hello-world-go-server
    cd hello-world-go-server

Step 2. Install Go, then run the server using:
  
    go run helloworld.go


  
To build it as a standalone executable:
  
    go build helloworld.go

### Using Docker

Install Docker, then run:
  
    docker-compose up -d

### Kubernetes
<i>Note: hello-world-k8s.yaml contains configs for metallb (as load balancer for bare metal deployments). Make sure to edit/omit it according to your needs.</i>
  
Step 1. Set up your K8s cluster (e.g. minikube, microk8s, etc.)
  
Step 2. Edit hello-world-k8s.yaml

Step 3. Deploy
  
    kubectl apply -f hello-world-k8s.yaml -n development

Step 4. Try to access the IP address allocated to it

    kubectl get svc -A
  
    curl -I http://<assigned-ip>/
