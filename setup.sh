#!/bin/bash

echo "Creating nicolerenee k8s namespace"
kubectl create ns nicolerenee

echo "Setting up existing helm releases"
for i in {1..20}
do
  helm install ./chart --name test$i --namespace nicolerenee --set name=test$i > /dev/null 2>&1
done
