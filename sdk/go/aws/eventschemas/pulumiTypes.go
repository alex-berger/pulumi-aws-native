// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventschemas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html
type DiscovererTagsEntry struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html#cfn-eventschemas-discoverer-tagsentry-key
	Key string `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html#cfn-eventschemas-discoverer-tagsentry-value
	Value string `pulumi:"value"`
}

// DiscovererTagsEntryInput is an input type that accepts DiscovererTagsEntryArgs and DiscovererTagsEntryOutput values.
// You can construct a concrete instance of `DiscovererTagsEntryInput` via:
//
//          DiscovererTagsEntryArgs{...}
type DiscovererTagsEntryInput interface {
	pulumi.Input

	ToDiscovererTagsEntryOutput() DiscovererTagsEntryOutput
	ToDiscovererTagsEntryOutputWithContext(context.Context) DiscovererTagsEntryOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html
type DiscovererTagsEntryArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html#cfn-eventschemas-discoverer-tagsentry-key
	Key pulumi.StringInput `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html#cfn-eventschemas-discoverer-tagsentry-value
	Value pulumi.StringInput `pulumi:"value"`
}

func (DiscovererTagsEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiscovererTagsEntry)(nil)).Elem()
}

func (i DiscovererTagsEntryArgs) ToDiscovererTagsEntryOutput() DiscovererTagsEntryOutput {
	return i.ToDiscovererTagsEntryOutputWithContext(context.Background())
}

func (i DiscovererTagsEntryArgs) ToDiscovererTagsEntryOutputWithContext(ctx context.Context) DiscovererTagsEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiscovererTagsEntryOutput)
}

// DiscovererTagsEntryArrayInput is an input type that accepts DiscovererTagsEntryArray and DiscovererTagsEntryArrayOutput values.
// You can construct a concrete instance of `DiscovererTagsEntryArrayInput` via:
//
//          DiscovererTagsEntryArray{ DiscovererTagsEntryArgs{...} }
type DiscovererTagsEntryArrayInput interface {
	pulumi.Input

	ToDiscovererTagsEntryArrayOutput() DiscovererTagsEntryArrayOutput
	ToDiscovererTagsEntryArrayOutputWithContext(context.Context) DiscovererTagsEntryArrayOutput
}

type DiscovererTagsEntryArray []DiscovererTagsEntryInput

func (DiscovererTagsEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiscovererTagsEntry)(nil)).Elem()
}

func (i DiscovererTagsEntryArray) ToDiscovererTagsEntryArrayOutput() DiscovererTagsEntryArrayOutput {
	return i.ToDiscovererTagsEntryArrayOutputWithContext(context.Background())
}

func (i DiscovererTagsEntryArray) ToDiscovererTagsEntryArrayOutputWithContext(ctx context.Context) DiscovererTagsEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiscovererTagsEntryArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html
type DiscovererTagsEntryOutput struct{ *pulumi.OutputState }

func (DiscovererTagsEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiscovererTagsEntry)(nil)).Elem()
}

func (o DiscovererTagsEntryOutput) ToDiscovererTagsEntryOutput() DiscovererTagsEntryOutput {
	return o
}

func (o DiscovererTagsEntryOutput) ToDiscovererTagsEntryOutputWithContext(ctx context.Context) DiscovererTagsEntryOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html#cfn-eventschemas-discoverer-tagsentry-key
func (o DiscovererTagsEntryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DiscovererTagsEntry) string { return v.Key }).(pulumi.StringOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-discoverer-tagsentry.html#cfn-eventschemas-discoverer-tagsentry-value
func (o DiscovererTagsEntryOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DiscovererTagsEntry) string { return v.Value }).(pulumi.StringOutput)
}

type DiscovererTagsEntryArrayOutput struct{ *pulumi.OutputState }

func (DiscovererTagsEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiscovererTagsEntry)(nil)).Elem()
}

func (o DiscovererTagsEntryArrayOutput) ToDiscovererTagsEntryArrayOutput() DiscovererTagsEntryArrayOutput {
	return o
}

func (o DiscovererTagsEntryArrayOutput) ToDiscovererTagsEntryArrayOutputWithContext(ctx context.Context) DiscovererTagsEntryArrayOutput {
	return o
}

