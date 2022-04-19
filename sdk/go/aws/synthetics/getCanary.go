// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synthetics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Synthetics::Canary
func LookupCanary(ctx *pulumi.Context, args *LookupCanaryArgs, opts ...pulumi.InvokeOption) (*LookupCanaryResult, error) {
	var rv LookupCanaryResult
	err := ctx.Invoke("aws-native:synthetics:getCanary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCanaryArgs struct {
	// Name of the canary.
	Name string `pulumi:"name"`
}

type LookupCanaryResult struct {
	// Provide artifact configuration
	ArtifactConfig *CanaryArtifactConfig `pulumi:"artifactConfig"`
	// Provide the s3 bucket output location for test results
	ArtifactS3Location *string `pulumi:"artifactS3Location"`
	// Provide the canary script source
	Code *CanaryCode `pulumi:"code"`
	// Lambda Execution role used to run your canaries
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// Retention period of failed canary runs represented in number of days
	FailureRetentionPeriod *int `pulumi:"failureRetentionPeriod"`
	// Id of the canary
	Id *string `pulumi:"id"`
	// Provide canary run configuration
	RunConfig *CanaryRunConfig `pulumi:"runConfig"`
	// Runtime version of Synthetics Library
	RuntimeVersion *string `pulumi:"runtimeVersion"`
	// Frequency to run your canaries
	Schedule *CanarySchedule `pulumi:"schedule"`
	// Runs canary if set to True. Default is False
	StartCanaryAfterCreation *bool `pulumi:"startCanaryAfterCreation"`
	// State of the canary
	State *string `pulumi:"state"`
	// Retention period of successful canary runs represented in number of days
	SuccessRetentionPeriod *int        `pulumi:"successRetentionPeriod"`
	Tags                   []CanaryTag `pulumi:"tags"`
	// Provide VPC Configuration if enabled.
	VPCConfig *CanaryVPCConfig `pulumi:"vPCConfig"`
	// Visual reference configuration for visual testing
	VisualReference *CanaryVisualReference `pulumi:"visualReference"`
}

func LookupCanaryOutput(ctx *pulumi.Context, args LookupCanaryOutputArgs, opts ...pulumi.InvokeOption) LookupCanaryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCanaryResult, error) {
			args := v.(LookupCanaryArgs)
			r, err := LookupCanary(ctx, &args, opts...)
			var s LookupCanaryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCanaryResultOutput)
}

type LookupCanaryOutputArgs struct {
	// Name of the canary.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupCanaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCanaryArgs)(nil)).Elem()
}

type LookupCanaryResultOutput struct{ *pulumi.OutputState }

func (LookupCanaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCanaryResult)(nil)).Elem()
}

func (o LookupCanaryResultOutput) ToLookupCanaryResultOutput() LookupCanaryResultOutput {
	return o
}

func (o LookupCanaryResultOutput) ToLookupCanaryResultOutputWithContext(ctx context.Context) LookupCanaryResultOutput {
	return o
}

// Provide artifact configuration
func (o LookupCanaryResultOutput) ArtifactConfig() CanaryArtifactConfigPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryArtifactConfig { return v.ArtifactConfig }).(CanaryArtifactConfigPtrOutput)
}

// Provide the s3 bucket output location for test results
func (o LookupCanaryResultOutput) ArtifactS3Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.ArtifactS3Location }).(pulumi.StringPtrOutput)
}

// Provide the canary script source
func (o LookupCanaryResultOutput) Code() CanaryCodePtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryCode { return v.Code }).(CanaryCodePtrOutput)
}

// Lambda Execution role used to run your canaries
func (o LookupCanaryResultOutput) ExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.ExecutionRoleArn }).(pulumi.StringPtrOutput)
}

// Retention period of failed canary runs represented in number of days
func (o LookupCanaryResultOutput) FailureRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *int { return v.FailureRetentionPeriod }).(pulumi.IntPtrOutput)
}

// Id of the canary
func (o LookupCanaryResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Provide canary run configuration
func (o LookupCanaryResultOutput) RunConfig() CanaryRunConfigPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryRunConfig { return v.RunConfig }).(CanaryRunConfigPtrOutput)
}

// Runtime version of Synthetics Library
func (o LookupCanaryResultOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

// Frequency to run your canaries
func (o LookupCanaryResultOutput) Schedule() CanarySchedulePtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanarySchedule { return v.Schedule }).(CanarySchedulePtrOutput)
}

// Runs canary if set to True. Default is False
func (o LookupCanaryResultOutput) StartCanaryAfterCreation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *bool { return v.StartCanaryAfterCreation }).(pulumi.BoolPtrOutput)
}

// State of the canary
func (o LookupCanaryResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Retention period of successful canary runs represented in number of days
func (o LookupCanaryResultOutput) SuccessRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *int { return v.SuccessRetentionPeriod }).(pulumi.IntPtrOutput)
}

func (o LookupCanaryResultOutput) Tags() CanaryTagArrayOutput {
	return o.ApplyT(func(v LookupCanaryResult) []CanaryTag { return v.Tags }).(CanaryTagArrayOutput)
}

// Provide VPC Configuration if enabled.
func (o LookupCanaryResultOutput) VPCConfig() CanaryVPCConfigPtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryVPCConfig { return v.VPCConfig }).(CanaryVPCConfigPtrOutput)
}

// Visual reference configuration for visual testing
func (o LookupCanaryResultOutput) VisualReference() CanaryVisualReferencePtrOutput {
	return o.ApplyT(func(v LookupCanaryResult) *CanaryVisualReference { return v.VisualReference }).(CanaryVisualReferencePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCanaryResultOutput{})
}
