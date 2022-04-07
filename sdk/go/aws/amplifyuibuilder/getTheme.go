// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package amplifyuibuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::AmplifyUIBuilder::Theme Resource Type
func LookupTheme(ctx *pulumi.Context, args *LookupThemeArgs, opts ...pulumi.InvokeOption) (*LookupThemeResult, error) {
	var rv LookupThemeResult
	err := ctx.Invoke("aws-native:amplifyuibuilder:getTheme", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThemeArgs struct {
	AppId           string `pulumi:"appId"`
	EnvironmentName string `pulumi:"environmentName"`
	Id              string `pulumi:"id"`
}

type LookupThemeResult struct {
	AppId           *string       `pulumi:"appId"`
	CreatedAt       *string       `pulumi:"createdAt"`
	EnvironmentName *string       `pulumi:"environmentName"`
	Id              *string       `pulumi:"id"`
	ModifiedAt      *string       `pulumi:"modifiedAt"`
	Name            *string       `pulumi:"name"`
	Overrides       []ThemeValues `pulumi:"overrides"`
	Values          []ThemeValues `pulumi:"values"`
}

func LookupThemeOutput(ctx *pulumi.Context, args LookupThemeOutputArgs, opts ...pulumi.InvokeOption) LookupThemeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupThemeResult, error) {
			args := v.(LookupThemeArgs)
			r, err := LookupTheme(ctx, &args, opts...)
			var s LookupThemeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupThemeResultOutput)
}

type LookupThemeOutputArgs struct {
	AppId           pulumi.StringInput `pulumi:"appId"`
	EnvironmentName pulumi.StringInput `pulumi:"environmentName"`
	Id              pulumi.StringInput `pulumi:"id"`
}

func (LookupThemeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThemeArgs)(nil)).Elem()
}

type LookupThemeResultOutput struct{ *pulumi.OutputState }

func (LookupThemeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThemeResult)(nil)).Elem()
}

func (o LookupThemeResultOutput) ToLookupThemeResultOutput() LookupThemeResultOutput {
	return o
}

func (o LookupThemeResultOutput) ToLookupThemeResultOutputWithContext(ctx context.Context) LookupThemeResultOutput {
	return o
}

func (o LookupThemeResultOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o LookupThemeResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupThemeResultOutput) EnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.EnvironmentName }).(pulumi.StringPtrOutput)
}

func (o LookupThemeResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupThemeResultOutput) ModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.ModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupThemeResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupThemeResultOutput) Overrides() ThemeValuesArrayOutput {
	return o.ApplyT(func(v LookupThemeResult) []ThemeValues { return v.Overrides }).(ThemeValuesArrayOutput)
}

func (o LookupThemeResultOutput) Values() ThemeValuesArrayOutput {
	return o.ApplyT(func(v LookupThemeResult) []ThemeValues { return v.Values }).(ThemeValuesArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThemeResultOutput{})
}
