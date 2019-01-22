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
)

// 机器监控和实例监控都会产生Event，共用这么一个struct
type Event struct {
	Id          string            `json:"id"`
	Strategy    *Strategy         `json:"strategy"`
	Status      string            `json:"status"` // OK or PROBLEM
	Endpoint    string            `json:"endpoint"`
	LeftValue   float64           `json:"leftValue"`
	CurrentStep int               `json:"currentStep"`
	EventTime   int64             `json:"eventTime"`
	PushedTags  map[string]string `json:"pushedTags"`
}

func (this *Event) FormattedTime() string {
	return utils.UnixTsFormat(this.EventTime)
}

func (this *Event) String() string {
	return fmt.Sprintf(
		"<Endpoint:%s, Status:%s, Strategy:%v, LeftValue:%s, CurrentStep:%d, PushedTags:%v, TS:%s>",
		this.Endpoint,
		this.Status,
		this.Strategy,
		utils.ReadableFloat(this.LeftValue),
		this.CurrentStep,
		this.PushedTags,
		this.FormattedTime(),
	)
}

func (this *Event) StrategyId() string {
	if this.Strategy != nil {
		return this.Strategy.Id
	}

	return ""
}

func (this *Event) Priority() int {
	return this.Strategy.Priority
}

func (this *Event) Ttl() int {
	return this.Strategy.Ttl
}

func (this *Event) Note() string {
	return this.Strategy.Note
}

func (this *Event) Metric() string {
	return this.Strategy.Metric
}

func (this *Event) RightValue() float64 {
	return this.Strategy.RightValue
}

func (this *Event) Operator() string {
	return this.Strategy.Operator
}

func (this *Event) Func() string {
	return this.Strategy.Func
}

func (this *Event) MaxStep() int {
	return this.Strategy.MaxStep
}

func (this *Event) Counter() string {
	return fmt.Sprintf("%s/%s %s", this.Endpoint, this.Metric(), utils.SortedTags(this.PushedTags))
}
