// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CloudTrail supports data event logging for Amazon S3 objects and AWS Lambda functions. You can specify up to 250 resources for an individual event selector, but the total number of data resources cannot exceed 250 across all event selectors in a trail. This limit does not apply if you configure resource logging for all data events.
type TrailDataResource struct {
	// The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.
	Type string `pulumi:"type"`
	// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.
	Values []string `pulumi:"values"`
}

// TrailDataResourceInput is an input type that accepts TrailDataResourceArgs and TrailDataResourceOutput values.
// You can construct a concrete instance of `TrailDataResourceInput` via:
//
//          TrailDataResourceArgs{...}
type TrailDataResourceInput interface {
	pulumi.Input

	ToTrailDataResourceOutput() TrailDataResourceOutput
	ToTrailDataResourceOutputWithContext(context.Context) TrailDataResourceOutput
}

// CloudTrail supports data event logging for Amazon S3 objects and AWS Lambda functions. You can specify up to 250 resources for an individual event selector, but the total number of data resources cannot exceed 250 across all event selectors in a trail. This limit does not apply if you configure resource logging for all data events.
type TrailDataResourceArgs struct {
	// The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.
	Type pulumi.StringInput `pulumi:"type"`
	// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (TrailDataResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailDataResource)(nil)).Elem()
}

func (i TrailDataResourceArgs) ToTrailDataResourceOutput() TrailDataResourceOutput {
	return i.ToTrailDataResourceOutputWithContext(context.Background())
}

func (i TrailDataResourceArgs) ToTrailDataResourceOutputWithContext(ctx context.Context) TrailDataResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailDataResourceOutput)
}

// TrailDataResourceArrayInput is an input type that accepts TrailDataResourceArray and TrailDataResourceArrayOutput values.
// You can construct a concrete instance of `TrailDataResourceArrayInput` via:
//
//          TrailDataResourceArray{ TrailDataResourceArgs{...} }
type TrailDataResourceArrayInput interface {
	pulumi.Input

	ToTrailDataResourceArrayOutput() TrailDataResourceArrayOutput
	ToTrailDataResourceArrayOutputWithContext(context.Context) TrailDataResourceArrayOutput
}

type TrailDataResourceArray []TrailDataResourceInput

func (TrailDataResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailDataResource)(nil)).Elem()
}

func (i TrailDataResourceArray) ToTrailDataResourceArrayOutput() TrailDataResourceArrayOutput {
	return i.ToTrailDataResourceArrayOutputWithContext(context.Background())
}

func (i TrailDataResourceArray) ToTrailDataResourceArrayOutputWithContext(ctx context.Context) TrailDataResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailDataResourceArrayOutput)
}

// CloudTrail supports data event logging for Amazon S3 objects and AWS Lambda functions. You can specify up to 250 resources for an individual event selector, but the total number of data resources cannot exceed 250 across all event selectors in a trail. This limit does not apply if you configure resource logging for all data events.
type TrailDataResourceOutput struct{ *pulumi.OutputState }

func (TrailDataResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailDataResource)(nil)).Elem()
}

func (o TrailDataResourceOutput) ToTrailDataResourceOutput() TrailDataResourceOutput {
	return o
}

func (o TrailDataResourceOutput) ToTrailDataResourceOutputWithContext(ctx context.Context) TrailDataResourceOutput {
	return o
}

// The resource type in which you want to log data events. You can specify AWS::S3::Object or AWS::Lambda::Function resources.
func (o TrailDataResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TrailDataResource) string { return v.Type }).(pulumi.StringOutput)
}

// An array of Amazon Resource Name (ARN) strings or partial ARN strings for the specified objects.
func (o TrailDataResourceOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TrailDataResource) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type TrailDataResourceArrayOutput struct{ *pulumi.OutputState }

func (TrailDataResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailDataResource)(nil)).Elem()
}

func (o TrailDataResourceArrayOutput) ToTrailDataResourceArrayOutput() TrailDataResourceArrayOutput {
	return o
}

func (o TrailDataResourceArrayOutput) ToTrailDataResourceArrayOutputWithContext(ctx context.Context) TrailDataResourceArrayOutput {
	return o
}

func (o TrailDataResourceArrayOutput) Index(i pulumi.IntInput) TrailDataResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrailDataResource {
		return vs[0].([]TrailDataResource)[vs[1].(int)]
	}).(TrailDataResourceOutput)
}

