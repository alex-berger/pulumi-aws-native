// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DMS::ReplicationTask
func LookupReplicationTask(ctx *pulumi.Context, args *LookupReplicationTaskArgs, opts ...pulumi.InvokeOption) (*LookupReplicationTaskResult, error) {
	var rv LookupReplicationTaskResult
	err := ctx.Invoke("aws-native:dms:getReplicationTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationTaskArgs struct {
	Id string `pulumi:"id"`
}

type LookupReplicationTaskResult struct {
	CdcStartPosition          *string  `pulumi:"cdcStartPosition"`
	CdcStartTime              *float64 `pulumi:"cdcStartTime"`
	CdcStopPosition           *string  `pulumi:"cdcStopPosition"`
	Id                        *string  `pulumi:"id"`
	MigrationType             *string  `pulumi:"migrationType"`
	ReplicationTaskIdentifier *string  `pulumi:"replicationTaskIdentifier"`
	ReplicationTaskSettings   *string  `pulumi:"replicationTaskSettings"`
	TableMappings             *string  `pulumi:"tableMappings"`
	TaskData                  *string  `pulumi:"taskData"`
}

func LookupReplicationTaskOutput(ctx *pulumi.Context, args LookupReplicationTaskOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationTaskResult, error) {
			args := v.(LookupReplicationTaskArgs)
			r, err := LookupReplicationTask(ctx, &args, opts...)
			var s LookupReplicationTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationTaskResultOutput)
}

type LookupReplicationTaskOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupReplicationTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationTaskArgs)(nil)).Elem()
}

type LookupReplicationTaskResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationTaskResult)(nil)).Elem()
}

func (o LookupReplicationTaskResultOutput) ToLookupReplicationTaskResultOutput() LookupReplicationTaskResultOutput {
	return o
}

func (o LookupReplicationTaskResultOutput) ToLookupReplicationTaskResultOutputWithContext(ctx context.Context) LookupReplicationTaskResultOutput {
	return o
}

func (o LookupReplicationTaskResultOutput) CdcStartPosition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *string { return v.CdcStartPosition }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationTaskResultOutput) CdcStartTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *float64 { return v.CdcStartTime }).(pulumi.Float64PtrOutput)
}

func (o LookupReplicationTaskResultOutput) CdcStopPosition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *string { return v.CdcStopPosition }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationTaskResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationTaskResultOutput) MigrationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *string { return v.MigrationType }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationTaskResultOutput) ReplicationTaskIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *string { return v.ReplicationTaskIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationTaskResultOutput) ReplicationTaskSettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *string { return v.ReplicationTaskSettings }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationTaskResultOutput) TableMappings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *string { return v.TableMappings }).(pulumi.StringPtrOutput)
}

func (o LookupReplicationTaskResultOutput) TaskData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationTaskResult) *string { return v.TaskData }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationTaskResultOutput{})
}
