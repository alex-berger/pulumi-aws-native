// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::DomainName
func LookupDomainName(ctx *pulumi.Context, args *LookupDomainNameArgs, opts ...pulumi.InvokeOption) (*LookupDomainNameResult, error) {
	var rv LookupDomainNameResult
	err := ctx.Invoke("aws-native:appsync:getDomainName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainNameArgs struct {
	DomainName string `pulumi:"domainName"`
}

type LookupDomainNameResult struct {
	AppSyncDomainName *string `pulumi:"appSyncDomainName"`
	Description       *string `pulumi:"description"`
	HostedZoneId      *string `pulumi:"hostedZoneId"`
}

func LookupDomainNameOutput(ctx *pulumi.Context, args LookupDomainNameOutputArgs, opts ...pulumi.InvokeOption) LookupDomainNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainNameResult, error) {
			args := v.(LookupDomainNameArgs)
			r, err := LookupDomainName(ctx, &args, opts...)
			var s LookupDomainNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainNameResultOutput)
}

type LookupDomainNameOutputArgs struct {
	DomainName pulumi.StringInput `pulumi:"domainName"`
}

func (LookupDomainNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainNameArgs)(nil)).Elem()
}

type LookupDomainNameResultOutput struct{ *pulumi.OutputState }

func (LookupDomainNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainNameResult)(nil)).Elem()
}

func (o LookupDomainNameResultOutput) ToLookupDomainNameResultOutput() LookupDomainNameResultOutput {
	return o
}

func (o LookupDomainNameResultOutput) ToLookupDomainNameResultOutputWithContext(ctx context.Context) LookupDomainNameResultOutput {
	return o
}

func (o LookupDomainNameResultOutput) AppSyncDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainNameResult) *string { return v.AppSyncDomainName }).(pulumi.StringPtrOutput)
}

func (o LookupDomainNameResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainNameResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupDomainNameResultOutput) HostedZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainNameResult) *string { return v.HostedZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainNameResultOutput{})
}
