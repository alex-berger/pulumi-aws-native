// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Theme Resource Type.
func LookupTheme(ctx *pulumi.Context, args *LookupThemeArgs, opts ...pulumi.InvokeOption) (*LookupThemeResult, error) {
	var rv LookupThemeResult
	err := ctx.Invoke("aws-native:quicksight:getTheme", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThemeArgs struct {
	AwsAccountId string `pulumi:"awsAccountId"`
	ThemeId      string `pulumi:"themeId"`
}

type LookupThemeResult struct {
	// <p>The Amazon Resource Name (ARN) of the theme.</p>
	Arn *string `pulumi:"arn"`
	// <p>The date and time that the theme was created.</p>
	CreatedTime *string `pulumi:"createdTime"`
	// <p>The date and time that the theme was last updated.</p>
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// <p>A display name for the theme.</p>
	Name *string `pulumi:"name"`
	// <p>A valid grouping of resource permissions to apply to the new theme.
	// 			</p>
	Permissions []ThemeResourcePermission `pulumi:"permissions"`
	// <p>A map of the key-value pairs for the resource tag or tags that you want to add to the
	// 			resource.</p>
	Tags    []ThemeTag    `pulumi:"tags"`
	Type    *ThemeType    `pulumi:"type"`
	Version *ThemeVersion `pulumi:"version"`
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
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
	ThemeId      pulumi.StringInput `pulumi:"themeId"`
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

// <p>The Amazon Resource Name (ARN) of the theme.</p>
func (o LookupThemeResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>The date and time that the theme was created.</p>
func (o LookupThemeResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// <p>The date and time that the theme was last updated.</p>
func (o LookupThemeResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// <p>A display name for the theme.</p>
func (o LookupThemeResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// <p>A valid grouping of resource permissions to apply to the new theme.
// 			</p>
func (o LookupThemeResultOutput) Permissions() ThemeResourcePermissionArrayOutput {
	return o.ApplyT(func(v LookupThemeResult) []ThemeResourcePermission { return v.Permissions }).(ThemeResourcePermissionArrayOutput)
}

// <p>A map of the key-value pairs for the resource tag or tags that you want to add to the
// 			resource.</p>
func (o LookupThemeResultOutput) Tags() ThemeTagArrayOutput {
	return o.ApplyT(func(v LookupThemeResult) []ThemeTag { return v.Tags }).(ThemeTagArrayOutput)
}

func (o LookupThemeResultOutput) Type() ThemeTypePtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *ThemeType { return v.Type }).(ThemeTypePtrOutput)
}

func (o LookupThemeResultOutput) Version() ThemeVersionPtrOutput {
	return o.ApplyT(func(v LookupThemeResult) *ThemeVersion { return v.Version }).(ThemeVersionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThemeResultOutput{})
}
