// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DirectoryService::SimpleAD
func LookupSimpleAD(ctx *pulumi.Context, args *LookupSimpleADArgs, opts ...pulumi.InvokeOption) (*LookupSimpleADResult, error) {
	var rv LookupSimpleADResult
	err := ctx.Invoke("aws-native:directoryservice:getSimpleAD", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSimpleADArgs struct {
	Id string `pulumi:"id"`
}

type LookupSimpleADResult struct {
	Alias          *string  `pulumi:"alias"`
	DnsIpAddresses []string `pulumi:"dnsIpAddresses"`
	EnableSso      *bool    `pulumi:"enableSso"`
	Id             *string  `pulumi:"id"`
}

func LookupSimpleADOutput(ctx *pulumi.Context, args LookupSimpleADOutputArgs, opts ...pulumi.InvokeOption) LookupSimpleADResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimpleADResult, error) {
			args := v.(LookupSimpleADArgs)
			r, err := LookupSimpleAD(ctx, &args, opts...)
			var s LookupSimpleADResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSimpleADResultOutput)
}

type LookupSimpleADOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSimpleADOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimpleADArgs)(nil)).Elem()
}

type LookupSimpleADResultOutput struct{ *pulumi.OutputState }

func (LookupSimpleADResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimpleADResult)(nil)).Elem()
}

func (o LookupSimpleADResultOutput) ToLookupSimpleADResultOutput() LookupSimpleADResultOutput {
	return o
}

func (o LookupSimpleADResultOutput) ToLookupSimpleADResultOutputWithContext(ctx context.Context) LookupSimpleADResultOutput {
	return o
}

func (o LookupSimpleADResultOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimpleADResult) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o LookupSimpleADResultOutput) DnsIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSimpleADResult) []string { return v.DnsIpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupSimpleADResultOutput) EnableSso() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSimpleADResult) *bool { return v.EnableSso }).(pulumi.BoolPtrOutput)
}

func (o LookupSimpleADResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimpleADResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimpleADResultOutput{})
}
