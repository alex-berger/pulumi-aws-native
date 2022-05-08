// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devopsguru

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about notification channels you have configured with DevOps Guru.
type NotificationChannelConfig struct {
	Sns *NotificationChannelSnsChannelConfig `pulumi:"sns"`
}

// NotificationChannelConfigInput is an input type that accepts NotificationChannelConfigArgs and NotificationChannelConfigOutput values.
// You can construct a concrete instance of `NotificationChannelConfigInput` via:
//
//          NotificationChannelConfigArgs{...}
type NotificationChannelConfigInput interface {
	pulumi.Input

	ToNotificationChannelConfigOutput() NotificationChannelConfigOutput
	ToNotificationChannelConfigOutputWithContext(context.Context) NotificationChannelConfigOutput
}

// Information about notification channels you have configured with DevOps Guru.
type NotificationChannelConfigArgs struct {
	Sns NotificationChannelSnsChannelConfigPtrInput `pulumi:"sns"`
}

func (NotificationChannelConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelConfig)(nil)).Elem()
}

func (i NotificationChannelConfigArgs) ToNotificationChannelConfigOutput() NotificationChannelConfigOutput {
	return i.ToNotificationChannelConfigOutputWithContext(context.Background())
}

func (i NotificationChannelConfigArgs) ToNotificationChannelConfigOutputWithContext(ctx context.Context) NotificationChannelConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelConfigOutput)
}

// Information about notification channels you have configured with DevOps Guru.
type NotificationChannelConfigOutput struct{ *pulumi.OutputState }

func (NotificationChannelConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelConfig)(nil)).Elem()
}

func (o NotificationChannelConfigOutput) ToNotificationChannelConfigOutput() NotificationChannelConfigOutput {
	return o
}

func (o NotificationChannelConfigOutput) ToNotificationChannelConfigOutputWithContext(ctx context.Context) NotificationChannelConfigOutput {
	return o
}

func (o NotificationChannelConfigOutput) Sns() NotificationChannelSnsChannelConfigPtrOutput {
	return o.ApplyT(func(v NotificationChannelConfig) *NotificationChannelSnsChannelConfig { return v.Sns }).(NotificationChannelSnsChannelConfigPtrOutput)
}

// Information about a notification channel configured in DevOps Guru to send notifications when insights are created.
type NotificationChannelSnsChannelConfig struct {
	TopicArn *string `pulumi:"topicArn"`
}

// NotificationChannelSnsChannelConfigInput is an input type that accepts NotificationChannelSnsChannelConfigArgs and NotificationChannelSnsChannelConfigOutput values.
// You can construct a concrete instance of `NotificationChannelSnsChannelConfigInput` via:
//
//          NotificationChannelSnsChannelConfigArgs{...}
type NotificationChannelSnsChannelConfigInput interface {
	pulumi.Input

	ToNotificationChannelSnsChannelConfigOutput() NotificationChannelSnsChannelConfigOutput
	ToNotificationChannelSnsChannelConfigOutputWithContext(context.Context) NotificationChannelSnsChannelConfigOutput
}

// Information about a notification channel configured in DevOps Guru to send notifications when insights are created.
type NotificationChannelSnsChannelConfigArgs struct {
	TopicArn pulumi.StringPtrInput `pulumi:"topicArn"`
}

func (NotificationChannelSnsChannelConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelSnsChannelConfig)(nil)).Elem()
}

func (i NotificationChannelSnsChannelConfigArgs) ToNotificationChannelSnsChannelConfigOutput() NotificationChannelSnsChannelConfigOutput {
	return i.ToNotificationChannelSnsChannelConfigOutputWithContext(context.Background())
}

func (i NotificationChannelSnsChannelConfigArgs) ToNotificationChannelSnsChannelConfigOutputWithContext(ctx context.Context) NotificationChannelSnsChannelConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSnsChannelConfigOutput)
}

func (i NotificationChannelSnsChannelConfigArgs) ToNotificationChannelSnsChannelConfigPtrOutput() NotificationChannelSnsChannelConfigPtrOutput {
	return i.ToNotificationChannelSnsChannelConfigPtrOutputWithContext(context.Background())
}

func (i NotificationChannelSnsChannelConfigArgs) ToNotificationChannelSnsChannelConfigPtrOutputWithContext(ctx context.Context) NotificationChannelSnsChannelConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSnsChannelConfigOutput).ToNotificationChannelSnsChannelConfigPtrOutputWithContext(ctx)
}

