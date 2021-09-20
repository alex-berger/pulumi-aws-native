// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3outposts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BucketRuleStatus string

const (
	BucketRuleStatusEnabled  = BucketRuleStatus("Enabled")
	BucketRuleStatusDisabled = BucketRuleStatus("Disabled")
)

func (BucketRuleStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketRuleStatus)(nil)).Elem()
}

func (e BucketRuleStatus) ToBucketRuleStatusOutput() BucketRuleStatusOutput {
	return pulumi.ToOutput(e).(BucketRuleStatusOutput)
}

func (e BucketRuleStatus) ToBucketRuleStatusOutputWithContext(ctx context.Context) BucketRuleStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BucketRuleStatusOutput)
}

func (e BucketRuleStatus) ToBucketRuleStatusPtrOutput() BucketRuleStatusPtrOutput {
	return e.ToBucketRuleStatusPtrOutputWithContext(context.Background())
}

func (e BucketRuleStatus) ToBucketRuleStatusPtrOutputWithContext(ctx context.Context) BucketRuleStatusPtrOutput {
	return BucketRuleStatus(e).ToBucketRuleStatusOutputWithContext(ctx).ToBucketRuleStatusPtrOutputWithContext(ctx)
}

func (e BucketRuleStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BucketRuleStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BucketRuleStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BucketRuleStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BucketRuleStatusOutput struct{ *pulumi.OutputState }

func (BucketRuleStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketRuleStatus)(nil)).Elem()
}

func (o BucketRuleStatusOutput) ToBucketRuleStatusOutput() BucketRuleStatusOutput {
	return o
}

func (o BucketRuleStatusOutput) ToBucketRuleStatusOutputWithContext(ctx context.Context) BucketRuleStatusOutput {
	return o
}

func (o BucketRuleStatusOutput) ToBucketRuleStatusPtrOutput() BucketRuleStatusPtrOutput {
	return o.ToBucketRuleStatusPtrOutputWithContext(context.Background())
}

func (o BucketRuleStatusOutput) ToBucketRuleStatusPtrOutputWithContext(ctx context.Context) BucketRuleStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BucketRuleStatus) *BucketRuleStatus {
		return &v
	}).(BucketRuleStatusPtrOutput)
}

func (o BucketRuleStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BucketRuleStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BucketRuleStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BucketRuleStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BucketRuleStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BucketRuleStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BucketRuleStatusPtrOutput struct{ *pulumi.OutputState }

func (BucketRuleStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketRuleStatus)(nil)).Elem()
}

func (o BucketRuleStatusPtrOutput) ToBucketRuleStatusPtrOutput() BucketRuleStatusPtrOutput {
	return o
}

func (o BucketRuleStatusPtrOutput) ToBucketRuleStatusPtrOutputWithContext(ctx context.Context) BucketRuleStatusPtrOutput {
	return o
}

func (o BucketRuleStatusPtrOutput) Elem() BucketRuleStatusOutput {
	return o.ApplyT(func(v *BucketRuleStatus) BucketRuleStatus {
		if v != nil {
			return *v
		}
		var ret BucketRuleStatus
		return ret
	}).(BucketRuleStatusOutput)
}

func (o BucketRuleStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BucketRuleStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BucketRuleStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// BucketRuleStatusInput is an input type that accepts BucketRuleStatusArgs and BucketRuleStatusOutput values.
// You can construct a concrete instance of `BucketRuleStatusInput` via:
//
//          BucketRuleStatusArgs{...}
type BucketRuleStatusInput interface {
	pulumi.Input

	ToBucketRuleStatusOutput() BucketRuleStatusOutput
	ToBucketRuleStatusOutputWithContext(context.Context) BucketRuleStatusOutput
}

var bucketRuleStatusPtrType = reflect.TypeOf((**BucketRuleStatus)(nil)).Elem()

type BucketRuleStatusPtrInput interface {
	pulumi.Input

	ToBucketRuleStatusPtrOutput() BucketRuleStatusPtrOutput
	ToBucketRuleStatusPtrOutputWithContext(context.Context) BucketRuleStatusPtrOutput
}

type bucketRuleStatusPtr string

func BucketRuleStatusPtr(v string) BucketRuleStatusPtrInput {
	return (*bucketRuleStatusPtr)(&v)
}

func (*bucketRuleStatusPtr) ElementType() reflect.Type {
	return bucketRuleStatusPtrType
}

func (in *bucketRuleStatusPtr) ToBucketRuleStatusPtrOutput() BucketRuleStatusPtrOutput {
	return pulumi.ToOutput(in).(BucketRuleStatusPtrOutput)
}

func (in *bucketRuleStatusPtr) ToBucketRuleStatusPtrOutputWithContext(ctx context.Context) BucketRuleStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BucketRuleStatusPtrOutput)
}

