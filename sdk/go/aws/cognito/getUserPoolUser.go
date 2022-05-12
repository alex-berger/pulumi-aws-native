// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolUser
func LookupUserPoolUser(ctx *pulumi.Context, args *LookupUserPoolUserArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolUserResult, error) {
	var rv LookupUserPoolUserResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolUserArgs struct {
	Id string `pulumi:"id"`
}

type LookupUserPoolUserResult struct {
	Id *string `pulumi:"id"`
}

func LookupUserPoolUserOutput(ctx *pulumi.Context, args LookupUserPoolUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolUserResult, error) {
			args := v.(LookupUserPoolUserArgs)
			r, err := LookupUserPoolUser(ctx, &args, opts...)
			var s LookupUserPoolUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserPoolUserResultOutput)
}

type LookupUserPoolUserOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserPoolUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolUserArgs)(nil)).Elem()
}

type LookupUserPoolUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolUserResult)(nil)).Elem()
}

func (o LookupUserPoolUserResultOutput) ToLookupUserPoolUserResultOutput() LookupUserPoolUserResultOutput {
	return o
}

func (o LookupUserPoolUserResultOutput) ToLookupUserPoolUserResultOutputWithContext(ctx context.Context) LookupUserPoolUserResultOutput {
	return o
}

func (o LookupUserPoolUserResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolUserResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolUserResultOutput{})
}
