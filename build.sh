#!/bin/bash -eux

#for local use: eval $(minikube docker-env)
docker build -t positions .
kubectl create -f go-deployment.yaml
kubectl expose deployment positions --type=NodePort --name=positions --target-port=6000