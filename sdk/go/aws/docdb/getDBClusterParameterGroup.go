// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DocDB::DBClusterParameterGroup
func LookupDBClusterParameterGroup(ctx *pulumi.Context, args *LookupDBClusterParameterGroupArgs, opts ...pulumi.InvokeOption) (*LookupDBClusterParameterGroupResult, error) {
	var rv LookupDBClusterParameterGroupResult
	err := ctx.Invoke("aws-native:docdb:getDBClusterParameterGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDBClusterParameterGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupDBClusterParameterGroupResult struct {
	Id         *string                      `pulumi:"id"`
	Parameters interface{}                  `pulumi:"parameters"`
	Tags       []DBClusterParameterGroupTag `pulumi:"tags"`
}

func LookupDBClusterParameterGroupOutput(ctx *pulumi.Context, args LookupDBClusterParameterGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDBClusterParameterGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDBClusterParameterGroupResult, error) {
			args := v.(LookupDBClusterParameterGroupArgs)
			r, err := LookupDBClusterParameterGroup(ctx, &args, opts...)
			var s LookupDBClusterParameterGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDBClusterParameterGroupResultOutput)
}

type LookupDBClusterParameterGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDBClusterParameterGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBClusterParameterGroupArgs)(nil)).Elem()
}

type LookupDBClusterParameterGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDBClusterParameterGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBClusterParameterGroupResult)(nil)).Elem()
}

func (o LookupDBClusterParameterGroupResultOutput) ToLookupDBClusterParameterGroupResultOutput() LookupDBClusterParameterGroupResultOutput {
	return o
}

func (o LookupDBClusterParameterGroupResultOutput) ToLookupDBClusterParameterGroupResultOutputWithContext(ctx context.Context) LookupDBClusterParameterGroupResultOutput {
	return o
}

func (o LookupDBClusterParameterGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBClusterParameterGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDBClusterParameterGroupResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDBClusterParameterGroupResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupDBClusterParameterGroupResultOutput) Tags() DBClusterParameterGroupTagArrayOutput {
	return o.ApplyT(func(v LookupDBClusterParameterGroupResult) []DBClusterParameterGroupTag { return v.Tags }).(DBClusterParameterGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDBClusterParameterGroupResultOutput{})
}
