/**
 * Tencent is pleased to support the open source community by making polaris-go available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package unirate

import (
	"fmt"
	"time"

	"github.com/polarismesh/polaris-go/pkg/model"
)

const (
	defaultMaxQueuingTime = 1 * time.Second
)

// 匀速排队限流器配置
type Config struct {
	MaxQueuingTime *time.Duration `yaml:"maxQueuingTime" json:"maxQueuingTime"`
}

// 设置默认值
func (c *Config) SetDefault() {
	if nil == c.MaxQueuingTime {
		c.MaxQueuingTime = model.ToDurationPtr(defaultMaxQueuingTime)
	}
}

// 校验配置值
func (c *Config) Verify() error {
	if nil == c.MaxQueuingTime {
		return fmt.Errorf("MaxQueuingTime not configured")
	}
	if *c.MaxQueuingTime < 0 {
		return fmt.Errorf("Invalid maxQueuingTime: %v, it must greater than 0", *c.MaxQueuingTime)
	}
	return nil
}
