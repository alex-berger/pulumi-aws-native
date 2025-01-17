// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::ManagedPolicy
func LookupManagedPolicy(ctx *pulumi.Context, args *LookupManagedPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagedPolicyResult, error) {
	var rv LookupManagedPolicyResult
	err := ctx.Invoke("aws-native:iam:getManagedPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedPolicyArgs struct {
	Id string `pulumi:"id"`
}

type LookupManagedPolicyResult struct {
	Groups         []string    `pulumi:"groups"`
	Id             *string     `pulumi:"id"`
	PolicyDocument interface{} `pulumi:"policyDocument"`
	Roles          []string    `pulumi:"roles"`
	Users          []string    `pulumi:"users"`
}

func LookupManagedPolicyOutput(ctx *pulumi.Context, args LookupManagedPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupManagedPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedPolicyResult, error) {
			args := v.(LookupManagedPolicyArgs)
			r, err := LookupManagedPolicy(ctx, &args, opts...)
			var s LookupManagedPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedPolicyResultOutput)
}

type LookupManagedPolicyOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupManagedPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPolicyArgs)(nil)).Elem()
}

type LookupManagedPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupManagedPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedPolicyResult)(nil)).Elem()
}

func (o LookupManagedPolicyResultOutput) ToLookupManagedPolicyResultOutput() LookupManagedPolicyResultOutput {
	return o
}

func (o LookupManagedPolicyResultOutput) ToLookupManagedPolicyResultOutputWithContext(ctx context.Context) LookupManagedPolicyResultOutput {
	return o
}

func (o LookupManagedPolicyResultOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o LookupManagedPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupManagedPolicyResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func (o LookupManagedPolicyResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o LookupManagedPolicyResultOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagedPolicyResult) []string { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedPolicyResultOutput{})
}
