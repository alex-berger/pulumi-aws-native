// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::RDS::DBProxyTargetGroup
func LookupDBProxyTargetGroup(ctx *pulumi.Context, args *LookupDBProxyTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupDBProxyTargetGroupResult, error) {
	var rv LookupDBProxyTargetGroupResult
	err := ctx.Invoke("aws-native:rds:getDBProxyTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDBProxyTargetGroupArgs struct {
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn string `pulumi:"targetGroupArn"`
}

type LookupDBProxyTargetGroupResult struct {
	ConnectionPoolConfigurationInfo *DBProxyTargetGroupConnectionPoolConfigurationInfoFormat `pulumi:"connectionPoolConfigurationInfo"`
	DBClusterIdentifiers            []string                                                 `pulumi:"dBClusterIdentifiers"`
	DBInstanceIdentifiers           []string                                                 `pulumi:"dBInstanceIdentifiers"`
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn *string `pulumi:"targetGroupArn"`
}

func LookupDBProxyTargetGroupOutput(ctx *pulumi.Context, args LookupDBProxyTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDBProxyTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDBProxyTargetGroupResult, error) {
			args := v.(LookupDBProxyTargetGroupArgs)
			r, err := LookupDBProxyTargetGroup(ctx, &args, opts...)
			var s LookupDBProxyTargetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDBProxyTargetGroupResultOutput)
}

type LookupDBProxyTargetGroupOutputArgs struct {
	// The Amazon Resource Name (ARN) representing the target group.
	TargetGroupArn pulumi.StringInput `pulumi:"targetGroupArn"`
}

func (LookupDBProxyTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBProxyTargetGroupArgs)(nil)).Elem()
}

type LookupDBProxyTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDBProxyTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDBProxyTargetGroupResult)(nil)).Elem()
}

func (o LookupDBProxyTargetGroupResultOutput) ToLookupDBProxyTargetGroupResultOutput() LookupDBProxyTargetGroupResultOutput {
	return o
}

func (o LookupDBProxyTargetGroupResultOutput) ToLookupDBProxyTargetGroupResultOutputWithContext(ctx context.Context) LookupDBProxyTargetGroupResultOutput {
	return o
}

func (o LookupDBProxyTargetGroupResultOutput) ConnectionPoolConfigurationInfo() DBProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput {
	return o.ApplyT(func(v LookupDBProxyTargetGroupResult) *DBProxyTargetGroupConnectionPoolConfigurationInfoFormat {
		return v.ConnectionPoolConfigurationInfo
	}).(DBProxyTargetGroupConnectionPoolConfigurationInfoFormatPtrOutput)
}

func (o LookupDBProxyTargetGroupResultOutput) DBClusterIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDBProxyTargetGroupResult) []string { return v.DBClusterIdentifiers }).(pulumi.StringArrayOutput)
}

func (o LookupDBProxyTargetGroupResultOutput) DBInstanceIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDBProxyTargetGroupResult) []string { return v.DBInstanceIdentifiers }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) representing the target group.
func (o LookupDBProxyTargetGroupResultOutput) TargetGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDBProxyTargetGroupResult) *string { return v.TargetGroupArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDBProxyTargetGroupResultOutput{})
}
