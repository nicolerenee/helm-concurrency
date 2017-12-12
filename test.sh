#!/bin/bash


TILLER_NS=kube-system
TILLER_POD=`kubectl -n $TILLER_NS get pods \
       --selector=app=helm,name=tiller \
       -o jsonpath="{range .items[*]}{@.metadata.name}{end}" | head -n1`
kubectl port-forward -n $TILLER_NS $TILLER_POD 44134:44134 > /dev/null 2>&1 &
portforwardPID=$!

./setup.sh

echo "Starting go test..."
go test .

./cleanup.sh

kill $portforwardPID
