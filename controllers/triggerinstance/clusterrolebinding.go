/*
Copyright 2022 The KubeVela Authors.

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

package triggerinstance

import (
	"context"

	standardv1alpha1 "github.com/kubevela/kube-trigger/api/v1alpha1"
	"github.com/kubevela/kube-trigger/controllers/template"
	"github.com/kubevela/kube-trigger/controllers/utils"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func (r *Reconciler) createClusterRoleBinding(
	ctx context.Context,
	ki *standardv1alpha1.TriggerInstance,
	update bool,
) error {
	crb := template.GetClusterRoleBinding()

	// TODO(charlie0129): allow user to set custom privileges instead of cluster-admin.

	crb.Name = ki.Name
	crb.Namespace = ki.Namespace
	// It must have one subject.
	crb.Subjects[0].Name = ki.Name
	crb.Subjects[0].Namespace = ki.Namespace
	utils.SetOwnerReference(crb, ki)

	var err error
	if update {
		logger.Infof("updating ClusterRoleBinding: %s", types.NamespacedName{
			Namespace: crb.Namespace,
			Name:      crb.Name,
		}.String())
		err = r.Update(ctx, crb)
	} else {
		logger.Infof("creating new ClusterRoleBinding: %s", types.NamespacedName{
			Namespace: crb.Namespace,
			Name:      crb.Name,
		}.String())
		err = r.Create(ctx, crb)
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *Reconciler) reconcileClusterRoleBinding(ctx context.Context, ki *standardv1alpha1.TriggerInstance) error {
	crb := rbacv1.ClusterRoleBinding{}
	err := r.Get(ctx, utils.GetNamespacedName(ki), &crb)

	if err == nil {
		return r.createClusterRoleBinding(ctx, ki, true)
	}

	if errors.IsNotFound(err) {
		return r.createClusterRoleBinding(ctx, ki, false)
	}

	return err
}