// NotificationChannelSnsChannelConfigPtrInput is an input type that accepts NotificationChannelSnsChannelConfigArgs, NotificationChannelSnsChannelConfigPtr and NotificationChannelSnsChannelConfigPtrOutput values.
// You can construct a concrete instance of `NotificationChannelSnsChannelConfigPtrInput` via:
//
//          NotificationChannelSnsChannelConfigArgs{...}
//
//  or:
//
//          nil
type NotificationChannelSnsChannelConfigPtrInput interface {
	pulumi.Input

	ToNotificationChannelSnsChannelConfigPtrOutput() NotificationChannelSnsChannelConfigPtrOutput
	ToNotificationChannelSnsChannelConfigPtrOutputWithContext(context.Context) NotificationChannelSnsChannelConfigPtrOutput
}

type notificationChannelSnsChannelConfigPtrType NotificationChannelSnsChannelConfigArgs

func NotificationChannelSnsChannelConfigPtr(v *NotificationChannelSnsChannelConfigArgs) NotificationChannelSnsChannelConfigPtrInput {
	return (*notificationChannelSnsChannelConfigPtrType)(v)
}

func (*notificationChannelSnsChannelConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannelSnsChannelConfig)(nil)).Elem()
}

func (i *notificationChannelSnsChannelConfigPtrType) ToNotificationChannelSnsChannelConfigPtrOutput() NotificationChannelSnsChannelConfigPtrOutput {
	return i.ToNotificationChannelSnsChannelConfigPtrOutputWithContext(context.Background())
}

func (i *notificationChannelSnsChannelConfigPtrType) ToNotificationChannelSnsChannelConfigPtrOutputWithContext(ctx context.Context) NotificationChannelSnsChannelConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationChannelSnsChannelConfigPtrOutput)
}

// Information about a notification channel configured in DevOps Guru to send notifications when insights are created.
type NotificationChannelSnsChannelConfigOutput struct{ *pulumi.OutputState }

func (NotificationChannelSnsChannelConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationChannelSnsChannelConfig)(nil)).Elem()
}

func (o NotificationChannelSnsChannelConfigOutput) ToNotificationChannelSnsChannelConfigOutput() NotificationChannelSnsChannelConfigOutput {
	return o
}

func (o NotificationChannelSnsChannelConfigOutput) ToNotificationChannelSnsChannelConfigOutputWithContext(ctx context.Context) NotificationChannelSnsChannelConfigOutput {
	return o
}

func (o NotificationChannelSnsChannelConfigOutput) ToNotificationChannelSnsChannelConfigPtrOutput() NotificationChannelSnsChannelConfigPtrOutput {
	return o.ToNotificationChannelSnsChannelConfigPtrOutputWithContext(context.Background())
}

func (o NotificationChannelSnsChannelConfigOutput) ToNotificationChannelSnsChannelConfigPtrOutputWithContext(ctx context.Context) NotificationChannelSnsChannelConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationChannelSnsChannelConfig) *NotificationChannelSnsChannelConfig {
		return &v
	}).(NotificationChannelSnsChannelConfigPtrOutput)
}

func (o NotificationChannelSnsChannelConfigOutput) TopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationChannelSnsChannelConfig) *string { return v.TopicArn }).(pulumi.StringPtrOutput)
}

type NotificationChannelSnsChannelConfigPtrOutput struct{ *pulumi.OutputState }

func (NotificationChannelSnsChannelConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationChannelSnsChannelConfig)(nil)).Elem()
}

func (o NotificationChannelSnsChannelConfigPtrOutput) ToNotificationChannelSnsChannelConfigPtrOutput() NotificationChannelSnsChannelConfigPtrOutput {
	return o
}

func (o NotificationChannelSnsChannelConfigPtrOutput) ToNotificationChannelSnsChannelConfigPtrOutputWithContext(ctx context.Context) NotificationChannelSnsChannelConfigPtrOutput {
	return o
}

func (o NotificationChannelSnsChannelConfigPtrOutput) Elem() NotificationChannelSnsChannelConfigOutput {
	return o.ApplyT(func(v *NotificationChannelSnsChannelConfig) NotificationChannelSnsChannelConfig {
		if v != nil {
			return *v
		}
		var ret NotificationChannelSnsChannelConfig
		return ret
	}).(NotificationChannelSnsChannelConfigOutput)
}

func (o NotificationChannelSnsChannelConfigPtrOutput) TopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationChannelSnsChannelConfig) *string {
		if v == nil {
			return nil
		}
		return v.TopicArn
	}).(pulumi.StringPtrOutput)
}

