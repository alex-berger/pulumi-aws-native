// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GuardDuty::Master
func LookupMaster(ctx *pulumi.Context, args *LookupMasterArgs, opts ...pulumi.InvokeOption) (*LookupMasterResult, error) {
	var rv LookupMasterResult
	err := ctx.Invoke("aws-native:guardduty:getMaster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMasterArgs struct {
	MasterId string `pulumi:"masterId"`
}

type LookupMasterResult struct {
}

func LookupMasterOutput(ctx *pulumi.Context, args LookupMasterOutputArgs, opts ...pulumi.InvokeOption) LookupMasterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMasterResult, error) {
			args := v.(LookupMasterArgs)
			r, err := LookupMaster(ctx, &args, opts...)
			var s LookupMasterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMasterResultOutput)
}

type LookupMasterOutputArgs struct {
	MasterId pulumi.StringInput `pulumi:"masterId"`
}

func (LookupMasterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMasterArgs)(nil)).Elem()
}

type LookupMasterResultOutput struct{ *pulumi.OutputState }

func (LookupMasterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMasterResult)(nil)).Elem()
}

func (o LookupMasterResultOutput) ToLookupMasterResultOutput() LookupMasterResultOutput {
	return o
}

func (o LookupMasterResultOutput) ToLookupMasterResultOutputWithContext(ctx context.Context) LookupMasterResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LookupMasterResultOutput{})
}