// The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
type EndpointAccessType string

const (
	EndpointAccessTypeCustomerOwnedIp = EndpointAccessType("CustomerOwnedIp")
	EndpointAccessTypePrivate         = EndpointAccessType("Private")
)

func (EndpointAccessType) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAccessType)(nil)).Elem()
}

func (e EndpointAccessType) ToEndpointAccessTypeOutput() EndpointAccessTypeOutput {
	return pulumi.ToOutput(e).(EndpointAccessTypeOutput)
}

func (e EndpointAccessType) ToEndpointAccessTypeOutputWithContext(ctx context.Context) EndpointAccessTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointAccessTypeOutput)
}

func (e EndpointAccessType) ToEndpointAccessTypePtrOutput() EndpointAccessTypePtrOutput {
	return e.ToEndpointAccessTypePtrOutputWithContext(context.Background())
}

func (e EndpointAccessType) ToEndpointAccessTypePtrOutputWithContext(ctx context.Context) EndpointAccessTypePtrOutput {
	return EndpointAccessType(e).ToEndpointAccessTypeOutputWithContext(ctx).ToEndpointAccessTypePtrOutputWithContext(ctx)
}

func (e EndpointAccessType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointAccessType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointAccessType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointAccessType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointAccessTypeOutput struct{ *pulumi.OutputState }

func (EndpointAccessTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAccessType)(nil)).Elem()
}

func (o EndpointAccessTypeOutput) ToEndpointAccessTypeOutput() EndpointAccessTypeOutput {
	return o
}

func (o EndpointAccessTypeOutput) ToEndpointAccessTypeOutputWithContext(ctx context.Context) EndpointAccessTypeOutput {
	return o
}

func (o EndpointAccessTypeOutput) ToEndpointAccessTypePtrOutput() EndpointAccessTypePtrOutput {
	return o.ToEndpointAccessTypePtrOutputWithContext(context.Background())
}

func (o EndpointAccessTypeOutput) ToEndpointAccessTypePtrOutputWithContext(ctx context.Context) EndpointAccessTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointAccessType) *EndpointAccessType {
		return &v
	}).(EndpointAccessTypePtrOutput)
}

func (o EndpointAccessTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointAccessTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointAccessType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointAccessTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointAccessTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointAccessType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointAccessTypePtrOutput struct{ *pulumi.OutputState }

func (EndpointAccessTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAccessType)(nil)).Elem()
}

func (o EndpointAccessTypePtrOutput) ToEndpointAccessTypePtrOutput() EndpointAccessTypePtrOutput {
	return o
}

func (o EndpointAccessTypePtrOutput) ToEndpointAccessTypePtrOutputWithContext(ctx context.Context) EndpointAccessTypePtrOutput {
	return o
}

func (o EndpointAccessTypePtrOutput) Elem() EndpointAccessTypeOutput {
	return o.ApplyT(func(v *EndpointAccessType) EndpointAccessType {
		if v != nil {
			return *v
		}
		var ret EndpointAccessType
		return ret
	}).(EndpointAccessTypeOutput)
}

func (o EndpointAccessTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointAccessTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointAccessType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EndpointAccessTypeInput is an input type that accepts EndpointAccessTypeArgs and EndpointAccessTypeOutput values.
// You can construct a concrete instance of `EndpointAccessTypeInput` via:
//
//          EndpointAccessTypeArgs{...}
type EndpointAccessTypeInput interface {
	pulumi.Input

	ToEndpointAccessTypeOutput() EndpointAccessTypeOutput
	ToEndpointAccessTypeOutputWithContext(context.Context) EndpointAccessTypeOutput
}

var endpointAccessTypePtrType = reflect.TypeOf((**EndpointAccessType)(nil)).Elem()

type EndpointAccessTypePtrInput interface {
	pulumi.Input

	ToEndpointAccessTypePtrOutput() EndpointAccessTypePtrOutput
	ToEndpointAccessTypePtrOutputWithContext(context.Context) EndpointAccessTypePtrOutput
}

type endpointAccessTypePtr string

func EndpointAccessTypePtr(v string) EndpointAccessTypePtrInput {
	return (*endpointAccessTypePtr)(&v)
}

func (*endpointAccessTypePtr) ElementType() reflect.Type {
	return endpointAccessTypePtrType
}

func (in *endpointAccessTypePtr) ToEndpointAccessTypePtrOutput() EndpointAccessTypePtrOutput {
	return pulumi.ToOutput(in).(EndpointAccessTypePtrOutput)
}

func (in *endpointAccessTypePtr) ToEndpointAccessTypePtrOutputWithContext(ctx context.Context) EndpointAccessTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointAccessTypePtrOutput)
}