// CloudFormation resource for DevOps Guru to monitor
type ResourceCollectionCloudFormationCollectionFilter struct {
	// An array of CloudFormation stack names.
	StackNames []string `pulumi:"stackNames"`
}

// ResourceCollectionCloudFormationCollectionFilterInput is an input type that accepts ResourceCollectionCloudFormationCollectionFilterArgs and ResourceCollectionCloudFormationCollectionFilterOutput values.
// You can construct a concrete instance of `ResourceCollectionCloudFormationCollectionFilterInput` via:
//
//          ResourceCollectionCloudFormationCollectionFilterArgs{...}
type ResourceCollectionCloudFormationCollectionFilterInput interface {
	pulumi.Input

	ToResourceCollectionCloudFormationCollectionFilterOutput() ResourceCollectionCloudFormationCollectionFilterOutput
	ToResourceCollectionCloudFormationCollectionFilterOutputWithContext(context.Context) ResourceCollectionCloudFormationCollectionFilterOutput
}

// CloudFormation resource for DevOps Guru to monitor
type ResourceCollectionCloudFormationCollectionFilterArgs struct {
	// An array of CloudFormation stack names.
	StackNames pulumi.StringArrayInput `pulumi:"stackNames"`
}

func (ResourceCollectionCloudFormationCollectionFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollectionCloudFormationCollectionFilter)(nil)).Elem()
}

func (i ResourceCollectionCloudFormationCollectionFilterArgs) ToResourceCollectionCloudFormationCollectionFilterOutput() ResourceCollectionCloudFormationCollectionFilterOutput {
	return i.ToResourceCollectionCloudFormationCollectionFilterOutputWithContext(context.Background())
}

func (i ResourceCollectionCloudFormationCollectionFilterArgs) ToResourceCollectionCloudFormationCollectionFilterOutputWithContext(ctx context.Context) ResourceCollectionCloudFormationCollectionFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCollectionCloudFormationCollectionFilterOutput)
}

func (i ResourceCollectionCloudFormationCollectionFilterArgs) ToResourceCollectionCloudFormationCollectionFilterPtrOutput() ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return i.ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(context.Background())
}

func (i ResourceCollectionCloudFormationCollectionFilterArgs) ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(ctx context.Context) ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCollectionCloudFormationCollectionFilterOutput).ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(ctx)
}

// ResourceCollectionCloudFormationCollectionFilterPtrInput is an input type that accepts ResourceCollectionCloudFormationCollectionFilterArgs, ResourceCollectionCloudFormationCollectionFilterPtr and ResourceCollectionCloudFormationCollectionFilterPtrOutput values.
// You can construct a concrete instance of `ResourceCollectionCloudFormationCollectionFilterPtrInput` via:
//
//          ResourceCollectionCloudFormationCollectionFilterArgs{...}
//
//  or:
//
//          nil
type ResourceCollectionCloudFormationCollectionFilterPtrInput interface {
	pulumi.Input

	ToResourceCollectionCloudFormationCollectionFilterPtrOutput() ResourceCollectionCloudFormationCollectionFilterPtrOutput
	ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(context.Context) ResourceCollectionCloudFormationCollectionFilterPtrOutput
}

type resourceCollectionCloudFormationCollectionFilterPtrType ResourceCollectionCloudFormationCollectionFilterArgs

func ResourceCollectionCloudFormationCollectionFilterPtr(v *ResourceCollectionCloudFormationCollectionFilterArgs) ResourceCollectionCloudFormationCollectionFilterPtrInput {
	return (*resourceCollectionCloudFormationCollectionFilterPtrType)(v)
}

func (*resourceCollectionCloudFormationCollectionFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceCollectionCloudFormationCollectionFilter)(nil)).Elem()
}

func (i *resourceCollectionCloudFormationCollectionFilterPtrType) ToResourceCollectionCloudFormationCollectionFilterPtrOutput() ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return i.ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(context.Background())
}

func (i *resourceCollectionCloudFormationCollectionFilterPtrType) ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(ctx context.Context) ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCollectionCloudFormationCollectionFilterPtrOutput)
}

// CloudFormation resource for DevOps Guru to monitor
type ResourceCollectionCloudFormationCollectionFilterOutput struct{ *pulumi.OutputState }

func (ResourceCollectionCloudFormationCollectionFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollectionCloudFormationCollectionFilter)(nil)).Elem()
}