// The type of email sending events to publish to the event destination.
type TrailEventSelector struct {
	DataResources []TrailDataResource `pulumi:"dataResources"`
	// An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing "kms.amazonaws.com". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.
	ExcludeManagementEventSources []string `pulumi:"excludeManagementEventSources"`
	// Specify if you want your event selector to include management events for your trail.
	IncludeManagementEvents *bool `pulumi:"includeManagementEvents"`
	// Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.
	ReadWriteType *TrailEventSelectorReadWriteType `pulumi:"readWriteType"`
}

// TrailEventSelectorInput is an input type that accepts TrailEventSelectorArgs and TrailEventSelectorOutput values.
// You can construct a concrete instance of `TrailEventSelectorInput` via:
//
//          TrailEventSelectorArgs{...}
type TrailEventSelectorInput interface {
	pulumi.Input

	ToTrailEventSelectorOutput() TrailEventSelectorOutput
	ToTrailEventSelectorOutputWithContext(context.Context) TrailEventSelectorOutput
}

// The type of email sending events to publish to the event destination.
type TrailEventSelectorArgs struct {
	DataResources TrailDataResourceArrayInput `pulumi:"dataResources"`
	// An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing "kms.amazonaws.com". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.
	ExcludeManagementEventSources pulumi.StringArrayInput `pulumi:"excludeManagementEventSources"`
	// Specify if you want your event selector to include management events for your trail.
	IncludeManagementEvents pulumi.BoolPtrInput `pulumi:"includeManagementEvents"`
	// Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.
	ReadWriteType TrailEventSelectorReadWriteTypePtrInput `pulumi:"readWriteType"`
}

func (TrailEventSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailEventSelector)(nil)).Elem()
}

func (i TrailEventSelectorArgs) ToTrailEventSelectorOutput() TrailEventSelectorOutput {
	return i.ToTrailEventSelectorOutputWithContext(context.Background())
}

func (i TrailEventSelectorArgs) ToTrailEventSelectorOutputWithContext(ctx context.Context) TrailEventSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailEventSelectorOutput)
}

// TrailEventSelectorArrayInput is an input type that accepts TrailEventSelectorArray and TrailEventSelectorArrayOutput values.
// You can construct a concrete instance of `TrailEventSelectorArrayInput` via:
//
//          TrailEventSelectorArray{ TrailEventSelectorArgs{...} }
type TrailEventSelectorArrayInput interface {
	pulumi.Input

	ToTrailEventSelectorArrayOutput() TrailEventSelectorArrayOutput
	ToTrailEventSelectorArrayOutputWithContext(context.Context) TrailEventSelectorArrayOutput
}

type TrailEventSelectorArray []TrailEventSelectorInput

func (TrailEventSelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailEventSelector)(nil)).Elem()
}

func (i TrailEventSelectorArray) ToTrailEventSelectorArrayOutput() TrailEventSelectorArrayOutput {
	return i.ToTrailEventSelectorArrayOutputWithContext(context.Background())
}

func (i TrailEventSelectorArray) ToTrailEventSelectorArrayOutputWithContext(ctx context.Context) TrailEventSelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailEventSelectorArrayOutput)
}

// The type of email sending events to publish to the event destination.
type TrailEventSelectorOutput struct{ *pulumi.OutputState }

func (TrailEventSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailEventSelector)(nil)).Elem()
}

func (o TrailEventSelectorOutput) ToTrailEventSelectorOutput() TrailEventSelectorOutput {
	return o
}

func (o TrailEventSelectorOutput) ToTrailEventSelectorOutputWithContext(ctx context.Context) TrailEventSelectorOutput {
	return o
}

func (o TrailEventSelectorOutput) DataResources() TrailDataResourceArrayOutput {
	return o.ApplyT(func(v TrailEventSelector) []TrailDataResource { return v.DataResources }).(TrailDataResourceArrayOutput)
}

// An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing "kms.amazonaws.com". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.
func (o TrailEventSelectorOutput) ExcludeManagementEventSources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TrailEventSelector) []string { return v.ExcludeManagementEventSources }).(pulumi.StringArrayOutput)
}

// Specify if you want your event selector to include management events for your trail.
func (o TrailEventSelectorOutput) IncludeManagementEvents() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TrailEventSelector) *bool { return v.IncludeManagementEvents }).(pulumi.BoolPtrOutput)
}

// Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.
func (o TrailEventSelectorOutput) ReadWriteType() TrailEventSelectorReadWriteTypePtrOutput {
	return o.ApplyT(func(v TrailEventSelector) *TrailEventSelectorReadWriteType { return v.ReadWriteType }).(TrailEventSelectorReadWriteTypePtrOutput)
}

type TrailEventSelectorArrayOutput struct{ *pulumi.OutputState }

func (TrailEventSelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailEventSelector)(nil)).Elem()
}

func (o TrailEventSelectorArrayOutput) ToTrailEventSelectorArrayOutput() TrailEventSelectorArrayOutput {
	return o
}

func (o TrailEventSelectorArrayOutput) ToTrailEventSelectorArrayOutputWithContext(ctx context.Context) TrailEventSelectorArrayOutput {
	return o
}

func (o TrailEventSelectorArrayOutput) Index(i pulumi.IntInput) TrailEventSelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrailEventSelector {
		return vs[0].([]TrailEventSelector)[vs[1].(int)]
	}).(TrailEventSelectorOutput)
}

// A string that contains insight types that are logged on a trail.
type TrailInsightSelector struct {
	// The type of insight to log on a trail.
	InsightType *string `pulumi:"insightType"`
}

// TrailInsightSelectorInput is an input type that accepts TrailInsightSelectorArgs and TrailInsightSelectorOutput values.
// You can construct a concrete instance of `TrailInsightSelectorInput` via:
//
//          TrailInsightSelectorArgs{...}
type TrailInsightSelectorInput interface {
	pulumi.Input

	ToTrailInsightSelectorOutput() TrailInsightSelectorOutput
	ToTrailInsightSelectorOutputWithContext(context.Context) TrailInsightSelectorOutput
}

// A string that contains insight types that are logged on a trail.
type TrailInsightSelectorArgs struct {
	// The type of insight to log on a trail.
	InsightType pulumi.StringPtrInput `pulumi:"insightType"`
}

func (TrailInsightSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailInsightSelector)(nil)).Elem()
}

func (i TrailInsightSelectorArgs) ToTrailInsightSelectorOutput() TrailInsightSelectorOutput {
	return i.ToTrailInsightSelectorOutputWithContext(context.Background())
}

func (i TrailInsightSelectorArgs) ToTrailInsightSelectorOutputWithContext(ctx context.Context) TrailInsightSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailInsightSelectorOutput)
}

// TrailInsightSelectorArrayInput is an input type that accepts TrailInsightSelectorArray and TrailInsightSelectorArrayOutput values.
// You can construct a concrete instance of `TrailInsightSelectorArrayInput` via:
//
//          TrailInsightSelectorArray{ TrailInsightSelectorArgs{...} }
type TrailInsightSelectorArrayInput interface {
	pulumi.Input

	ToTrailInsightSelectorArrayOutput() TrailInsightSelectorArrayOutput
	ToTrailInsightSelectorArrayOutputWithContext(context.Context) TrailInsightSelectorArrayOutput
}

type TrailInsightSelectorArray []TrailInsightSelectorInput

func (TrailInsightSelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailInsightSelector)(nil)).Elem()
}

func (i TrailInsightSelectorArray) ToTrailInsightSelectorArrayOutput() TrailInsightSelectorArrayOutput {
	return i.ToTrailInsightSelectorArrayOutputWithContext(context.Background())
}

func (i TrailInsightSelectorArray) ToTrailInsightSelectorArrayOutputWithContext(ctx context.Context) TrailInsightSelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailInsightSelectorArrayOutput)
}

// A string that contains insight types that are logged on a trail.
type TrailInsightSelectorOutput struct{ *pulumi.OutputState }

func (TrailInsightSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailInsightSelector)(nil)).Elem()
}

func (o TrailInsightSelectorOutput) ToTrailInsightSelectorOutput() TrailInsightSelectorOutput {
	return o
}

func (o TrailInsightSelectorOutput) ToTrailInsightSelectorOutputWithContext(ctx context.Context) TrailInsightSelectorOutput {
	return o
}

// The type of insight to log on a trail.
func (o TrailInsightSelectorOutput) InsightType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrailInsightSelector) *string { return v.InsightType }).(pulumi.StringPtrOutput)
}

type TrailInsightSelectorArrayOutput struct{ *pulumi.OutputState }

func (TrailInsightSelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailInsightSelector)(nil)).Elem()
}

func (o TrailInsightSelectorArrayOutput) ToTrailInsightSelectorArrayOutput() TrailInsightSelectorArrayOutput {
	return o
}

