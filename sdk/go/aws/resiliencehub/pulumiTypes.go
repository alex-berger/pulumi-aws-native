// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resiliencehub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppPhysicalResourceId struct {
	AwsAccountId *string `pulumi:"awsAccountId"`
	AwsRegion    *string `pulumi:"awsRegion"`
	Identifier   string  `pulumi:"identifier"`
	Type         string  `pulumi:"type"`
}

// AppPhysicalResourceIdInput is an input type that accepts AppPhysicalResourceIdArgs and AppPhysicalResourceIdOutput values.
// You can construct a concrete instance of `AppPhysicalResourceIdInput` via:
//
//          AppPhysicalResourceIdArgs{...}
type AppPhysicalResourceIdInput interface {
	pulumi.Input

	ToAppPhysicalResourceIdOutput() AppPhysicalResourceIdOutput
	ToAppPhysicalResourceIdOutputWithContext(context.Context) AppPhysicalResourceIdOutput
}

type AppPhysicalResourceIdArgs struct {
	AwsAccountId pulumi.StringPtrInput `pulumi:"awsAccountId"`
	AwsRegion    pulumi.StringPtrInput `pulumi:"awsRegion"`
	Identifier   pulumi.StringInput    `pulumi:"identifier"`
	Type         pulumi.StringInput    `pulumi:"type"`
}

func (AppPhysicalResourceIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppPhysicalResourceId)(nil)).Elem()
}

func (i AppPhysicalResourceIdArgs) ToAppPhysicalResourceIdOutput() AppPhysicalResourceIdOutput {
	return i.ToAppPhysicalResourceIdOutputWithContext(context.Background())
}

func (i AppPhysicalResourceIdArgs) ToAppPhysicalResourceIdOutputWithContext(ctx context.Context) AppPhysicalResourceIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppPhysicalResourceIdOutput)
}

type AppPhysicalResourceIdOutput struct{ *pulumi.OutputState }

func (AppPhysicalResourceIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppPhysicalResourceId)(nil)).Elem()
}

func (o AppPhysicalResourceIdOutput) ToAppPhysicalResourceIdOutput() AppPhysicalResourceIdOutput {
	return o
}

func (o AppPhysicalResourceIdOutput) ToAppPhysicalResourceIdOutputWithContext(ctx context.Context) AppPhysicalResourceIdOutput {
	return o
}

func (o AppPhysicalResourceIdOutput) AwsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppPhysicalResourceId) *string { return v.AwsAccountId }).(pulumi.StringPtrOutput)
}

func (o AppPhysicalResourceIdOutput) AwsRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppPhysicalResourceId) *string { return v.AwsRegion }).(pulumi.StringPtrOutput)
}

func (o AppPhysicalResourceIdOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v AppPhysicalResourceId) string { return v.Identifier }).(pulumi.StringOutput)
}

func (o AppPhysicalResourceIdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AppPhysicalResourceId) string { return v.Type }).(pulumi.StringOutput)
}

// Resource mapping is used to map logical resources from template to physical resource
type AppResourceMapping struct {
	LogicalStackName   *string               `pulumi:"logicalStackName"`
	MappingType        string                `pulumi:"mappingType"`
	PhysicalResourceId AppPhysicalResourceId `pulumi:"physicalResourceId"`
	ResourceName       *string               `pulumi:"resourceName"`
}

// AppResourceMappingInput is an input type that accepts AppResourceMappingArgs and AppResourceMappingOutput values.
// You can construct a concrete instance of `AppResourceMappingInput` via:
//
//          AppResourceMappingArgs{...}
type AppResourceMappingInput interface {
	pulumi.Input

	ToAppResourceMappingOutput() AppResourceMappingOutput
	ToAppResourceMappingOutputWithContext(context.Context) AppResourceMappingOutput
}

// Resource mapping is used to map logical resources from template to physical resource
type AppResourceMappingArgs struct {
	LogicalStackName   pulumi.StringPtrInput      `pulumi:"logicalStackName"`
	MappingType        pulumi.StringInput         `pulumi:"mappingType"`
	PhysicalResourceId AppPhysicalResourceIdInput `pulumi:"physicalResourceId"`
	ResourceName       pulumi.StringPtrInput      `pulumi:"resourceName"`
}

func (AppResourceMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourceMapping)(nil)).Elem()
}