func (o ResourceCollectionCloudFormationCollectionFilterOutput) ToResourceCollectionCloudFormationCollectionFilterOutput() ResourceCollectionCloudFormationCollectionFilterOutput {
	return o
}

func (o ResourceCollectionCloudFormationCollectionFilterOutput) ToResourceCollectionCloudFormationCollectionFilterOutputWithContext(ctx context.Context) ResourceCollectionCloudFormationCollectionFilterOutput {
	return o
}

func (o ResourceCollectionCloudFormationCollectionFilterOutput) ToResourceCollectionCloudFormationCollectionFilterPtrOutput() ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return o.ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(context.Background())
}

func (o ResourceCollectionCloudFormationCollectionFilterOutput) ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(ctx context.Context) ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceCollectionCloudFormationCollectionFilter) *ResourceCollectionCloudFormationCollectionFilter {
		return &v
	}).(ResourceCollectionCloudFormationCollectionFilterPtrOutput)
}

// An array of CloudFormation stack names.
func (o ResourceCollectionCloudFormationCollectionFilterOutput) StackNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceCollectionCloudFormationCollectionFilter) []string { return v.StackNames }).(pulumi.StringArrayOutput)
}

type ResourceCollectionCloudFormationCollectionFilterPtrOutput struct{ *pulumi.OutputState }

func (ResourceCollectionCloudFormationCollectionFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceCollectionCloudFormationCollectionFilter)(nil)).Elem()
}

func (o ResourceCollectionCloudFormationCollectionFilterPtrOutput) ToResourceCollectionCloudFormationCollectionFilterPtrOutput() ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return o
}

func (o ResourceCollectionCloudFormationCollectionFilterPtrOutput) ToResourceCollectionCloudFormationCollectionFilterPtrOutputWithContext(ctx context.Context) ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return o
}

func (o ResourceCollectionCloudFormationCollectionFilterPtrOutput) Elem() ResourceCollectionCloudFormationCollectionFilterOutput {
	return o.ApplyT(func(v *ResourceCollectionCloudFormationCollectionFilter) ResourceCollectionCloudFormationCollectionFilter {
		if v != nil {
			return *v
		}
		var ret ResourceCollectionCloudFormationCollectionFilter
		return ret
	}).(ResourceCollectionCloudFormationCollectionFilterOutput)
}

// An array of CloudFormation stack names.
func (o ResourceCollectionCloudFormationCollectionFilterPtrOutput) StackNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceCollectionCloudFormationCollectionFilter) []string {
		if v == nil {
			return nil
		}
		return v.StackNames
	}).(pulumi.StringArrayOutput)
}

// Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
type ResourceCollectionFilter struct {
	CloudFormation *ResourceCollectionCloudFormationCollectionFilter `pulumi:"cloudFormation"`
	Tags           []ResourceCollectionTagCollection                 `pulumi:"tags"`
}

// ResourceCollectionFilterInput is an input type that accepts ResourceCollectionFilterArgs and ResourceCollectionFilterOutput values.
// You can construct a concrete instance of `ResourceCollectionFilterInput` via:
//
//          ResourceCollectionFilterArgs{...}
type ResourceCollectionFilterInput interface {
	pulumi.Input

	ToResourceCollectionFilterOutput() ResourceCollectionFilterOutput
	ToResourceCollectionFilterOutputWithContext(context.Context) ResourceCollectionFilterOutput
}

// Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
type ResourceCollectionFilterArgs struct {
	CloudFormation ResourceCollectionCloudFormationCollectionFilterPtrInput `pulumi:"cloudFormation"`
	Tags           ResourceCollectionTagCollectionArrayInput                `pulumi:"tags"`
}

func (ResourceCollectionFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollectionFilter)(nil)).Elem()
}

func (i ResourceCollectionFilterArgs) ToResourceCollectionFilterOutput() ResourceCollectionFilterOutput {
	return i.ToResourceCollectionFilterOutputWithContext(context.Background())
}

func (i ResourceCollectionFilterArgs) ToResourceCollectionFilterOutputWithContext(ctx context.Context) ResourceCollectionFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCollectionFilterOutput)
}

// Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
type ResourceCollectionFilterOutput struct{ *pulumi.OutputState }

func (ResourceCollectionFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollectionFilter)(nil)).Elem()
}

func (o ResourceCollectionFilterOutput) ToResourceCollectionFilterOutput() ResourceCollectionFilterOutput {
	return o
}

func (o ResourceCollectionFilterOutput) ToResourceCollectionFilterOutputWithContext(ctx context.Context) ResourceCollectionFilterOutput {
	return o
}

