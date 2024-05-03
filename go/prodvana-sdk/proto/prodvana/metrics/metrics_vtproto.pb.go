// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.6.0
// source: prodvana/metrics/metrics.proto

package metrics

import (
	durationpb1 "github.com/planetscale/vtprotobuf/types/known/durationpb"
	timestamppb1 "github.com/planetscale/vtprotobuf/types/known/timestamppb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *DeploymentMetrics_Deltas) CloneVT() *DeploymentMetrics_Deltas {
	if m == nil {
		return (*DeploymentMetrics_Deltas)(nil)
	}
	r := new(DeploymentMetrics_Deltas)
	r.TotalDeployments = m.TotalDeployments
	r.SuccessfulDeployments = m.SuccessfulDeployments
	r.SuccessfulDeploymentsPerDay = m.SuccessfulDeploymentsPerDay
	r.FailedDeployments = m.FailedDeployments
	r.MedianSuccessfulDeploymentDuration = m.MedianSuccessfulDeploymentDuration
	r.DeploymentFailureRate = m.DeploymentFailureRate
	r.TotalRollbacks = m.TotalRollbacks
	r.SuccessfulRollbacks = m.SuccessfulRollbacks
	r.MedianSuccessfulRollbackDuration = m.MedianSuccessfulRollbackDuration
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeploymentMetrics_Deltas) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeploymentMetrics_Values) CloneVT() *DeploymentMetrics_Values {
	if m == nil {
		return (*DeploymentMetrics_Values)(nil)
	}
	r := new(DeploymentMetrics_Values)
	r.TotalDeployments = m.TotalDeployments
	r.SuccessfulDeployments = m.SuccessfulDeployments
	r.SuccessfulDeploymentsPerDay = m.SuccessfulDeploymentsPerDay
	r.FailedDeployments = m.FailedDeployments
	r.MedianSuccessfulDeploymentDuration = (*durationpb.Duration)((*durationpb1.Duration)(m.MedianSuccessfulDeploymentDuration).CloneVT())
	r.DeploymentFailureRate = m.DeploymentFailureRate
	r.TotalRollbacks = m.TotalRollbacks
	r.SuccessfulRollbacks = m.SuccessfulRollbacks
	r.MedianSuccessfulRollbackDuration = (*durationpb.Duration)((*durationpb1.Duration)(m.MedianSuccessfulRollbackDuration).CloneVT())
	r.LastSuccessfulDeploymentCompletionTime = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.LastSuccessfulDeploymentCompletionTime).CloneVT())
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeploymentMetrics_Values) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeploymentMetrics_TimeseriesMetrics) CloneVT() *DeploymentMetrics_TimeseriesMetrics {
	if m == nil {
		return (*DeploymentMetrics_TimeseriesMetrics)(nil)
	}
	r := new(DeploymentMetrics_TimeseriesMetrics)
	r.Timestamp = (*timestamppb.Timestamp)((*timestamppb1.Timestamp)(m.Timestamp).CloneVT())
	r.SuccessfulDeployments = m.SuccessfulDeployments
	r.FailedDeployments = m.FailedDeployments
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeploymentMetrics_TimeseriesMetrics) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *DeploymentMetrics) CloneVT() *DeploymentMetrics {
	if m == nil {
		return (*DeploymentMetrics)(nil)
	}
	r := new(DeploymentMetrics)
	r.Values = m.Values.CloneVT()
	r.Deltas = m.Deltas.CloneVT()
	if rhs := m.Timeseries; rhs != nil {
		tmpContainer := make([]*DeploymentMetrics_TimeseriesMetrics, len(rhs))
		for k, v := range rhs {
			tmpContainer[k] = v.CloneVT()
		}
		r.Timeseries = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *DeploymentMetrics) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *Cost) CloneVT() *Cost {
	if m == nil {
		return (*Cost)(nil)
	}
	r := new(Cost)
	r.Configured = m.Configured
	r.Cost = m.Cost
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *Cost) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *DeploymentMetrics_Deltas) StableEqualVT(that *DeploymentMetrics_Deltas) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TotalDeployments != that.TotalDeployments {
		return false
	}
	if this.SuccessfulDeployments != that.SuccessfulDeployments {
		return false
	}
	if this.SuccessfulDeploymentsPerDay != that.SuccessfulDeploymentsPerDay {
		return false
	}
	if this.FailedDeployments != that.FailedDeployments {
		return false
	}
	if this.MedianSuccessfulDeploymentDuration != that.MedianSuccessfulDeploymentDuration {
		return false
	}
	if this.DeploymentFailureRate != that.DeploymentFailureRate {
		return false
	}
	if this.TotalRollbacks != that.TotalRollbacks {
		return false
	}
	if this.SuccessfulRollbacks != that.SuccessfulRollbacks {
		return false
	}
	if this.MedianSuccessfulRollbackDuration != that.MedianSuccessfulRollbackDuration {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentMetrics_Deltas) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentMetrics_Deltas)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *DeploymentMetrics_Values) StableEqualVT(that *DeploymentMetrics_Values) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TotalDeployments != that.TotalDeployments {
		return false
	}
	if this.SuccessfulDeployments != that.SuccessfulDeployments {
		return false
	}
	if this.SuccessfulDeploymentsPerDay != that.SuccessfulDeploymentsPerDay {
		return false
	}
	if this.FailedDeployments != that.FailedDeployments {
		return false
	}
	if !(*durationpb1.Duration)(this.MedianSuccessfulDeploymentDuration).StableEqualVT((*durationpb1.Duration)(that.MedianSuccessfulDeploymentDuration)) {
		return false
	}
	if this.DeploymentFailureRate != that.DeploymentFailureRate {
		return false
	}
	if this.TotalRollbacks != that.TotalRollbacks {
		return false
	}
	if this.SuccessfulRollbacks != that.SuccessfulRollbacks {
		return false
	}
	if !(*durationpb1.Duration)(this.MedianSuccessfulRollbackDuration).StableEqualVT((*durationpb1.Duration)(that.MedianSuccessfulRollbackDuration)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastSuccessfulDeploymentCompletionTime).StableEqualVT((*timestamppb1.Timestamp)(that.LastSuccessfulDeploymentCompletionTime)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentMetrics_Values) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentMetrics_Values)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *DeploymentMetrics_TimeseriesMetrics) StableEqualVT(that *DeploymentMetrics_TimeseriesMetrics) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.Timestamp).StableEqualVT((*timestamppb1.Timestamp)(that.Timestamp)) {
		return false
	}
	if this.SuccessfulDeployments != that.SuccessfulDeployments {
		return false
	}
	if this.FailedDeployments != that.FailedDeployments {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentMetrics_TimeseriesMetrics) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentMetrics_TimeseriesMetrics)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *DeploymentMetrics) StableEqualVT(that *DeploymentMetrics) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Values.StableEqualVT(that.Values) {
		return false
	}
	if len(this.Timeseries) != len(that.Timeseries) {
		return false
	}
	for i, vx := range this.Timeseries {
		vy := that.Timeseries[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &DeploymentMetrics_TimeseriesMetrics{}
			}
			if q == nil {
				q = &DeploymentMetrics_TimeseriesMetrics{}
			}
			if !p.StableEqualVT(q) {
				return false
			}
		}
	}
	if !this.Deltas.StableEqualVT(that.Deltas) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentMetrics) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentMetrics)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *DeploymentMetrics_Deltas) EqualVT(that *DeploymentMetrics_Deltas) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TotalDeployments != that.TotalDeployments {
		return false
	}
	if this.SuccessfulDeployments != that.SuccessfulDeployments {
		return false
	}
	if this.SuccessfulDeploymentsPerDay != that.SuccessfulDeploymentsPerDay {
		return false
	}
	if this.FailedDeployments != that.FailedDeployments {
		return false
	}
	if this.MedianSuccessfulDeploymentDuration != that.MedianSuccessfulDeploymentDuration {
		return false
	}
	if this.DeploymentFailureRate != that.DeploymentFailureRate {
		return false
	}
	if this.TotalRollbacks != that.TotalRollbacks {
		return false
	}
	if this.SuccessfulRollbacks != that.SuccessfulRollbacks {
		return false
	}
	if this.MedianSuccessfulRollbackDuration != that.MedianSuccessfulRollbackDuration {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentMetrics_Deltas) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentMetrics_Deltas)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeploymentMetrics_Values) EqualVT(that *DeploymentMetrics_Values) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.TotalDeployments != that.TotalDeployments {
		return false
	}
	if this.SuccessfulDeployments != that.SuccessfulDeployments {
		return false
	}
	if this.SuccessfulDeploymentsPerDay != that.SuccessfulDeploymentsPerDay {
		return false
	}
	if this.FailedDeployments != that.FailedDeployments {
		return false
	}
	if !(*durationpb1.Duration)(this.MedianSuccessfulDeploymentDuration).EqualVT((*durationpb1.Duration)(that.MedianSuccessfulDeploymentDuration)) {
		return false
	}
	if this.DeploymentFailureRate != that.DeploymentFailureRate {
		return false
	}
	if this.TotalRollbacks != that.TotalRollbacks {
		return false
	}
	if this.SuccessfulRollbacks != that.SuccessfulRollbacks {
		return false
	}
	if !(*durationpb1.Duration)(this.MedianSuccessfulRollbackDuration).EqualVT((*durationpb1.Duration)(that.MedianSuccessfulRollbackDuration)) {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.LastSuccessfulDeploymentCompletionTime).EqualVT((*timestamppb1.Timestamp)(that.LastSuccessfulDeploymentCompletionTime)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentMetrics_Values) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentMetrics_Values)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeploymentMetrics_TimeseriesMetrics) EqualVT(that *DeploymentMetrics_TimeseriesMetrics) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !(*timestamppb1.Timestamp)(this.Timestamp).EqualVT((*timestamppb1.Timestamp)(that.Timestamp)) {
		return false
	}
	if this.SuccessfulDeployments != that.SuccessfulDeployments {
		return false
	}
	if this.FailedDeployments != that.FailedDeployments {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentMetrics_TimeseriesMetrics) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentMetrics_TimeseriesMetrics)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *DeploymentMetrics) EqualVT(that *DeploymentMetrics) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if !this.Values.EqualVT(that.Values) {
		return false
	}
	if len(this.Timeseries) != len(that.Timeseries) {
		return false
	}
	for i, vx := range this.Timeseries {
		vy := that.Timeseries[i]
		if p, q := vx, vy; p != q {
			if p == nil {
				p = &DeploymentMetrics_TimeseriesMetrics{}
			}
			if q == nil {
				q = &DeploymentMetrics_TimeseriesMetrics{}
			}
			if !p.EqualVT(q) {
				return false
			}
		}
	}
	if !this.Deltas.EqualVT(that.Deltas) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *DeploymentMetrics) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*DeploymentMetrics)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *Cost) StableEqualVT(that *Cost) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Configured != that.Configured {
		return false
	}
	if this.Cost != that.Cost {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Cost) StableEqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Cost)
	if !ok {
		return false
	}
	return this.StableEqualVT(that)
}
func (this *Cost) EqualVT(that *Cost) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Configured != that.Configured {
		return false
	}
	if this.Cost != that.Cost {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *Cost) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*Cost)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}