// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
func LookupModelExplainabilityJobDefinition(ctx *pulumi.Context, args *LookupModelExplainabilityJobDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupModelExplainabilityJobDefinitionResult, error) {
	var rv LookupModelExplainabilityJobDefinitionResult
	err := ctx.Invoke("aws-native:sagemaker:getModelExplainabilityJobDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelExplainabilityJobDefinitionArgs struct {
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn string `pulumi:"jobDefinitionArn"`
}

type LookupModelExplainabilityJobDefinitionResult struct {
	// The time at which the job definition was created.
	CreationTime *string `pulumi:"creationTime"`
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn *string `pulumi:"jobDefinitionArn"`
}

func LookupModelExplainabilityJobDefinitionOutput(ctx *pulumi.Context, args LookupModelExplainabilityJobDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupModelExplainabilityJobDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelExplainabilityJobDefinitionResult, error) {
			args := v.(LookupModelExplainabilityJobDefinitionArgs)
			r, err := LookupModelExplainabilityJobDefinition(ctx, &args, opts...)
			var s LookupModelExplainabilityJobDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelExplainabilityJobDefinitionResultOutput)
}

type LookupModelExplainabilityJobDefinitionOutputArgs struct {
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn pulumi.StringInput `pulumi:"jobDefinitionArn"`
}

func (LookupModelExplainabilityJobDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelExplainabilityJobDefinitionArgs)(nil)).Elem()
}

type LookupModelExplainabilityJobDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupModelExplainabilityJobDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelExplainabilityJobDefinitionResult)(nil)).Elem()
}

func (o LookupModelExplainabilityJobDefinitionResultOutput) ToLookupModelExplainabilityJobDefinitionResultOutput() LookupModelExplainabilityJobDefinitionResultOutput {
	return o
}

func (o LookupModelExplainabilityJobDefinitionResultOutput) ToLookupModelExplainabilityJobDefinitionResultOutputWithContext(ctx context.Context) LookupModelExplainabilityJobDefinitionResultOutput {
	return o
}

// The time at which the job definition was created.
func (o LookupModelExplainabilityJobDefinitionResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelExplainabilityJobDefinitionResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of job definition.
func (o LookupModelExplainabilityJobDefinitionResultOutput) JobDefinitionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelExplainabilityJobDefinitionResult) *string { return v.JobDefinitionArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelExplainabilityJobDefinitionResultOutput{})
}