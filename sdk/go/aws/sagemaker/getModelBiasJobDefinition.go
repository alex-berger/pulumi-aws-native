// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition
func LookupModelBiasJobDefinition(ctx *pulumi.Context, args *LookupModelBiasJobDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupModelBiasJobDefinitionResult, error) {
	var rv LookupModelBiasJobDefinitionResult
	err := ctx.Invoke("aws-native:sagemaker:getModelBiasJobDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelBiasJobDefinitionArgs struct {
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn string `pulumi:"jobDefinitionArn"`
}

type LookupModelBiasJobDefinitionResult struct {
	// The time at which the job definition was created.
	CreationTime *string `pulumi:"creationTime"`
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn *string `pulumi:"jobDefinitionArn"`
}

func LookupModelBiasJobDefinitionOutput(ctx *pulumi.Context, args LookupModelBiasJobDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupModelBiasJobDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelBiasJobDefinitionResult, error) {
			args := v.(LookupModelBiasJobDefinitionArgs)
			r, err := LookupModelBiasJobDefinition(ctx, &args, opts...)
			var s LookupModelBiasJobDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelBiasJobDefinitionResultOutput)
}

type LookupModelBiasJobDefinitionOutputArgs struct {
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn pulumi.StringInput `pulumi:"jobDefinitionArn"`
}

func (LookupModelBiasJobDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelBiasJobDefinitionArgs)(nil)).Elem()
}

type LookupModelBiasJobDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupModelBiasJobDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelBiasJobDefinitionResult)(nil)).Elem()
}

func (o LookupModelBiasJobDefinitionResultOutput) ToLookupModelBiasJobDefinitionResultOutput() LookupModelBiasJobDefinitionResultOutput {
	return o
}

func (o LookupModelBiasJobDefinitionResultOutput) ToLookupModelBiasJobDefinitionResultOutputWithContext(ctx context.Context) LookupModelBiasJobDefinitionResultOutput {
	return o
}

// The time at which the job definition was created.
func (o LookupModelBiasJobDefinitionResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelBiasJobDefinitionResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of job definition.
func (o LookupModelBiasJobDefinitionResultOutput) JobDefinitionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModelBiasJobDefinitionResult) *string { return v.JobDefinitionArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelBiasJobDefinitionResultOutput{})
}
