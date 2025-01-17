// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package applicationautoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApplicationAutoScaling::ScalableTarget
func LookupScalableTarget(ctx *pulumi.Context, args *LookupScalableTargetArgs, opts ...pulumi.InvokeOption) (*LookupScalableTargetResult, error) {
	var rv LookupScalableTargetResult
	err := ctx.Invoke("aws-native:applicationautoscaling:getScalableTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScalableTargetArgs struct {
	Id string `pulumi:"id"`
}

type LookupScalableTargetResult struct {
	Id               *string                         `pulumi:"id"`
	MaxCapacity      *int                            `pulumi:"maxCapacity"`
	MinCapacity      *int                            `pulumi:"minCapacity"`
	RoleARN          *string                         `pulumi:"roleARN"`
	ScheduledActions []ScalableTargetScheduledAction `pulumi:"scheduledActions"`
	SuspendedState   *ScalableTargetSuspendedState   `pulumi:"suspendedState"`
}

func LookupScalableTargetOutput(ctx *pulumi.Context, args LookupScalableTargetOutputArgs, opts ...pulumi.InvokeOption) LookupScalableTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScalableTargetResult, error) {
			args := v.(LookupScalableTargetArgs)
			r, err := LookupScalableTarget(ctx, &args, opts...)
			var s LookupScalableTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScalableTargetResultOutput)
}

type LookupScalableTargetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupScalableTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalableTargetArgs)(nil)).Elem()
}

type LookupScalableTargetResultOutput struct{ *pulumi.OutputState }

func (LookupScalableTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalableTargetResult)(nil)).Elem()
}

func (o LookupScalableTargetResultOutput) ToLookupScalableTargetResultOutput() LookupScalableTargetResultOutput {
	return o
}

func (o LookupScalableTargetResultOutput) ToLookupScalableTargetResultOutputWithContext(ctx context.Context) LookupScalableTargetResultOutput {
	return o
}

func (o LookupScalableTargetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalableTargetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupScalableTargetResultOutput) MaxCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalableTargetResult) *int { return v.MaxCapacity }).(pulumi.IntPtrOutput)
}

func (o LookupScalableTargetResultOutput) MinCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupScalableTargetResult) *int { return v.MinCapacity }).(pulumi.IntPtrOutput)
}

func (o LookupScalableTargetResultOutput) RoleARN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalableTargetResult) *string { return v.RoleARN }).(pulumi.StringPtrOutput)
}

func (o LookupScalableTargetResultOutput) ScheduledActions() ScalableTargetScheduledActionArrayOutput {
	return o.ApplyT(func(v LookupScalableTargetResult) []ScalableTargetScheduledAction { return v.ScheduledActions }).(ScalableTargetScheduledActionArrayOutput)
}

func (o LookupScalableTargetResultOutput) SuspendedState() ScalableTargetSuspendedStatePtrOutput {
	return o.ApplyT(func(v LookupScalableTargetResult) *ScalableTargetSuspendedState { return v.SuspendedState }).(ScalableTargetSuspendedStatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScalableTargetResultOutput{})
}
