/*
 * Copyright (c) 2020 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package runtime
package runtime

import (
	"github.com/baidu/openless/pkg/funclet/runtime/api"
	"github.com/baidu/openless/pkg/util/logs"
)

type RuntimeManagerInterface interface {
	api.ContainerManager
	ResourceManager
}

type Manager struct {
	ResourceManager
	api.ContainerManager
}

type RuntimeManagerParameters struct {
	ContainerNum int
	RuntimeCmd   string
	Option       *ResourceOption
	Logger       *logs.Logger
}

func NewRuntimeManager(p *RuntimeManagerParameters) (rm RuntimeManagerInterface, err error) {
	runtimeManager := Manager{}

	resourceCtrl, err := NewResourceManager(p.Option, p.ContainerNum)
	if err != nil {
		return nil, err
	}

	containerCtrl, err := NewContainerRuntime(api.RuntimeTypeRunc, p.RuntimeCmd, p.Logger)
	if err != nil {
		return nil, err
	}

	runtimeManager.ResourceManager = resourceCtrl
	runtimeManager.ContainerManager = containerCtrl
	return &runtimeManager, nil
}
