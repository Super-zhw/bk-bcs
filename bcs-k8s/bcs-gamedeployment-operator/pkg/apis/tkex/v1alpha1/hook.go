/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package v1alpha1

import (
	hookv1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
)

func (g *GameDeployment) GetPreDeleteHook() *hookv1alpha1.HookStep {
	return g.Spec.PreDeleteUpdateStrategy.Hook
}

func (g *GameDeploymentStatus) GetPreDeleteHookConditions() []hookv1alpha1.PreDeleteHookCondition {
	return g.PreDeleteHookConditions
}

func (g *GameDeploymentStatus) SetPreDeleteHookConditions(newConditions []hookv1alpha1.PreDeleteHookCondition) {
	g.PreDeleteHookConditions = newConditions
}

func (g *GameDeployment) GetPreInplaceHook() *hookv1alpha1.HookStep {
	return g.Spec.PreInplaceUpdateStrategy.Hook
}

func (g *GameDeploymentStatus) GetPreInplaceHookConditions() []hookv1alpha1.PreInplaceHookCondition {
	return g.PreInplaceHookConditions
}

func (g *GameDeploymentStatus) SetPreInplaceHookConditions(newConditions []hookv1alpha1.PreInplaceHookCondition) {
	g.PreInplaceHookConditions = newConditions
}