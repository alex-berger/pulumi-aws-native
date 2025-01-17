// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::DeviceFarm::VPCEConfiguration creates a new Device Farm VPCE Configuration
func LookupVPCEConfiguration(ctx *pulumi.Context, args *LookupVPCEConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupVPCEConfigurationResult, error) {
	var rv LookupVPCEConfigurationResult
	err := ctx.Invoke("aws-native:devicefarm:getVPCEConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVPCEConfigurationArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupVPCEConfigurationResult struct {
	Arn                          *string                `pulumi:"arn"`
	ServiceDnsName               *string                `pulumi:"serviceDnsName"`
	Tags                         []VPCEConfigurationTag `pulumi:"tags"`
	VpceConfigurationDescription *string                `pulumi:"vpceConfigurationDescription"`
	VpceConfigurationName        *string                `pulumi:"vpceConfigurationName"`
	VpceServiceName              *string                `pulumi:"vpceServiceName"`
}

func LookupVPCEConfigurationOutput(ctx *pulumi.Context, args LookupVPCEConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupVPCEConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVPCEConfigurationResult, error) {
			args := v.(LookupVPCEConfigurationArgs)
			r, err := LookupVPCEConfiguration(ctx, &args, opts...)
			var s LookupVPCEConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVPCEConfigurationResultOutput)
}

type LookupVPCEConfigurationOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupVPCEConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCEConfigurationArgs)(nil)).Elem()
}

type LookupVPCEConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupVPCEConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVPCEConfigurationResult)(nil)).Elem()
}

func (o LookupVPCEConfigurationResultOutput) ToLookupVPCEConfigurationResultOutput() LookupVPCEConfigurationResultOutput {
	return o
}

func (o LookupVPCEConfigurationResultOutput) ToLookupVPCEConfigurationResultOutputWithContext(ctx context.Context) LookupVPCEConfigurationResultOutput {
	return o
}

func (o LookupVPCEConfigurationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEConfigurationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupVPCEConfigurationResultOutput) ServiceDnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEConfigurationResult) *string { return v.ServiceDnsName }).(pulumi.StringPtrOutput)
}

func (o LookupVPCEConfigurationResultOutput) Tags() VPCEConfigurationTagArrayOutput {
	return o.ApplyT(func(v LookupVPCEConfigurationResult) []VPCEConfigurationTag { return v.Tags }).(VPCEConfigurationTagArrayOutput)
}

func (o LookupVPCEConfigurationResultOutput) VpceConfigurationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEConfigurationResult) *string { return v.VpceConfigurationDescription }).(pulumi.StringPtrOutput)
}

func (o LookupVPCEConfigurationResultOutput) VpceConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEConfigurationResult) *string { return v.VpceConfigurationName }).(pulumi.StringPtrOutput)
}

func (o LookupVPCEConfigurationResultOutput) VpceServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVPCEConfigurationResult) *string { return v.VpceServiceName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVPCEConfigurationResultOutput{})
}
