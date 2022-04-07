// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::TrafficMirrorFilter
func LookupTrafficMirrorFilter(ctx *pulumi.Context, args *LookupTrafficMirrorFilterArgs, opts ...pulumi.InvokeOption) (*LookupTrafficMirrorFilterResult, error) {
	var rv LookupTrafficMirrorFilterResult
	err := ctx.Invoke("aws-native:ec2:getTrafficMirrorFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrafficMirrorFilterArgs struct {
	Id string `pulumi:"id"`
}

type LookupTrafficMirrorFilterResult struct {
	Id              *string                  `pulumi:"id"`
	NetworkServices []string                 `pulumi:"networkServices"`
	Tags            []TrafficMirrorFilterTag `pulumi:"tags"`
}

func LookupTrafficMirrorFilterOutput(ctx *pulumi.Context, args LookupTrafficMirrorFilterOutputArgs, opts ...pulumi.InvokeOption) LookupTrafficMirrorFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrafficMirrorFilterResult, error) {
			args := v.(LookupTrafficMirrorFilterArgs)
			r, err := LookupTrafficMirrorFilter(ctx, &args, opts...)
			var s LookupTrafficMirrorFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrafficMirrorFilterResultOutput)
}

type LookupTrafficMirrorFilterOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTrafficMirrorFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficMirrorFilterArgs)(nil)).Elem()
}

type LookupTrafficMirrorFilterResultOutput struct{ *pulumi.OutputState }

func (LookupTrafficMirrorFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficMirrorFilterResult)(nil)).Elem()
}

func (o LookupTrafficMirrorFilterResultOutput) ToLookupTrafficMirrorFilterResultOutput() LookupTrafficMirrorFilterResultOutput {
	return o
}

func (o LookupTrafficMirrorFilterResultOutput) ToLookupTrafficMirrorFilterResultOutputWithContext(ctx context.Context) LookupTrafficMirrorFilterResultOutput {
	return o
}

func (o LookupTrafficMirrorFilterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorFilterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTrafficMirrorFilterResultOutput) NetworkServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrafficMirrorFilterResult) []string { return v.NetworkServices }).(pulumi.StringArrayOutput)
}

func (o LookupTrafficMirrorFilterResultOutput) Tags() TrafficMirrorFilterTagArrayOutput {
	return o.ApplyT(func(v LookupTrafficMirrorFilterResult) []TrafficMirrorFilterTag { return v.Tags }).(TrafficMirrorFilterTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrafficMirrorFilterResultOutput{})
}
