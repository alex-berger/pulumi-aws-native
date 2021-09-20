// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cassandra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TableClusteringKeyColumnOrderBy string

const (
	TableClusteringKeyColumnOrderByAsc  = TableClusteringKeyColumnOrderBy("ASC")
	TableClusteringKeyColumnOrderByDesc = TableClusteringKeyColumnOrderBy("DESC")
)

func (TableClusteringKeyColumnOrderBy) ElementType() reflect.Type {
	return reflect.TypeOf((*TableClusteringKeyColumnOrderBy)(nil)).Elem()
}

func (e TableClusteringKeyColumnOrderBy) ToTableClusteringKeyColumnOrderByOutput() TableClusteringKeyColumnOrderByOutput {
	return pulumi.ToOutput(e).(TableClusteringKeyColumnOrderByOutput)
}

func (e TableClusteringKeyColumnOrderBy) ToTableClusteringKeyColumnOrderByOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableClusteringKeyColumnOrderByOutput)
}

func (e TableClusteringKeyColumnOrderBy) ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput {
	return e.ToTableClusteringKeyColumnOrderByPtrOutputWithContext(context.Background())
}

func (e TableClusteringKeyColumnOrderBy) ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByPtrOutput {
	return TableClusteringKeyColumnOrderBy(e).ToTableClusteringKeyColumnOrderByOutputWithContext(ctx).ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx)
}

func (e TableClusteringKeyColumnOrderBy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableClusteringKeyColumnOrderBy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableClusteringKeyColumnOrderBy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableClusteringKeyColumnOrderBy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableClusteringKeyColumnOrderByOutput struct{ *pulumi.OutputState }

func (TableClusteringKeyColumnOrderByOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableClusteringKeyColumnOrderBy)(nil)).Elem()
}

func (o TableClusteringKeyColumnOrderByOutput) ToTableClusteringKeyColumnOrderByOutput() TableClusteringKeyColumnOrderByOutput {
	return o
}

func (o TableClusteringKeyColumnOrderByOutput) ToTableClusteringKeyColumnOrderByOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByOutput {
	return o
}

func (o TableClusteringKeyColumnOrderByOutput) ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput {
	return o.ToTableClusteringKeyColumnOrderByPtrOutputWithContext(context.Background())
}

func (o TableClusteringKeyColumnOrderByOutput) ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableClusteringKeyColumnOrderBy) *TableClusteringKeyColumnOrderBy {
		return &v
	}).(TableClusteringKeyColumnOrderByPtrOutput)
}

func (o TableClusteringKeyColumnOrderByOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableClusteringKeyColumnOrderByOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableClusteringKeyColumnOrderBy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableClusteringKeyColumnOrderByOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableClusteringKeyColumnOrderByOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableClusteringKeyColumnOrderBy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableClusteringKeyColumnOrderByPtrOutput struct{ *pulumi.OutputState }

func (TableClusteringKeyColumnOrderByPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableClusteringKeyColumnOrderBy)(nil)).Elem()
}

func (o TableClusteringKeyColumnOrderByPtrOutput) ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput {
	return o
}

func (o TableClusteringKeyColumnOrderByPtrOutput) ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByPtrOutput {
	return o
}

func (o TableClusteringKeyColumnOrderByPtrOutput) Elem() TableClusteringKeyColumnOrderByOutput {
	return o.ApplyT(func(v *TableClusteringKeyColumnOrderBy) TableClusteringKeyColumnOrderBy {
		if v != nil {
			return *v
		}
		var ret TableClusteringKeyColumnOrderBy
		return ret
	}).(TableClusteringKeyColumnOrderByOutput)
}

func (o TableClusteringKeyColumnOrderByPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableClusteringKeyColumnOrderByPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableClusteringKeyColumnOrderBy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableClusteringKeyColumnOrderByInput is an input type that accepts TableClusteringKeyColumnOrderByArgs and TableClusteringKeyColumnOrderByOutput values.
// You can construct a concrete instance of `TableClusteringKeyColumnOrderByInput` via:
//
//          TableClusteringKeyColumnOrderByArgs{...}
type TableClusteringKeyColumnOrderByInput interface {
	pulumi.Input

	ToTableClusteringKeyColumnOrderByOutput() TableClusteringKeyColumnOrderByOutput
	ToTableClusteringKeyColumnOrderByOutputWithContext(context.Context) TableClusteringKeyColumnOrderByOutput
}

var tableClusteringKeyColumnOrderByPtrType = reflect.TypeOf((**TableClusteringKeyColumnOrderBy)(nil)).Elem()

type TableClusteringKeyColumnOrderByPtrInput interface {
	pulumi.Input

	ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput
	ToTableClusteringKeyColumnOrderByPtrOutputWithContext(context.Context) TableClusteringKeyColumnOrderByPtrOutput
}

type tableClusteringKeyColumnOrderByPtr string

func TableClusteringKeyColumnOrderByPtr(v string) TableClusteringKeyColumnOrderByPtrInput {
	return (*tableClusteringKeyColumnOrderByPtr)(&v)
}

func (*tableClusteringKeyColumnOrderByPtr) ElementType() reflect.Type {
	return tableClusteringKeyColumnOrderByPtrType
}

func (in *tableClusteringKeyColumnOrderByPtr) ToTableClusteringKeyColumnOrderByPtrOutput() TableClusteringKeyColumnOrderByPtrOutput {
	return pulumi.ToOutput(in).(TableClusteringKeyColumnOrderByPtrOutput)
}

func (in *tableClusteringKeyColumnOrderByPtr) ToTableClusteringKeyColumnOrderByPtrOutputWithContext(ctx context.Context) TableClusteringKeyColumnOrderByPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableClusteringKeyColumnOrderByPtrOutput)
}

