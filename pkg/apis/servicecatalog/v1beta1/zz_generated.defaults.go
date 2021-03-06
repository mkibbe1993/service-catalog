// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by defaulter-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&ClusterServiceBroker{}, func(obj interface{}) { SetObjectDefaults_ClusterServiceBroker(obj.(*ClusterServiceBroker)) })
	scheme.AddTypeDefaultingFunc(&ClusterServiceBrokerList{}, func(obj interface{}) { SetObjectDefaults_ClusterServiceBrokerList(obj.(*ClusterServiceBrokerList)) })
	scheme.AddTypeDefaultingFunc(&ServiceBinding{}, func(obj interface{}) { SetObjectDefaults_ServiceBinding(obj.(*ServiceBinding)) })
	scheme.AddTypeDefaultingFunc(&ServiceBindingList{}, func(obj interface{}) { SetObjectDefaults_ServiceBindingList(obj.(*ServiceBindingList)) })
	scheme.AddTypeDefaultingFunc(&ServiceBroker{}, func(obj interface{}) { SetObjectDefaults_ServiceBroker(obj.(*ServiceBroker)) })
	scheme.AddTypeDefaultingFunc(&ServiceBrokerList{}, func(obj interface{}) { SetObjectDefaults_ServiceBrokerList(obj.(*ServiceBrokerList)) })
	scheme.AddTypeDefaultingFunc(&ServiceInstance{}, func(obj interface{}) { SetObjectDefaults_ServiceInstance(obj.(*ServiceInstance)) })
	scheme.AddTypeDefaultingFunc(&ServiceInstanceList{}, func(obj interface{}) { SetObjectDefaults_ServiceInstanceList(obj.(*ServiceInstanceList)) })
	return nil
}

func SetObjectDefaults_ClusterServiceBroker(in *ClusterServiceBroker) {
	SetDefaults_ClusterServiceBrokerSpec(&in.Spec)
}

func SetObjectDefaults_ClusterServiceBrokerList(in *ClusterServiceBrokerList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ClusterServiceBroker(a)
	}
}

func SetObjectDefaults_ServiceBinding(in *ServiceBinding) {
	SetDefaults_ServiceBinding(in)
	SetDefaults_ServiceBindingSpec(&in.Spec)
}

func SetObjectDefaults_ServiceBindingList(in *ServiceBindingList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ServiceBinding(a)
	}
}

func SetObjectDefaults_ServiceBroker(in *ServiceBroker) {
	SetDefaults_ServiceBrokerSpec(&in.Spec)
}

func SetObjectDefaults_ServiceBrokerList(in *ServiceBrokerList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ServiceBroker(a)
	}
}

func SetObjectDefaults_ServiceInstance(in *ServiceInstance) {
	SetDefaults_ServiceInstanceSpec(&in.Spec)
}

func SetObjectDefaults_ServiceInstanceList(in *ServiceInstanceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ServiceInstance(a)
	}
}
