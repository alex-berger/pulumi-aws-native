// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::S3::MultiRegionAccessPoint is an Amazon S3 resource type that dynamically routes S3 requests to easily satisfy geographic compliance requirements based on customer-defined routing policies.
func LookupMultiRegionAccessPoint(ctx *pulumi.Context, args *LookupMultiRegionAccessPointArgs, opts ...pulumi.InvokeOption) (*LookupMultiRegionAccessPointResult, error) {
	var rv LookupMultiRegionAccessPointResult
	err := ctx.Invoke("aws-native:s3:getMultiRegionAccessPoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMultiRegionAccessPointArgs struct {
	// The name you want to assign to this Multi Region Access Point.
	Name string `pulumi:"name"`
}

type LookupMultiRegionAccessPointResult struct {
	// The alias is a unique identifier to, and is part of the public DNS name for this Multi Region Access Point
	Alias *string `pulumi:"alias"`
	// The timestamp of the when the Multi Region Access Point is created
	CreatedAt *string `pulumi:"createdAt"`
}

func LookupMultiRegionAccessPointOutput(ctx *pulumi.Context, args LookupMultiRegionAccessPointOutputArgs, opts ...pulumi.InvokeOption) LookupMultiRegionAccessPointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMultiRegionAccessPointResult, error) {
			args := v.(LookupMultiRegionAccessPointArgs)
			r, err := LookupMultiRegionAccessPoint(ctx, &args, opts...)
			var s LookupMultiRegionAccessPointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMultiRegionAccessPointResultOutput)
}

type LookupMultiRegionAccessPointOutputArgs struct {
	// The name you want to assign to this Multi Region Access Point.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupMultiRegionAccessPointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMultiRegionAccessPointArgs)(nil)).Elem()
}

type LookupMultiRegionAccessPointResultOutput struct{ *pulumi.OutputState }

func (LookupMultiRegionAccessPointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMultiRegionAccessPointResult)(nil)).Elem()
}

func (o LookupMultiRegionAccessPointResultOutput) ToLookupMultiRegionAccessPointResultOutput() LookupMultiRegionAccessPointResultOutput {
	return o
}

func (o LookupMultiRegionAccessPointResultOutput) ToLookupMultiRegionAccessPointResultOutputWithContext(ctx context.Context) LookupMultiRegionAccessPointResultOutput {
	return o
}

// The alias is a unique identifier to, and is part of the public DNS name for this Multi Region Access Point
func (o LookupMultiRegionAccessPointResultOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMultiRegionAccessPointResult) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

// The timestamp of the when the Multi Region Access Point is created
func (o LookupMultiRegionAccessPointResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMultiRegionAccessPointResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMultiRegionAccessPointResultOutput{})
}
