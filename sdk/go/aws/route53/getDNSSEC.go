// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource used to control (enable/disable) DNSSEC in a specific hosted zone.
func LookupDNSSEC(ctx *pulumi.Context, args *LookupDNSSECArgs, opts ...pulumi.InvokeOption) (*LookupDNSSECResult, error) {
	var rv LookupDNSSECResult
	err := ctx.Invoke("aws-native:route53:getDNSSEC", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDNSSECArgs struct {
	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId string `pulumi:"hostedZoneId"`
}

type LookupDNSSECResult struct {
}

func LookupDNSSECOutput(ctx *pulumi.Context, args LookupDNSSECOutputArgs, opts ...pulumi.InvokeOption) LookupDNSSECResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDNSSECResult, error) {
			args := v.(LookupDNSSECArgs)
			r, err := LookupDNSSEC(ctx, &args, opts...)
			var s LookupDNSSECResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDNSSECResultOutput)
}

type LookupDNSSECOutputArgs struct {
	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId pulumi.StringInput `pulumi:"hostedZoneId"`
}

func (LookupDNSSECOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDNSSECArgs)(nil)).Elem()
}

type LookupDNSSECResultOutput struct{ *pulumi.OutputState }

func (LookupDNSSECResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDNSSECResult)(nil)).Elem()
}

func (o LookupDNSSECResultOutput) ToLookupDNSSECResultOutput() LookupDNSSECResultOutput {
	return o
}

func (o LookupDNSSECResultOutput) ToLookupDNSSECResultOutputWithContext(ctx context.Context) LookupDNSSECResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LookupDNSSECResultOutput{})
}
