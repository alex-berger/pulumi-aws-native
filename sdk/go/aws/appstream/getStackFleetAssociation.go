// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::StackFleetAssociation
func LookupStackFleetAssociation(ctx *pulumi.Context, args *LookupStackFleetAssociationArgs, opts ...pulumi.InvokeOption) (*LookupStackFleetAssociationResult, error) {
	var rv LookupStackFleetAssociationResult
	err := ctx.Invoke("aws-native:appstream:getStackFleetAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStackFleetAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupStackFleetAssociationResult struct {
	FleetName *string `pulumi:"fleetName"`
	Id        *string `pulumi:"id"`
	StackName *string `pulumi:"stackName"`
}

func LookupStackFleetAssociationOutput(ctx *pulumi.Context, args LookupStackFleetAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupStackFleetAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStackFleetAssociationResult, error) {
			args := v.(LookupStackFleetAssociationArgs)
			r, err := LookupStackFleetAssociation(ctx, &args, opts...)
			var s LookupStackFleetAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStackFleetAssociationResultOutput)
}

type LookupStackFleetAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupStackFleetAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackFleetAssociationArgs)(nil)).Elem()
}

type LookupStackFleetAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupStackFleetAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackFleetAssociationResult)(nil)).Elem()
}

func (o LookupStackFleetAssociationResultOutput) ToLookupStackFleetAssociationResultOutput() LookupStackFleetAssociationResultOutput {
	return o
}

func (o LookupStackFleetAssociationResultOutput) ToLookupStackFleetAssociationResultOutputWithContext(ctx context.Context) LookupStackFleetAssociationResultOutput {
	return o
}

func (o LookupStackFleetAssociationResultOutput) FleetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackFleetAssociationResult) *string { return v.FleetName }).(pulumi.StringPtrOutput)
}

func (o LookupStackFleetAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackFleetAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupStackFleetAssociationResultOutput) StackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackFleetAssociationResult) *string { return v.StackName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackFleetAssociationResultOutput{})
}