func (o ResourceCollectionFilterOutput) CloudFormation() ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return o.ApplyT(func(v ResourceCollectionFilter) *ResourceCollectionCloudFormationCollectionFilter {
		return v.CloudFormation
	}).(ResourceCollectionCloudFormationCollectionFilterPtrOutput)
}

func (o ResourceCollectionFilterOutput) Tags() ResourceCollectionTagCollectionArrayOutput {
	return o.ApplyT(func(v ResourceCollectionFilter) []ResourceCollectionTagCollection { return v.Tags }).(ResourceCollectionTagCollectionArrayOutput)
}

type ResourceCollectionFilterPtrOutput struct{ *pulumi.OutputState }

func (ResourceCollectionFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceCollectionFilter)(nil)).Elem()
}

func (o ResourceCollectionFilterPtrOutput) ToResourceCollectionFilterPtrOutput() ResourceCollectionFilterPtrOutput {
	return o
}

func (o ResourceCollectionFilterPtrOutput) ToResourceCollectionFilterPtrOutputWithContext(ctx context.Context) ResourceCollectionFilterPtrOutput {
	return o
}

func (o ResourceCollectionFilterPtrOutput) Elem() ResourceCollectionFilterOutput {
	return o.ApplyT(func(v *ResourceCollectionFilter) ResourceCollectionFilter {
		if v != nil {
			return *v
		}
		var ret ResourceCollectionFilter
		return ret
	}).(ResourceCollectionFilterOutput)
}

func (o ResourceCollectionFilterPtrOutput) CloudFormation() ResourceCollectionCloudFormationCollectionFilterPtrOutput {
	return o.ApplyT(func(v *ResourceCollectionFilter) *ResourceCollectionCloudFormationCollectionFilter {
		if v == nil {
			return nil
		}
		return v.CloudFormation
	}).(ResourceCollectionCloudFormationCollectionFilterPtrOutput)
}

func (o ResourceCollectionFilterPtrOutput) Tags() ResourceCollectionTagCollectionArrayOutput {
	return o.ApplyT(func(v *ResourceCollectionFilter) []ResourceCollectionTagCollection {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(ResourceCollectionTagCollectionArrayOutput)
}

// Tagged resource for DevOps Guru to monitor
type ResourceCollectionTagCollection struct {
	// A Tag key for DevOps Guru app boundary.
	AppBoundaryKey *string `pulumi:"appBoundaryKey"`
	// Tag values of DevOps Guru app boundary.
	TagValues []string `pulumi:"tagValues"`
}

// ResourceCollectionTagCollectionInput is an input type that accepts ResourceCollectionTagCollectionArgs and ResourceCollectionTagCollectionOutput values.
// You can construct a concrete instance of `ResourceCollectionTagCollectionInput` via:
//
//          ResourceCollectionTagCollectionArgs{...}
type ResourceCollectionTagCollectionInput interface {
	pulumi.Input

	ToResourceCollectionTagCollectionOutput() ResourceCollectionTagCollectionOutput
	ToResourceCollectionTagCollectionOutputWithContext(context.Context) ResourceCollectionTagCollectionOutput
}

// Tagged resource for DevOps Guru to monitor
type ResourceCollectionTagCollectionArgs struct {
	// A Tag key for DevOps Guru app boundary.
	AppBoundaryKey pulumi.StringPtrInput `pulumi:"appBoundaryKey"`
	// Tag values of DevOps Guru app boundary.
	TagValues pulumi.StringArrayInput `pulumi:"tagValues"`
}

func (ResourceCollectionTagCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollectionTagCollection)(nil)).Elem()
}

func (i ResourceCollectionTagCollectionArgs) ToResourceCollectionTagCollectionOutput() ResourceCollectionTagCollectionOutput {
	return i.ToResourceCollectionTagCollectionOutputWithContext(context.Background())
}

func (i ResourceCollectionTagCollectionArgs) ToResourceCollectionTagCollectionOutputWithContext(ctx context.Context) ResourceCollectionTagCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCollectionTagCollectionOutput)
}

// ResourceCollectionTagCollectionArrayInput is an input type that accepts ResourceCollectionTagCollectionArray and ResourceCollectionTagCollectionArrayOutput values.
// You can construct a concrete instance of `ResourceCollectionTagCollectionArrayInput` via:
//
//          ResourceCollectionTagCollectionArray{ ResourceCollectionTagCollectionArgs{...} }
type ResourceCollectionTagCollectionArrayInput interface {
	pulumi.Input

	ToResourceCollectionTagCollectionArrayOutput() ResourceCollectionTagCollectionArrayOutput
	ToResourceCollectionTagCollectionArrayOutputWithContext(context.Context) ResourceCollectionTagCollectionArrayOutput
}

