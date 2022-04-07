// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AutoScaling::ScalingPolicy
func LookupScalingPolicy(ctx *pulumi.Context, args *LookupScalingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupScalingPolicyResult, error) {
	var rv LookupScalingPolicyResult
	err := ctx.Invoke("aws-native:autoscaling:getScalingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScalingPolicyArgs struct {
	Id string `pulumi:"id"`
}

type LookupScalingPolicyResult struct {
	AdjustmentType                 *string                                      `pulumi:"adjustmentType"`
	AutoScalingGroupName           *string                                      `pulumi:"autoScalingGroupName"`
	Cooldown                       *string                                      `pulumi:"cooldown"`
	EstimatedInstanceWarmup        *int                                         `pulumi:"estimatedInstanceWarmup"`
	Id                             *string                                      `pulumi:"id"`
	MetricAggregationType          *string                                      `pulumi:"metricAggregationType"`
	MinAdjustmentMagnitude         *int                                         `pulumi:"minAdjustmentMagnitude"`
	PolicyType                     *string                                      `pulumi:"policyType"`
	PredictiveScalingConfiguration *ScalingPolicyPredictiveScalingConfiguration `pulumi:"predictiveScalingConfiguration"`
	ScalingAdjustment              *int                                         `pulumi:"scalingAdjustment"`
	StepAdjustments                []ScalingPolicyStepAdjustment                `pulumi:"stepAdjustments"`
	TargetTrackingConfiguration    *ScalingPolicyTargetTrackingConfiguration    `pulumi:"targetTrackingConfiguration"`
}

func LookupScalingPolicyOutput(ctx *pulumi.Context, args LookupScalingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupScalingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScalingPolicyResult, error) {
			args := v.(LookupScalingPolicyArgs)
			r, err := LookupScalingPolicy(ctx, &args, opts...)
			var s LookupScalingPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScalingPolicyResultOutput)
}

type LookupScalingPolicyOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupScalingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPolicyArgs)(nil)).Elem()
}

type LookupScalingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupScalingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPolicyResult)(nil)).Elem()
}

func (o LookupScalingPolicyResultOutput) ToLookupScalingPolicyResultOutput() LookupScalingPolicyResultOutput {
	return o
}

func (o LookupScalingPolicyResultOutput) ToLookupScalingPolicyResultOutputWithContext(ctx context.Context) LookupScalingPolicyResultOutput {
	return o
}

func (o LookupScalingPolicyResultOutput) AdjustmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *string { return v.AdjustmentType }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPolicyResultOutput) AutoScalingGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *string { return v.AutoScalingGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPolicyResultOutput) Cooldown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *string { return v.Cooldown }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPolicyResultOutput) EstimatedInstanceWarmup() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *int { return v.EstimatedInstanceWarmup }).(pulumi.IntPtrOutput)
}

func (o LookupScalingPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPolicyResultOutput) MetricAggregationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *string { return v.MetricAggregationType }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPolicyResultOutput) MinAdjustmentMagnitude() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *int { return v.MinAdjustmentMagnitude }).(pulumi.IntPtrOutput)
}

func (o LookupScalingPolicyResultOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *string { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPolicyResultOutput) PredictiveScalingConfiguration() ScalingPolicyPredictiveScalingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *ScalingPolicyPredictiveScalingConfiguration {
		return v.PredictiveScalingConfiguration
	}).(ScalingPolicyPredictiveScalingConfigurationPtrOutput)
}

func (o LookupScalingPolicyResultOutput) ScalingAdjustment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *int { return v.ScalingAdjustment }).(pulumi.IntPtrOutput)
}

func (o LookupScalingPolicyResultOutput) StepAdjustments() ScalingPolicyStepAdjustmentArrayOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) []ScalingPolicyStepAdjustment { return v.StepAdjustments }).(ScalingPolicyStepAdjustmentArrayOutput)
}

func (o LookupScalingPolicyResultOutput) TargetTrackingConfiguration() ScalingPolicyTargetTrackingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupScalingPolicyResult) *ScalingPolicyTargetTrackingConfiguration {
		return v.TargetTrackingConfiguration
	}).(ScalingPolicyTargetTrackingConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScalingPolicyResultOutput{})
}
