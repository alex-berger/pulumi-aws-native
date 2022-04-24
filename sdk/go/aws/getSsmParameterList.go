// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSsmParameterList(ctx *pulumi.Context, args *GetSsmParameterListArgs, opts ...pulumi.InvokeOption) (*GetSsmParameterListResult, error) {
	var rv GetSsmParameterListResult
	err := ctx.Invoke("aws-native:index:getSsmParameterList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSsmParameterListArgs struct {
	Name string `pulumi:"name"`
}

type GetSsmParameterListResult struct {
	Value []string `pulumi:"value"`
}

func GetSsmParameterListOutput(ctx *pulumi.Context, args GetSsmParameterListOutputArgs, opts ...pulumi.InvokeOption) GetSsmParameterListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSsmParameterListResult, error) {
			args := v.(GetSsmParameterListArgs)
			r, err := GetSsmParameterList(ctx, &args, opts...)
			var s GetSsmParameterListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSsmParameterListResultOutput)
}

type GetSsmParameterListOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetSsmParameterListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSsmParameterListArgs)(nil)).Elem()
}

type GetSsmParameterListResultOutput struct{ *pulumi.OutputState }

func (GetSsmParameterListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSsmParameterListResult)(nil)).Elem()
}

func (o GetSsmParameterListResultOutput) ToGetSsmParameterListResultOutput() GetSsmParameterListResultOutput {
	return o
}

func (o GetSsmParameterListResultOutput) ToGetSsmParameterListResultOutputWithContext(ctx context.Context) GetSsmParameterListResultOutput {
	return o
}

func (o GetSsmParameterListResultOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSsmParameterListResult) []string { return v.Value }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSsmParameterListResultOutput{})
}
