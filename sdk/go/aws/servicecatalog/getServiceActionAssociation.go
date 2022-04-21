// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::ServiceCatalog::ServiceActionAssociation
func LookupServiceActionAssociation(ctx *pulumi.Context, args *LookupServiceActionAssociationArgs, opts ...pulumi.InvokeOption) (*LookupServiceActionAssociationResult, error) {
	var rv LookupServiceActionAssociationResult
	err := ctx.Invoke("aws-native:servicecatalog:getServiceActionAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceActionAssociationArgs struct {
	ProductId              string `pulumi:"productId"`
	ProvisioningArtifactId string `pulumi:"provisioningArtifactId"`
	ServiceActionId        string `pulumi:"serviceActionId"`
}

type LookupServiceActionAssociationResult struct {
}

func LookupServiceActionAssociationOutput(ctx *pulumi.Context, args LookupServiceActionAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupServiceActionAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceActionAssociationResult, error) {
			args := v.(LookupServiceActionAssociationArgs)
			r, err := LookupServiceActionAssociation(ctx, &args, opts...)
			var s LookupServiceActionAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceActionAssociationResultOutput)
}

type LookupServiceActionAssociationOutputArgs struct {
	ProductId              pulumi.StringInput `pulumi:"productId"`
	ProvisioningArtifactId pulumi.StringInput `pulumi:"provisioningArtifactId"`
	ServiceActionId        pulumi.StringInput `pulumi:"serviceActionId"`
}

func (LookupServiceActionAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceActionAssociationArgs)(nil)).Elem()
}

type LookupServiceActionAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupServiceActionAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceActionAssociationResult)(nil)).Elem()
}

func (o LookupServiceActionAssociationResultOutput) ToLookupServiceActionAssociationResultOutput() LookupServiceActionAssociationResultOutput {
	return o
}

func (o LookupServiceActionAssociationResultOutput) ToLookupServiceActionAssociationResultOutputWithContext(ctx context.Context) LookupServiceActionAssociationResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LookupServiceActionAssociationResultOutput{})
}
