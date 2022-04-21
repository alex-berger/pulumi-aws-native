// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directoryservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MicrosoftADVpcSettings struct {
	SubnetIds []string `pulumi:"subnetIds"`
	VpcId     string   `pulumi:"vpcId"`
}

// MicrosoftADVpcSettingsInput is an input type that accepts MicrosoftADVpcSettingsArgs and MicrosoftADVpcSettingsOutput values.
// You can construct a concrete instance of `MicrosoftADVpcSettingsInput` via:
//
//          MicrosoftADVpcSettingsArgs{...}
type MicrosoftADVpcSettingsInput interface {
	pulumi.Input

	ToMicrosoftADVpcSettingsOutput() MicrosoftADVpcSettingsOutput
	ToMicrosoftADVpcSettingsOutputWithContext(context.Context) MicrosoftADVpcSettingsOutput
}

type MicrosoftADVpcSettingsArgs struct {
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
	VpcId     pulumi.StringInput      `pulumi:"vpcId"`
}

func (MicrosoftADVpcSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MicrosoftADVpcSettings)(nil)).Elem()
}

func (i MicrosoftADVpcSettingsArgs) ToMicrosoftADVpcSettingsOutput() MicrosoftADVpcSettingsOutput {
	return i.ToMicrosoftADVpcSettingsOutputWithContext(context.Background())
}

func (i MicrosoftADVpcSettingsArgs) ToMicrosoftADVpcSettingsOutputWithContext(ctx context.Context) MicrosoftADVpcSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MicrosoftADVpcSettingsOutput)
}

type MicrosoftADVpcSettingsOutput struct{ *pulumi.OutputState }

func (MicrosoftADVpcSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MicrosoftADVpcSettings)(nil)).Elem()
}

func (o MicrosoftADVpcSettingsOutput) ToMicrosoftADVpcSettingsOutput() MicrosoftADVpcSettingsOutput {
	return o
}

func (o MicrosoftADVpcSettingsOutput) ToMicrosoftADVpcSettingsOutputWithContext(ctx context.Context) MicrosoftADVpcSettingsOutput {
	return o
}

func (o MicrosoftADVpcSettingsOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MicrosoftADVpcSettings) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o MicrosoftADVpcSettingsOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v MicrosoftADVpcSettings) string { return v.VpcId }).(pulumi.StringOutput)
}

type SimpleADVpcSettings struct {
	SubnetIds []string `pulumi:"subnetIds"`
	VpcId     string   `pulumi:"vpcId"`
}

// SimpleADVpcSettingsInput is an input type that accepts SimpleADVpcSettingsArgs and SimpleADVpcSettingsOutput values.
// You can construct a concrete instance of `SimpleADVpcSettingsInput` via:
//
//          SimpleADVpcSettingsArgs{...}
type SimpleADVpcSettingsInput interface {
	pulumi.Input

	ToSimpleADVpcSettingsOutput() SimpleADVpcSettingsOutput
	ToSimpleADVpcSettingsOutputWithContext(context.Context) SimpleADVpcSettingsOutput
}

type SimpleADVpcSettingsArgs struct {
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
	VpcId     pulumi.StringInput      `pulumi:"vpcId"`
}

func (SimpleADVpcSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleADVpcSettings)(nil)).Elem()
}

func (i SimpleADVpcSettingsArgs) ToSimpleADVpcSettingsOutput() SimpleADVpcSettingsOutput {
	return i.ToSimpleADVpcSettingsOutputWithContext(context.Background())
}

func (i SimpleADVpcSettingsArgs) ToSimpleADVpcSettingsOutputWithContext(ctx context.Context) SimpleADVpcSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleADVpcSettingsOutput)
}

type SimpleADVpcSettingsOutput struct{ *pulumi.OutputState }

func (SimpleADVpcSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleADVpcSettings)(nil)).Elem()
}

func (o SimpleADVpcSettingsOutput) ToSimpleADVpcSettingsOutput() SimpleADVpcSettingsOutput {
	return o
}

func (o SimpleADVpcSettingsOutput) ToSimpleADVpcSettingsOutputWithContext(ctx context.Context) SimpleADVpcSettingsOutput {
	return o
}

func (o SimpleADVpcSettingsOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SimpleADVpcSettings) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o SimpleADVpcSettingsOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v SimpleADVpcSettings) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MicrosoftADVpcSettingsInput)(nil)).Elem(), MicrosoftADVpcSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SimpleADVpcSettingsInput)(nil)).Elem(), SimpleADVpcSettingsArgs{})
	pulumi.RegisterOutputType(MicrosoftADVpcSettingsOutput{})
	pulumi.RegisterOutputType(SimpleADVpcSettingsOutput{})
}
