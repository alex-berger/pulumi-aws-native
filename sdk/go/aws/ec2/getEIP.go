// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::EIP
func LookupEIP(ctx *pulumi.Context, args *LookupEIPArgs, opts ...pulumi.InvokeOption) (*LookupEIPResult, error) {
	var rv LookupEIPResult
	err := ctx.Invoke("aws-native:ec2:getEIP", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEIPArgs struct {
	Id string `pulumi:"id"`
}

type LookupEIPResult struct {
	AllocationId   *string  `pulumi:"allocationId"`
	Id             *string  `pulumi:"id"`
	InstanceId     *string  `pulumi:"instanceId"`
	PublicIpv4Pool *string  `pulumi:"publicIpv4Pool"`
	Tags           []EIPTag `pulumi:"tags"`
}

func LookupEIPOutput(ctx *pulumi.Context, args LookupEIPOutputArgs, opts ...pulumi.InvokeOption) LookupEIPResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEIPResult, error) {
			args := v.(LookupEIPArgs)
			r, err := LookupEIP(ctx, &args, opts...)
			var s LookupEIPResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEIPResultOutput)
}

type LookupEIPOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEIPOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEIPArgs)(nil)).Elem()
}

type LookupEIPResultOutput struct{ *pulumi.OutputState }

func (LookupEIPResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEIPResult)(nil)).Elem()
}

func (o LookupEIPResultOutput) ToLookupEIPResultOutput() LookupEIPResultOutput {
	return o
}

func (o LookupEIPResultOutput) ToLookupEIPResultOutputWithContext(ctx context.Context) LookupEIPResultOutput {
	return o
}

func (o LookupEIPResultOutput) AllocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEIPResult) *string { return v.AllocationId }).(pulumi.StringPtrOutput)
}

func (o LookupEIPResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEIPResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEIPResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEIPResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o LookupEIPResultOutput) PublicIpv4Pool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEIPResult) *string { return v.PublicIpv4Pool }).(pulumi.StringPtrOutput)
}

func (o LookupEIPResultOutput) Tags() EIPTagArrayOutput {
	return o.ApplyT(func(v LookupEIPResult) []EIPTag { return v.Tags }).(EIPTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEIPResultOutput{})
}
