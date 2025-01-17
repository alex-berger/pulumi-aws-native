// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::LakeFormation::DataLakeSettings
func LookupDataLakeSettings(ctx *pulumi.Context, args *LookupDataLakeSettingsArgs, opts ...pulumi.InvokeOption) (*LookupDataLakeSettingsResult, error) {
	var rv LookupDataLakeSettingsResult
	err := ctx.Invoke("aws-native:lakeformation:getDataLakeSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataLakeSettingsArgs struct {
	Id string `pulumi:"id"`
}

type LookupDataLakeSettingsResult struct {
	Admins                *DataLakeSettingsAdmins `pulumi:"admins"`
	Id                    *string                 `pulumi:"id"`
	TrustedResourceOwners []string                `pulumi:"trustedResourceOwners"`
}

func LookupDataLakeSettingsOutput(ctx *pulumi.Context, args LookupDataLakeSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupDataLakeSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataLakeSettingsResult, error) {
			args := v.(LookupDataLakeSettingsArgs)
			r, err := LookupDataLakeSettings(ctx, &args, opts...)
			var s LookupDataLakeSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataLakeSettingsResultOutput)
}

type LookupDataLakeSettingsOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDataLakeSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataLakeSettingsArgs)(nil)).Elem()
}

type LookupDataLakeSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupDataLakeSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataLakeSettingsResult)(nil)).Elem()
}

func (o LookupDataLakeSettingsResultOutput) ToLookupDataLakeSettingsResultOutput() LookupDataLakeSettingsResultOutput {
	return o
}

func (o LookupDataLakeSettingsResultOutput) ToLookupDataLakeSettingsResultOutputWithContext(ctx context.Context) LookupDataLakeSettingsResultOutput {
	return o
}

func (o LookupDataLakeSettingsResultOutput) Admins() DataLakeSettingsAdminsPtrOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) *DataLakeSettingsAdmins { return v.Admins }).(DataLakeSettingsAdminsPtrOutput)
}

func (o LookupDataLakeSettingsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDataLakeSettingsResultOutput) TrustedResourceOwners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) []string { return v.TrustedResourceOwners }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataLakeSettingsResultOutput{})
}