func (i AppResourceMappingArgs) ToAppResourceMappingOutput() AppResourceMappingOutput {
	return i.ToAppResourceMappingOutputWithContext(context.Background())
}

func (i AppResourceMappingArgs) ToAppResourceMappingOutputWithContext(ctx context.Context) AppResourceMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourceMappingOutput)
}

// AppResourceMappingArrayInput is an input type that accepts AppResourceMappingArray and AppResourceMappingArrayOutput values.
// You can construct a concrete instance of `AppResourceMappingArrayInput` via:
//
//          AppResourceMappingArray{ AppResourceMappingArgs{...} }
type AppResourceMappingArrayInput interface {
	pulumi.Input

	ToAppResourceMappingArrayOutput() AppResourceMappingArrayOutput
	ToAppResourceMappingArrayOutputWithContext(context.Context) AppResourceMappingArrayOutput
}

type AppResourceMappingArray []AppResourceMappingInput

func (AppResourceMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppResourceMapping)(nil)).Elem()
}

func (i AppResourceMappingArray) ToAppResourceMappingArrayOutput() AppResourceMappingArrayOutput {
	return i.ToAppResourceMappingArrayOutputWithContext(context.Background())
}

func (i AppResourceMappingArray) ToAppResourceMappingArrayOutputWithContext(ctx context.Context) AppResourceMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResourceMappingArrayOutput)
}

// Resource mapping is used to map logical resources from template to physical resource
type AppResourceMappingOutput struct{ *pulumi.OutputState }

func (AppResourceMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppResourceMapping)(nil)).Elem()
}

func (o AppResourceMappingOutput) ToAppResourceMappingOutput() AppResourceMappingOutput {
	return o
}

func (o AppResourceMappingOutput) ToAppResourceMappingOutputWithContext(ctx context.Context) AppResourceMappingOutput {
	return o
}

func (o AppResourceMappingOutput) LogicalStackName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourceMapping) *string { return v.LogicalStackName }).(pulumi.StringPtrOutput)
}

func (o AppResourceMappingOutput) MappingType() pulumi.StringOutput {
	return o.ApplyT(func(v AppResourceMapping) string { return v.MappingType }).(pulumi.StringOutput)
}

func (o AppResourceMappingOutput) PhysicalResourceId() AppPhysicalResourceIdOutput {
	return o.ApplyT(func(v AppResourceMapping) AppPhysicalResourceId { return v.PhysicalResourceId }).(AppPhysicalResourceIdOutput)
}

func (o AppResourceMappingOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppResourceMapping) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

type AppResourceMappingArrayOutput struct{ *pulumi.OutputState }

func (AppResourceMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AppResourceMapping)(nil)).Elem()
}

func (o AppResourceMappingArrayOutput) ToAppResourceMappingArrayOutput() AppResourceMappingArrayOutput {
	return o
}

func (o AppResourceMappingArrayOutput) ToAppResourceMappingArrayOutputWithContext(ctx context.Context) AppResourceMappingArrayOutput {
	return o
}

func (o AppResourceMappingArrayOutput) Index(i pulumi.IntInput) AppResourceMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AppResourceMapping {
		return vs[0].([]AppResourceMapping)[vs[1].(int)]
	}).(AppResourceMappingOutput)
}

type AppTagMap struct {
}

// AppTagMapInput is an input type that accepts AppTagMap and AppTagMapOutput values.
// You can construct a concrete instance of `AppTagMapInput` via:
//
//          AppTagMap{ "key": AppTagArgs{...} }
type AppTagMapInput interface {
	pulumi.Input

	ToAppTagMapOutput() AppTagMapOutput
	ToAppTagMapOutputWithContext(context.Context) AppTagMapOutput
}

type AppTagMapArgs struct {
}

func (AppTagMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppTagMap)(nil)).Elem()
}

func (i AppTagMapArgs) ToAppTagMapOutput() AppTagMapOutput {
	return i.ToAppTagMapOutputWithContext(context.Background())
}

func (i AppTagMapArgs) ToAppTagMapOutputWithContext(ctx context.Context) AppTagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppTagMapOutput)
}

func (i AppTagMapArgs) ToAppTagMapPtrOutput() AppTagMapPtrOutput {
	return i.ToAppTagMapPtrOutputWithContext(context.Background())
}

