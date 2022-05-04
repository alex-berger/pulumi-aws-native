// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func MapCfnPropsToResource(ctx *pulumi.Context, args *MapCfnPropsToResourceArgs, opts ...pulumi.InvokeOption) (*MapCfnPropsToResourceResult, error) {
	var rv MapCfnPropsToResourceResult
	err := ctx.Invoke("aws-native:index:mapCfnPropsToResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type MapCfnPropsToResourceArgs struct {
	Cf    string      `pulumi:"cf"`
	Props interface{} `pulumi:"props"`
}

type MapCfnPropsToResourceResult struct {
	IsSupported bool        `pulumi:"isSupported"`
	Outputs     []string    `pulumi:"outputs"`
	Props       interface{} `pulumi:"props"`
}

func MapCfnPropsToResourceOutput(ctx *pulumi.Context, args MapCfnPropsToResourceOutputArgs, opts ...pulumi.InvokeOption) MapCfnPropsToResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (MapCfnPropsToResourceResult, error) {
			args := v.(MapCfnPropsToResourceArgs)
			r, err := MapCfnPropsToResource(ctx, &args, opts...)
			var s MapCfnPropsToResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(MapCfnPropsToResourceResultOutput)
}

type MapCfnPropsToResourceOutputArgs struct {
	Cf    pulumi.StringInput `pulumi:"cf"`
	Props pulumi.Input       `pulumi:"props"`
}

func (MapCfnPropsToResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MapCfnPropsToResourceArgs)(nil)).Elem()
}

type MapCfnPropsToResourceResultOutput struct{ *pulumi.OutputState }

func (MapCfnPropsToResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MapCfnPropsToResourceResult)(nil)).Elem()
}

func (o MapCfnPropsToResourceResultOutput) ToMapCfnPropsToResourceResultOutput() MapCfnPropsToResourceResultOutput {
	return o
}

func (o MapCfnPropsToResourceResultOutput) ToMapCfnPropsToResourceResultOutputWithContext(ctx context.Context) MapCfnPropsToResourceResultOutput {
	return o
}

func (o MapCfnPropsToResourceResultOutput) IsSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v MapCfnPropsToResourceResult) bool { return v.IsSupported }).(pulumi.BoolOutput)
}

func (o MapCfnPropsToResourceResultOutput) Outputs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MapCfnPropsToResourceResult) []string { return v.Outputs }).(pulumi.StringArrayOutput)
}

func (o MapCfnPropsToResourceResultOutput) Props() pulumi.AnyOutput {
	return o.ApplyT(func(v MapCfnPropsToResourceResult) interface{} { return v.Props }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(MapCfnPropsToResourceResultOutput{})
}
