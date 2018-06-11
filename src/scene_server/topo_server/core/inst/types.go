/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package inst

import (
	frtypes "configcenter/src/common/mapstr"

	"configcenter/src/scene_server/topo_server/core/model"
	"configcenter/src/scene_server/topo_server/core/types"
)

// Inst the inst interface
type Inst interface {
	model.Operation
	GetObject() model.Object

	GetInstID() int
	GetInstName() string

	SetValue(key string, value interface{}) error

	SetValues(values frtypes.MapStr) error

	GetValues() (frtypes.MapStr, error)

	ToMapStr() (frtypes.MapStr, error)
}

// Factory used to all inst
type Factory interface {
	CreateInst(params types.LogicParams, obj model.Object) Inst
}