type EndpointStatus string

const (
	EndpointStatusAvailable = EndpointStatus("Available")
	EndpointStatusPending   = EndpointStatus("Pending")
	EndpointStatusDeleting  = EndpointStatus("Deleting")
)

func (EndpointStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointStatus)(nil)).Elem()
}

func (e EndpointStatus) ToEndpointStatusOutput() EndpointStatusOutput {
	return pulumi.ToOutput(e).(EndpointStatusOutput)
}

func (e EndpointStatus) ToEndpointStatusOutputWithContext(ctx context.Context) EndpointStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointStatusOutput)
}

func (e EndpointStatus) ToEndpointStatusPtrOutput() EndpointStatusPtrOutput {
	return e.ToEndpointStatusPtrOutputWithContext(context.Background())
}

func (e EndpointStatus) ToEndpointStatusPtrOutputWithContext(ctx context.Context) EndpointStatusPtrOutput {
	return EndpointStatus(e).ToEndpointStatusOutputWithContext(ctx).ToEndpointStatusPtrOutputWithContext(ctx)
}

func (e EndpointStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointStatusOutput struct{ *pulumi.OutputState }

func (EndpointStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointStatus)(nil)).Elem()
}

func (o EndpointStatusOutput) ToEndpointStatusOutput() EndpointStatusOutput {
	return o
}

func (o EndpointStatusOutput) ToEndpointStatusOutputWithContext(ctx context.Context) EndpointStatusOutput {
	return o
}

func (o EndpointStatusOutput) ToEndpointStatusPtrOutput() EndpointStatusPtrOutput {
	return o.ToEndpointStatusPtrOutputWithContext(context.Background())
}

func (o EndpointStatusOutput) ToEndpointStatusPtrOutputWithContext(ctx context.Context) EndpointStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointStatus) *EndpointStatus {
		return &v
	}).(EndpointStatusPtrOutput)
}

func (o EndpointStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointStatusPtrOutput struct{ *pulumi.OutputState }

func (EndpointStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointStatus)(nil)).Elem()
}

func (o EndpointStatusPtrOutput) ToEndpointStatusPtrOutput() EndpointStatusPtrOutput {
	return o
}

func (o EndpointStatusPtrOutput) ToEndpointStatusPtrOutputWithContext(ctx context.Context) EndpointStatusPtrOutput {
	return o
}

func (o EndpointStatusPtrOutput) Elem() EndpointStatusOutput {
	return o.ApplyT(func(v *EndpointStatus) EndpointStatus {
		if v != nil {
			return *v
		}
		var ret EndpointStatus
		return ret
	}).(EndpointStatusOutput)
}

func (o EndpointStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EndpointStatusInput is an input type that accepts EndpointStatusArgs and EndpointStatusOutput values.
// You can construct a concrete instance of `EndpointStatusInput` via:
//
//          EndpointStatusArgs{...}
type EndpointStatusInput interface {
	pulumi.Input

	ToEndpointStatusOutput() EndpointStatusOutput
	ToEndpointStatusOutputWithContext(context.Context) EndpointStatusOutput
}

var endpointStatusPtrType = reflect.TypeOf((**EndpointStatus)(nil)).Elem()

type EndpointStatusPtrInput interface {
	pulumi.Input

	ToEndpointStatusPtrOutput() EndpointStatusPtrOutput
	ToEndpointStatusPtrOutputWithContext(context.Context) EndpointStatusPtrOutput
}

type endpointStatusPtr string

func EndpointStatusPtr(v string) EndpointStatusPtrInput {
	return (*endpointStatusPtr)(&v)
}

func (*endpointStatusPtr) ElementType() reflect.Type {
	return endpointStatusPtrType
}

func (in *endpointStatusPtr) ToEndpointStatusPtrOutput() EndpointStatusPtrOutput {
	return pulumi.ToOutput(in).(EndpointStatusPtrOutput)
}

func (in *endpointStatusPtr) ToEndpointStatusPtrOutputWithContext(ctx context.Context) EndpointStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BucketRuleStatusOutput{})
	pulumi.RegisterOutputType(BucketRuleStatusPtrOutput{})
	pulumi.RegisterOutputType(EndpointAccessTypeOutput{})
	pulumi.RegisterOutputType(EndpointAccessTypePtrOutput{})
	pulumi.RegisterOutputType(EndpointStatusOutput{})
	pulumi.RegisterOutputType(EndpointStatusPtrOutput{})
}