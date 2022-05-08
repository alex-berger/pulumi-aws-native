// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SigningProfilePlatformId string

const (
	SigningProfilePlatformIdAWSLambdaSHA384ECDSA = SigningProfilePlatformId("AWSLambda-SHA384-ECDSA")
)

func (SigningProfilePlatformId) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningProfilePlatformId)(nil)).Elem()
}

func (e SigningProfilePlatformId) ToSigningProfilePlatformIdOutput() SigningProfilePlatformIdOutput {
	return pulumi.ToOutput(e).(SigningProfilePlatformIdOutput)
}

func (e SigningProfilePlatformId) ToSigningProfilePlatformIdOutputWithContext(ctx context.Context) SigningProfilePlatformIdOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SigningProfilePlatformIdOutput)
}

func (e SigningProfilePlatformId) ToSigningProfilePlatformIdPtrOutput() SigningProfilePlatformIdPtrOutput {
	return e.ToSigningProfilePlatformIdPtrOutputWithContext(context.Background())
}

func (e SigningProfilePlatformId) ToSigningProfilePlatformIdPtrOutputWithContext(ctx context.Context) SigningProfilePlatformIdPtrOutput {
	return SigningProfilePlatformId(e).ToSigningProfilePlatformIdOutputWithContext(ctx).ToSigningProfilePlatformIdPtrOutputWithContext(ctx)
}

func (e SigningProfilePlatformId) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SigningProfilePlatformId) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SigningProfilePlatformId) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SigningProfilePlatformId) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SigningProfilePlatformIdOutput struct{ *pulumi.OutputState }

func (SigningProfilePlatformIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningProfilePlatformId)(nil)).Elem()
}

func (o SigningProfilePlatformIdOutput) ToSigningProfilePlatformIdOutput() SigningProfilePlatformIdOutput {
	return o
}

func (o SigningProfilePlatformIdOutput) ToSigningProfilePlatformIdOutputWithContext(ctx context.Context) SigningProfilePlatformIdOutput {
	return o
}

func (o SigningProfilePlatformIdOutput) ToSigningProfilePlatformIdPtrOutput() SigningProfilePlatformIdPtrOutput {
	return o.ToSigningProfilePlatformIdPtrOutputWithContext(context.Background())
}

func (o SigningProfilePlatformIdOutput) ToSigningProfilePlatformIdPtrOutputWithContext(ctx context.Context) SigningProfilePlatformIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SigningProfilePlatformId) *SigningProfilePlatformId {
		return &v
	}).(SigningProfilePlatformIdPtrOutput)
}

func (o SigningProfilePlatformIdOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SigningProfilePlatformIdOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SigningProfilePlatformId) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SigningProfilePlatformIdOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SigningProfilePlatformIdOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SigningProfilePlatformId) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SigningProfilePlatformIdPtrOutput struct{ *pulumi.OutputState }

func (SigningProfilePlatformIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningProfilePlatformId)(nil)).Elem()
}

func (o SigningProfilePlatformIdPtrOutput) ToSigningProfilePlatformIdPtrOutput() SigningProfilePlatformIdPtrOutput {
	return o
}

func (o SigningProfilePlatformIdPtrOutput) ToSigningProfilePlatformIdPtrOutputWithContext(ctx context.Context) SigningProfilePlatformIdPtrOutput {
	return o
}

func (o SigningProfilePlatformIdPtrOutput) Elem() SigningProfilePlatformIdOutput {
	return o.ApplyT(func(v *SigningProfilePlatformId) SigningProfilePlatformId {
		if v != nil {
			return *v
		}
		var ret SigningProfilePlatformId
		return ret
	}).(SigningProfilePlatformIdOutput)
}

func (o SigningProfilePlatformIdPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SigningProfilePlatformIdPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SigningProfilePlatformId) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SigningProfilePlatformIdInput is an input type that accepts SigningProfilePlatformIdArgs and SigningProfilePlatformIdOutput values.
// You can construct a concrete instance of `SigningProfilePlatformIdInput` via:
//
//          SigningProfilePlatformIdArgs{...}
type SigningProfilePlatformIdInput interface {
	pulumi.Input

	ToSigningProfilePlatformIdOutput() SigningProfilePlatformIdOutput
	ToSigningProfilePlatformIdOutputWithContext(context.Context) SigningProfilePlatformIdOutput
}

var signingProfilePlatformIdPtrType = reflect.TypeOf((**SigningProfilePlatformId)(nil)).Elem()

type SigningProfilePlatformIdPtrInput interface {
	pulumi.Input

	ToSigningProfilePlatformIdPtrOutput() SigningProfilePlatformIdPtrOutput
	ToSigningProfilePlatformIdPtrOutputWithContext(context.Context) SigningProfilePlatformIdPtrOutput
}

