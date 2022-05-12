// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ServiceCatalog::TagOptionAssociation
func LookupTagOptionAssociation(ctx *pulumi.Context, args *LookupTagOptionAssociationArgs, opts ...pulumi.InvokeOption) (*LookupTagOptionAssociationResult, error) {
	var rv LookupTagOptionAssociationResult
	err := ctx.Invoke("aws-native:servicecatalog:getTagOptionAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagOptionAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupTagOptionAssociationResult struct {
	Id *string `pulumi:"id"`
}

func LookupTagOptionAssociationOutput(ctx *pulumi.Context, args LookupTagOptionAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupTagOptionAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagOptionAssociationResult, error) {
			args := v.(LookupTagOptionAssociationArgs)
			r, err := LookupTagOptionAssociation(ctx, &args, opts...)
			var s LookupTagOptionAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagOptionAssociationResultOutput)
}

type LookupTagOptionAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTagOptionAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagOptionAssociationArgs)(nil)).Elem()
}

type LookupTagOptionAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupTagOptionAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagOptionAssociationResult)(nil)).Elem()
}

func (o LookupTagOptionAssociationResultOutput) ToLookupTagOptionAssociationResultOutput() LookupTagOptionAssociationResultOutput {
	return o
}

func (o LookupTagOptionAssociationResultOutput) ToLookupTagOptionAssociationResultOutputWithContext(ctx context.Context) LookupTagOptionAssociationResultOutput {
	return o
}

func (o LookupTagOptionAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTagOptionAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagOptionAssociationResultOutput{})
}