func (i AppTagMapArgs) ToAppTagMapPtrOutputWithContext(ctx context.Context) AppTagMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppTagMapOutput).ToAppTagMapPtrOutputWithContext(ctx)
}

// AppTagMapPtrInput is an input type that accepts AppTagMapArgs, AppTagMapPtr and AppTagMapPtrOutput values.
// You can construct a concrete instance of `AppTagMapPtrInput` via:
//
//          AppTagMapArgs{...}
//
//  or:
//
//          nil
type AppTagMapPtrInput interface {
	pulumi.Input

	ToAppTagMapPtrOutput() AppTagMapPtrOutput
	ToAppTagMapPtrOutputWithContext(context.Context) AppTagMapPtrOutput
}

type appTagMapPtrType AppTagMapArgs

func AppTagMapPtr(v *AppTagMapArgs) AppTagMapPtrInput {
	return (*appTagMapPtrType)(v)
}

func (*appTagMapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppTagMap)(nil)).Elem()
}

func (i *appTagMapPtrType) ToAppTagMapPtrOutput() AppTagMapPtrOutput {
	return i.ToAppTagMapPtrOutputWithContext(context.Background())
}

func (i *appTagMapPtrType) ToAppTagMapPtrOutputWithContext(ctx context.Context) AppTagMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppTagMapPtrOutput)
}

type AppTagMapOutput struct{ *pulumi.OutputState }

func (AppTagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppTagMap)(nil)).Elem()
}

func (o AppTagMapOutput) ToAppTagMapOutput() AppTagMapOutput {
	return o
}

func (o AppTagMapOutput) ToAppTagMapOutputWithContext(ctx context.Context) AppTagMapOutput {
	return o
}

func (o AppTagMapOutput) ToAppTagMapPtrOutput() AppTagMapPtrOutput {
	return o.ToAppTagMapPtrOutputWithContext(context.Background())
}

func (o AppTagMapOutput) ToAppTagMapPtrOutputWithContext(ctx context.Context) AppTagMapPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppTagMap) *AppTagMap {
		return &v
	}).(AppTagMapPtrOutput)
}

type AppTagMapPtrOutput struct{ *pulumi.OutputState }

func (AppTagMapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppTagMap)(nil)).Elem()
}

func (o AppTagMapPtrOutput) ToAppTagMapPtrOutput() AppTagMapPtrOutput {
	return o
}

func (o AppTagMapPtrOutput) ToAppTagMapPtrOutputWithContext(ctx context.Context) AppTagMapPtrOutput {
	return o
}

func (o AppTagMapPtrOutput) Elem() AppTagMapOutput {
	return o.ApplyT(func(v *AppTagMap) AppTagMap {
		if v != nil {
			return *v
		}
		var ret AppTagMap
		return ret
	}).(AppTagMapOutput)
}

type ResiliencyPolicyPolicyMap struct {
}

// ResiliencyPolicyPolicyMapInput is an input type that accepts ResiliencyPolicyPolicyMap and ResiliencyPolicyPolicyMapOutput values.
// You can construct a concrete instance of `ResiliencyPolicyPolicyMapInput` via:
//
//          ResiliencyPolicyPolicyMap{ "key": ResiliencyPolicyPolicyArgs{...} }
type ResiliencyPolicyPolicyMapInput interface {
	pulumi.Input

	ToResiliencyPolicyPolicyMapOutput() ResiliencyPolicyPolicyMapOutput
	ToResiliencyPolicyPolicyMapOutputWithContext(context.Context) ResiliencyPolicyPolicyMapOutput
}

type ResiliencyPolicyPolicyMapArgs struct {
}

func (ResiliencyPolicyPolicyMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResiliencyPolicyPolicyMap)(nil)).Elem()
}

func (i ResiliencyPolicyPolicyMapArgs) ToResiliencyPolicyPolicyMapOutput() ResiliencyPolicyPolicyMapOutput {
	return i.ToResiliencyPolicyPolicyMapOutputWithContext(context.Background())
}

func (i ResiliencyPolicyPolicyMapArgs) ToResiliencyPolicyPolicyMapOutputWithContext(ctx context.Context) ResiliencyPolicyPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResiliencyPolicyPolicyMapOutput)
}

type ResiliencyPolicyPolicyMapOutput struct{ *pulumi.OutputState }

func (ResiliencyPolicyPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResiliencyPolicyPolicyMap)(nil)).Elem()
}

