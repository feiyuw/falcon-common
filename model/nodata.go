// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"

	"dev.dingxiang-inc.com/aladdin/falcon-common/utils"
	ttime "github.com/toolkits/time"
)

type NodataItem struct {
	Counter string `json:"counter"`
	Ts      int64  `json:"ts"`
	FStatus string `json:"fstatus"`
	FTs     int64  `json:"fts"`
}

func (this *NodataItem) String() string {
	return fmt.Sprintf("{NodataItem counter:%s ts:%s fecthStatus:%s fetchTs:%s}",
		this.Counter, ttime.FormatTs(this.Ts), this.FStatus, ttime.FormatTs(this.FTs))
}

type NodataConfig struct {
	Name     string            `json:"name"`
	Endpoint string            `json:"endpoint"`
	Metric   string            `json:"metric"`
	Tags     map[string]string `json:"tags"`
	Type     string            `json:"type"`
	Step     int64             `json:"step"`
	Mock     float64           `json:"mock"`
}

func NewNodataConfig(name string, endpoint string, metric string, tags map[string]string, dstype string, step int64, mock float64) *NodataConfig {
	return &NodataConfig{name, endpoint, metric, tags, dstype, step, mock}
}

func (this *NodataConfig) String() string {
	return fmt.Sprintf("{NodataConfig name:%s, endpoint:%s, metric:%s, tags:%s, type:%s, step:%d, mock:%f}",
		this.Name, this.Endpoint, this.Metric, utils.SortedTags(this.Tags), this.Type, this.Step, this.Mock)
}

type NodatasResponse struct {
	NodataConfigs map[string]*NodataConfig `json:"nodataConfigs"`
}
