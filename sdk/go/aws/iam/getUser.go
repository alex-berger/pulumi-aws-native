// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::User
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("aws-native:iam:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	Id string `pulumi:"id"`
}

type LookupUserResult struct {
	Arn                 *string           `pulumi:"arn"`
	Groups              []string          `pulumi:"groups"`
	Id                  *string           `pulumi:"id"`
	LoginProfile        *UserLoginProfile `pulumi:"loginProfile"`
	ManagedPolicyArns   []string          `pulumi:"managedPolicyArns"`
	Path                *string           `pulumi:"path"`
	PermissionsBoundary *string           `pulumi:"permissionsBoundary"`
	Policies            []UserPolicy      `pulumi:"policies"`
	Tags                []UserTag         `pulumi:"tags"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

type LookupUserOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o LookupUserResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) LoginProfile() UserLoginProfilePtrOutput {
	return o.ApplyT(func(v LookupUserResult) *UserLoginProfile { return v.LoginProfile }).(UserLoginProfilePtrOutput)
}

func (o LookupUserResultOutput) ManagedPolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []string { return v.ManagedPolicyArns }).(pulumi.StringArrayOutput)
}

func (o LookupUserResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) PermissionsBoundary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.PermissionsBoundary }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) Policies() UserPolicyArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []UserPolicy { return v.Policies }).(UserPolicyArrayOutput)
}

func (o LookupUserResultOutput) Tags() UserTagArrayOutput {
	return o.ApplyT(func(v LookupUserResult) []UserTag { return v.Tags }).(UserTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