type signingProfilePlatformIdPtr string

func SigningProfilePlatformIdPtr(v string) SigningProfilePlatformIdPtrInput {
	return (*signingProfilePlatformIdPtr)(&v)
}

func (*signingProfilePlatformIdPtr) ElementType() reflect.Type {
	return signingProfilePlatformIdPtrType
}

func (in *signingProfilePlatformIdPtr) ToSigningProfilePlatformIdPtrOutput() SigningProfilePlatformIdPtrOutput {
	return pulumi.ToOutput(in).(SigningProfilePlatformIdPtrOutput)
}

func (in *signingProfilePlatformIdPtr) ToSigningProfilePlatformIdPtrOutputWithContext(ctx context.Context) SigningProfilePlatformIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SigningProfilePlatformIdPtrOutput)
}

type SigningProfileSignatureValidityPeriodType string

const (
	SigningProfileSignatureValidityPeriodTypeDays   = SigningProfileSignatureValidityPeriodType("DAYS")
	SigningProfileSignatureValidityPeriodTypeMonths = SigningProfileSignatureValidityPeriodType("MONTHS")
	SigningProfileSignatureValidityPeriodTypeYears  = SigningProfileSignatureValidityPeriodType("YEARS")
)

func (SigningProfileSignatureValidityPeriodType) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningProfileSignatureValidityPeriodType)(nil)).Elem()
}

func (e SigningProfileSignatureValidityPeriodType) ToSigningProfileSignatureValidityPeriodTypeOutput() SigningProfileSignatureValidityPeriodTypeOutput {
	return pulumi.ToOutput(e).(SigningProfileSignatureValidityPeriodTypeOutput)
}

func (e SigningProfileSignatureValidityPeriodType) ToSigningProfileSignatureValidityPeriodTypeOutputWithContext(ctx context.Context) SigningProfileSignatureValidityPeriodTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SigningProfileSignatureValidityPeriodTypeOutput)
}

func (e SigningProfileSignatureValidityPeriodType) ToSigningProfileSignatureValidityPeriodTypePtrOutput() SigningProfileSignatureValidityPeriodTypePtrOutput {
	return e.ToSigningProfileSignatureValidityPeriodTypePtrOutputWithContext(context.Background())
}

func (e SigningProfileSignatureValidityPeriodType) ToSigningProfileSignatureValidityPeriodTypePtrOutputWithContext(ctx context.Context) SigningProfileSignatureValidityPeriodTypePtrOutput {
	return SigningProfileSignatureValidityPeriodType(e).ToSigningProfileSignatureValidityPeriodTypeOutputWithContext(ctx).ToSigningProfileSignatureValidityPeriodTypePtrOutputWithContext(ctx)
}

func (e SigningProfileSignatureValidityPeriodType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SigningProfileSignatureValidityPeriodType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SigningProfileSignatureValidityPeriodType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SigningProfileSignatureValidityPeriodType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SigningProfileSignatureValidityPeriodTypeOutput struct{ *pulumi.OutputState }

func (SigningProfileSignatureValidityPeriodTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SigningProfileSignatureValidityPeriodType)(nil)).Elem()
}

func (o SigningProfileSignatureValidityPeriodTypeOutput) ToSigningProfileSignatureValidityPeriodTypeOutput() SigningProfileSignatureValidityPeriodTypeOutput {
	return o
}

func (o SigningProfileSignatureValidityPeriodTypeOutput) ToSigningProfileSignatureValidityPeriodTypeOutputWithContext(ctx context.Context) SigningProfileSignatureValidityPeriodTypeOutput {
	return o
}

func (o SigningProfileSignatureValidityPeriodTypeOutput) ToSigningProfileSignatureValidityPeriodTypePtrOutput() SigningProfileSignatureValidityPeriodTypePtrOutput {
	return o.ToSigningProfileSignatureValidityPeriodTypePtrOutputWithContext(context.Background())
}

func (o SigningProfileSignatureValidityPeriodTypeOutput) ToSigningProfileSignatureValidityPeriodTypePtrOutputWithContext(ctx context.Context) SigningProfileSignatureValidityPeriodTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SigningProfileSignatureValidityPeriodType) *SigningProfileSignatureValidityPeriodType {
		return &v
	}).(SigningProfileSignatureValidityPeriodTypePtrOutput)
}

func (o SigningProfileSignatureValidityPeriodTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SigningProfileSignatureValidityPeriodTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SigningProfileSignatureValidityPeriodType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SigningProfileSignatureValidityPeriodTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SigningProfileSignatureValidityPeriodTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SigningProfileSignatureValidityPeriodType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SigningProfileSignatureValidityPeriodTypePtrOutput struct{ *pulumi.OutputState }

func (SigningProfileSignatureValidityPeriodTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SigningProfileSignatureValidityPeriodType)(nil)).Elem()
}

