/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - API 网关(BlueKing - APIGateway) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *     http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"micro-gateway/api/v1beta1"
	gatewayv1beta1 "micro-gateway/api/v1beta1"
	"micro-gateway/pkg/registry"
)

// BkGatewayPluginMetadataReconciler reconciles a BkGatewayPluginMetadata object
type BkGatewayPluginMetadataReconciler struct {
	adapater *registryAdapter
	Handler  registry.KubeEventHandler
	client.Client
	Scheme *runtime.Scheme
}

//nolint:lll
//+kubebuilder:rbac:groups=gateway.bk.tencent.com,resources=bkgatewaypluginmetadatas,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the BkGatewayPluginMetadata object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *BkGatewayPluginMetadataReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.V(1).Info("bk gateway stage trigger", "obj", req)
	r.adapater.Reconcile(ctx, req, &v1beta1.BkGatewayPluginMetadata{}, logger)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BkGatewayPluginMetadataReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.adapater = &registryAdapter{
		resMetaMap: make(map[types.NamespacedName]*registry.ResourceMetadata),
		handler:    r.Handler,
		Client:     r.Client,
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&gatewayv1beta1.BkGatewayPluginMetadata{}).
		Complete(r)
}