// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.
func LookupKeySigningKey(ctx *pulumi.Context, args *LookupKeySigningKeyArgs, opts ...pulumi.InvokeOption) (*LookupKeySigningKeyResult, error) {
	var rv LookupKeySigningKeyResult
	err := ctx.Invoke("aws-native:route53:getKeySigningKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKeySigningKeyArgs struct {
	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId string `pulumi:"hostedZoneId"`
	// An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
	Name string `pulumi:"name"`
}

type LookupKeySigningKeyResult struct {
	// A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
	Status *KeySigningKeyStatus `pulumi:"status"`
}

func LookupKeySigningKeyOutput(ctx *pulumi.Context, args LookupKeySigningKeyOutputArgs, opts ...pulumi.InvokeOption) LookupKeySigningKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKeySigningKeyResult, error) {
			args := v.(LookupKeySigningKeyArgs)
			r, err := LookupKeySigningKey(ctx, &args, opts...)
			var s LookupKeySigningKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKeySigningKeyResultOutput)
}

type LookupKeySigningKeyOutputArgs struct {
	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId pulumi.StringInput `pulumi:"hostedZoneId"`
	// An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupKeySigningKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeySigningKeyArgs)(nil)).Elem()
}

type LookupKeySigningKeyResultOutput struct{ *pulumi.OutputState }

func (LookupKeySigningKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeySigningKeyResult)(nil)).Elem()
}

func (o LookupKeySigningKeyResultOutput) ToLookupKeySigningKeyResultOutput() LookupKeySigningKeyResultOutput {
	return o
}

func (o LookupKeySigningKeyResultOutput) ToLookupKeySigningKeyResultOutputWithContext(ctx context.Context) LookupKeySigningKeyResultOutput {
	return o
}

// A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
func (o LookupKeySigningKeyResultOutput) Status() KeySigningKeyStatusPtrOutput {
	return o.ApplyT(func(v LookupKeySigningKeyResult) *KeySigningKeyStatus { return v.Status }).(KeySigningKeyStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeySigningKeyResultOutput{})
}
