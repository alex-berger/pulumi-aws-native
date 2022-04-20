// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalogappregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation
func LookupResourceAssociation(ctx *pulumi.Context, args *LookupResourceAssociationArgs, opts ...pulumi.InvokeOption) (*LookupResourceAssociationResult, error) {
	var rv LookupResourceAssociationResult
	err := ctx.Invoke("aws-native:servicecatalogappregistry:getResourceAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupResourceAssociationResult struct {
	// The name or the Id of the Application.
	Application    *string `pulumi:"application"`
	ApplicationArn *string `pulumi:"applicationArn"`
	Id             *string `pulumi:"id"`
	// The name or the Id of the Resource.
	Resource    *string `pulumi:"resource"`
	ResourceArn *string `pulumi:"resourceArn"`
	// The type of the CFN Resource for now it's enum CFN_STACK.
	ResourceType *ResourceAssociationResourceType `pulumi:"resourceType"`
}

func LookupResourceAssociationOutput(ctx *pulumi.Context, args LookupResourceAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupResourceAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceAssociationResult, error) {
			args := v.(LookupResourceAssociationArgs)
			r, err := LookupResourceAssociation(ctx, &args, opts...)
			var s LookupResourceAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceAssociationResultOutput)
}

type LookupResourceAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResourceAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceAssociationArgs)(nil)).Elem()
}

type LookupResourceAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupResourceAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceAssociationResult)(nil)).Elem()
}

func (o LookupResourceAssociationResultOutput) ToLookupResourceAssociationResultOutput() LookupResourceAssociationResultOutput {
	return o
}

func (o LookupResourceAssociationResultOutput) ToLookupResourceAssociationResultOutputWithContext(ctx context.Context) LookupResourceAssociationResultOutput {
	return o
}

// The name or the Id of the Application.
func (o LookupResourceAssociationResultOutput) Application() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceAssociationResult) *string { return v.Application }).(pulumi.StringPtrOutput)
}

func (o LookupResourceAssociationResultOutput) ApplicationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceAssociationResult) *string { return v.ApplicationArn }).(pulumi.StringPtrOutput)
}

func (o LookupResourceAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name or the Id of the Resource.
func (o LookupResourceAssociationResultOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceAssociationResult) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

func (o LookupResourceAssociationResultOutput) ResourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceAssociationResult) *string { return v.ResourceArn }).(pulumi.StringPtrOutput)
}

// The type of the CFN Resource for now it's enum CFN_STACK.
func (o LookupResourceAssociationResultOutput) ResourceType() ResourceAssociationResourceTypePtrOutput {
	return o.ApplyT(func(v LookupResourceAssociationResult) *ResourceAssociationResourceType { return v.ResourceType }).(ResourceAssociationResourceTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceAssociationResultOutput{})
}
