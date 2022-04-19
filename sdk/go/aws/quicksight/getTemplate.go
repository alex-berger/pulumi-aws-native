// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Template Resource Type.
func LookupTemplate(ctx *pulumi.Context, args *LookupTemplateArgs, opts ...pulumi.InvokeOption) (*LookupTemplateResult, error) {
	var rv LookupTemplateResult
	err := ctx.Invoke("aws-native:quicksight:getTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateArgs struct {
	AwsAccountId string `pulumi:"awsAccountId"`
	TemplateId   string `pulumi:"templateId"`
}

type LookupTemplateResult struct {
	// <p>The Amazon Resource Name (ARN) of the template.</p>
	Arn *string `pulumi:"arn"`
	// <p>A display name for the template.</p>
	Name *string `pulumi:"name"`
	// <p>A list of resource permissions to be set on the template. </p>
	Permissions []TemplateResourcePermission `pulumi:"permissions"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
	Tags []TemplateTag `pulumi:"tags"`
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
	AwsAccountId pulumi.StringInput `pulumi:"awsAccountId"`
	TemplateId   pulumi.StringInput `pulumi:"templateId"`
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

// <p>The Amazon Resource Name (ARN) of the template.</p>
func (o LookupTemplateResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>A display name for the template.</p>
func (o LookupTemplateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// <p>A list of resource permissions to be set on the template. </p>
func (o LookupTemplateResultOutput) Permissions() TemplateResourcePermissionArrayOutput {
	return o.ApplyT(func(v LookupTemplateResult) []TemplateResourcePermission { return v.Permissions }).(TemplateResourcePermissionArrayOutput)
}

// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>
func (o LookupTemplateResultOutput) Tags() TemplateTagArrayOutput {
	return o.ApplyT(func(v LookupTemplateResult) []TemplateTag { return v.Tags }).(TemplateTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateResultOutput{})
}