func (o ResiliencyPolicyPolicyMapOutput) ToResiliencyPolicyPolicyMapOutput() ResiliencyPolicyPolicyMapOutput {
	return o
}

func (o ResiliencyPolicyPolicyMapOutput) ToResiliencyPolicyPolicyMapOutputWithContext(ctx context.Context) ResiliencyPolicyPolicyMapOutput {
	return o
}

type ResiliencyPolicyPolicyMapPtrOutput struct{ *pulumi.OutputState }

func (ResiliencyPolicyPolicyMapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResiliencyPolicyPolicyMap)(nil)).Elem()
}

func (o ResiliencyPolicyPolicyMapPtrOutput) ToResiliencyPolicyPolicyMapPtrOutput() ResiliencyPolicyPolicyMapPtrOutput {
	return o
}

func (o ResiliencyPolicyPolicyMapPtrOutput) ToResiliencyPolicyPolicyMapPtrOutputWithContext(ctx context.Context) ResiliencyPolicyPolicyMapPtrOutput {
	return o
}

func (o ResiliencyPolicyPolicyMapPtrOutput) Elem() ResiliencyPolicyPolicyMapOutput {
	return o.ApplyT(func(v *ResiliencyPolicyPolicyMap) ResiliencyPolicyPolicyMap {
		if v != nil {
			return *v
		}
		var ret ResiliencyPolicyPolicyMap
		return ret
	}).(ResiliencyPolicyPolicyMapOutput)
}

type ResiliencyPolicyTagMap struct {
}

// ResiliencyPolicyTagMapInput is an input type that accepts ResiliencyPolicyTagMap and ResiliencyPolicyTagMapOutput values.
// You can construct a concrete instance of `ResiliencyPolicyTagMapInput` via:
//
//          ResiliencyPolicyTagMap{ "key": ResiliencyPolicyTagArgs{...} }
type ResiliencyPolicyTagMapInput interface {
	pulumi.Input

	ToResiliencyPolicyTagMapOutput() ResiliencyPolicyTagMapOutput
	ToResiliencyPolicyTagMapOutputWithContext(context.Context) ResiliencyPolicyTagMapOutput
}

type ResiliencyPolicyTagMapArgs struct {
}

func (ResiliencyPolicyTagMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResiliencyPolicyTagMap)(nil)).Elem()
}

func (i ResiliencyPolicyTagMapArgs) ToResiliencyPolicyTagMapOutput() ResiliencyPolicyTagMapOutput {
	return i.ToResiliencyPolicyTagMapOutputWithContext(context.Background())
}

func (i ResiliencyPolicyTagMapArgs) ToResiliencyPolicyTagMapOutputWithContext(ctx context.Context) ResiliencyPolicyTagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResiliencyPolicyTagMapOutput)
}

func (i ResiliencyPolicyTagMapArgs) ToResiliencyPolicyTagMapPtrOutput() ResiliencyPolicyTagMapPtrOutput {
	return i.ToResiliencyPolicyTagMapPtrOutputWithContext(context.Background())
}

func (i ResiliencyPolicyTagMapArgs) ToResiliencyPolicyTagMapPtrOutputWithContext(ctx context.Context) ResiliencyPolicyTagMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResiliencyPolicyTagMapOutput).ToResiliencyPolicyTagMapPtrOutputWithContext(ctx)
}

// ResiliencyPolicyTagMapPtrInput is an input type that accepts ResiliencyPolicyTagMapArgs, ResiliencyPolicyTagMapPtr and ResiliencyPolicyTagMapPtrOutput values.
// You can construct a concrete instance of `ResiliencyPolicyTagMapPtrInput` via:
//
//          ResiliencyPolicyTagMapArgs{...}
//
//  or:
//
//          nil
type ResiliencyPolicyTagMapPtrInput interface {
	pulumi.Input

	ToResiliencyPolicyTagMapPtrOutput() ResiliencyPolicyTagMapPtrOutput
	ToResiliencyPolicyTagMapPtrOutputWithContext(context.Context) ResiliencyPolicyTagMapPtrOutput
}

type resiliencyPolicyTagMapPtrType ResiliencyPolicyTagMapArgs

func ResiliencyPolicyTagMapPtr(v *ResiliencyPolicyTagMapArgs) ResiliencyPolicyTagMapPtrInput {
	return (*resiliencyPolicyTagMapPtrType)(v)
}

func (*resiliencyPolicyTagMapPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResiliencyPolicyTagMap)(nil)).Elem()
}

func (i *resiliencyPolicyTagMapPtrType) ToResiliencyPolicyTagMapPtrOutput() ResiliencyPolicyTagMapPtrOutput {
	return i.ToResiliencyPolicyTagMapPtrOutputWithContext(context.Background())
}

func (i *resiliencyPolicyTagMapPtrType) ToResiliencyPolicyTagMapPtrOutputWithContext(ctx context.Context) ResiliencyPolicyTagMapPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResiliencyPolicyTagMapPtrOutput)
}

type ResiliencyPolicyTagMapOutput struct{ *pulumi.OutputState }

func (ResiliencyPolicyTagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResiliencyPolicyTagMap)(nil)).Elem()
}

func (o ResiliencyPolicyTagMapOutput) ToResiliencyPolicyTagMapOutput() ResiliencyPolicyTagMapOutput {
	return o
}

func (o ResiliencyPolicyTagMapOutput) ToResiliencyPolicyTagMapOutputWithContext(ctx context.Context) ResiliencyPolicyTagMapOutput {
	return o
}

func (o ResiliencyPolicyTagMapOutput) ToResiliencyPolicyTagMapPtrOutput() ResiliencyPolicyTagMapPtrOutput {
	return o.ToResiliencyPolicyTagMapPtrOutputWithContext(context.Background())
}

func (o ResiliencyPolicyTagMapOutput) ToResiliencyPolicyTagMapPtrOutputWithContext(ctx context.Context) ResiliencyPolicyTagMapPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResiliencyPolicyTagMap) *ResiliencyPolicyTagMap {
		return &v
	}).(ResiliencyPolicyTagMapPtrOutput)
}

type ResiliencyPolicyTagMapPtrOutput struct{ *pulumi.OutputState }

func (ResiliencyPolicyTagMapPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResiliencyPolicyTagMap)(nil)).Elem()
}

func (o ResiliencyPolicyTagMapPtrOutput) ToResiliencyPolicyTagMapPtrOutput() ResiliencyPolicyTagMapPtrOutput {
	return o
}

func (o ResiliencyPolicyTagMapPtrOutput) ToResiliencyPolicyTagMapPtrOutputWithContext(ctx context.Context) ResiliencyPolicyTagMapPtrOutput {
	return o
}

func (o ResiliencyPolicyTagMapPtrOutput) Elem() ResiliencyPolicyTagMapOutput {
	return o.ApplyT(func(v *ResiliencyPolicyTagMap) ResiliencyPolicyTagMap {
		if v != nil {
			return *v
		}
		var ret ResiliencyPolicyTagMap
		return ret
	}).(ResiliencyPolicyTagMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppPhysicalResourceIdInput)(nil)).Elem(), AppPhysicalResourceIdArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppResourceMappingInput)(nil)).Elem(), AppResourceMappingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppResourceMappingArrayInput)(nil)).Elem(), AppResourceMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppTagMapInput)(nil)).Elem(), AppTagMapArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppTagMapPtrInput)(nil)).Elem(), AppTagMapArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResiliencyPolicyPolicyMapInput)(nil)).Elem(), ResiliencyPolicyPolicyMapArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResiliencyPolicyTagMapInput)(nil)).Elem(), ResiliencyPolicyTagMapArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResiliencyPolicyTagMapPtrInput)(nil)).Elem(), ResiliencyPolicyTagMapArgs{})
	pulumi.RegisterOutputType(AppPhysicalResourceIdOutput{})
	pulumi.RegisterOutputType(AppResourceMappingOutput{})
	pulumi.RegisterOutputType(AppResourceMappingArrayOutput{})
	pulumi.RegisterOutputType(AppTagMapOutput{})
	pulumi.RegisterOutputType(AppTagMapPtrOutput{})
	pulumi.RegisterOutputType(ResiliencyPolicyPolicyMapOutput{})
	pulumi.RegisterOutputType(ResiliencyPolicyPolicyMapPtrOutput{})
	pulumi.RegisterOutputType(ResiliencyPolicyTagMapOutput{})
	pulumi.RegisterOutputType(ResiliencyPolicyTagMapPtrOutput{})
}
