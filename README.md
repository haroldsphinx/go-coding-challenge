This is a simple GoLang Challenge that detects a user's remote IP Address

# Installation Instructions

- Setup Docker - https://docs.docker.com/compose/gettingstarted/
- Confirgure all your environment variables in the .env files e.g PORT
- Build Image - docker build -t <image_name>
- Run the image docker run -d -p $HOST_PORT:$COntainerPort <image_name>


# How to make this application Highly Available and Scalable?
There are several ways to approach Availability and Scalability in Containers:

# Using Kubernetes
1. The Kubernetes Way (This is assuming you already have a Kubernetes Cluster setup)
- kubectl run go-coding-app --image=haroldsphinx/go-coding-app --port=8083
- kubectl get pods (to see your container running)
- kubectl expose deployment go-coding-app --type=LoadBalancer --port=9000 ==target-port=8083
- kubetctl get svc (To get the service detail including the external IP Address)
- kubectl set image deployment/go-coding-app go-coding-app=haroldsphinx/go-coding-app (Set Rolling Update)

# Deploying to Serverless

You can follow the documentation based on you chosen cloud provider service