// Server-side encryption type
type TableEncryptionType string

const (
	TableEncryptionTypeAwsOwnedKmsKey        = TableEncryptionType("AWS_OWNED_KMS_KEY")
	TableEncryptionTypeCustomerManagedKmsKey = TableEncryptionType("CUSTOMER_MANAGED_KMS_KEY")
)

func (TableEncryptionType) ElementType() reflect.Type {
	return reflect.TypeOf((*TableEncryptionType)(nil)).Elem()
}

func (e TableEncryptionType) ToTableEncryptionTypeOutput() TableEncryptionTypeOutput {
	return pulumi.ToOutput(e).(TableEncryptionTypeOutput)
}

func (e TableEncryptionType) ToTableEncryptionTypeOutputWithContext(ctx context.Context) TableEncryptionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableEncryptionTypeOutput)
}

func (e TableEncryptionType) ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput {
	return e.ToTableEncryptionTypePtrOutputWithContext(context.Background())
}

func (e TableEncryptionType) ToTableEncryptionTypePtrOutputWithContext(ctx context.Context) TableEncryptionTypePtrOutput {
	return TableEncryptionType(e).ToTableEncryptionTypeOutputWithContext(ctx).ToTableEncryptionTypePtrOutputWithContext(ctx)
}

func (e TableEncryptionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableEncryptionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableEncryptionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableEncryptionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableEncryptionTypeOutput struct{ *pulumi.OutputState }

func (TableEncryptionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableEncryptionType)(nil)).Elem()
}

func (o TableEncryptionTypeOutput) ToTableEncryptionTypeOutput() TableEncryptionTypeOutput {
	return o
}

func (o TableEncryptionTypeOutput) ToTableEncryptionTypeOutputWithContext(ctx context.Context) TableEncryptionTypeOutput {
	return o
}

func (o TableEncryptionTypeOutput) ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput {
	return o.ToTableEncryptionTypePtrOutputWithContext(context.Background())
}

func (o TableEncryptionTypeOutput) ToTableEncryptionTypePtrOutputWithContext(ctx context.Context) TableEncryptionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableEncryptionType) *TableEncryptionType {
		return &v
	}).(TableEncryptionTypePtrOutput)
}

func (o TableEncryptionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableEncryptionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableEncryptionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableEncryptionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableEncryptionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableEncryptionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableEncryptionTypePtrOutput struct{ *pulumi.OutputState }

func (TableEncryptionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableEncryptionType)(nil)).Elem()
}

func (o TableEncryptionTypePtrOutput) ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput {
	return o
}

func (o TableEncryptionTypePtrOutput) ToTableEncryptionTypePtrOutputWithContext(ctx context.Context) TableEncryptionTypePtrOutput {
	return o
}

func (o TableEncryptionTypePtrOutput) Elem() TableEncryptionTypeOutput {
	return o.ApplyT(func(v *TableEncryptionType) TableEncryptionType {
		if v != nil {
			return *v
		}
		var ret TableEncryptionType
		return ret
	}).(TableEncryptionTypeOutput)
}

