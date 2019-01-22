package model

import (
	"fmt"
)

type Aggregator struct {
	// TODO: for specified cluster or endpoint?
	Id         string `json:"_id"`
	Endpoint   string `json:"endpoint"`
	Metric     string `json:"metric"`
	Tags       string `json:"tags"`
	DsType     string `json:"dstype"`
	Step       int32  `json:"step"`
	Expression string `json:"expression"`
}

func (this *Aggregator) String() string {
	return fmt.Sprintf(
		"<id:%s,endpoint:%s,metric:%s,tags:%s,dstype:%s,step:%d,expression:%s>",
		this.Id,
		this.Endpoint,
		this.Metric,
		this.Tags,
		this.DsType,
		this.Step,
		this.Expression,
	)
}

type AggregatorsResponse struct {
	Aggregators map[string]*Aggregator `json:"aggregators"`
}
