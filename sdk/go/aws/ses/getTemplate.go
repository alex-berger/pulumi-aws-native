// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SES::Template
func LookupTemplate(ctx *pulumi.Context, args *LookupTemplateArgs, opts ...pulumi.InvokeOption) (*LookupTemplateResult, error) {
	var rv LookupTemplateResult
	err := ctx.Invoke("aws-native:ses:getTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateArgs struct {
	Id string `pulumi:"id"`
}

type LookupTemplateResult struct {
	Id       *string       `pulumi:"id"`
	Template *TemplateType `pulumi:"template"`
}

func LookupTemplateOutput(ctx *pulumi.Context, args LookupTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplateResult, error) {
			args := v.(LookupTemplateArgs)
			r, err := LookupTemplate(ctx, &args, opts...)
			var s LookupTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplateResultOutput)
}

type LookupTemplateOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateArgs)(nil)).Elem()
}

type LookupTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateResult)(nil)).Elem()
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutput() LookupTemplateResultOutput {
	return o
}

func (o LookupTemplateResultOutput) ToLookupTemplateResultOutputWithContext(ctx context.Context) LookupTemplateResultOutput {
	return o
}

func (o LookupTemplateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateResultOutput) Template() TemplateTypePtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *TemplateType { return v.Template }).(TemplateTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateResultOutput{})
}