type ResourceCollectionTagCollectionArray []ResourceCollectionTagCollectionInput

func (ResourceCollectionTagCollectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceCollectionTagCollection)(nil)).Elem()
}

func (i ResourceCollectionTagCollectionArray) ToResourceCollectionTagCollectionArrayOutput() ResourceCollectionTagCollectionArrayOutput {
	return i.ToResourceCollectionTagCollectionArrayOutputWithContext(context.Background())
}

func (i ResourceCollectionTagCollectionArray) ToResourceCollectionTagCollectionArrayOutputWithContext(ctx context.Context) ResourceCollectionTagCollectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCollectionTagCollectionArrayOutput)
}

// Tagged resource for DevOps Guru to monitor
type ResourceCollectionTagCollectionOutput struct{ *pulumi.OutputState }

func (ResourceCollectionTagCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollectionTagCollection)(nil)).Elem()
}

func (o ResourceCollectionTagCollectionOutput) ToResourceCollectionTagCollectionOutput() ResourceCollectionTagCollectionOutput {
	return o
}

func (o ResourceCollectionTagCollectionOutput) ToResourceCollectionTagCollectionOutputWithContext(ctx context.Context) ResourceCollectionTagCollectionOutput {
	return o
}

// A Tag key for DevOps Guru app boundary.
func (o ResourceCollectionTagCollectionOutput) AppBoundaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceCollectionTagCollection) *string { return v.AppBoundaryKey }).(pulumi.StringPtrOutput)
}

// Tag values of DevOps Guru app boundary.
func (o ResourceCollectionTagCollectionOutput) TagValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceCollectionTagCollection) []string { return v.TagValues }).(pulumi.StringArrayOutput)
}

type ResourceCollectionTagCollectionArrayOutput struct{ *pulumi.OutputState }

func (ResourceCollectionTagCollectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceCollectionTagCollection)(nil)).Elem()
}

func (o ResourceCollectionTagCollectionArrayOutput) ToResourceCollectionTagCollectionArrayOutput() ResourceCollectionTagCollectionArrayOutput {
	return o
}

func (o ResourceCollectionTagCollectionArrayOutput) ToResourceCollectionTagCollectionArrayOutputWithContext(ctx context.Context) ResourceCollectionTagCollectionArrayOutput {
	return o
}

func (o ResourceCollectionTagCollectionArrayOutput) Index(i pulumi.IntInput) ResourceCollectionTagCollectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceCollectionTagCollection {
		return vs[0].([]ResourceCollectionTagCollection)[vs[1].(int)]
	}).(ResourceCollectionTagCollectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationChannelConfigInput)(nil)).Elem(), NotificationChannelConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationChannelSnsChannelConfigInput)(nil)).Elem(), NotificationChannelSnsChannelConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationChannelSnsChannelConfigPtrInput)(nil)).Elem(), NotificationChannelSnsChannelConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCollectionCloudFormationCollectionFilterInput)(nil)).Elem(), ResourceCollectionCloudFormationCollectionFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCollectionCloudFormationCollectionFilterPtrInput)(nil)).Elem(), ResourceCollectionCloudFormationCollectionFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCollectionFilterInput)(nil)).Elem(), ResourceCollectionFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCollectionTagCollectionInput)(nil)).Elem(), ResourceCollectionTagCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceCollectionTagCollectionArrayInput)(nil)).Elem(), ResourceCollectionTagCollectionArray{})
	pulumi.RegisterOutputType(NotificationChannelConfigOutput{})
	pulumi.RegisterOutputType(NotificationChannelSnsChannelConfigOutput{})
	pulumi.RegisterOutputType(NotificationChannelSnsChannelConfigPtrOutput{})
	pulumi.RegisterOutputType(ResourceCollectionCloudFormationCollectionFilterOutput{})
	pulumi.RegisterOutputType(ResourceCollectionCloudFormationCollectionFilterPtrOutput{})
	pulumi.RegisterOutputType(ResourceCollectionFilterOutput{})
	pulumi.RegisterOutputType(ResourceCollectionFilterPtrOutput{})
	pulumi.RegisterOutputType(ResourceCollectionTagCollectionOutput{})
	pulumi.RegisterOutputType(ResourceCollectionTagCollectionArrayOutput{})
}
