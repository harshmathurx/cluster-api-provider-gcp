/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import (
	"fmt"
	"reflect"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var gcpclustertemplatelog = logf.Log.WithName("gcpclustertemplate-resource")

func (r *GCPClusterTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update,path=/mutate-infrastructure-cluster-x-k8s-io-v1alpha4-gcpclustertemplate,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=gcpclustertemplates,versions=v1alpha4,name=default.gcpclustertemplate.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

var _ webhook.Defaulter = &GCPClusterTemplate{}

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (r *GCPClusterTemplate) Default() {
	gcpclustertemplatelog.Info("default", "name", r.Name)
}

//+kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1alpha4-gcpclustertemplate,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=gcpclustertemplates,versions=v1alpha4,name=validation.gcpclustertemplate.infrastructure.cluster.x-k8s.io,sideEffects=None,admissionReviewVersions=v1beta1

var _ webhook.Validator = &GCPClusterTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPClusterTemplate) ValidateCreate() error {
	gcpclustertemplatelog.Info("validate create", "name", r.Name)

	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPClusterTemplate) ValidateUpdate(oldRaw runtime.Object) error {
	old, ok := oldRaw.(*GCPClusterTemplate)
	if !ok {
		return apierrors.NewBadRequest(fmt.Sprintf("expected an GCPClusterTemplate but got a %T", oldRaw))
	}

	if !reflect.DeepEqual(r.Spec, old.Spec) {
		return apierrors.NewBadRequest("GCPClusterTemplate.Spec is immutable")
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPClusterTemplate) ValidateDelete() error {
	gcpclustertemplatelog.Info("validate delete", "name", r.Name)
	return nil
}
