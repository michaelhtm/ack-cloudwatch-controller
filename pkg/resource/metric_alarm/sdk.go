// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package metric_alarm

import (
	"context"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithy "github.com/aws/smithy-go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/cloudwatch-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &svcsdk.Client{}
	_ = &svcapitypes.MetricAlarm{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
	_ = &aws.Config{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	input.AlarmNames = []string{*r.ko.Spec.Name}

	var resp *svcsdk.DescribeAlarmsOutput
	resp, err = rm.sdkapi.DescribeAlarms(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeAlarms", err)
	if err != nil {
		var awsErr smithy.APIError
		if errors.As(err, &awsErr) && awsErr.ErrorCode() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.MetricAlarms {
		if elem.ActionsEnabled != nil {
			ko.Spec.ActionsEnabled = elem.ActionsEnabled
		} else {
			ko.Spec.ActionsEnabled = nil
		}
		if elem.AlarmActions != nil {
			ko.Spec.AlarmActions = aws.StringSlice(elem.AlarmActions)
		} else {
			ko.Spec.AlarmActions = nil
		}
		if elem.AlarmDescription != nil {
			ko.Spec.AlarmDescription = elem.AlarmDescription
		} else {
			ko.Spec.AlarmDescription = nil
		}
		if elem.ComparisonOperator != "" {
			ko.Spec.ComparisonOperator = aws.String(string(elem.ComparisonOperator))
		} else {
			ko.Spec.ComparisonOperator = nil
		}
		if elem.DatapointsToAlarm != nil {
			datapointsToAlarmCopy := int64(*elem.DatapointsToAlarm)
			ko.Spec.DatapointsToAlarm = &datapointsToAlarmCopy
		} else {
			ko.Spec.DatapointsToAlarm = nil
		}
		if elem.Dimensions != nil {
			f8 := []*svcapitypes.Dimension{}
			for _, f8iter := range elem.Dimensions {
				f8elem := &svcapitypes.Dimension{}
				if f8iter.Name != nil {
					f8elem.Name = f8iter.Name
				}
				if f8iter.Value != nil {
					f8elem.Value = f8iter.Value
				}
				f8 = append(f8, f8elem)
			}
			ko.Spec.Dimensions = f8
		} else {
			ko.Spec.Dimensions = nil
		}
		if elem.EvaluateLowSampleCountPercentile != nil {
			ko.Spec.EvaluateLowSampleCountPercentile = elem.EvaluateLowSampleCountPercentile
		} else {
			ko.Spec.EvaluateLowSampleCountPercentile = nil
		}
		if elem.EvaluationPeriods != nil {
			evaluationPeriodsCopy := int64(*elem.EvaluationPeriods)
			ko.Spec.EvaluationPeriods = &evaluationPeriodsCopy
		} else {
			ko.Spec.EvaluationPeriods = nil
		}
		if elem.ExtendedStatistic != nil {
			ko.Spec.ExtendedStatistic = elem.ExtendedStatistic
		} else {
			ko.Spec.ExtendedStatistic = nil
		}
		if elem.InsufficientDataActions != nil {
			ko.Spec.InsufficientDataActions = aws.StringSlice(elem.InsufficientDataActions)
		} else {
			ko.Spec.InsufficientDataActions = nil
		}
		if elem.MetricName != nil {
			ko.Spec.MetricName = elem.MetricName
		} else {
			ko.Spec.MetricName = nil
		}
		if elem.Metrics != nil {
			f15 := []*svcapitypes.MetricDataQuery{}
			for _, f15iter := range elem.Metrics {
				f15elem := &svcapitypes.MetricDataQuery{}
				if f15iter.AccountId != nil {
					f15elem.AccountID = f15iter.AccountId
				}
				if f15iter.Expression != nil {
					f15elem.Expression = f15iter.Expression
				}
				if f15iter.Id != nil {
					f15elem.ID = f15iter.Id
				}
				if f15iter.Label != nil {
					f15elem.Label = f15iter.Label
				}
				if f15iter.MetricStat != nil {
					f15elemf4 := &svcapitypes.MetricStat{}
					if f15iter.MetricStat.Metric != nil {
						f15elemf4f0 := &svcapitypes.Metric{}
						if f15iter.MetricStat.Metric.Dimensions != nil {
							f15elemf4f0f0 := []*svcapitypes.Dimension{}
							for _, f15elemf4f0f0iter := range f15iter.MetricStat.Metric.Dimensions {
								f15elemf4f0f0elem := &svcapitypes.Dimension{}
								if f15elemf4f0f0iter.Name != nil {
									f15elemf4f0f0elem.Name = f15elemf4f0f0iter.Name
								}
								if f15elemf4f0f0iter.Value != nil {
									f15elemf4f0f0elem.Value = f15elemf4f0f0iter.Value
								}
								f15elemf4f0f0 = append(f15elemf4f0f0, f15elemf4f0f0elem)
							}
							f15elemf4f0.Dimensions = f15elemf4f0f0
						}
						if f15iter.MetricStat.Metric.MetricName != nil {
							f15elemf4f0.MetricName = f15iter.MetricStat.Metric.MetricName
						}
						if f15iter.MetricStat.Metric.Namespace != nil {
							f15elemf4f0.Namespace = f15iter.MetricStat.Metric.Namespace
						}
						f15elemf4.Metric = f15elemf4f0
					}
					if f15iter.MetricStat.Period != nil {
						periodCopy := int64(*f15iter.MetricStat.Period)
						f15elemf4.Period = &periodCopy
					}
					if f15iter.MetricStat.Stat != nil {
						f15elemf4.Stat = f15iter.MetricStat.Stat
					}
					if f15iter.MetricStat.Unit != "" {
						f15elemf4.Unit = aws.String(string(f15iter.MetricStat.Unit))
					}
					f15elem.MetricStat = f15elemf4
				}
				if f15iter.Period != nil {
					periodCopy := int64(*f15iter.Period)
					f15elem.Period = &periodCopy
				}
				if f15iter.ReturnData != nil {
					f15elem.ReturnData = f15iter.ReturnData
				}
				f15 = append(f15, f15elem)
			}
			ko.Spec.Metrics = f15
		} else {
			ko.Spec.Metrics = nil
		}
		if elem.Namespace != nil {
			ko.Spec.Namespace = elem.Namespace
		} else {
			ko.Spec.Namespace = nil
		}
		if elem.OKActions != nil {
			ko.Spec.OKActions = aws.StringSlice(elem.OKActions)
		} else {
			ko.Spec.OKActions = nil
		}
		if elem.Period != nil {
			periodCopy := int64(*elem.Period)
			ko.Spec.Period = &periodCopy
		} else {
			ko.Spec.Period = nil
		}
		if elem.Statistic != "" {
			ko.Spec.Statistic = aws.String(string(elem.Statistic))
		} else {
			ko.Spec.Statistic = nil
		}
		if elem.Threshold != nil {
			ko.Spec.Threshold = elem.Threshold
		} else {
			ko.Spec.Threshold = nil
		}
		if elem.ThresholdMetricId != nil {
			ko.Spec.ThresholdMetricID = elem.ThresholdMetricId
		} else {
			ko.Spec.ThresholdMetricID = nil
		}
		if elem.TreatMissingData != nil {
			ko.Spec.TreatMissingData = elem.TreatMissingData
		} else {
			ko.Spec.TreatMissingData = nil
		}
		if elem.Unit != "" {
			ko.Spec.Unit = aws.String(string(elem.Unit))
		} else {
			ko.Spec.Unit = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return false
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeAlarmsInput, error) {
	res := &svcsdk.DescribeAlarmsInput{}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.PutMetricAlarmOutput
	_ = resp
	resp, err = rm.sdkapi.PutMetricAlarm(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "PutMetricAlarm", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.PutMetricAlarmInput, error) {
	res := &svcsdk.PutMetricAlarmInput{}

	if r.ko.Spec.ActionsEnabled != nil {
		res.ActionsEnabled = r.ko.Spec.ActionsEnabled
	}
	if r.ko.Spec.AlarmActions != nil {
		res.AlarmActions = aws.ToStringSlice(r.ko.Spec.AlarmActions)
	}
	if r.ko.Spec.AlarmDescription != nil {
		res.AlarmDescription = r.ko.Spec.AlarmDescription
	}
	if r.ko.Spec.Name != nil {
		res.AlarmName = r.ko.Spec.Name
	}
	if r.ko.Spec.ComparisonOperator != nil {
		res.ComparisonOperator = svcsdktypes.ComparisonOperator(*r.ko.Spec.ComparisonOperator)
	}
	if r.ko.Spec.DatapointsToAlarm != nil {
		datapointsToAlarmCopy0 := *r.ko.Spec.DatapointsToAlarm
		if datapointsToAlarmCopy0 > math.MaxInt32 || datapointsToAlarmCopy0 < math.MinInt32 {
			return nil, fmt.Errorf("error: field DatapointsToAlarm is of type int32")
		}
		datapointsToAlarmCopy := int32(datapointsToAlarmCopy0)
		res.DatapointsToAlarm = &datapointsToAlarmCopy
	}
	if r.ko.Spec.Dimensions != nil {
		f6 := []svcsdktypes.Dimension{}
		for _, f6iter := range r.ko.Spec.Dimensions {
			f6elem := &svcsdktypes.Dimension{}
			if f6iter.Name != nil {
				f6elem.Name = f6iter.Name
			}
			if f6iter.Value != nil {
				f6elem.Value = f6iter.Value
			}
			f6 = append(f6, *f6elem)
		}
		res.Dimensions = f6
	}
	if r.ko.Spec.EvaluateLowSampleCountPercentile != nil {
		res.EvaluateLowSampleCountPercentile = r.ko.Spec.EvaluateLowSampleCountPercentile
	}
	if r.ko.Spec.EvaluationPeriods != nil {
		evaluationPeriodsCopy0 := *r.ko.Spec.EvaluationPeriods
		if evaluationPeriodsCopy0 > math.MaxInt32 || evaluationPeriodsCopy0 < math.MinInt32 {
			return nil, fmt.Errorf("error: field EvaluationPeriods is of type int32")
		}
		evaluationPeriodsCopy := int32(evaluationPeriodsCopy0)
		res.EvaluationPeriods = &evaluationPeriodsCopy
	}
	if r.ko.Spec.ExtendedStatistic != nil {
		res.ExtendedStatistic = r.ko.Spec.ExtendedStatistic
	}
	if r.ko.Spec.InsufficientDataActions != nil {
		res.InsufficientDataActions = aws.ToStringSlice(r.ko.Spec.InsufficientDataActions)
	}
	if r.ko.Spec.MetricName != nil {
		res.MetricName = r.ko.Spec.MetricName
	}
	if r.ko.Spec.Metrics != nil {
		f12 := []svcsdktypes.MetricDataQuery{}
		for _, f12iter := range r.ko.Spec.Metrics {
			f12elem := &svcsdktypes.MetricDataQuery{}
			if f12iter.AccountID != nil {
				f12elem.AccountId = f12iter.AccountID
			}
			if f12iter.Expression != nil {
				f12elem.Expression = f12iter.Expression
			}
			if f12iter.ID != nil {
				f12elem.Id = f12iter.ID
			}
			if f12iter.Label != nil {
				f12elem.Label = f12iter.Label
			}
			if f12iter.MetricStat != nil {
				f12elemf4 := &svcsdktypes.MetricStat{}
				if f12iter.MetricStat.Metric != nil {
					f12elemf4f0 := &svcsdktypes.Metric{}
					if f12iter.MetricStat.Metric.Dimensions != nil {
						f12elemf4f0f0 := []svcsdktypes.Dimension{}
						for _, f12elemf4f0f0iter := range f12iter.MetricStat.Metric.Dimensions {
							f12elemf4f0f0elem := &svcsdktypes.Dimension{}
							if f12elemf4f0f0iter.Name != nil {
								f12elemf4f0f0elem.Name = f12elemf4f0f0iter.Name
							}
							if f12elemf4f0f0iter.Value != nil {
								f12elemf4f0f0elem.Value = f12elemf4f0f0iter.Value
							}
							f12elemf4f0f0 = append(f12elemf4f0f0, *f12elemf4f0f0elem)
						}
						f12elemf4f0.Dimensions = f12elemf4f0f0
					}
					if f12iter.MetricStat.Metric.MetricName != nil {
						f12elemf4f0.MetricName = f12iter.MetricStat.Metric.MetricName
					}
					if f12iter.MetricStat.Metric.Namespace != nil {
						f12elemf4f0.Namespace = f12iter.MetricStat.Metric.Namespace
					}
					f12elemf4.Metric = f12elemf4f0
				}
				if f12iter.MetricStat.Period != nil {
					periodCopy0 := *f12iter.MetricStat.Period
					if periodCopy0 > math.MaxInt32 || periodCopy0 < math.MinInt32 {
						return nil, fmt.Errorf("error: field Period is of type int32")
					}
					periodCopy := int32(periodCopy0)
					f12elemf4.Period = &periodCopy
				}
				if f12iter.MetricStat.Stat != nil {
					f12elemf4.Stat = f12iter.MetricStat.Stat
				}
				if f12iter.MetricStat.Unit != nil {
					f12elemf4.Unit = svcsdktypes.StandardUnit(*f12iter.MetricStat.Unit)
				}
				f12elem.MetricStat = f12elemf4
			}
			if f12iter.Period != nil {
				periodCopy0 := *f12iter.Period
				if periodCopy0 > math.MaxInt32 || periodCopy0 < math.MinInt32 {
					return nil, fmt.Errorf("error: field Period is of type int32")
				}
				periodCopy := int32(periodCopy0)
				f12elem.Period = &periodCopy
			}
			if f12iter.ReturnData != nil {
				f12elem.ReturnData = f12iter.ReturnData
			}
			f12 = append(f12, *f12elem)
		}
		res.Metrics = f12
	}
	if r.ko.Spec.Namespace != nil {
		res.Namespace = r.ko.Spec.Namespace
	}
	if r.ko.Spec.OKActions != nil {
		res.OKActions = aws.ToStringSlice(r.ko.Spec.OKActions)
	}
	if r.ko.Spec.Period != nil {
		periodCopy0 := *r.ko.Spec.Period
		if periodCopy0 > math.MaxInt32 || periodCopy0 < math.MinInt32 {
			return nil, fmt.Errorf("error: field Period is of type int32")
		}
		periodCopy := int32(periodCopy0)
		res.Period = &periodCopy
	}
	if r.ko.Spec.Statistic != nil {
		res.Statistic = svcsdktypes.Statistic(*r.ko.Spec.Statistic)
	}
	if r.ko.Spec.Tags != nil {
		f17 := []svcsdktypes.Tag{}
		for _, f17iter := range r.ko.Spec.Tags {
			f17elem := &svcsdktypes.Tag{}
			if f17iter.Key != nil {
				f17elem.Key = f17iter.Key
			}
			if f17iter.Value != nil {
				f17elem.Value = f17iter.Value
			}
			f17 = append(f17, *f17elem)
		}
		res.Tags = f17
	}
	if r.ko.Spec.Threshold != nil {
		res.Threshold = r.ko.Spec.Threshold
	}
	if r.ko.Spec.ThresholdMetricID != nil {
		res.ThresholdMetricId = r.ko.Spec.ThresholdMetricID
	}
	if r.ko.Spec.TreatMissingData != nil {
		res.TreatMissingData = r.ko.Spec.TreatMissingData
	}
	if r.ko.Spec.Unit != nil {
		res.Unit = svcsdktypes.StandardUnit(*r.ko.Spec.Unit)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.PutMetricAlarmOutput
	_ = resp
	resp, err = rm.sdkapi.PutMetricAlarm(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "PutMetricAlarm", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.PutMetricAlarmInput, error) {
	res := &svcsdk.PutMetricAlarmInput{}

	if r.ko.Spec.ActionsEnabled != nil {
		res.ActionsEnabled = r.ko.Spec.ActionsEnabled
	}
	if r.ko.Spec.AlarmActions != nil {
		res.AlarmActions = aws.ToStringSlice(r.ko.Spec.AlarmActions)
	}
	if r.ko.Spec.AlarmDescription != nil {
		res.AlarmDescription = r.ko.Spec.AlarmDescription
	}
	if r.ko.Spec.Name != nil {
		res.AlarmName = r.ko.Spec.Name
	}
	if r.ko.Spec.ComparisonOperator != nil {
		res.ComparisonOperator = svcsdktypes.ComparisonOperator(*r.ko.Spec.ComparisonOperator)
	}
	if r.ko.Spec.DatapointsToAlarm != nil {
		datapointsToAlarmCopy0 := *r.ko.Spec.DatapointsToAlarm
		if datapointsToAlarmCopy0 > math.MaxInt32 || datapointsToAlarmCopy0 < math.MinInt32 {
			return nil, fmt.Errorf("error: field DatapointsToAlarm is of type int32")
		}
		datapointsToAlarmCopy := int32(datapointsToAlarmCopy0)
		res.DatapointsToAlarm = &datapointsToAlarmCopy
	}
	if r.ko.Spec.Dimensions != nil {
		f6 := []svcsdktypes.Dimension{}
		for _, f6iter := range r.ko.Spec.Dimensions {
			f6elem := &svcsdktypes.Dimension{}
			if f6iter.Name != nil {
				f6elem.Name = f6iter.Name
			}
			if f6iter.Value != nil {
				f6elem.Value = f6iter.Value
			}
			f6 = append(f6, *f6elem)
		}
		res.Dimensions = f6
	}
	if r.ko.Spec.EvaluateLowSampleCountPercentile != nil {
		res.EvaluateLowSampleCountPercentile = r.ko.Spec.EvaluateLowSampleCountPercentile
	}
	if r.ko.Spec.EvaluationPeriods != nil {
		evaluationPeriodsCopy0 := *r.ko.Spec.EvaluationPeriods
		if evaluationPeriodsCopy0 > math.MaxInt32 || evaluationPeriodsCopy0 < math.MinInt32 {
			return nil, fmt.Errorf("error: field EvaluationPeriods is of type int32")
		}
		evaluationPeriodsCopy := int32(evaluationPeriodsCopy0)
		res.EvaluationPeriods = &evaluationPeriodsCopy
	}
	if r.ko.Spec.ExtendedStatistic != nil {
		res.ExtendedStatistic = r.ko.Spec.ExtendedStatistic
	}
	if r.ko.Spec.InsufficientDataActions != nil {
		res.InsufficientDataActions = aws.ToStringSlice(r.ko.Spec.InsufficientDataActions)
	}
	if r.ko.Spec.MetricName != nil {
		res.MetricName = r.ko.Spec.MetricName
	}
	if r.ko.Spec.Metrics != nil {
		f12 := []svcsdktypes.MetricDataQuery{}
		for _, f12iter := range r.ko.Spec.Metrics {
			f12elem := &svcsdktypes.MetricDataQuery{}
			if f12iter.AccountID != nil {
				f12elem.AccountId = f12iter.AccountID
			}
			if f12iter.Expression != nil {
				f12elem.Expression = f12iter.Expression
			}
			if f12iter.ID != nil {
				f12elem.Id = f12iter.ID
			}
			if f12iter.Label != nil {
				f12elem.Label = f12iter.Label
			}
			if f12iter.MetricStat != nil {
				f12elemf4 := &svcsdktypes.MetricStat{}
				if f12iter.MetricStat.Metric != nil {
					f12elemf4f0 := &svcsdktypes.Metric{}
					if f12iter.MetricStat.Metric.Dimensions != nil {
						f12elemf4f0f0 := []svcsdktypes.Dimension{}
						for _, f12elemf4f0f0iter := range f12iter.MetricStat.Metric.Dimensions {
							f12elemf4f0f0elem := &svcsdktypes.Dimension{}
							if f12elemf4f0f0iter.Name != nil {
								f12elemf4f0f0elem.Name = f12elemf4f0f0iter.Name
							}
							if f12elemf4f0f0iter.Value != nil {
								f12elemf4f0f0elem.Value = f12elemf4f0f0iter.Value
							}
							f12elemf4f0f0 = append(f12elemf4f0f0, *f12elemf4f0f0elem)
						}
						f12elemf4f0.Dimensions = f12elemf4f0f0
					}
					if f12iter.MetricStat.Metric.MetricName != nil {
						f12elemf4f0.MetricName = f12iter.MetricStat.Metric.MetricName
					}
					if f12iter.MetricStat.Metric.Namespace != nil {
						f12elemf4f0.Namespace = f12iter.MetricStat.Metric.Namespace
					}
					f12elemf4.Metric = f12elemf4f0
				}
				if f12iter.MetricStat.Period != nil {
					periodCopy0 := *f12iter.MetricStat.Period
					if periodCopy0 > math.MaxInt32 || periodCopy0 < math.MinInt32 {
						return nil, fmt.Errorf("error: field Period is of type int32")
					}
					periodCopy := int32(periodCopy0)
					f12elemf4.Period = &periodCopy
				}
				if f12iter.MetricStat.Stat != nil {
					f12elemf4.Stat = f12iter.MetricStat.Stat
				}
				if f12iter.MetricStat.Unit != nil {
					f12elemf4.Unit = svcsdktypes.StandardUnit(*f12iter.MetricStat.Unit)
				}
				f12elem.MetricStat = f12elemf4
			}
			if f12iter.Period != nil {
				periodCopy0 := *f12iter.Period
				if periodCopy0 > math.MaxInt32 || periodCopy0 < math.MinInt32 {
					return nil, fmt.Errorf("error: field Period is of type int32")
				}
				periodCopy := int32(periodCopy0)
				f12elem.Period = &periodCopy
			}
			if f12iter.ReturnData != nil {
				f12elem.ReturnData = f12iter.ReturnData
			}
			f12 = append(f12, *f12elem)
		}
		res.Metrics = f12
	}
	if r.ko.Spec.Namespace != nil {
		res.Namespace = r.ko.Spec.Namespace
	}
	if r.ko.Spec.OKActions != nil {
		res.OKActions = aws.ToStringSlice(r.ko.Spec.OKActions)
	}
	if r.ko.Spec.Period != nil {
		periodCopy0 := *r.ko.Spec.Period
		if periodCopy0 > math.MaxInt32 || periodCopy0 < math.MinInt32 {
			return nil, fmt.Errorf("error: field Period is of type int32")
		}
		periodCopy := int32(periodCopy0)
		res.Period = &periodCopy
	}
	if r.ko.Spec.Statistic != nil {
		res.Statistic = svcsdktypes.Statistic(*r.ko.Spec.Statistic)
	}
	if r.ko.Spec.Tags != nil {
		f17 := []svcsdktypes.Tag{}
		for _, f17iter := range r.ko.Spec.Tags {
			f17elem := &svcsdktypes.Tag{}
			if f17iter.Key != nil {
				f17elem.Key = f17iter.Key
			}
			if f17iter.Value != nil {
				f17elem.Value = f17iter.Value
			}
			f17 = append(f17, *f17elem)
		}
		res.Tags = f17
	}
	if r.ko.Spec.Threshold != nil {
		res.Threshold = r.ko.Spec.Threshold
	}
	if r.ko.Spec.ThresholdMetricID != nil {
		res.ThresholdMetricId = r.ko.Spec.ThresholdMetricID
	}
	if r.ko.Spec.TreatMissingData != nil {
		res.TreatMissingData = r.ko.Spec.TreatMissingData
	}
	if r.ko.Spec.Unit != nil {
		res.Unit = svcsdktypes.StandardUnit(*r.ko.Spec.Unit)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	input.AlarmNames = []string{*r.ko.Spec.Name}

	var resp *svcsdk.DeleteAlarmsOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteAlarms(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteAlarms", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteAlarmsInput, error) {
	res := &svcsdk.DeleteAlarmsInput{}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.MetricAlarm,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
