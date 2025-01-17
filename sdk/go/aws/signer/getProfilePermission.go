// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupProfilePermission(ctx *pulumi.Context, args *LookupProfilePermissionArgs, opts ...pulumi.InvokeOption) (*LookupProfilePermissionResult, error) {
	var rv LookupProfilePermissionResult
	err := ctx.Invoke("aws-native:signer:getProfilePermission", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfilePermissionArgs struct {
	ProfileName string `pulumi:"profileName"`
	StatementId string `pulumi:"statementId"`
}

type LookupProfilePermissionResult struct {
}

func LookupProfilePermissionOutput(ctx *pulumi.Context, args LookupProfilePermissionOutputArgs, opts ...pulumi.InvokeOption) LookupProfilePermissionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfilePermissionResult, error) {
			args := v.(LookupProfilePermissionArgs)
			r, err := LookupProfilePermission(ctx, &args, opts...)
			var s LookupProfilePermissionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfilePermissionResultOutput)
}

type LookupProfilePermissionOutputArgs struct {
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	StatementId pulumi.StringInput `pulumi:"statementId"`
}

func (LookupProfilePermissionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfilePermissionArgs)(nil)).Elem()
}

type LookupProfilePermissionResultOutput struct{ *pulumi.OutputState }

func (LookupProfilePermissionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfilePermissionResult)(nil)).Elem()
}

func (o LookupProfilePermissionResultOutput) ToLookupProfilePermissionResultOutput() LookupProfilePermissionResultOutput {
	return o
}

func (o LookupProfilePermissionResultOutput) ToLookupProfilePermissionResultOutputWithContext(ctx context.Context) LookupProfilePermissionResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LookupProfilePermissionResultOutput{})
}