func (o TableEncryptionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableEncryptionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableEncryptionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableEncryptionTypeInput is an input type that accepts TableEncryptionTypeArgs and TableEncryptionTypeOutput values.
// You can construct a concrete instance of `TableEncryptionTypeInput` via:
//
//          TableEncryptionTypeArgs{...}
type TableEncryptionTypeInput interface {
	pulumi.Input

	ToTableEncryptionTypeOutput() TableEncryptionTypeOutput
	ToTableEncryptionTypeOutputWithContext(context.Context) TableEncryptionTypeOutput
}

var tableEncryptionTypePtrType = reflect.TypeOf((**TableEncryptionType)(nil)).Elem()

type TableEncryptionTypePtrInput interface {
	pulumi.Input

	ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput
	ToTableEncryptionTypePtrOutputWithContext(context.Context) TableEncryptionTypePtrOutput
}

type tableEncryptionTypePtr string

func TableEncryptionTypePtr(v string) TableEncryptionTypePtrInput {
	return (*tableEncryptionTypePtr)(&v)
}

func (*tableEncryptionTypePtr) ElementType() reflect.Type {
	return tableEncryptionTypePtrType
}

func (in *tableEncryptionTypePtr) ToTableEncryptionTypePtrOutput() TableEncryptionTypePtrOutput {
	return pulumi.ToOutput(in).(TableEncryptionTypePtrOutput)
}

func (in *tableEncryptionTypePtr) ToTableEncryptionTypePtrOutputWithContext(ctx context.Context) TableEncryptionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableEncryptionTypePtrOutput)
}

// Capacity mode for the specified table
type TableMode string

const (
	TableModeProvisioned = TableMode("PROVISIONED")
	TableModeOnDemand    = TableMode("ON_DEMAND")
)

func (TableMode) ElementType() reflect.Type {
	return reflect.TypeOf((*TableMode)(nil)).Elem()
}

func (e TableMode) ToTableModeOutput() TableModeOutput {
	return pulumi.ToOutput(e).(TableModeOutput)
}

func (e TableMode) ToTableModeOutputWithContext(ctx context.Context) TableModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableModeOutput)
}

func (e TableMode) ToTableModePtrOutput() TableModePtrOutput {
	return e.ToTableModePtrOutputWithContext(context.Background())
}

func (e TableMode) ToTableModePtrOutputWithContext(ctx context.Context) TableModePtrOutput {
	return TableMode(e).ToTableModeOutputWithContext(ctx).ToTableModePtrOutputWithContext(ctx)
}

func (e TableMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableModeOutput struct{ *pulumi.OutputState }

func (TableModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableMode)(nil)).Elem()
}

func (o TableModeOutput) ToTableModeOutput() TableModeOutput {
	return o
}

func (o TableModeOutput) ToTableModeOutputWithContext(ctx context.Context) TableModeOutput {
	return o
}

func (o TableModeOutput) ToTableModePtrOutput() TableModePtrOutput {
	return o.ToTableModePtrOutputWithContext(context.Background())
}

func (o TableModeOutput) ToTableModePtrOutputWithContext(ctx context.Context) TableModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableMode) *TableMode {
		return &v
	}).(TableModePtrOutput)
}

func (o TableModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableModePtrOutput struct{ *pulumi.OutputState }

func (TableModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableMode)(nil)).Elem()
}

func (o TableModePtrOutput) ToTableModePtrOutput() TableModePtrOutput {
	return o
}

func (o TableModePtrOutput) ToTableModePtrOutputWithContext(ctx context.Context) TableModePtrOutput {
	return o
}

func (o TableModePtrOutput) Elem() TableModeOutput {
	return o.ApplyT(func(v *TableMode) TableMode {
		if v != nil {
			return *v
		}
		var ret TableMode
		return ret
	}).(TableModeOutput)
}

func (o TableModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableModeInput is an input type that accepts TableModeArgs and TableModeOutput values.
// You can construct a concrete instance of `TableModeInput` via:
//
//          TableModeArgs{...}
type TableModeInput interface {
	pulumi.Input

	ToTableModeOutput() TableModeOutput
	ToTableModeOutputWithContext(context.Context) TableModeOutput
}

var tableModePtrType = reflect.TypeOf((**TableMode)(nil)).Elem()

type TableModePtrInput interface {
	pulumi.Input

	ToTableModePtrOutput() TableModePtrOutput
	ToTableModePtrOutputWithContext(context.Context) TableModePtrOutput
}

type tableModePtr string

func TableModePtr(v string) TableModePtrInput {
	return (*tableModePtr)(&v)
}

func (*tableModePtr) ElementType() reflect.Type {
	return tableModePtrType
}

func (in *tableModePtr) ToTableModePtrOutput() TableModePtrOutput {
	return pulumi.ToOutput(in).(TableModePtrOutput)
}

func (in *tableModePtr) ToTableModePtrOutputWithContext(ctx context.Context) TableModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableModePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(TableClusteringKeyColumnOrderByOutput{})
	pulumi.RegisterOutputType(TableClusteringKeyColumnOrderByPtrOutput{})
	pulumi.RegisterOutputType(TableEncryptionTypeOutput{})
	pulumi.RegisterOutputType(TableEncryptionTypePtrOutput{})
	pulumi.RegisterOutputType(TableModeOutput{})
	pulumi.RegisterOutputType(TableModePtrOutput{})
}