func (o TrailInsightSelectorArrayOutput) ToTrailInsightSelectorArrayOutputWithContext(ctx context.Context) TrailInsightSelectorArrayOutput {
	return o
}

func (o TrailInsightSelectorArrayOutput) Index(i pulumi.IntInput) TrailInsightSelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrailInsightSelector {
		return vs[0].([]TrailInsightSelector)[vs[1].(int)]
	}).(TrailInsightSelectorOutput)
}

// An arbitrary set of tags (key-value pairs) for this trail.
type TrailTag struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// TrailTagInput is an input type that accepts TrailTagArgs and TrailTagOutput values.
// You can construct a concrete instance of `TrailTagInput` via:
//
//          TrailTagArgs{...}
type TrailTagInput interface {
	pulumi.Input

	ToTrailTagOutput() TrailTagOutput
	ToTrailTagOutputWithContext(context.Context) TrailTagOutput
}

// An arbitrary set of tags (key-value pairs) for this trail.
type TrailTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (TrailTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailTag)(nil)).Elem()
}

func (i TrailTagArgs) ToTrailTagOutput() TrailTagOutput {
	return i.ToTrailTagOutputWithContext(context.Background())
}

func (i TrailTagArgs) ToTrailTagOutputWithContext(ctx context.Context) TrailTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailTagOutput)
}

// TrailTagArrayInput is an input type that accepts TrailTagArray and TrailTagArrayOutput values.
// You can construct a concrete instance of `TrailTagArrayInput` via:
//
//          TrailTagArray{ TrailTagArgs{...} }
type TrailTagArrayInput interface {
	pulumi.Input

	ToTrailTagArrayOutput() TrailTagArrayOutput
	ToTrailTagArrayOutputWithContext(context.Context) TrailTagArrayOutput
}

type TrailTagArray []TrailTagInput

func (TrailTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailTag)(nil)).Elem()
}

func (i TrailTagArray) ToTrailTagArrayOutput() TrailTagArrayOutput {
	return i.ToTrailTagArrayOutputWithContext(context.Background())
}

func (i TrailTagArray) ToTrailTagArrayOutputWithContext(ctx context.Context) TrailTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailTagArrayOutput)
}

// An arbitrary set of tags (key-value pairs) for this trail.
type TrailTagOutput struct{ *pulumi.OutputState }

func (TrailTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailTag)(nil)).Elem()
}

func (o TrailTagOutput) ToTrailTagOutput() TrailTagOutput {
	return o
}

func (o TrailTagOutput) ToTrailTagOutputWithContext(ctx context.Context) TrailTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o TrailTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v TrailTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o TrailTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TrailTag) string { return v.Value }).(pulumi.StringOutput)
}

type TrailTagArrayOutput struct{ *pulumi.OutputState }

func (TrailTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrailTag)(nil)).Elem()
}

func (o TrailTagArrayOutput) ToTrailTagArrayOutput() TrailTagArrayOutput {
	return o
}

func (o TrailTagArrayOutput) ToTrailTagArrayOutputWithContext(ctx context.Context) TrailTagArrayOutput {
	return o
}

func (o TrailTagArrayOutput) Index(i pulumi.IntInput) TrailTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrailTag {
		return vs[0].([]TrailTag)[vs[1].(int)]
	}).(TrailTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrailDataResourceInput)(nil)).Elem(), TrailDataResourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailDataResourceArrayInput)(nil)).Elem(), TrailDataResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailEventSelectorInput)(nil)).Elem(), TrailEventSelectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailEventSelectorArrayInput)(nil)).Elem(), TrailEventSelectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailInsightSelectorInput)(nil)).Elem(), TrailInsightSelectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailInsightSelectorArrayInput)(nil)).Elem(), TrailInsightSelectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailTagInput)(nil)).Elem(), TrailTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailTagArrayInput)(nil)).Elem(), TrailTagArray{})
	pulumi.RegisterOutputType(TrailDataResourceOutput{})
	pulumi.RegisterOutputType(TrailDataResourceArrayOutput{})
	pulumi.RegisterOutputType(TrailEventSelectorOutput{})
	pulumi.RegisterOutputType(TrailEventSelectorArrayOutput{})
	pulumi.RegisterOutputType(TrailInsightSelectorOutput{})
	pulumi.RegisterOutputType(TrailInsightSelectorArrayOutput{})
	pulumi.RegisterOutputType(TrailTagOutput{})
	pulumi.RegisterOutputType(TrailTagArrayOutput{})
}
