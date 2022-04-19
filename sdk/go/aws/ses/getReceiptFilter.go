// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SES::ReceiptFilter
func LookupReceiptFilter(ctx *pulumi.Context, args *LookupReceiptFilterArgs, opts ...pulumi.InvokeOption) (*LookupReceiptFilterResult, error) {
	var rv LookupReceiptFilterResult
	err := ctx.Invoke("aws-native:ses:getReceiptFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReceiptFilterArgs struct {
	Id string `pulumi:"id"`
}

type LookupReceiptFilterResult struct {
	Id *string `pulumi:"id"`
}

func LookupReceiptFilterOutput(ctx *pulumi.Context, args LookupReceiptFilterOutputArgs, opts ...pulumi.InvokeOption) LookupReceiptFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReceiptFilterResult, error) {
			args := v.(LookupReceiptFilterArgs)
			r, err := LookupReceiptFilter(ctx, &args, opts...)
			var s LookupReceiptFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReceiptFilterResultOutput)
}

type LookupReceiptFilterOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupReceiptFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReceiptFilterArgs)(nil)).Elem()
}

type LookupReceiptFilterResultOutput struct{ *pulumi.OutputState }

func (LookupReceiptFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReceiptFilterResult)(nil)).Elem()
}

func (o LookupReceiptFilterResultOutput) ToLookupReceiptFilterResultOutput() LookupReceiptFilterResultOutput {
	return o
}

func (o LookupReceiptFilterResultOutput) ToLookupReceiptFilterResultOutputWithContext(ctx context.Context) LookupReceiptFilterResultOutput {
	return o
}

func (o LookupReceiptFilterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReceiptFilterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReceiptFilterResultOutput{})
}
