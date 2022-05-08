// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::App
func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	var rv LookupAppResult
	err := ctx.Invoke("aws-native:sagemaker:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	// The name of the app.
	AppName string `pulumi:"appName"`
	// The type of app.
	AppType AppType `pulumi:"appType"`
	// The domain ID.
	DomainId string `pulumi:"domainId"`
	// The user profile name.
	UserProfileName string `pulumi:"userProfileName"`
}

type LookupAppResult struct {
	// The Amazon Resource Name (ARN) of the app.
	AppArn *string `pulumi:"appArn"`
	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
	ResourceSpec *AppResourceSpec `pulumi:"resourceSpec"`
}

func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppResult, error) {
			args := v.(LookupAppArgs)
			r, err := LookupApp(ctx, &args, opts...)
			var s LookupAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	// The name of the app.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The type of app.
	AppType AppTypeInput `pulumi:"appType"`
	// The domain ID.
	DomainId pulumi.StringInput `pulumi:"domainId"`
	// The user profile name.
	UserProfileName pulumi.StringInput `pulumi:"userProfileName"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}

type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the app.
func (o LookupAppResultOutput) AppArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.AppArn }).(pulumi.StringPtrOutput)
}

// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
func (o LookupAppResultOutput) ResourceSpec() AppResourceSpecPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *AppResourceSpec { return v.ResourceSpec }).(AppResourceSpecPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
