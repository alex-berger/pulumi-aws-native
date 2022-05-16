// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package evidently

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Evidently::Experiment.
func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	var rv LookupExperimentResult
	err := ctx.Invoke("aws-native:evidently:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupExperimentResult struct {
	Arn               *string                         `pulumi:"arn"`
	Description       *string                         `pulumi:"description"`
	MetricGoals       []ExperimentMetricGoalObject    `pulumi:"metricGoals"`
	OnlineAbConfig    *ExperimentOnlineAbConfigObject `pulumi:"onlineAbConfig"`
	RandomizationSalt *string                         `pulumi:"randomizationSalt"`
	// Start Experiment. Default is False
	RunningStatus *ExperimentRunningStatusObject `pulumi:"runningStatus"`
	SamplingRate  *int                           `pulumi:"samplingRate"`
	// An array of key-value pairs to apply to this resource.
	Tags       []ExperimentTag             `pulumi:"tags"`
	Treatments []ExperimentTreatmentObject `pulumi:"treatments"`
}

func LookupExperimentOutput(ctx *pulumi.Context, args LookupExperimentOutputArgs, opts ...pulumi.InvokeOption) LookupExperimentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExperimentResult, error) {
			args := v.(LookupExperimentArgs)
			r, err := LookupExperiment(ctx, &args, opts...)
			var s LookupExperimentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExperimentResultOutput)
}

type LookupExperimentOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupExperimentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentArgs)(nil)).Elem()
}

type LookupExperimentResultOutput struct{ *pulumi.OutputState }

func (LookupExperimentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentResult)(nil)).Elem()
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutput() LookupExperimentResultOutput {
	return o
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutputWithContext(ctx context.Context) LookupExperimentResultOutput {
	return o
}

func (o LookupExperimentResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupExperimentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupExperimentResultOutput) MetricGoals() ExperimentMetricGoalObjectArrayOutput {
	return o.ApplyT(func(v LookupExperimentResult) []ExperimentMetricGoalObject { return v.MetricGoals }).(ExperimentMetricGoalObjectArrayOutput)
}

func (o LookupExperimentResultOutput) OnlineAbConfig() ExperimentOnlineAbConfigObjectPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *ExperimentOnlineAbConfigObject { return v.OnlineAbConfig }).(ExperimentOnlineAbConfigObjectPtrOutput)
}

func (o LookupExperimentResultOutput) RandomizationSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *string { return v.RandomizationSalt }).(pulumi.StringPtrOutput)
}

// Start Experiment. Default is False
func (o LookupExperimentResultOutput) RunningStatus() ExperimentRunningStatusObjectPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *ExperimentRunningStatusObject { return v.RunningStatus }).(ExperimentRunningStatusObjectPtrOutput)
}

func (o LookupExperimentResultOutput) SamplingRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *int { return v.SamplingRate }).(pulumi.IntPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupExperimentResultOutput) Tags() ExperimentTagArrayOutput {
	return o.ApplyT(func(v LookupExperimentResult) []ExperimentTag { return v.Tags }).(ExperimentTagArrayOutput)
}

func (o LookupExperimentResultOutput) Treatments() ExperimentTreatmentObjectArrayOutput {
	return o.ApplyT(func(v LookupExperimentResult) []ExperimentTreatmentObject { return v.Treatments }).(ExperimentTreatmentObjectArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExperimentResultOutput{})
}