// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.
type KeyKeySpec string

const (
	KeyKeySpecSymmetricDefault = KeyKeySpec("SYMMETRIC_DEFAULT")
	KeyKeySpecRsa2048          = KeyKeySpec("RSA_2048")
	KeyKeySpecRsa3072          = KeyKeySpec("RSA_3072")
	KeyKeySpecRsa4096          = KeyKeySpec("RSA_4096")
	KeyKeySpecEccNistP256      = KeyKeySpec("ECC_NIST_P256")
	KeyKeySpecEccNistP384      = KeyKeySpec("ECC_NIST_P384")
	KeyKeySpecEccNistP521      = KeyKeySpec("ECC_NIST_P521")
	KeyKeySpecEccSecgP256k1    = KeyKeySpec("ECC_SECG_P256K1")
)

func (KeyKeySpec) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyKeySpec)(nil)).Elem()
}

func (e KeyKeySpec) ToKeyKeySpecOutput() KeyKeySpecOutput {
	return pulumi.ToOutput(e).(KeyKeySpecOutput)
}

func (e KeyKeySpec) ToKeyKeySpecOutputWithContext(ctx context.Context) KeyKeySpecOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyKeySpecOutput)
}

func (e KeyKeySpec) ToKeyKeySpecPtrOutput() KeyKeySpecPtrOutput {
	return e.ToKeyKeySpecPtrOutputWithContext(context.Background())
}

func (e KeyKeySpec) ToKeyKeySpecPtrOutputWithContext(ctx context.Context) KeyKeySpecPtrOutput {
	return KeyKeySpec(e).ToKeyKeySpecOutputWithContext(ctx).ToKeyKeySpecPtrOutputWithContext(ctx)
}

func (e KeyKeySpec) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyKeySpec) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyKeySpec) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyKeySpec) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyKeySpecOutput struct{ *pulumi.OutputState }

func (KeyKeySpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyKeySpec)(nil)).Elem()
}

func (o KeyKeySpecOutput) ToKeyKeySpecOutput() KeyKeySpecOutput {
	return o
}

func (o KeyKeySpecOutput) ToKeyKeySpecOutputWithContext(ctx context.Context) KeyKeySpecOutput {
	return o
}

func (o KeyKeySpecOutput) ToKeyKeySpecPtrOutput() KeyKeySpecPtrOutput {
	return o.ToKeyKeySpecPtrOutputWithContext(context.Background())
}

func (o KeyKeySpecOutput) ToKeyKeySpecPtrOutputWithContext(ctx context.Context) KeyKeySpecPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyKeySpec) *KeyKeySpec {
		return &v
	}).(KeyKeySpecPtrOutput)
}

func (o KeyKeySpecOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyKeySpecOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyKeySpec) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyKeySpecOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyKeySpecOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyKeySpec) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyKeySpecPtrOutput struct{ *pulumi.OutputState }

func (KeyKeySpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyKeySpec)(nil)).Elem()
}

func (o KeyKeySpecPtrOutput) ToKeyKeySpecPtrOutput() KeyKeySpecPtrOutput {
	return o
}

func (o KeyKeySpecPtrOutput) ToKeyKeySpecPtrOutputWithContext(ctx context.Context) KeyKeySpecPtrOutput {
	return o
}

func (o KeyKeySpecPtrOutput) Elem() KeyKeySpecOutput {
	return o.ApplyT(func(v *KeyKeySpec) KeyKeySpec {
		if v != nil {
			return *v
		}
		var ret KeyKeySpec
		return ret
	}).(KeyKeySpecOutput)
}

func (o KeyKeySpecPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyKeySpecPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyKeySpec) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KeyKeySpecInput is an input type that accepts KeyKeySpecArgs and KeyKeySpecOutput values.
// You can construct a concrete instance of `KeyKeySpecInput` via:
//
//          KeyKeySpecArgs{...}
type KeyKeySpecInput interface {
	pulumi.Input

	ToKeyKeySpecOutput() KeyKeySpecOutput
	ToKeyKeySpecOutputWithContext(context.Context) KeyKeySpecOutput
}

var keyKeySpecPtrType = reflect.TypeOf((**KeyKeySpec)(nil)).Elem()

type KeyKeySpecPtrInput interface {
	pulumi.Input

	ToKeyKeySpecPtrOutput() KeyKeySpecPtrOutput
	ToKeyKeySpecPtrOutputWithContext(context.Context) KeyKeySpecPtrOutput
}

type keyKeySpecPtr string

func KeyKeySpecPtr(v string) KeyKeySpecPtrInput {
	return (*keyKeySpecPtr)(&v)
}

func (*keyKeySpecPtr) ElementType() reflect.Type {
	return keyKeySpecPtrType
}

func (in *keyKeySpecPtr) ToKeyKeySpecPtrOutput() KeyKeySpecPtrOutput {
	return pulumi.ToOutput(in).(KeyKeySpecPtrOutput)
}

func (in *keyKeySpecPtr) ToKeyKeySpecPtrOutputWithContext(ctx context.Context) KeyKeySpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyKeySpecPtrOutput)
}

// Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.
type KeyKeyUsage string

const (
	KeyKeyUsageEncryptDecrypt = KeyKeyUsage("ENCRYPT_DECRYPT")
	KeyKeyUsageSignVerify     = KeyKeyUsage("SIGN_VERIFY")
)

func (KeyKeyUsage) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyKeyUsage)(nil)).Elem()
}

func (e KeyKeyUsage) ToKeyKeyUsageOutput() KeyKeyUsageOutput {
	return pulumi.ToOutput(e).(KeyKeyUsageOutput)
}

func (e KeyKeyUsage) ToKeyKeyUsageOutputWithContext(ctx context.Context) KeyKeyUsageOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeyKeyUsageOutput)
}

func (e KeyKeyUsage) ToKeyKeyUsagePtrOutput() KeyKeyUsagePtrOutput {
	return e.ToKeyKeyUsagePtrOutputWithContext(context.Background())
}

func (e KeyKeyUsage) ToKeyKeyUsagePtrOutputWithContext(ctx context.Context) KeyKeyUsagePtrOutput {
	return KeyKeyUsage(e).ToKeyKeyUsageOutputWithContext(ctx).ToKeyKeyUsagePtrOutputWithContext(ctx)
}

func (e KeyKeyUsage) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyKeyUsage) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeyKeyUsage) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeyKeyUsage) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeyKeyUsageOutput struct{ *pulumi.OutputState }

func (KeyKeyUsageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyKeyUsage)(nil)).Elem()
}

func (o KeyKeyUsageOutput) ToKeyKeyUsageOutput() KeyKeyUsageOutput {
	return o
}

func (o KeyKeyUsageOutput) ToKeyKeyUsageOutputWithContext(ctx context.Context) KeyKeyUsageOutput {
	return o
}

func (o KeyKeyUsageOutput) ToKeyKeyUsagePtrOutput() KeyKeyUsagePtrOutput {
	return o.ToKeyKeyUsagePtrOutputWithContext(context.Background())
}

func (o KeyKeyUsageOutput) ToKeyKeyUsagePtrOutputWithContext(ctx context.Context) KeyKeyUsagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyKeyUsage) *KeyKeyUsage {
		return &v
	}).(KeyKeyUsagePtrOutput)
}

func (o KeyKeyUsageOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeyKeyUsageOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyKeyUsage) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeyKeyUsageOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyKeyUsageOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeyKeyUsage) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeyKeyUsagePtrOutput struct{ *pulumi.OutputState }

func (KeyKeyUsagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyKeyUsage)(nil)).Elem()
}

func (o KeyKeyUsagePtrOutput) ToKeyKeyUsagePtrOutput() KeyKeyUsagePtrOutput {
	return o
}

func (o KeyKeyUsagePtrOutput) ToKeyKeyUsagePtrOutputWithContext(ctx context.Context) KeyKeyUsagePtrOutput {
	return o
}

func (o KeyKeyUsagePtrOutput) Elem() KeyKeyUsageOutput {
	return o.ApplyT(func(v *KeyKeyUsage) KeyKeyUsage {
		if v != nil {
			return *v
		}
		var ret KeyKeyUsage
		return ret
	}).(KeyKeyUsageOutput)
}

func (o KeyKeyUsagePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeyKeyUsagePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeyKeyUsage) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// KeyKeyUsageInput is an input type that accepts KeyKeyUsageArgs and KeyKeyUsageOutput values.
// You can construct a concrete instance of `KeyKeyUsageInput` via:
//
//          KeyKeyUsageArgs{...}
type KeyKeyUsageInput interface {
	pulumi.Input

	ToKeyKeyUsageOutput() KeyKeyUsageOutput
	ToKeyKeyUsageOutputWithContext(context.Context) KeyKeyUsageOutput
}

var keyKeyUsagePtrType = reflect.TypeOf((**KeyKeyUsage)(nil)).Elem()

type KeyKeyUsagePtrInput interface {
	pulumi.Input

	ToKeyKeyUsagePtrOutput() KeyKeyUsagePtrOutput
	ToKeyKeyUsagePtrOutputWithContext(context.Context) KeyKeyUsagePtrOutput
}

type keyKeyUsagePtr string

func KeyKeyUsagePtr(v string) KeyKeyUsagePtrInput {
	return (*keyKeyUsagePtr)(&v)
}

func (*keyKeyUsagePtr) ElementType() reflect.Type {
	return keyKeyUsagePtrType
}

func (in *keyKeyUsagePtr) ToKeyKeyUsagePtrOutput() KeyKeyUsagePtrOutput {
	return pulumi.ToOutput(in).(KeyKeyUsagePtrOutput)
}

func (in *keyKeyUsagePtr) ToKeyKeyUsagePtrOutputWithContext(ctx context.Context) KeyKeyUsagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeyKeyUsagePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyKeySpecOutput{})
	pulumi.RegisterOutputType(KeyKeySpecPtrOutput{})
	pulumi.RegisterOutputType(KeyKeyUsageOutput{})
	pulumi.RegisterOutputType(KeyKeyUsagePtrOutput{})
}