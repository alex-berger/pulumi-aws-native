// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::UserToGroupAddition
func LookupUserToGroupAddition(ctx *pulumi.Context, args *LookupUserToGroupAdditionArgs, opts ...pulumi.InvokeOption) (*LookupUserToGroupAdditionResult, error) {
	var rv LookupUserToGroupAdditionResult
	err := ctx.Invoke("aws-native:iam:getUserToGroupAddition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserToGroupAdditionArgs struct {
	Id string `pulumi:"id"`
}

type LookupUserToGroupAdditionResult struct {
	GroupName *string  `pulumi:"groupName"`
	Id        *string  `pulumi:"id"`
	Users     []string `pulumi:"users"`
}

func LookupUserToGroupAdditionOutput(ctx *pulumi.Context, args LookupUserToGroupAdditionOutputArgs, opts ...pulumi.InvokeOption) LookupUserToGroupAdditionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserToGroupAdditionResult, error) {
			args := v.(LookupUserToGroupAdditionArgs)
			r, err := LookupUserToGroupAddition(ctx, &args, opts...)
			var s LookupUserToGroupAdditionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserToGroupAdditionResultOutput)
}

type LookupUserToGroupAdditionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserToGroupAdditionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserToGroupAdditionArgs)(nil)).Elem()
}

type LookupUserToGroupAdditionResultOutput struct{ *pulumi.OutputState }

func (LookupUserToGroupAdditionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserToGroupAdditionResult)(nil)).Elem()
}

func (o LookupUserToGroupAdditionResultOutput) ToLookupUserToGroupAdditionResultOutput() LookupUserToGroupAdditionResultOutput {
	return o
}

func (o LookupUserToGroupAdditionResultOutput) ToLookupUserToGroupAdditionResultOutputWithContext(ctx context.Context) LookupUserToGroupAdditionResultOutput {
	return o
}

func (o LookupUserToGroupAdditionResultOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserToGroupAdditionResult) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

func (o LookupUserToGroupAdditionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserToGroupAdditionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupUserToGroupAdditionResultOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserToGroupAdditionResult) []string { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserToGroupAdditionResultOutput{})
}
