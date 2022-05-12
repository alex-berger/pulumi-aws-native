// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudWatch::Alarm
func LookupAlarm(ctx *pulumi.Context, args *LookupAlarmArgs, opts ...pulumi.InvokeOption) (*LookupAlarmResult, error) {
	var rv LookupAlarmResult
	err := ctx.Invoke("aws-native:cloudwatch:getAlarm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlarmArgs struct {
	Id string `pulumi:"id"`
}

type LookupAlarmResult struct {
	ActionsEnabled                   *bool                  `pulumi:"actionsEnabled"`
	AlarmActions                     []string               `pulumi:"alarmActions"`
	AlarmDescription                 *string                `pulumi:"alarmDescription"`
	Arn                              *string                `pulumi:"arn"`
	ComparisonOperator               *string                `pulumi:"comparisonOperator"`
	DatapointsToAlarm                *int                   `pulumi:"datapointsToAlarm"`
	Dimensions                       []AlarmDimension       `pulumi:"dimensions"`
	EvaluateLowSampleCountPercentile *string                `pulumi:"evaluateLowSampleCountPercentile"`
	EvaluationPeriods                *int                   `pulumi:"evaluationPeriods"`
	ExtendedStatistic                *string                `pulumi:"extendedStatistic"`
	Id                               *string                `pulumi:"id"`
	InsufficientDataActions          []string               `pulumi:"insufficientDataActions"`
	MetricName                       *string                `pulumi:"metricName"`
	Metrics                          []AlarmMetricDataQuery `pulumi:"metrics"`
	Namespace                        *string                `pulumi:"namespace"`
	OKActions                        []string               `pulumi:"oKActions"`
	Period                           *int                   `pulumi:"period"`
	Statistic                        *string                `pulumi:"statistic"`
	Threshold                        *float64               `pulumi:"threshold"`
	ThresholdMetricId                *string                `pulumi:"thresholdMetricId"`
	TreatMissingData                 *string                `pulumi:"treatMissingData"`
	Unit                             *string                `pulumi:"unit"`
}

func LookupAlarmOutput(ctx *pulumi.Context, args LookupAlarmOutputArgs, opts ...pulumi.InvokeOption) LookupAlarmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlarmResult, error) {
			args := v.(LookupAlarmArgs)
			r, err := LookupAlarm(ctx, &args, opts...)
			var s LookupAlarmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlarmResultOutput)
}

type LookupAlarmOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAlarmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmArgs)(nil)).Elem()
}

type LookupAlarmResultOutput struct{ *pulumi.OutputState }

func (LookupAlarmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmResult)(nil)).Elem()
}

func (o LookupAlarmResultOutput) ToLookupAlarmResultOutput() LookupAlarmResultOutput {
	return o
}

func (o LookupAlarmResultOutput) ToLookupAlarmResultOutputWithContext(ctx context.Context) LookupAlarmResultOutput {
	return o
}

func (o LookupAlarmResultOutput) ActionsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *bool { return v.ActionsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupAlarmResultOutput) AlarmActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmResult) []string { return v.AlarmActions }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmResultOutput) AlarmDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.AlarmDescription }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) ComparisonOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.ComparisonOperator }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) DatapointsToAlarm() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *int { return v.DatapointsToAlarm }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmResultOutput) Dimensions() AlarmDimensionArrayOutput {
	return o.ApplyT(func(v LookupAlarmResult) []AlarmDimension { return v.Dimensions }).(AlarmDimensionArrayOutput)
}

func (o LookupAlarmResultOutput) EvaluateLowSampleCountPercentile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.EvaluateLowSampleCountPercentile }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) EvaluationPeriods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *int { return v.EvaluationPeriods }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmResultOutput) ExtendedStatistic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.ExtendedStatistic }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) InsufficientDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmResult) []string { return v.InsufficientDataActions }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmResultOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) Metrics() AlarmMetricDataQueryArrayOutput {
	return o.ApplyT(func(v LookupAlarmResult) []AlarmMetricDataQuery { return v.Metrics }).(AlarmMetricDataQueryArrayOutput)
}

func (o LookupAlarmResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) OKActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmResult) []string { return v.OKActions }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmResultOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *int { return v.Period }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmResultOutput) Statistic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.Statistic }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

func (o LookupAlarmResultOutput) ThresholdMetricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.ThresholdMetricId }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) TreatMissingData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.TreatMissingData }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmResultOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmResult) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlarmResultOutput{})
}
