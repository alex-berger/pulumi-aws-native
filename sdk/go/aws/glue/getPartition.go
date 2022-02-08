// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::Partition
func LookupPartition(ctx *pulumi.Context, args *LookupPartitionArgs, opts ...pulumi.InvokeOption) (*LookupPartitionResult, error) {
	var rv LookupPartitionResult
	err := ctx.Invoke("aws-native:glue:getPartition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartitionArgs struct {
	Id string `pulumi:"id"`
}

type LookupPartitionResult struct {
	Id             *string             `pulumi:"id"`
	PartitionInput *PartitionInputType `pulumi:"partitionInput"`
}

func LookupPartitionOutput(ctx *pulumi.Context, args LookupPartitionOutputArgs, opts ...pulumi.InvokeOption) LookupPartitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartitionResult, error) {
			args := v.(LookupPartitionArgs)
			r, err := LookupPartition(ctx, &args, opts...)
			return *r, err
		}).(LookupPartitionResultOutput)
}

type LookupPartitionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPartitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartitionArgs)(nil)).Elem()
}

type LookupPartitionResultOutput struct{ *pulumi.OutputState }

func (LookupPartitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartitionResult)(nil)).Elem()
}

func (o LookupPartitionResultOutput) ToLookupPartitionResultOutput() LookupPartitionResultOutput {
	return o
}

func (o LookupPartitionResultOutput) ToLookupPartitionResultOutputWithContext(ctx context.Context) LookupPartitionResultOutput {
	return o
}

func (o LookupPartitionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartitionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPartitionResultOutput) PartitionInput() PartitionInputTypePtrOutput {
	return o.ApplyT(func(v LookupPartitionResult) *PartitionInputType { return v.PartitionInput }).(PartitionInputTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartitionResultOutput{})
}