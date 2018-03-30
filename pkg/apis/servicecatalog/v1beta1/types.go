/*
Copyright 2016 The Kubernetes Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterServiceBroker represents an entity that provides
// ClusterServiceClasses for use in the service catalog.
// +k8s:openapi-gen=x-kubernetes-print-columns:custom-columns=NAME:.metadata.name,URL:.spec.url
type ClusterServiceBroker struct {
	metav1.TypeMeta `json:",inline"`

	// Non-namespaced.  The name of this resource in etcd is in ObjectMeta.Name.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the broker.
	// +optional
	Spec ClusterServiceBrokerSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current status of a broker.
	// +optional
	Status ClusterServiceBrokerStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterServiceBrokerList is a list of Brokers.
type ClusterServiceBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ClusterServiceBroker `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceBroker represents an entity that provides
// ServiceClasses for use in the service catalog.
// +k8s:openapi-gen=x-kubernetes-print-columns:custom-columns=NAME:.metadata.name,URL:.spec.url
type ServiceBroker struct {
	metav1.TypeMeta `json:",inline"`

	// The name of this resource in etcd is in ObjectMeta.Name.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the broker.
	// +optional
	Spec ServiceBrokerSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current status of a broker.
	// +optional
	Status ServiceBrokerStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceBrokerList is a list of Brokers.
type ServiceBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ServiceBroker `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// CommonServiceBrokerSpec represents a description of a Broker.
type CommonServiceBrokerSpec struct {
	// URL is the address used to communicate with the ServiceBroker.
	URL string `json:"url" protobuf:"bytes,1,opt,name=url"`

	// InsecureSkipTLSVerify disables TLS certificate verification when communicating with this Broker.
	// This is strongly discouraged.  You should use the CABundle instead.
	// +optional
	InsecureSkipTLSVerify bool `json:"insecureSkipTLSVerify,omitempty" protobuf:"varint,2,opt,name=insecureSkipTLSVerify"`

	// CABundle is a PEM encoded CA bundle which will be used to validate a Broker's serving certificate.
	// +optional
	CABundle []byte `json:"caBundle,omitempty" protobuf:"bytes,3,opt,name=caBundle"`

	// RelistBehavior specifies the type of relist behavior the catalog should
	// exhibit when relisting ServiceClasses available from a broker.
	// +optional
	RelistBehavior ServiceBrokerRelistBehavior `json:"relistBehavior" protobuf:"bytes,4,opt,name=relistBehavior"`

	// RelistDuration is the frequency by which a controller will relist the
	// broker when the RelistBehavior is set to ServiceBrokerRelistBehaviorDuration.
	// Users are cautioned against configuring low values for the RelistDuration,
	// as this can easily overload the controller manager in an environment with
	// many brokers. The actual interval is intrinsically governed by the
	// configured resync interval of the controller, which acts as a minimum bound.
	// For example, with a resync interval of 5m and a RelistDuration of 2m, relists
	// will occur at the resync interval of 5m.
	RelistDuration *metav1.Duration `json:"relistDuration,omitempty" protobuf:"bytes,5,opt,name=relistDuration"`

	// RelistRequests is a strictly increasing, non-negative integer counter that
	// can be manually incremented by a user to manually trigger a relist.
	// +optional
	RelistRequests int64 `json:"relistRequests" protobuf:"varint,6,opt,name=relistRequests"`
}

// ClusterServiceBrokerSpec represents a description of a Broker.
type ClusterServiceBrokerSpec struct {
	CommonServiceBrokerSpec `json:",inline" protobuf:"bytes,1,opt,name=commonServiceBrokerSpec"`

	// AuthInfo contains the data that the service catalog should use to authenticate
	// with the ClusterServiceBroker.
	AuthInfo *ClusterServiceBrokerAuthInfo `json:"authInfo,omitempty" protobuf:"bytes,2,opt,name=authInfo"`
}

// ServiceBrokerSpec represents a description of a Broker.
type ServiceBrokerSpec struct {
	CommonServiceBrokerSpec `json:",inline" protobuf:"bytes,1,opt,name=commonServiceBrokerSpec"`

	// AuthInfo contains the data that the service catalog should use to authenticate
	// with the ServiceBroker.
	AuthInfo *ServiceBrokerAuthInfo `json:"authInfo,omitempty" protobuf:"bytes,2,opt,name=authInfo"`
}

// ServiceBrokerRelistBehavior represents a type of broker relist behavior.
type ServiceBrokerRelistBehavior string

const (
	// ServiceBrokerRelistBehaviorDuration indicates that the broker will be
	// relisted automatically after the specified duration has passed.
	ServiceBrokerRelistBehaviorDuration ServiceBrokerRelistBehavior = "Duration"

	// ServiceBrokerRelistBehaviorManual indicates that the broker is only
	// relisted when the spec of the broker changes.
	ServiceBrokerRelistBehaviorManual ServiceBrokerRelistBehavior = "Manual"
)

// ClusterServiceBrokerAuthInfo is a union type that contains information on
// one of the authentication methods the the service catalog and brokers may
// support, according to the OpenServiceBroker API specification
// (https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md).
type ClusterServiceBrokerAuthInfo struct {
	// ClusterBasicAuthConfigprovides configuration for basic authentication.
	Basic *ClusterBasicAuthConfig `json:"basic,omitempty" protobuf:"bytes,1,opt,name=basic"`
	// ClusterBearerTokenAuthConfig provides configuration to send an opaque value as a bearer token.
	// The value is referenced from the 'token' field of the given secret.  This value should only
	// contain the token value and not the `Bearer` scheme.
	Bearer *ClusterBearerTokenAuthConfig `json:"bearer,omitempty" protobuf:"bytes,2,opt,name=bearer"`
}

// ClusterBasicAuthConfig provides config for the basic authentication of
// cluster scoped brokers.
type ClusterBasicAuthConfig struct {
	// SecretRef is a reference to a Secret containing information the
	// catalog should use to authenticate to this ServiceBroker.
	//
	// Required at least one of the fields:
	// - Secret.Data["username"] - username used for authentication
	// - Secret.Data["password"] - password or token needed for authentication
	SecretRef *ObjectReference `json:"secretRef,omitempty" protobuf:"bytes,1,opt,name=secretRef"`
}

// ClusterBearerTokenAuthConfig provides config for the bearer token
// authentication of cluster scoped brokers.
type ClusterBearerTokenAuthConfig struct {
	// SecretRef is a reference to a Secret containing information the
	// catalog should use to authenticate to this ServiceBroker.
	//
	// Required field:
	// - Secret.Data["token"] - bearer token for authentication
	SecretRef *ObjectReference `json:"secretRef,omitempty" protobuf:"bytes,1,opt,name=secretRef"`
}

// ServiceBrokerAuthInfo is a union type that contains information on
// one of the authentication methods the the service catalog and brokers may
// support, according to the OpenServiceBroker API specification
// (https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md).
type ServiceBrokerAuthInfo struct {
	// BasicAuthConfig provides configuration for basic authentication.
	Basic *BasicAuthConfig `json:"basic,omitempty" protobuf:"bytes,1,opt,name=basic"`
	// BearerTokenAuthConfig provides configuration to send an opaque value as a bearer token.
	// The value is referenced from the 'token' field of the given secret.  This value should only
	// contain the token value and not the `Bearer` scheme.
	Bearer *BearerTokenAuthConfig `json:"bearer,omitempty" protobuf:"bytes,2,opt,name=bearer"`
}

// BasicAuthConfig provides config for the basic authentication of
// cluster scoped brokers.
type BasicAuthConfig struct {
	// SecretRef is a reference to a Secret containing information the
	// catalog should use to authenticate to this ServiceBroker.
	//
	// Required at least one of the fields:
	// - Secret.Data["username"] - username used for authentication
	// - Secret.Data["password"] - password or token needed for authentication
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,1,opt,name=secretRef"`
}

// BearerTokenAuthConfig provides config for the bearer token
// authentication of cluster scoped brokers.
type BearerTokenAuthConfig struct {
	// SecretRef is a reference to a Secret containing information the
	// catalog should use to authenticate to this ServiceBroker.
	//
	// Required field:
	// - Secret.Data["token"] - bearer token for authentication
	SecretRef *LocalObjectReference `json:"secretRef,omitempty" protobuf:"bytes,1,opt,name=secretRef"`
}

const (
	// BasicAuthUsernameKey is the key of the username for SecretTypeBasicAuth secrets
	BasicAuthUsernameKey = "username"
	// BasicAuthPasswordKey is the key of the password or token for SecretTypeBasicAuth secrets
	BasicAuthPasswordKey = "password"

	// BearerTokenKey is the key of the bearer token for SecretTypeBearerTokenAuth secrets
	BearerTokenKey = "token"
)

// CommonServiceBrokerStatus represents the current status of a Broker.
type CommonServiceBrokerStatus struct {
	Conditions []ServiceBrokerCondition `json:"conditions" protobuf:"bytes,1,rep,name=conditions"`

	// ReconciledGeneration is the 'Generation' of the ClusterServiceBrokerSpec that
	// was last processed by the controller. The reconciled generation is updated
	// even if the controller failed to process the spec.
	ReconciledGeneration int64 `json:"reconciledGeneration" protobuf:"varint,2,opt,name=reconciledGeneration"`

	// OperationStartTime is the time at which the current operation began.
	OperationStartTime *metav1.Time `json:"operationStartTime,omitempty" protobuf:"bytes,3,opt,name=operationStartTime"`

	// LastCatalogRetrievalTime is the time the Catalog was last fetched from
	// the Service Broker
	LastCatalogRetrievalTime *metav1.Time `json:"lastCatalogRetrievalTime,omitempty" protobuf:"bytes,4,opt,name=lastCatalogRetrievalTime"`
}

// ClusterServiceBrokerStatus represents the current status of a
// ClusterServiceBroker.
type ClusterServiceBrokerStatus struct {
	CommonServiceBrokerStatus `json:",inline" protobuf:"bytes,1,opt,name=commonServiceBrokerStatus"`
}

// ServiceBrokerStatus the current status of a ServiceBroker.
type ServiceBrokerStatus struct {
	CommonServiceBrokerStatus `json:",inline" protobuf:"bytes,1,opt,name=commonServiceBrokerStatus"`
}

// ServiceBrokerCondition contains condition information for a Broker.
type ServiceBrokerCondition struct {
	// Type of the condition, currently ('Ready').
	Type ServiceBrokerConditionType `json:"type" protobuf:"bytes,1,opt,name=type"`

	// Status of the condition, one of ('True', 'False', 'Unknown').
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,3,opt,name=lastTransitionTime"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	Reason string `json:"reason" protobuf:"bytes,4,opt,name=reason"`

	// Message is a human readable description of the details of the last
	// transition, complementing reason.
	Message string `json:"message" protobuf:"bytes,5,opt,name=message"`
}

// ServiceBrokerConditionType represents a broker condition value.
type ServiceBrokerConditionType string

const (
	// ServiceBrokerConditionReady represents the fact that a given broker condition
	// is in ready state.
	ServiceBrokerConditionReady ServiceBrokerConditionType = "Ready"

	// ServiceBrokerConditionFailed represents information about a final failure
	// that should not be retried.
	ServiceBrokerConditionFailed ServiceBrokerConditionType = "Failed"
)

// ConditionStatus represents a condition's status.
type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in
// the condition; "ConditionFalse" means a resource is not in the condition;
// "ConditionUnknown" means kubernetes can't decide if a resource is in the
// condition or not. In the future, we could add other intermediate
// conditions, e.g. ConditionDegraded.
const (
	// ConditionTrue represents the fact that a given condition is true
	ConditionTrue ConditionStatus = "True"

	// ConditionFalse represents the fact that a given condition is false
	ConditionFalse ConditionStatus = "False"

	// ConditionUnknown represents the fact that a given condition is unknown
	ConditionUnknown ConditionStatus = "Unknown"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterServiceClassList is a list of ClusterServiceClasses.
type ClusterServiceClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ClusterServiceClass `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterServiceClass represents an offering in the service catalog.
// +k8s:openapi-gen=x-kubernetes-print-columns:custom-columns=NAME:.metadata.name,EXTERNAL NAME:.spec.externalName,BROKER:.spec.clusterServiceBrokerName,BINDABLE:.spec.bindable,PLAN UPDATABLE:.spec.planUpdatable
type ClusterServiceClass struct {
	metav1.TypeMeta `json:",inline"`

	// Non-namespaced.  The name of this resource in etcd is in ObjectMeta.Name.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the cluster service class.
	// +optional
	Spec ClusterServiceClassSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current status of the cluster service class.
	// +optional
	Status ClusterServiceClassStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceClassList is a list of ServiceClasses.
type ServiceClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ServiceClass `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceClass represents a namespaced offering in the service catalog.
type ServiceClass struct {
	metav1.TypeMeta `json:",inline"`

	// The name of this resource in etcd is in ObjectMeta.Name.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the service class.
	// +optional
	Spec ServiceClassSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current status of a service class.
	// +optional
	Status ServiceClassStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ServiceClassStatus represents status information about a ServiceClass.
type ServiceClassStatus struct {
	CommonServiceClassStatus `json:",inline" protobuf:"bytes,1,opt,name=commonServiceClassStatus"`
}

// ClusterServiceClassStatus represents status information about a
// ClusterServiceClass.
type ClusterServiceClassStatus struct {
	CommonServiceClassStatus `json:",inline" protobuf:"bytes,2,opt,name=commonServiceClassStatus"`
}

// CommonServiceClassStatus represents common status information between
// cluster scoped and namespace scoped ServiceClasses.
type CommonServiceClassStatus struct {
	// RemovedFromBrokerCatalog indicates that the broker removed the service from its
	// catalog.
	RemovedFromBrokerCatalog bool `json:"removedFromBrokerCatalog" protobuf:"varint,1,opt,name=removedFromBrokerCatalog"`
}

// CommonServiceClassSpec represents details about a ServiceClass
type CommonServiceClassSpec struct {
	// ExternalName is the name of this object that the Service Broker
	// exposed this Service Class as. Mutable.
	ExternalName string `json:"externalName" protobuf:"bytes,1,opt,name=externalName"`

	// ExternalID is the identity of this object for use with the OSB API.
	//
	// Immutable.
	ExternalID string `json:"externalID" protobuf:"bytes,2,opt,name=externalID"`

	// Description is a short description of this ServiceClass.
	Description string `json:"description" protobuf:"bytes,3,opt,name=description"`

	// Bindable indicates whether a user can create bindings to an
	// ServiceInstance provisioned from this service. ServicePlan
	// has an optional field called Bindable which overrides the value of
	// this field.
	Bindable bool `json:"bindable" protobuf:"varint,4,opt,name=bindable"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// BindingRetrievable indicates whether fetching a binding via a GET on
	// its endpoint is supported for all plans.
	BindingRetrievable bool `json:"bindingRetrievable" protobuf:"varint,10,opt,name=bindingRetrievable"`

	// PlanUpdatable indicates whether instances provisioned from this
	// ServiceClass may change ServicePlans after being
	// provisioned.
	PlanUpdatable bool `json:"planUpdatable" protobuf:"varint,6,opt,name=planUpdatable"`

	// ExternalMetadata is a blob of information about the
	// ServiceClass, meant to be user-facing content and display
	// instructions. This field may contain platform-specific conventional
	// values.
	ExternalMetadata *runtime.RawExtension `json:"externalMetadata,omitempty" protobuf:"bytes,7,opt,name=externalMetadata"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// Tags is a list of strings that represent different classification
	// attributes of the ServiceClass.  These are used in Cloud
	// Foundry in a way similar to Kubernetes labels, but they currently
	// have no special meaning in Kubernetes.
	Tags []string `json:"tags,omitempty" protobuf:"bytes,8,rep,name=tags"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// Requires exposes a list of Cloud Foundry-specific 'permissions'
	// that must be granted to an instance of this service within Cloud
	// Foundry.  These 'permissions' have no meaning within Kubernetes and an
	// ServiceInstance provisioned from this ServiceClass will not
	// work correctly.
	Requires []string `json:"requires,omitempty" protobuf:"bytes,9,rep,name=requires"`
}

// ClusterServiceClassSpec represents the details about a ClusterServiceClass
type ClusterServiceClassSpec struct {
	CommonServiceClassSpec `json:",inline" protobuf:"bytes,1,opt,name=commonServiceClassSpec"`

	// ClusterServiceBrokerName is the reference to the Broker that provides this
	// ClusterServiceClass.
	//
	// Immutable.
	ClusterServiceBrokerName string `json:"clusterServiceBrokerName" protobuf:"bytes,2,opt,name=clusterServiceBrokerName"`
}

// ServiceClassSpec represents the details about a ServiceClass
type ServiceClassSpec struct {
	CommonServiceClassSpec `json:",inline" protobuf:"bytes,1,opt,name=commonServiceClassSpec"`

	// ServiceBrokerName is the reference to the Broker that provides this
	// ServiceClass.
	//
	// Immutable.
	ServiceBrokerName string `json:"serviceBrokerName" protobuf:"bytes,2,opt,name=serviceBrokerName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterServicePlanList is a list of ClusterServicePlans.
type ClusterServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ClusterServicePlan `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterServicePlan represents a tier of a ServiceClass.
// +k8s:openapi-gen=x-kubernetes-print-columns:custom-columns=NAME:.metadata.name,EXTERNAL NAME:.spec.externalName,BROKER:.spec.clusterServiceBrokerName,CLASS:.spec.clusterServiceClassRef.name
type ClusterServicePlan struct {
	metav1.TypeMeta `json:",inline"`

	// Non-namespaced.  The name of this resource in etcd is in ObjectMeta.Name.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the service plan.
	// +optional
	Spec ClusterServicePlanSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current status of the service plan.
	// +optional
	Status ClusterServicePlanStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// CommonServicePlanSpec represents details that are shared by both
// a ClusterServicePlan and a namespaced ServicePlan
type CommonServicePlanSpec struct {
	// ExternalName is the name of this object that the Service Broker
	// exposed this Service Plan as. Mutable.
	ExternalName string `json:"externalName" protobuf:"bytes,1,opt,name=externalName"`

	// ExternalID is the identity of this object for use with the OSB API.
	//
	// Immutable.
	ExternalID string `json:"externalID" protobuf:"bytes,2,opt,name=externalID"`

	// Description is a short description of this ServicePlan.
	Description string `json:"description" protobuf:"bytes,3,opt,name=description"`

	// Bindable indicates whether a user can create bindings to an
	// ServiceInstance using this ServicePlan.  If set, overrides
	// the value of the corresponding ServiceClassSpec Bindable field.
	Bindable *bool `json:"bindable,omitempty" protobuf:"varint,4,opt,name=bindable"`

	// Free indicates whether this plan is available at no cost.
	Free bool `json:"free" protobuf:"varint,5,opt,name=free"`

	// ExternalMetadata is a blob of information about the plan, meant to be
	// user-facing content and display instructions.  This field may contain
	// platform-specific conventional values.
	ExternalMetadata *runtime.RawExtension `json:"externalMetadata,omitempty" protobuf:"bytes,6,opt,name=externalMetadata"` //TODO(mkibbe): why *runtime.RawExtension as opposed to runtime.RawExtension?

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// ServiceInstanceCreateParameterSchema is the schema for the parameters
	// that may be supplied when provisioning a new ServiceInstance on this plan.
	ServiceInstanceCreateParameterSchema *runtime.RawExtension `json:"instanceCreateParameterSchema,omitempty" protobuf:"bytes,7,opt,name=instanceCreateParameterSchema"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// ServiceInstanceUpdateParameterSchema is the schema for the parameters
	// that may be updated once an ServiceInstance has been provisioned on
	// this plan. This field only has meaning if the corresponding ServiceClassSpec is
	// PlanUpdatable.
	ServiceInstanceUpdateParameterSchema *runtime.RawExtension `json:"instanceUpdateParameterSchema,omitempty" protobuf:"bytes,8,opt,name=instanceUpdateParameterSchema"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// ServiceBindingCreateParameterSchema is the schema for the parameters that
	// may be supplied binding to a ServiceInstance on this plan.
	ServiceBindingCreateParameterSchema *runtime.RawExtension `json:"serviceBindingCreateParameterSchema,omitempty" protobuf:"bytes,9,opt,name=serviceBindingCreateParameterSchema"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.when a bind operation stored in the
	// Secret when binding to a ServiceInstance on this plan.
	// The ResponseSchema feature gate needs to be enabled for this field to
	// be populated.
	//
	// ServiceBindingCreateResponseSchema is the schema for the response that
	// will be returned by the broker when binding to a ServiceInstance on this plan.
	// The schema also contains the sub-schema for the credentials part of the
	// broker's response, which allows clients to see what the credentials
	// will look like even before the binding operation is performed.
	ServiceBindingCreateResponseSchema *runtime.RawExtension `json:"serviceBindingCreateResponseSchema,omitempty" protobuf:"bytes,10,opt,name=serviceBindingCreateResponseSchema"`
}

// ClusterServicePlanSpec represents details about a ClusterServicePlan.
type ClusterServicePlanSpec struct {
	// CommonServicePlanSpec contains the common details of this ClusterServicePlan
	CommonServicePlanSpec `json:",inline" protobuf:"bytes,3,opt,name=commonServicePlanSpec"`

	// ClusterServiceBrokerName is the name of the ClusterServiceBroker
	// that offers this ClusterServicePlan.
	ClusterServiceBrokerName string `json:"clusterServiceBrokerName" protobuf:"bytes,1,opt,name=clusterServiceBrokerName"`

	// ClusterServiceClassRef is a reference to the service class that
	// owns this plan.
	ClusterServiceClassRef ClusterObjectReference `json:"clusterServiceClassRef" protobuf:"bytes,2,opt,name=clusterServiceClassRef"`
}

// ClusterServicePlanStatus represents status information about a
// ClusterServicePlan.
type ClusterServicePlanStatus struct {
	CommonServicePlanStatus `json:",inline" protobuf:"bytes,1,opt,name=commonServicePlanStatus"`
}

// CommonServicePlanStatus represents status information about a
// ClusterServicePlan or a ServicePlan.
type CommonServicePlanStatus struct {
	// RemovedFromBrokerCatalog indicates that the broker removed the plan
	// from its catalog.
	RemovedFromBrokerCatalog bool `json:"removedFromBrokerCatalog" protobuf:"varint,1,opt,name=removedFromBrokerCatalog"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServicePlanList is a list of rServicePlans.
type ServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ServicePlan `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServicePlan represents a tier of a ServiceClass.
// +k8s:openapi-gen=x-kubernetes-print-columns:custom-columns=NAME:.metadata.name,EXTERNAL NAME:.spec.externalName,BROKER:.spec.serviceBrokerName,CLASS:.spec.serviceClassRef.name
type ServicePlan struct {
	metav1.TypeMeta `json:",inline"`

	// Non-namespaced.  The name of this resource in etcd is in ObjectMeta.Name.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the service plan.
	// +optional
	Spec ServicePlanSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current status of the service plan.
	// +optional
	Status ServicePlanStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ServicePlanSpec represents details about a ServicePlan.
type ServicePlanSpec struct {
	// CommonServicePlanSpec contains the common details of this ServicePlan
	CommonServicePlanSpec `json:",inline" protobuf:"bytes,1,opt,name=commonServicePlanSpec"`

	// ServiceBrokerName is the name of the ServiceBroker
	// that offers this ServicePlan.
	ServiceBrokerName string `json:"serviceBrokerName" protobuf:"bytes,2,opt,name=serviceBrokerName"`

	// ServiceClassRef is a reference to the service class that
	// owns this plan.
	ServiceClassRef LocalObjectReference `json:"serviceClassRef" protobuf:"bytes,3,opt,name=serviceClassRef"`
}

// ServicePlanStatus represents status information about a
// ServicePlan.
type ServicePlanStatus struct {
	CommonServicePlanStatus `json:",inline" protobuf:"bytes,1,opt,name=commonServicePlanStatus"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceInstanceList is a list of instances.
type ServiceInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ServiceInstance `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// UserInfo holds information about the user that last changed a resource's spec.
type UserInfo struct {
	Username string                `json:"username" protobuf:"bytes,1,opt,name=username"`
	UID      string                `json:"uid" protobuf:"bytes,2,opt,name=uid"`
	Groups   []string              `json:"groups,omitempty" protobuf:"bytes,3,rep,name=groups"`
	Extra    map[string]ExtraValue `json:"extra,omitempty" protobuf:"bytes,4,rep,name=extra"`
}

// ExtraValue contains additional information about a user that may be
// provided by the authenticator.
type ExtraValue []string

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceInstance represents a provisioned instance of a ServiceClass.
// Currently, the spec field cannot be changed once a ServiceInstance is
// created.  Spec changes submitted by users will be ignored.
//
// In the future, this will be allowed and will represent the intention that
// the ServiceInstance should have the plan and/or parameters updated at the
// ClusterServiceBroker.
// +k8s:openapi-gen=x-kubernetes-print-columns:custom-columns=NAME:.metadata.name,CLASS:.spec.clusterServiceClassExternalName,PLAN:.spec.clusterServicePlanExternalName
type ServiceInstance struct {
	metav1.TypeMeta `json:",inline"`

	// The name of this resource in etcd is in ObjectMeta.Name.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the behavior of the service instance.
	// +optional
	Spec ServiceInstanceSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current status of a service instance.
	// +optional
	Status ServiceInstanceStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// PlanReference defines the user specification for the desired
// ServicePlan and ServiceClass. Because there are multiple ways to
// specify the desired Class/Plan, this structure specifies the
// allowed ways to specify the intent.
//
// Currently supported ways:
//  - ClusterServiceClassExternalName and ClusterServicePlanExternalName
//  - ClusterServiceClassName and ClusterServicePlanName
//
// For both of these ways, if a ClusterServiceClass only has one plan
// then leaving the *ServicePlanName is optional.
type PlanReference struct {
	// ClusterServiceClassExternalName is the human-readable name of the
	// service as reported by the broker. Note that if the broker changes
	// the name of the ClusterServiceClass, it will not be reflected here,
	// and to see the current name of the ClusterServiceClass, you should
	// follow the ClusterServiceClassRef below.
	//
	// Immutable.
	ClusterServiceClassExternalName string `json:"clusterServiceClassExternalName,omitempty" protobuf:"bytes,1,opt,name=clusterServiceClassExternalName"`
	// ClusterServicePlanExternalName is the human-readable name of the plan
	// as reported by the broker. Note that if the broker changes the name
	// of the ClusterServicePlan, it will not be reflected here, and to see
	// the current name of the ClusterServicePlan, you should follow the
	// ClusterServicePlanRef below.
	ClusterServicePlanExternalName string `json:"clusterServicePlanExternalName,omitempty" protobuf:"bytes,2,opt,name=clusterServicePlanExternalName"`

	// ClusterServiceClassName is the kubernetes name of the
	// ClusterServiceClass.
	//
	// Immutable.
	ClusterServiceClassName string `json:"clusterServiceClassName,omitempty" protobuf:"bytes,3,opt,name=clusterServiceClassName"`
	// ClusterServicePlanName is kubernetes name of the ClusterServicePlan.
	ClusterServicePlanName string `json:"clusterServicePlanName,omitempty" protobuf:"bytes,4,opt,name=clusterServicePlanName"`
}

// ServiceInstanceSpec represents the desired state of an Instance.
type ServiceInstanceSpec struct {
	// Specification of what ServiceClass/ServicePlan is being provisioned.
	PlanReference `json:",inline" protobuf:"bytes,8,opt,name=planReference"`

	// ClusterServiceClassRef is a reference to the ClusterServiceClass
	// that the user selected.
	// This is set by the controller based on
	// ClusterServiceClassExternalName
	ClusterServiceClassRef *ClusterObjectReference `json:"clusterServiceClassRef,omitempty" protobuf:"bytes,1,opt,name=clusterServiceClassRef"`
	// ClusterServicePlanRef is a reference to the ClusterServicePlan
	// that the user selected.
	// This is set by the controller based on
	// ClusterServicePlanExternalName
	ClusterServicePlanRef *ClusterObjectReference `json:"clusterServicePlanRef,omitempty" protobuf:"bytes,2,opt,name=clusterServicePlanRef"`

	// Parameters is a set of the parameters to be passed to the underlying
	// broker. The inline YAML/JSON payload to be translated into equivalent
	// JSON object. If a top-level parameter name exists in multiples sources
	// among `Parameters` and `ParametersFrom` fields, it is considered to be
	// a user error in the specification.
	//
	// The Parameters field is NOT secret or secured in any way and should
	// NEVER be used to hold sensitive information. To set parameters that
	// contain secret information, you should ALWAYS store that information
	// in a Secret and use the ParametersFrom field.
	//
	// +optional
	Parameters *runtime.RawExtension `json:"parameters,omitempty" protobuf:"bytes,3,opt,name=parameters"`

	// List of sources to populate parameters.
	// If a top-level parameter name exists in multiples sources among
	// `Parameters` and `ParametersFrom` fields, it is
	// considered to be a user error in the specification
	// +optional
	ParametersFrom []ParametersFromSource `json:"parametersFrom,omitempty" protobuf:"bytes,4,rep,name=parametersFrom"`

	// ExternalID is the identity of this object for use with the OSB SB API.
	//
	// Immutable.
	// +optional
	ExternalID string `json:"externalID" protobuf:"bytes,5,opt,name=externalID"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// UserInfo contains information about the user that last modified this
	// instance. This field is set by the API server and not settable by the
	// end-user. User-provided values for this field are not saved.
	// +optional
	UserInfo *UserInfo `json:"userInfo,omitempty" protobuf:"bytes,6,opt,name=userInfo"`

	// UpdateRequests is a strictly increasing, non-negative integer counter that
	// can be manually incremented by a user to manually trigger an update. This
	// allows for parameters to be updated with any out-of-band changes that have
	// been made to the secrets from which the parameters are sourced.
	// +optional
	UpdateRequests int64 `json:"updateRequests" protobuf:"varint,7,opt,name=updateRequests"`
}

// ServiceInstanceStatus represents the current status of an Instance.
type ServiceInstanceStatus struct {
	// Conditions is an array of ServiceInstanceConditions capturing aspects of an
	// ServiceInstance's status.
	Conditions []ServiceInstanceCondition `json:"conditions" protobuf:"bytes,1,rep,name=conditions"`

	// AsyncOpInProgress is set to true if there is an ongoing async operation
	// against this Service Instance in progress.
	AsyncOpInProgress bool `json:"asyncOpInProgress" protobuf:"varint,2,opt,name=asyncOpInProgress"`

	// OrphanMitigationInProgress is set to true if there is an ongoing orphan
	// mitigation operation against this ServiceInstance in progress.
	OrphanMitigationInProgress bool `json:"orphanMitigationInProgress" protobuf:"varint,3,opt,name=orphanMitigationInProgress"`

	// LastOperation is the string that the broker may have returned when
	// an async operation started, it should be sent back to the broker
	// on poll requests as a query param.
	LastOperation *string `json:"lastOperation,omitempty" protobuf:"bytes,4,opt,name=lastOperation"`

	// DashboardURL is the URL of a web-based management user interface for
	// the service instance.
	DashboardURL *string `json:"dashboardURL,omitempty" protobuf:"bytes,5,opt,name=dashboardURL"`

	// CurrentOperation is the operation the Controller is currently performing
	// on the ServiceInstance.
	CurrentOperation ServiceInstanceOperation `json:"currentOperation,omitempty" protobuf:"bytes,6,opt,name=currentOperation"`

	// ReconciledGeneration is the 'Generation' of the serviceInstanceSpec that
	// was last processed by the controller. The reconciled generation is updated
	// even if the controller failed to process the spec.
	// Deprecated: use ObservedGeneration with conditions set to true to find
	// whether generation was reconciled.
	ReconciledGeneration int64 `json:"reconciledGeneration" protobuf:"varint,7,opt,name=reconciledGeneration"`

	// ObservedGeneration is the 'Generation' of the serviceInstanceSpec that
	// was last processed by the controller. The observed generation is updated
	// whenever the status is updated regardless of operation result.
	ObservedGeneration int64 `json:"observedGeneration" protobuf:"varint,8,opt,name=observedGeneration"`

	// OperationStartTime is the time at which the current operation began.
	OperationStartTime *metav1.Time `json:"operationStartTime,omitempty" protobuf:"bytes,9,opt,name=operationStartTime"`

	// InProgressProperties is the properties state of the ServiceInstance when
	// a Provision or Update is in progress. If the current operation is a
	// Deprovision, this will be nil.
	InProgressProperties *ServiceInstancePropertiesState `json:"inProgressProperties,omitempty" protobuf:"bytes,10,opt,name=inProgressProperties"`

	// ExternalProperties is the properties state of the ServiceInstance which the
	// broker knows about.
	ExternalProperties *ServiceInstancePropertiesState `json:"externalProperties,omitempty" protobuf:"bytes,11,opt,name=externalProperties"`

	// ProvisionStatus describes whether the instance is in the provisioned state.
	ProvisionStatus ServiceInstanceProvisionStatus `json:"provisionStatus" protobuf:"bytes,12,opt,name=provisionStatus"`

	// DeprovisionStatus describes what has been done to deprovision the
	// ServiceInstance.
	DeprovisionStatus ServiceInstanceDeprovisionStatus `json:"deprovisionStatus" protobuf:"bytes,13,opt,name=deprovisionStatus"`
}

// ServiceInstanceCondition contains condition information about an Instance.
type ServiceInstanceCondition struct {
	// Type of the condition, currently ('Ready').
	Type ServiceInstanceConditionType `json:"type" protobuf:"bytes,1,opt,name=type"`

	// Status of the condition, one of ('True', 'False', 'Unknown').
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,3,opt,name=lastTransitionTime"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	Reason string `json:"reason" protobuf:"bytes,4,opt,name=reason"`

	// Message is a human readable description of the details of the last
	// transition, complementing reason.
	Message string `json:"message" protobuf:"bytes,5,opt,name=message"`
}

// ServiceInstanceConditionType represents a instance condition value.
type ServiceInstanceConditionType string

const (
	// ServiceInstanceConditionReady represents that a given InstanceCondition is in
	// ready state.
	ServiceInstanceConditionReady ServiceInstanceConditionType = "Ready"

	// ServiceInstanceConditionFailed represents information about a final failure
	// that should not be retried.
	ServiceInstanceConditionFailed ServiceInstanceConditionType = "Failed"

	// ServiceInstanceConditionOrphanMitigation represents information about an
	// orphan mitigation that is required after failed provisioning.
	ServiceInstanceConditionOrphanMitigation ServiceInstanceConditionType = "OrphanMitigation"
)

// ServiceInstanceOperation represents a type of operation the controller can
// be performing for a service instance in the OSB API.
type ServiceInstanceOperation string

const (
	// ServiceInstanceOperationProvision indicates that the ServiceInstance is
	// being Provisioned.
	ServiceInstanceOperationProvision ServiceInstanceOperation = "Provision"
	// ServiceInstanceOperationUpdate indicates that the ServiceInstance is
	// being Updated.
	ServiceInstanceOperationUpdate ServiceInstanceOperation = "Update"
	// ServiceInstanceOperationDeprovision indicates that the ServiceInstance is
	// being Deprovisioned.
	ServiceInstanceOperationDeprovision ServiceInstanceOperation = "Deprovision"
)

// ServiceInstancePropertiesState is the state of a ServiceInstance that
// the ClusterServiceBroker knows about.
type ServiceInstancePropertiesState struct {
	// ClusterServicePlanExternalName is the name of the plan that the
	// broker knows this ServiceInstance to be on. This is the human
	// readable plan name from the OSB API.
	ClusterServicePlanExternalName string `json:"clusterServicePlanExternalName" protobuf:"bytes,1,opt,name=clusterServicePlanExternalName"`

	// ClusterServicePlanExternalID is the external ID of the plan that the
	// broker knows this ServiceInstance to be on.
	ClusterServicePlanExternalID string `json:"clusterServicePlanExternalID" protobuf:"bytes,2,opt,name=clusterServicePlanExternalID"`

	// Parameters is a blob of the parameters and their values that the broker
	// knows about for this ServiceInstance.  If a parameter was sourced from
	// a secret, its value will be "<redacted>" in this blob.
	Parameters *runtime.RawExtension `json:"parameters,omitempty" protobuf:"bytes,3,opt,name=parameters"`

	// ParametersChecksum is the checksum of the parameters that were sent.
	ParametersChecksum string `json:"parameterChecksum,omitempty" protobuf:"bytes,4,opt,name=parameterChecksum"` //TODO(mkibbe): Parameter(s)Checksum?

	// UserInfo is information about the user that made the request.
	UserInfo *UserInfo `json:"userInfo,omitempty" protobuf:"bytes,5,opt,name=userInfo"`
}

// ServiceInstanceDeprovisionStatus is the status of deprovisioning a
// ServiceInstance
type ServiceInstanceDeprovisionStatus string

const (
	// ServiceInstanceDeprovisionStatusNotRequired indicates that a provision
	// request has not been sent for the ServiceInstance, so no deprovision
	// request needs to be made.
	ServiceInstanceDeprovisionStatusNotRequired ServiceInstanceDeprovisionStatus = "NotRequired"
	// ServiceInstanceDeprovisionStatusRequired indicates that a provision
	// request has been sent for the ServiceInstance. A deprovision request
	// must be made before deleting the ServiceInstance.
	ServiceInstanceDeprovisionStatusRequired ServiceInstanceDeprovisionStatus = "Required"
	// ServiceInstanceDeprovisionStatusSucceeded indicates that a deprovision
	// request has been sent for the ServiceInstance, and the request was
	// successful.
	ServiceInstanceDeprovisionStatusSucceeded ServiceInstanceDeprovisionStatus = "Succeeded"
	// ServiceInstanceDeprovisionStatusFailed indicates that deprovision
	// requests have been sent for the ServiceInstance but they failed. The
	// controller has given up on sending more deprovision requests.
	ServiceInstanceDeprovisionStatusFailed ServiceInstanceDeprovisionStatus = "Failed"
)

// ServiceInstanceProvisionStatus is the status of provisioning a
// ServiceInstance
type ServiceInstanceProvisionStatus string

const (
	// ServiceInstanceProvisionStatusProvisioned indicates that the instance
	// was provisioned.
	ServiceInstanceProvisionStatusProvisioned ServiceInstanceProvisionStatus = "Provisioned"
	// ServiceInstanceProvisionStatusNotProvisioned indicates that the instance
	// was not ever provisioned or was deprovisioned.
	ServiceInstanceProvisionStatusNotProvisioned ServiceInstanceProvisionStatus = "NotProvisioned"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceBindingList is a list of ServiceBindings.
type ServiceBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ServiceBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceBinding represents a "used by" relationship between an application and an
// ServiceInstance.
// +k8s:openapi-gen=x-kubernetes-print-columns:custom-columns=NAME:.metadata.name,INSTANCE:.spec.instanceRef.name,SECRET:.spec.secretName
type ServiceBinding struct {
	metav1.TypeMeta `json:",inline"`

	// The name of this resource in etcd is in ObjectMeta.Name.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec represents the desired state of a ServiceBinding.
	// +optional
	Spec ServiceBindingSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Status represents the current status of a ServiceBinding.
	// +optional
	Status ServiceBindingStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ServiceBindingSpec represents the desired state of a
// ServiceBinding.
//
// The spec field cannot be changed after a ServiceBinding is
// created.  Changes submitted to the spec field will be ignored.
type ServiceBindingSpec struct {
	// ServiceInstanceRef is the reference to the Instance this ServiceBinding is to.
	//
	// Immutable.
	ServiceInstanceRef LocalObjectReference `json:"instanceRef" protobuf:"bytes,1,opt,name=instanceRef"`

	// Parameters is a set of the parameters to be passed to the underlying
	// broker. The inline YAML/JSON payload to be translated into equivalent
	// JSON object. If a top-level parameter name exists in multiples sources
	// among `Parameters` and `ParametersFrom` fields, it is considered to be
	// a user error in the specification.
	//
	// The Parameters field is NOT secret or secured in any way and should
	// NEVER be used to hold sensitive information. To set parameters that
	// contain secret information, you should ALWAYS store that information
	// in a Secret and use the ParametersFrom field.
	//
	// +optional
	Parameters *runtime.RawExtension `json:"parameters,omitempty" protobuf:"bytes,2,opt,name=parameters"`

	// List of sources to populate parameters.
	// If a top-level parameter name exists in multiples sources among
	// `Parameters` and `ParametersFrom` fields, it is
	// considered to be a user error in the specification.
	// +optional
	ParametersFrom []ParametersFromSource `json:"parametersFrom,omitempty" protobuf:"bytes,3,rep,name=parametersFrom"`

	// SecretName is the name of the secret to create in the ServiceBinding's
	// namespace that will hold the credentials associated with the ServiceBinding.
	SecretName string `json:"secretName,omitempty" protobuf:"bytes,4,opt,name=secretName"`

	// List of transformations that should be applied to the credentials
	// associated with the ServiceBinding before they are inserted into the Secret.
	SecretTransform []SecretTransform `json:"secretTransform,omitempty" protobuf:"bytes,5,rep,name=secretTransform"`

	// ExternalID is the identity of this object for use with the OSB API.
	//
	// Immutable.
	// +optional
	ExternalID string `json:"externalID" protobuf:"bytes,6,opt,name=externalID"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// UserInfo contains information about the user that last modified this
	// ServiceBinding. This field is set by the API server and not
	// settable by the end-user. User-provided values for this field are not saved.
	// +optional
	UserInfo *UserInfo `json:"userInfo,omitempty" protobuf:"bytes,7,opt,name=userInfo"`
}

// ServiceBindingStatus represents the current status of a ServiceBinding.
type ServiceBindingStatus struct {
	Conditions []ServiceBindingCondition `json:"conditions" protobuf:"bytes,1,rep,name=conditions"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// AsyncOpInProgress is set to true if there is an ongoing async operation
	// against this ServiceBinding in progress.
	AsyncOpInProgress bool `json:"asyncOpInProgress" protobuf:"varint,2,opt,name=asyncOpInProgress"`

	// Currently, this field is ALPHA: it may change or disappear at any time
	// and its data will not be migrated.
	//
	// LastOperation is the string that the broker may have returned when
	// an async operation started, it should be sent back to the broker
	// on poll requests as a query param.
	LastOperation *string `json:"lastOperation,omitempty" protobuf:"bytes,3,opt,name=lastOperation"`

	// CurrentOperation is the operation the Controller is currently performing
	// on the ServiceBinding.
	CurrentOperation ServiceBindingOperation `json:"currentOperation,omitempty" protobuf:"bytes,4,opt,name=currentOperation"`

	// ReconciledGeneration is the 'Generation' of the
	// ServiceBindingSpec that was last processed by the controller.
	// The reconciled generation is updated even if the controller failed to
	// process the spec.
	ReconciledGeneration int64 `json:"reconciledGeneration" protobuf:"varint,5,opt,name=reconciledGeneration"`

	// OperationStartTime is the time at which the current operation began.
	OperationStartTime *metav1.Time `json:"operationStartTime,omitempty" protobuf:"bytes,6,opt,name=operationStartTime"`

	// InProgressProperties is the properties state of the
	// ServiceBinding when a Bind is in progress. If the current
	// operation is an Unbind, this will be nil.
	InProgressProperties *ServiceBindingPropertiesState `json:"inProgressProperties,omitempty" protobuf:"bytes,7,opt,name=inProgressProperties"`

	// ExternalProperties is the properties state of the
	// ServiceBinding which the broker knows about.
	ExternalProperties *ServiceBindingPropertiesState `json:"externalProperties,omitempty" protobuf:"bytes,8,opt,name=externalProperties"`

	// OrphanMitigationInProgress is a flag that represents whether orphan
	// mitigation is in progress.
	OrphanMitigationInProgress bool `json:"orphanMitigationInProgress" protobuf:"varint,9,opt,name=orphanMitigationInProgress"`

	// UnbindStatus describes what has been done to unbind the ServiceBinding.
	UnbindStatus ServiceBindingUnbindStatus `json:"unbindStatus" protobuf:"bytes,10,opt,name=unbindStatus"`
}

// ServiceBindingCondition condition information for a ServiceBinding.
type ServiceBindingCondition struct {
	// Type of the condition, currently ('Ready').
	Type ServiceBindingConditionType `json:"type" protobuf:"bytes,1,opt,name=type"`

	// Status of the condition, one of ('True', 'False', 'Unknown').
	Status ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,3,opt,name=lastTransitionTime"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	Reason string `json:"reason" protobuf:"bytes,4,opt,name=reason"`

	// Message is a human readable description of the details of the last
	// transition, complementing reason.
	Message string `json:"message" protobuf:"bytes,5,opt,name=message"`
}

// ServiceBindingConditionType represents a ServiceBindingCondition value.
type ServiceBindingConditionType string

const (
	// ServiceBindingConditionReady represents a binding condition is in ready state.
	ServiceBindingConditionReady ServiceBindingConditionType = "Ready"

	// ServiceBindingConditionFailed represents a ServiceBindingCondition that has failed
	// completely and should not be retried.
	ServiceBindingConditionFailed ServiceBindingConditionType = "Failed"
)

// ServiceBindingOperation represents a type of operation
// the controller can be performing for a binding in the OSB API.
type ServiceBindingOperation string

const (
	// ServiceBindingOperationBind indicates that the
	// ServiceBinding is being bound.
	ServiceBindingOperationBind ServiceBindingOperation = "Bind"
	// ServiceBindingOperationUnbind indicates that the
	// ServiceBinding is being unbound.
	ServiceBindingOperationUnbind ServiceBindingOperation = "Unbind"
)

// ServiceBindingUnbindStatus is the status of unbinding a Binding
type ServiceBindingUnbindStatus string

const (
	// ServiceBindingUnbindStatusNotRequired indicates that a binding request
	// has not been sent for the ServiceBinding, so no unbinding request
	// needs to be made.
	ServiceBindingUnbindStatusNotRequired ServiceBindingUnbindStatus = "NotRequired"
	// ServiceBindingUnbindStatusRequired indicates that a binding request has
	// been sent for the ServiceBinding. An unbind request must be made before
	// deleting the ServiceBinding.
	ServiceBindingUnbindStatusRequired ServiceBindingUnbindStatus = "Required"
	// ServiceBindingUnbindStatusSucceeded indicates that a unbind request has
	// been sent for the ServiceBinding, and the request was successful.
	ServiceBindingUnbindStatusSucceeded ServiceBindingUnbindStatus = "Succeeded"
	// ServiceBindingUnbindStatusFailed indicates that unbind requests have
	// been sent for the ServiceBinding but they failed. The controller has
	// given up on sending more unbind requests.
	ServiceBindingUnbindStatusFailed ServiceBindingUnbindStatus = "Failed"
)

// These are external finalizer values to service catalog, must be qualified name.
const (
	FinalizerServiceCatalog string = "kubernetes-incubator/service-catalog"
)

// ServiceBindingPropertiesState is the state of a
// ServiceBinding that the ClusterServiceBroker knows about.
type ServiceBindingPropertiesState struct {
	// Parameters is a blob of the parameters and their values that the broker
	// knows about for this ServiceBinding.  If a parameter was
	// sourced from a secret, its value will be "<redacted>" in this blob.
	Parameters *runtime.RawExtension `json:"parameters,omitempty" protobuf:"bytes,1,opt,name=parameters"`

	// ParametersChecksum is the checksum of the parameters that were sent.
	ParametersChecksum string `json:"parameterChecksum,omitempty" protobuf:"bytes,2,opt,name=parameterChecksum"`

	// UserInfo is information about the user that made the request.
	UserInfo *UserInfo `json:"userInfo,omitempty" protobuf:"bytes,3,opt,name=userInfo"`
}

// ParametersFromSource represents the source of a set of Parameters
type ParametersFromSource struct {
	// The Secret key to select from.
	// The value must be a JSON object.
	//+optional
	SecretKeyRef *SecretKeyReference `json:"secretKeyRef,omitempty" protobuf:"bytes,1,opt,name=secretKeyRef"`
}

// SecretKeyReference references a key of a Secret.
type SecretKeyReference struct {
	// The name of the secret in the pod's namespace to select from.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// The key of the secret to select from.  Must be a valid secret key.
	Key string `json:"key" protobuf:"bytes,2,opt,name=key"`
}

// ObjectReference contains enough information to let you locate the
// referenced object.
type ObjectReference struct {
	// Namespace of the referent.
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,1,opt,name=namespace"`
	// Name of the referent.
	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
}

// LocalObjectReference contains enough information to let you locate the
// referenced object inside the same namespace.
type LocalObjectReference struct {
	// Name of the referent.
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// ClusterObjectReference contains enough information to let you locate the
// cluster-scoped referenced object.
type ClusterObjectReference struct {
	// Name of the referent.
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// SecretTransform is a single transformation that is applied to the
// credentials returned from the broker before they are inserted into
// the Secret associated with the ServiceBinding.
// Because different brokers providing the same type of service may
// each return a different credentials structure, users can specify
// the transformations that should be applied to the Secret to adapt
// its entries to whatever the service consumer expects.
// For example, the credentials returned by the broker may include the
// key "USERNAME", but the consumer requires the username to be
// exposed under the key "DB_USER" instead. To have the Service
// Catalog transform the Secret, the following SecretTransform must
// be specified in ServiceBinding.spec.secretTransform:
// - {"renameKey": {"from": "USERNAME", "to": "DB_USER"}}
type SecretTransform struct {
	RenameKey *RenameKeyTransform `json:"renameKey,omitempty" protobuf:"bytes,1,opt,name=renameKey"`
}

// RenameKeyTransform specifies that one of the credentials keys returned
// from the broker should be renamed and stored under a different key
// in the Secret.
// For example, given the following credentials entry:
//     "USERNAME": "johndoe"
// and the following RenameKeyTransform:
//     {"from": "USERNAME", "to": "DB_USER"}
// the following entry will appear in the Secret:
//     "DB_USER": "johndoe"
type RenameKeyTransform struct {
	From string `json:"from" protobuf:"bytes,1,opt,name=from"`
	To   string `json:"to" protobuf:"bytes,2,opt,name=to"`
}