func (o DiscovererTagsEntryArrayOutput) Index(i pulumi.IntInput) DiscovererTagsEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiscovererTagsEntry {
		return vs[0].([]DiscovererTagsEntry)[vs[1].(int)]
	}).(DiscovererTagsEntryOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html
type RegistryTagsEntry struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html#cfn-eventschemas-registry-tagsentry-key
	Key string `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html#cfn-eventschemas-registry-tagsentry-value
	Value string `pulumi:"value"`
}

// RegistryTagsEntryInput is an input type that accepts RegistryTagsEntryArgs and RegistryTagsEntryOutput values.
// You can construct a concrete instance of `RegistryTagsEntryInput` via:
//
//          RegistryTagsEntryArgs{...}
type RegistryTagsEntryInput interface {
	pulumi.Input

	ToRegistryTagsEntryOutput() RegistryTagsEntryOutput
	ToRegistryTagsEntryOutputWithContext(context.Context) RegistryTagsEntryOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html
type RegistryTagsEntryArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html#cfn-eventschemas-registry-tagsentry-key
	Key pulumi.StringInput `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html#cfn-eventschemas-registry-tagsentry-value
	Value pulumi.StringInput `pulumi:"value"`
}

func (RegistryTagsEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryTagsEntry)(nil)).Elem()
}

func (i RegistryTagsEntryArgs) ToRegistryTagsEntryOutput() RegistryTagsEntryOutput {
	return i.ToRegistryTagsEntryOutputWithContext(context.Background())
}

func (i RegistryTagsEntryArgs) ToRegistryTagsEntryOutputWithContext(ctx context.Context) RegistryTagsEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryTagsEntryOutput)
}

// RegistryTagsEntryArrayInput is an input type that accepts RegistryTagsEntryArray and RegistryTagsEntryArrayOutput values.
// You can construct a concrete instance of `RegistryTagsEntryArrayInput` via:
//
//          RegistryTagsEntryArray{ RegistryTagsEntryArgs{...} }
type RegistryTagsEntryArrayInput interface {
	pulumi.Input

	ToRegistryTagsEntryArrayOutput() RegistryTagsEntryArrayOutput
	ToRegistryTagsEntryArrayOutputWithContext(context.Context) RegistryTagsEntryArrayOutput
}

type RegistryTagsEntryArray []RegistryTagsEntryInput

func (RegistryTagsEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryTagsEntry)(nil)).Elem()
}

func (i RegistryTagsEntryArray) ToRegistryTagsEntryArrayOutput() RegistryTagsEntryArrayOutput {
	return i.ToRegistryTagsEntryArrayOutputWithContext(context.Background())
}

func (i RegistryTagsEntryArray) ToRegistryTagsEntryArrayOutputWithContext(ctx context.Context) RegistryTagsEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryTagsEntryArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html
type RegistryTagsEntryOutput struct{ *pulumi.OutputState }

func (RegistryTagsEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryTagsEntry)(nil)).Elem()
}

func (o RegistryTagsEntryOutput) ToRegistryTagsEntryOutput() RegistryTagsEntryOutput {
	return o
}

func (o RegistryTagsEntryOutput) ToRegistryTagsEntryOutputWithContext(ctx context.Context) RegistryTagsEntryOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html#cfn-eventschemas-registry-tagsentry-key
func (o RegistryTagsEntryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v RegistryTagsEntry) string { return v.Key }).(pulumi.StringOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-registry-tagsentry.html#cfn-eventschemas-registry-tagsentry-value
func (o RegistryTagsEntryOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v RegistryTagsEntry) string { return v.Value }).(pulumi.StringOutput)
}

type RegistryTagsEntryArrayOutput struct{ *pulumi.OutputState }

func (RegistryTagsEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryTagsEntry)(nil)).Elem()
}

func (o RegistryTagsEntryArrayOutput) ToRegistryTagsEntryArrayOutput() RegistryTagsEntryArrayOutput {
	return o
}

func (o RegistryTagsEntryArrayOutput) ToRegistryTagsEntryArrayOutputWithContext(ctx context.Context) RegistryTagsEntryArrayOutput {
	return o
}

