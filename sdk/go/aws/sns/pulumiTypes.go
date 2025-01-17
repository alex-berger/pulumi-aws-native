// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TopicSubscription struct {
	Endpoint string `pulumi:"endpoint"`
	Protocol string `pulumi:"protocol"`
}

// TopicSubscriptionInput is an input type that accepts TopicSubscriptionArgs and TopicSubscriptionOutput values.
// You can construct a concrete instance of `TopicSubscriptionInput` via:
//
//          TopicSubscriptionArgs{...}
type TopicSubscriptionInput interface {
	pulumi.Input

	ToTopicSubscriptionOutput() TopicSubscriptionOutput
	ToTopicSubscriptionOutputWithContext(context.Context) TopicSubscriptionOutput
}

type TopicSubscriptionArgs struct {
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	Protocol pulumi.StringInput `pulumi:"protocol"`
}

func (TopicSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicSubscription)(nil)).Elem()
}

func (i TopicSubscriptionArgs) ToTopicSubscriptionOutput() TopicSubscriptionOutput {
	return i.ToTopicSubscriptionOutputWithContext(context.Background())
}

func (i TopicSubscriptionArgs) ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionOutput)
}

// TopicSubscriptionArrayInput is an input type that accepts TopicSubscriptionArray and TopicSubscriptionArrayOutput values.
// You can construct a concrete instance of `TopicSubscriptionArrayInput` via:
//
//          TopicSubscriptionArray{ TopicSubscriptionArgs{...} }
type TopicSubscriptionArrayInput interface {
	pulumi.Input

	ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput
	ToTopicSubscriptionArrayOutputWithContext(context.Context) TopicSubscriptionArrayOutput
}

type TopicSubscriptionArray []TopicSubscriptionInput

func (TopicSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicSubscription)(nil)).Elem()
}

func (i TopicSubscriptionArray) ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput {
	return i.ToTopicSubscriptionArrayOutputWithContext(context.Background())
}

func (i TopicSubscriptionArray) ToTopicSubscriptionArrayOutputWithContext(ctx context.Context) TopicSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionArrayOutput)
}

type TopicSubscriptionOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicSubscription)(nil)).Elem()
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionOutput() TopicSubscriptionOutput {
	return o
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput {
	return o
}

func (o TopicSubscriptionOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v TopicSubscription) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o TopicSubscriptionOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v TopicSubscription) string { return v.Protocol }).(pulumi.StringOutput)
}

type TopicSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicSubscription)(nil)).Elem()
}

func (o TopicSubscriptionArrayOutput) ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput {
	return o
}

func (o TopicSubscriptionArrayOutput) ToTopicSubscriptionArrayOutputWithContext(ctx context.Context) TopicSubscriptionArrayOutput {
	return o
}

func (o TopicSubscriptionArrayOutput) Index(i pulumi.IntInput) TopicSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicSubscription {
		return vs[0].([]TopicSubscription)[vs[1].(int)]
	}).(TopicSubscriptionOutput)
}

type TopicTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// TopicTagInput is an input type that accepts TopicTagArgs and TopicTagOutput values.
// You can construct a concrete instance of `TopicTagInput` via:
//
//          TopicTagArgs{...}
type TopicTagInput interface {
	pulumi.Input

	ToTopicTagOutput() TopicTagOutput
	ToTopicTagOutputWithContext(context.Context) TopicTagOutput
}

type TopicTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TopicTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicTag)(nil)).Elem()
}

func (i TopicTagArgs) ToTopicTagOutput() TopicTagOutput {
	return i.ToTopicTagOutputWithContext(context.Background())
}

func (i TopicTagArgs) ToTopicTagOutputWithContext(ctx context.Context) TopicTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicTagOutput)
}

// TopicTagArrayInput is an input type that accepts TopicTagArray and TopicTagArrayOutput values.
// You can construct a concrete instance of `TopicTagArrayInput` via:
//
//          TopicTagArray{ TopicTagArgs{...} }
type TopicTagArrayInput interface {
	pulumi.Input

	ToTopicTagArrayOutput() TopicTagArrayOutput
	ToTopicTagArrayOutputWithContext(context.Context) TopicTagArrayOutput
}

type TopicTagArray []TopicTagInput

func (TopicTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicTag)(nil)).Elem()
}

func (i TopicTagArray) ToTopicTagArrayOutput() TopicTagArrayOutput {
	return i.ToTopicTagArrayOutputWithContext(context.Background())
}

func (i TopicTagArray) ToTopicTagArrayOutputWithContext(ctx context.Context) TopicTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicTagArrayOutput)
}

type TopicTagOutput struct{ *pulumi.OutputState }

func (TopicTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicTag)(nil)).Elem()
}

func (o TopicTagOutput) ToTopicTagOutput() TopicTagOutput {
	return o
}

func (o TopicTagOutput) ToTopicTagOutputWithContext(ctx context.Context) TopicTagOutput {
	return o
}

func (o TopicTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v TopicTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o TopicTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TopicTag) string { return v.Value }).(pulumi.StringOutput)
}

type TopicTagArrayOutput struct{ *pulumi.OutputState }

func (TopicTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicTag)(nil)).Elem()
}

func (o TopicTagArrayOutput) ToTopicTagArrayOutput() TopicTagArrayOutput {
	return o
}

func (o TopicTagArrayOutput) ToTopicTagArrayOutputWithContext(ctx context.Context) TopicTagArrayOutput {
	return o
}

func (o TopicTagArrayOutput) Index(i pulumi.IntInput) TopicTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicTag {
		return vs[0].([]TopicTag)[vs[1].(int)]
	}).(TopicTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicSubscriptionInput)(nil)).Elem(), TopicSubscriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicSubscriptionArrayInput)(nil)).Elem(), TopicSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicTagInput)(nil)).Elem(), TopicTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicTagArrayInput)(nil)).Elem(), TopicTagArray{})
	pulumi.RegisterOutputType(TopicSubscriptionOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(TopicTagOutput{})
	pulumi.RegisterOutputType(TopicTagArrayOutput{})
}
