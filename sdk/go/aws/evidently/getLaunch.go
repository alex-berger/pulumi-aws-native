// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package evidently

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Evidently::Launch.
func LookupLaunch(ctx *pulumi.Context, args *LookupLaunchArgs, opts ...pulumi.InvokeOption) (*LookupLaunchResult, error) {
	var rv LookupLaunchResult
	err := ctx.Invoke("aws-native:evidently:getLaunch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLaunchArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupLaunchResult struct {
	Arn                   *string                        `pulumi:"arn"`
	Description           *string                        `pulumi:"description"`
	Groups                []LaunchGroupObject            `pulumi:"groups"`
	MetricMonitors        []LaunchMetricDefinitionObject `pulumi:"metricMonitors"`
	RandomizationSalt     *string                        `pulumi:"randomizationSalt"`
	ScheduledSplitsConfig []LaunchStepConfig             `pulumi:"scheduledSplitsConfig"`
	// An array of key-value pairs to apply to this resource.
	Tags []LaunchTag `pulumi:"tags"`
}

func LookupLaunchOutput(ctx *pulumi.Context, args LookupLaunchOutputArgs, opts ...pulumi.InvokeOption) LookupLaunchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLaunchResult, error) {
			args := v.(LookupLaunchArgs)
			r, err := LookupLaunch(ctx, &args, opts...)
			var s LookupLaunchResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLaunchResultOutput)
}

type LookupLaunchOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupLaunchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLaunchArgs)(nil)).Elem()
}

type LookupLaunchResultOutput struct{ *pulumi.OutputState }

func (LookupLaunchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLaunchResult)(nil)).Elem()
}

func (o LookupLaunchResultOutput) ToLookupLaunchResultOutput() LookupLaunchResultOutput {
	return o
}

func (o LookupLaunchResultOutput) ToLookupLaunchResultOutputWithContext(ctx context.Context) LookupLaunchResultOutput {
	return o
}

func (o LookupLaunchResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLaunchResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupLaunchResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLaunchResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLaunchResultOutput) Groups() LaunchGroupObjectArrayOutput {
	return o.ApplyT(func(v LookupLaunchResult) []LaunchGroupObject { return v.Groups }).(LaunchGroupObjectArrayOutput)
}

func (o LookupLaunchResultOutput) MetricMonitors() LaunchMetricDefinitionObjectArrayOutput {
	return o.ApplyT(func(v LookupLaunchResult) []LaunchMetricDefinitionObject { return v.MetricMonitors }).(LaunchMetricDefinitionObjectArrayOutput)
}

func (o LookupLaunchResultOutput) RandomizationSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLaunchResult) *string { return v.RandomizationSalt }).(pulumi.StringPtrOutput)
}

func (o LookupLaunchResultOutput) ScheduledSplitsConfig() LaunchStepConfigArrayOutput {
	return o.ApplyT(func(v LookupLaunchResult) []LaunchStepConfig { return v.ScheduledSplitsConfig }).(LaunchStepConfigArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupLaunchResultOutput) Tags() LaunchTagArrayOutput {
	return o.ApplyT(func(v LookupLaunchResult) []LaunchTag { return v.Tags }).(LaunchTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLaunchResultOutput{})
}