func (o RegistryTagsEntryArrayOutput) Index(i pulumi.IntInput) RegistryTagsEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryTagsEntry {
		return vs[0].([]RegistryTagsEntry)[vs[1].(int)]
	}).(RegistryTagsEntryOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html
type SchemaTagsEntry struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html#cfn-eventschemas-schema-tagsentry-key
	Key string `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html#cfn-eventschemas-schema-tagsentry-value
	Value string `pulumi:"value"`
}

// SchemaTagsEntryInput is an input type that accepts SchemaTagsEntryArgs and SchemaTagsEntryOutput values.
// You can construct a concrete instance of `SchemaTagsEntryInput` via:
//
//          SchemaTagsEntryArgs{...}
type SchemaTagsEntryInput interface {
	pulumi.Input

	ToSchemaTagsEntryOutput() SchemaTagsEntryOutput
	ToSchemaTagsEntryOutputWithContext(context.Context) SchemaTagsEntryOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html
type SchemaTagsEntryArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html#cfn-eventschemas-schema-tagsentry-key
	Key pulumi.StringInput `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html#cfn-eventschemas-schema-tagsentry-value
	Value pulumi.StringInput `pulumi:"value"`
}

func (SchemaTagsEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaTagsEntry)(nil)).Elem()
}

func (i SchemaTagsEntryArgs) ToSchemaTagsEntryOutput() SchemaTagsEntryOutput {
	return i.ToSchemaTagsEntryOutputWithContext(context.Background())
}

func (i SchemaTagsEntryArgs) ToSchemaTagsEntryOutputWithContext(ctx context.Context) SchemaTagsEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaTagsEntryOutput)
}

// SchemaTagsEntryArrayInput is an input type that accepts SchemaTagsEntryArray and SchemaTagsEntryArrayOutput values.
// You can construct a concrete instance of `SchemaTagsEntryArrayInput` via:
//
//          SchemaTagsEntryArray{ SchemaTagsEntryArgs{...} }
type SchemaTagsEntryArrayInput interface {
	pulumi.Input

	ToSchemaTagsEntryArrayOutput() SchemaTagsEntryArrayOutput
	ToSchemaTagsEntryArrayOutputWithContext(context.Context) SchemaTagsEntryArrayOutput
}

type SchemaTagsEntryArray []SchemaTagsEntryInput

func (SchemaTagsEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SchemaTagsEntry)(nil)).Elem()
}

func (i SchemaTagsEntryArray) ToSchemaTagsEntryArrayOutput() SchemaTagsEntryArrayOutput {
	return i.ToSchemaTagsEntryArrayOutputWithContext(context.Background())
}

func (i SchemaTagsEntryArray) ToSchemaTagsEntryArrayOutputWithContext(ctx context.Context) SchemaTagsEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaTagsEntryArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html
type SchemaTagsEntryOutput struct{ *pulumi.OutputState }

func (SchemaTagsEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaTagsEntry)(nil)).Elem()
}

func (o SchemaTagsEntryOutput) ToSchemaTagsEntryOutput() SchemaTagsEntryOutput {
	return o
}

func (o SchemaTagsEntryOutput) ToSchemaTagsEntryOutputWithContext(ctx context.Context) SchemaTagsEntryOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html#cfn-eventschemas-schema-tagsentry-key
func (o SchemaTagsEntryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v SchemaTagsEntry) string { return v.Key }).(pulumi.StringOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventschemas-schema-tagsentry.html#cfn-eventschemas-schema-tagsentry-value
func (o SchemaTagsEntryOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SchemaTagsEntry) string { return v.Value }).(pulumi.StringOutput)
}

type SchemaTagsEntryArrayOutput struct{ *pulumi.OutputState }

func (SchemaTagsEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SchemaTagsEntry)(nil)).Elem()
}

func (o SchemaTagsEntryArrayOutput) ToSchemaTagsEntryArrayOutput() SchemaTagsEntryArrayOutput {
	return o
}

func (o SchemaTagsEntryArrayOutput) ToSchemaTagsEntryArrayOutputWithContext(ctx context.Context) SchemaTagsEntryArrayOutput {
	return o
}

func (o SchemaTagsEntryArrayOutput) Index(i pulumi.IntInput) SchemaTagsEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SchemaTagsEntry {
		return vs[0].([]SchemaTagsEntry)[vs[1].(int)]
	}).(SchemaTagsEntryOutput)
}

func init() {
	pulumi.RegisterOutputType(DiscovererTagsEntryOutput{})
	pulumi.RegisterOutputType(DiscovererTagsEntryArrayOutput{})
	pulumi.RegisterOutputType(RegistryTagsEntryOutput{})
	pulumi.RegisterOutputType(RegistryTagsEntryArrayOutput{})
	pulumi.RegisterOutputType(SchemaTagsEntryOutput{})
	pulumi.RegisterOutputType(SchemaTagsEntryArrayOutput{})
}