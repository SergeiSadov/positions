#!/bin/bash -eux

helm repo add elastic https://Helm.elastic.co
curl -O https://raw.githubusercontent.com/elastic/Helm-charts/master/elasticsearch/examples/minikube/values.yaml
helm install --name elasticsearch elastic/elasticsearch -f ./values.yaml
kubectl port-forward svc/elasticsearch-master 9200
helm install --name kibana elastic/kibana
kubectl port-forward deployment/kibana-kibana 5601
helm install --name metricbeat elastic/metricbeat