func (o SigningProfileSignatureValidityPeriodTypePtrOutput) ToSigningProfileSignatureValidityPeriodTypePtrOutput() SigningProfileSignatureValidityPeriodTypePtrOutput {
	return o
}

func (o SigningProfileSignatureValidityPeriodTypePtrOutput) ToSigningProfileSignatureValidityPeriodTypePtrOutputWithContext(ctx context.Context) SigningProfileSignatureValidityPeriodTypePtrOutput {
	return o
}

func (o SigningProfileSignatureValidityPeriodTypePtrOutput) Elem() SigningProfileSignatureValidityPeriodTypeOutput {
	return o.ApplyT(func(v *SigningProfileSignatureValidityPeriodType) SigningProfileSignatureValidityPeriodType {
		if v != nil {
			return *v
		}
		var ret SigningProfileSignatureValidityPeriodType
		return ret
	}).(SigningProfileSignatureValidityPeriodTypeOutput)
}

func (o SigningProfileSignatureValidityPeriodTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SigningProfileSignatureValidityPeriodTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SigningProfileSignatureValidityPeriodType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SigningProfileSignatureValidityPeriodTypeInput is an input type that accepts SigningProfileSignatureValidityPeriodTypeArgs and SigningProfileSignatureValidityPeriodTypeOutput values.
// You can construct a concrete instance of `SigningProfileSignatureValidityPeriodTypeInput` via:
//
//          SigningProfileSignatureValidityPeriodTypeArgs{...}
type SigningProfileSignatureValidityPeriodTypeInput interface {
	pulumi.Input

	ToSigningProfileSignatureValidityPeriodTypeOutput() SigningProfileSignatureValidityPeriodTypeOutput
	ToSigningProfileSignatureValidityPeriodTypeOutputWithContext(context.Context) SigningProfileSignatureValidityPeriodTypeOutput
}

var signingProfileSignatureValidityPeriodTypePtrType = reflect.TypeOf((**SigningProfileSignatureValidityPeriodType)(nil)).Elem()

type SigningProfileSignatureValidityPeriodTypePtrInput interface {
	pulumi.Input

	ToSigningProfileSignatureValidityPeriodTypePtrOutput() SigningProfileSignatureValidityPeriodTypePtrOutput
	ToSigningProfileSignatureValidityPeriodTypePtrOutputWithContext(context.Context) SigningProfileSignatureValidityPeriodTypePtrOutput
}

type signingProfileSignatureValidityPeriodTypePtr string

func SigningProfileSignatureValidityPeriodTypePtr(v string) SigningProfileSignatureValidityPeriodTypePtrInput {
	return (*signingProfileSignatureValidityPeriodTypePtr)(&v)
}

func (*signingProfileSignatureValidityPeriodTypePtr) ElementType() reflect.Type {
	return signingProfileSignatureValidityPeriodTypePtrType
}

func (in *signingProfileSignatureValidityPeriodTypePtr) ToSigningProfileSignatureValidityPeriodTypePtrOutput() SigningProfileSignatureValidityPeriodTypePtrOutput {
	return pulumi.ToOutput(in).(SigningProfileSignatureValidityPeriodTypePtrOutput)
}

func (in *signingProfileSignatureValidityPeriodTypePtr) ToSigningProfileSignatureValidityPeriodTypePtrOutputWithContext(ctx context.Context) SigningProfileSignatureValidityPeriodTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SigningProfileSignatureValidityPeriodTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SigningProfilePlatformIdInput)(nil)).Elem(), SigningProfilePlatformId("AWSLambda-SHA384-ECDSA"))
	pulumi.RegisterInputType(reflect.TypeOf((*SigningProfilePlatformIdPtrInput)(nil)).Elem(), SigningProfilePlatformId("AWSLambda-SHA384-ECDSA"))
	pulumi.RegisterInputType(reflect.TypeOf((*SigningProfileSignatureValidityPeriodTypeInput)(nil)).Elem(), SigningProfileSignatureValidityPeriodType("DAYS"))
	pulumi.RegisterInputType(reflect.TypeOf((*SigningProfileSignatureValidityPeriodTypePtrInput)(nil)).Elem(), SigningProfileSignatureValidityPeriodType("DAYS"))
	pulumi.RegisterOutputType(SigningProfilePlatformIdOutput{})
	pulumi.RegisterOutputType(SigningProfilePlatformIdPtrOutput{})
	pulumi.RegisterOutputType(SigningProfileSignatureValidityPeriodTypeOutput{})
	pulumi.RegisterOutputType(SigningProfileSignatureValidityPeriodTypePtrOutput{})
}
