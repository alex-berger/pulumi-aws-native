// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codedeploy

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CodeDeploy::DeploymentConfig
func LookupDeploymentConfig(ctx *pulumi.Context, args *LookupDeploymentConfigArgs, opts ...pulumi.InvokeOption) (*LookupDeploymentConfigResult, error) {
	var rv LookupDeploymentConfigResult
	err := ctx.Invoke("aws-native:codedeploy:getDeploymentConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeploymentConfigArgs struct {
	Id string `pulumi:"id"`
}

type LookupDeploymentConfigResult struct {
	Id *string `pulumi:"id"`
}

func LookupDeploymentConfigOutput(ctx *pulumi.Context, args LookupDeploymentConfigOutputArgs, opts ...pulumi.InvokeOption) LookupDeploymentConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeploymentConfigResult, error) {
			args := v.(LookupDeploymentConfigArgs)
			r, err := LookupDeploymentConfig(ctx, &args, opts...)
			var s LookupDeploymentConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeploymentConfigResultOutput)
}

type LookupDeploymentConfigOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDeploymentConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentConfigArgs)(nil)).Elem()
}

type LookupDeploymentConfigResultOutput struct{ *pulumi.OutputState }

func (LookupDeploymentConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeploymentConfigResult)(nil)).Elem()
}

func (o LookupDeploymentConfigResultOutput) ToLookupDeploymentConfigResultOutput() LookupDeploymentConfigResultOutput {
	return o
}

func (o LookupDeploymentConfigResultOutput) ToLookupDeploymentConfigResultOutputWithContext(ctx context.Context) LookupDeploymentConfigResultOutput {
	return o
}

func (o LookupDeploymentConfigResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeploymentConfigResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeploymentConfigResultOutput{})
}
