// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type QueueTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// QueueTagInput is an input type that accepts QueueTagArgs and QueueTagOutput values.
// You can construct a concrete instance of `QueueTagInput` via:
//
//          QueueTagArgs{...}
type QueueTagInput interface {
	pulumi.Input

	ToQueueTagOutput() QueueTagOutput
	ToQueueTagOutputWithContext(context.Context) QueueTagOutput
}

type QueueTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (QueueTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueTag)(nil)).Elem()
}

func (i QueueTagArgs) ToQueueTagOutput() QueueTagOutput {
	return i.ToQueueTagOutputWithContext(context.Background())
}

func (i QueueTagArgs) ToQueueTagOutputWithContext(ctx context.Context) QueueTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueTagOutput)
}

// QueueTagArrayInput is an input type that accepts QueueTagArray and QueueTagArrayOutput values.
// You can construct a concrete instance of `QueueTagArrayInput` via:
//
//          QueueTagArray{ QueueTagArgs{...} }
type QueueTagArrayInput interface {
	pulumi.Input

	ToQueueTagArrayOutput() QueueTagArrayOutput
	ToQueueTagArrayOutputWithContext(context.Context) QueueTagArrayOutput
}

type QueueTagArray []QueueTagInput

func (QueueTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueueTag)(nil)).Elem()
}

func (i QueueTagArray) ToQueueTagArrayOutput() QueueTagArrayOutput {
	return i.ToQueueTagArrayOutputWithContext(context.Background())
}

func (i QueueTagArray) ToQueueTagArrayOutputWithContext(ctx context.Context) QueueTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueTagArrayOutput)
}

type QueueTagOutput struct{ *pulumi.OutputState }

func (QueueTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueTag)(nil)).Elem()
}

func (o QueueTagOutput) ToQueueTagOutput() QueueTagOutput {
	return o
}

func (o QueueTagOutput) ToQueueTagOutputWithContext(ctx context.Context) QueueTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o QueueTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v QueueTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o QueueTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v QueueTag) string { return v.Value }).(pulumi.StringOutput)
}

type QueueTagArrayOutput struct{ *pulumi.OutputState }

func (QueueTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueueTag)(nil)).Elem()
}

func (o QueueTagArrayOutput) ToQueueTagArrayOutput() QueueTagArrayOutput {
	return o
}

func (o QueueTagArrayOutput) ToQueueTagArrayOutputWithContext(ctx context.Context) QueueTagArrayOutput {
	return o
}

func (o QueueTagArrayOutput) Index(i pulumi.IntInput) QueueTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueueTag {
		return vs[0].([]QueueTag)[vs[1].(int)]
	}).(QueueTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueTagInput)(nil)).Elem(), QueueTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*QueueTagArrayInput)(nil)).Elem(), QueueTagArray{})
	pulumi.RegisterOutputType(QueueTagOutput{})
	pulumi.RegisterOutputType(QueueTagArrayOutput{})
}
