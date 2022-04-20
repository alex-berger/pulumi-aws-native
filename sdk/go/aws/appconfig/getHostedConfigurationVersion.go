// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::HostedConfigurationVersion
func LookupHostedConfigurationVersion(ctx *pulumi.Context, args *LookupHostedConfigurationVersionArgs, opts ...pulumi.InvokeOption) (*LookupHostedConfigurationVersionResult, error) {
	var rv LookupHostedConfigurationVersionResult
	err := ctx.Invoke("aws-native:appconfig:getHostedConfigurationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostedConfigurationVersionArgs struct {
	Id string `pulumi:"id"`
}

type LookupHostedConfigurationVersionResult struct {
	Id *string `pulumi:"id"`
}

func LookupHostedConfigurationVersionOutput(ctx *pulumi.Context, args LookupHostedConfigurationVersionOutputArgs, opts ...pulumi.InvokeOption) LookupHostedConfigurationVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostedConfigurationVersionResult, error) {
			args := v.(LookupHostedConfigurationVersionArgs)
			r, err := LookupHostedConfigurationVersion(ctx, &args, opts...)
			var s LookupHostedConfigurationVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostedConfigurationVersionResultOutput)
}

type LookupHostedConfigurationVersionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupHostedConfigurationVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostedConfigurationVersionArgs)(nil)).Elem()
}

type LookupHostedConfigurationVersionResultOutput struct{ *pulumi.OutputState }

func (LookupHostedConfigurationVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostedConfigurationVersionResult)(nil)).Elem()
}

func (o LookupHostedConfigurationVersionResultOutput) ToLookupHostedConfigurationVersionResultOutput() LookupHostedConfigurationVersionResultOutput {
	return o
}

func (o LookupHostedConfigurationVersionResultOutput) ToLookupHostedConfigurationVersionResultOutputWithContext(ctx context.Context) LookupHostedConfigurationVersionResultOutput {
	return o
}

func (o LookupHostedConfigurationVersionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostedConfigurationVersionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostedConfigurationVersionResultOutput{})
}
