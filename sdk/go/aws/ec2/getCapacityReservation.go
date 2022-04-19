// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::CapacityReservation
func LookupCapacityReservation(ctx *pulumi.Context, args *LookupCapacityReservationArgs, opts ...pulumi.InvokeOption) (*LookupCapacityReservationResult, error) {
	var rv LookupCapacityReservationResult
	err := ctx.Invoke("aws-native:ec2:getCapacityReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityReservationArgs struct {
	Id string `pulumi:"id"`
}

type LookupCapacityReservationResult struct {
	AvailableInstanceCount *int    `pulumi:"availableInstanceCount"`
	EndDate                *string `pulumi:"endDate"`
	EndDateType            *string `pulumi:"endDateType"`
	Id                     *string `pulumi:"id"`
	InstanceCount          *int    `pulumi:"instanceCount"`
	TotalInstanceCount     *int    `pulumi:"totalInstanceCount"`
}

func LookupCapacityReservationOutput(ctx *pulumi.Context, args LookupCapacityReservationOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityReservationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityReservationResult, error) {
			args := v.(LookupCapacityReservationArgs)
			r, err := LookupCapacityReservation(ctx, &args, opts...)
			var s LookupCapacityReservationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityReservationResultOutput)
}

type LookupCapacityReservationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCapacityReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationArgs)(nil)).Elem()
}

type LookupCapacityReservationResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationResult)(nil)).Elem()
}

func (o LookupCapacityReservationResultOutput) ToLookupCapacityReservationResultOutput() LookupCapacityReservationResultOutput {
	return o
}

func (o LookupCapacityReservationResultOutput) ToLookupCapacityReservationResultOutputWithContext(ctx context.Context) LookupCapacityReservationResultOutput {
	return o
}

func (o LookupCapacityReservationResultOutput) AvailableInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) *int { return v.AvailableInstanceCount }).(pulumi.IntPtrOutput)
}

func (o LookupCapacityReservationResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o LookupCapacityReservationResultOutput) EndDateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) *string { return v.EndDateType }).(pulumi.StringPtrOutput)
}

func (o LookupCapacityReservationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCapacityReservationResultOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

func (o LookupCapacityReservationResultOutput) TotalInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) *int { return v.TotalInstanceCount }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityReservationResultOutput{})
}
