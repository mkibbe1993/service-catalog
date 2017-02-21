#!/bin/bash
# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail
set -x

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

which kubectl
kubectl

# Clean up old containers if still around
docker rm -f etcd apiserver > /dev/null 2>&1 || true

# Start etcd, our DB
docker run --name etcd -d --net host quay.io/coreos/etcd > /dev/null

# And now our API Server
docker run --name apiserver -d --net host \
	-v ${ROOT}:/go/src/github.com/kubernetes-incubator/service-catalog \
	-v ${ROOT}/.var/run/kubernetes-service-catalog:/var/run/kubernetes-service-catalog \
	-v ${ROOT}/.kube:/root/.kube \
	scbuildimage \
	bin/apiserver -v 10 --etcd-servers http://localhost:2379 > /dev/null

sleep 10

APISERVER_IP="$(docker logs apiserver 2>&1 | tr '\r\n' ' ' | sed 's/.*Choosing IP \([0-9.]*\).*/\1/')"
echo "API Server IP: $APISERVER_IP"

# Waot for apiserver to be up and running
while ! curl -k "http://${APISERVER_IP}:6443" > /dev/null 2>&1 ; do
	sleep 1
done

which kubectl

# Setup our credentials
kubectl config set-credentials service-catalog-creds --username=admin --password=admin
kubectl config set-cluster service-catalog-cluster --server="https://${APISERVER_IP}:6443" --certificate-authority=/var/run/kubernetes-service-catalog/apiserver.crt
kubectl config set-context service-catalog-ctx --cluster=service-catalog-cluster --user=service-catalog-creds
kubectl config use-context service-catalog-ctx

# create a few resources
kubectl create -f contrib/examples/apiserver/broker.yaml
kubectl create -f contrib/examples/apiserver/serviceclass.yaml
kubectl create -f contrib/examples/apiserver/instance.yaml
kubectl create -f contrib/examples/apiserver/binding.yaml

kubectl get broker test-broker -o yaml
kubectl get serviceclass test-serviceclass -o yaml
kubectl get instance test-instance --namespace test-ns -o yaml
kubectl get binding test-binding --namespace test-ns -o yaml

kubectl delete -f contrib/examples/apiserver/broker.yaml
kubectl delete -f contrib/examples/apiserver/serviceclass.yaml
kubectl delete -f contrib/examples/apiserver/instance.yaml
kubectl delete -f contrib/examples/apiserver/binding.yaml


