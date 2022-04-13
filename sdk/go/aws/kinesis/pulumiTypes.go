// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream. Removing this property from your stack template and updating your stack disables encryption.
type StreamEncryption struct {
	// The encryption type to use. The only valid value is KMS.
	EncryptionType StreamEncryptionEncryptionType `pulumi:"encryptionType"`
	// The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.
	KeyId string `pulumi:"keyId"`
}

// StreamEncryptionInput is an input type that accepts StreamEncryptionArgs and StreamEncryptionOutput values.
// You can construct a concrete instance of `StreamEncryptionInput` via:
//
//          StreamEncryptionArgs{...}
type StreamEncryptionInput interface {
	pulumi.Input

	ToStreamEncryptionOutput() StreamEncryptionOutput
	ToStreamEncryptionOutputWithContext(context.Context) StreamEncryptionOutput
}

// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream. Removing this property from your stack template and updating your stack disables encryption.
type StreamEncryptionArgs struct {
	// The encryption type to use. The only valid value is KMS.
	EncryptionType StreamEncryptionEncryptionTypeInput `pulumi:"encryptionType"`
	// The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.
	KeyId pulumi.StringInput `pulumi:"keyId"`
}

func (StreamEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamEncryption)(nil)).Elem()
}

func (i StreamEncryptionArgs) ToStreamEncryptionOutput() StreamEncryptionOutput {
	return i.ToStreamEncryptionOutputWithContext(context.Background())
}

func (i StreamEncryptionArgs) ToStreamEncryptionOutputWithContext(ctx context.Context) StreamEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamEncryptionOutput)
}

func (i StreamEncryptionArgs) ToStreamEncryptionPtrOutput() StreamEncryptionPtrOutput {
	return i.ToStreamEncryptionPtrOutputWithContext(context.Background())
}

func (i StreamEncryptionArgs) ToStreamEncryptionPtrOutputWithContext(ctx context.Context) StreamEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamEncryptionOutput).ToStreamEncryptionPtrOutputWithContext(ctx)
}

// StreamEncryptionPtrInput is an input type that accepts StreamEncryptionArgs, StreamEncryptionPtr and StreamEncryptionPtrOutput values.
// You can construct a concrete instance of `StreamEncryptionPtrInput` via:
//
//          StreamEncryptionArgs{...}
//
//  or:
//
//          nil
type StreamEncryptionPtrInput interface {
	pulumi.Input

	ToStreamEncryptionPtrOutput() StreamEncryptionPtrOutput
	ToStreamEncryptionPtrOutputWithContext(context.Context) StreamEncryptionPtrOutput
}

type streamEncryptionPtrType StreamEncryptionArgs

func StreamEncryptionPtr(v *StreamEncryptionArgs) StreamEncryptionPtrInput {
	return (*streamEncryptionPtrType)(v)
}

func (*streamEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamEncryption)(nil)).Elem()
}

func (i *streamEncryptionPtrType) ToStreamEncryptionPtrOutput() StreamEncryptionPtrOutput {
	return i.ToStreamEncryptionPtrOutputWithContext(context.Background())
}

func (i *streamEncryptionPtrType) ToStreamEncryptionPtrOutputWithContext(ctx context.Context) StreamEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamEncryptionPtrOutput)
}

// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream. Removing this property from your stack template and updating your stack disables encryption.
type StreamEncryptionOutput struct{ *pulumi.OutputState }

func (StreamEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamEncryption)(nil)).Elem()
}

func (o StreamEncryptionOutput) ToStreamEncryptionOutput() StreamEncryptionOutput {
	return o
}

func (o StreamEncryptionOutput) ToStreamEncryptionOutputWithContext(ctx context.Context) StreamEncryptionOutput {
	return o
}

func (o StreamEncryptionOutput) ToStreamEncryptionPtrOutput() StreamEncryptionPtrOutput {
	return o.ToStreamEncryptionPtrOutputWithContext(context.Background())
}

func (o StreamEncryptionOutput) ToStreamEncryptionPtrOutputWithContext(ctx context.Context) StreamEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamEncryption) *StreamEncryption {
		return &v
	}).(StreamEncryptionPtrOutput)
}

// The encryption type to use. The only valid value is KMS.
func (o StreamEncryptionOutput) EncryptionType() StreamEncryptionEncryptionTypeOutput {
	return o.ApplyT(func(v StreamEncryption) StreamEncryptionEncryptionType { return v.EncryptionType }).(StreamEncryptionEncryptionTypeOutput)
}

// The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.
func (o StreamEncryptionOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v StreamEncryption) string { return v.KeyId }).(pulumi.StringOutput)
}

type StreamEncryptionPtrOutput struct{ *pulumi.OutputState }

func (StreamEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamEncryption)(nil)).Elem()
}

func (o StreamEncryptionPtrOutput) ToStreamEncryptionPtrOutput() StreamEncryptionPtrOutput {
	return o
}

func (o StreamEncryptionPtrOutput) ToStreamEncryptionPtrOutputWithContext(ctx context.Context) StreamEncryptionPtrOutput {
	return o
}

func (o StreamEncryptionPtrOutput) Elem() StreamEncryptionOutput {
	return o.ApplyT(func(v *StreamEncryption) StreamEncryption {
		if v != nil {
			return *v
		}
		var ret StreamEncryption
		return ret
	}).(StreamEncryptionOutput)
}

// The encryption type to use. The only valid value is KMS.
func (o StreamEncryptionPtrOutput) EncryptionType() StreamEncryptionEncryptionTypePtrOutput {
	return o.ApplyT(func(v *StreamEncryption) *StreamEncryptionEncryptionType {
		if v == nil {
			return nil
		}
		return &v.EncryptionType
	}).(StreamEncryptionEncryptionTypePtrOutput)
}

// The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.
func (o StreamEncryptionPtrOutput) KeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamEncryption) *string {
		if v == nil {
			return nil
		}
		return &v.KeyId
	}).(pulumi.StringPtrOutput)
}

// When specified, enables or updates the mode of stream. Default is PROVISIONED.
type StreamModeDetails struct {
	// The mode of the stream
	StreamMode StreamModeDetailsStreamMode `pulumi:"streamMode"`
}

// StreamModeDetailsInput is an input type that accepts StreamModeDetailsArgs and StreamModeDetailsOutput values.
// You can construct a concrete instance of `StreamModeDetailsInput` via:
//
//          StreamModeDetailsArgs{...}
type StreamModeDetailsInput interface {
	pulumi.Input

	ToStreamModeDetailsOutput() StreamModeDetailsOutput
	ToStreamModeDetailsOutputWithContext(context.Context) StreamModeDetailsOutput
}

// When specified, enables or updates the mode of stream. Default is PROVISIONED.
type StreamModeDetailsArgs struct {
	// The mode of the stream
	StreamMode StreamModeDetailsStreamModeInput `pulumi:"streamMode"`
}

func (StreamModeDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamModeDetails)(nil)).Elem()
}

func (i StreamModeDetailsArgs) ToStreamModeDetailsOutput() StreamModeDetailsOutput {
	return i.ToStreamModeDetailsOutputWithContext(context.Background())
}

func (i StreamModeDetailsArgs) ToStreamModeDetailsOutputWithContext(ctx context.Context) StreamModeDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamModeDetailsOutput)
}

func (i StreamModeDetailsArgs) ToStreamModeDetailsPtrOutput() StreamModeDetailsPtrOutput {
	return i.ToStreamModeDetailsPtrOutputWithContext(context.Background())
}

func (i StreamModeDetailsArgs) ToStreamModeDetailsPtrOutputWithContext(ctx context.Context) StreamModeDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamModeDetailsOutput).ToStreamModeDetailsPtrOutputWithContext(ctx)
}

// StreamModeDetailsPtrInput is an input type that accepts StreamModeDetailsArgs, StreamModeDetailsPtr and StreamModeDetailsPtrOutput values.
// You can construct a concrete instance of `StreamModeDetailsPtrInput` via:
//
//          StreamModeDetailsArgs{...}
//
//  or:
//
//          nil
type StreamModeDetailsPtrInput interface {
	pulumi.Input

	ToStreamModeDetailsPtrOutput() StreamModeDetailsPtrOutput
	ToStreamModeDetailsPtrOutputWithContext(context.Context) StreamModeDetailsPtrOutput
}

type streamModeDetailsPtrType StreamModeDetailsArgs

func StreamModeDetailsPtr(v *StreamModeDetailsArgs) StreamModeDetailsPtrInput {
	return (*streamModeDetailsPtrType)(v)
}

func (*streamModeDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamModeDetails)(nil)).Elem()
}

func (i *streamModeDetailsPtrType) ToStreamModeDetailsPtrOutput() StreamModeDetailsPtrOutput {
	return i.ToStreamModeDetailsPtrOutputWithContext(context.Background())
}

func (i *streamModeDetailsPtrType) ToStreamModeDetailsPtrOutputWithContext(ctx context.Context) StreamModeDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamModeDetailsPtrOutput)
}

// When specified, enables or updates the mode of stream. Default is PROVISIONED.
type StreamModeDetailsOutput struct{ *pulumi.OutputState }

func (StreamModeDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamModeDetails)(nil)).Elem()
}

func (o StreamModeDetailsOutput) ToStreamModeDetailsOutput() StreamModeDetailsOutput {
	return o
}

func (o StreamModeDetailsOutput) ToStreamModeDetailsOutputWithContext(ctx context.Context) StreamModeDetailsOutput {
	return o
}

func (o StreamModeDetailsOutput) ToStreamModeDetailsPtrOutput() StreamModeDetailsPtrOutput {
	return o.ToStreamModeDetailsPtrOutputWithContext(context.Background())
}

func (o StreamModeDetailsOutput) ToStreamModeDetailsPtrOutputWithContext(ctx context.Context) StreamModeDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamModeDetails) *StreamModeDetails {
		return &v
	}).(StreamModeDetailsPtrOutput)
}

// The mode of the stream
func (o StreamModeDetailsOutput) StreamMode() StreamModeDetailsStreamModeOutput {
	return o.ApplyT(func(v StreamModeDetails) StreamModeDetailsStreamMode { return v.StreamMode }).(StreamModeDetailsStreamModeOutput)
}

type StreamModeDetailsPtrOutput struct{ *pulumi.OutputState }

func (StreamModeDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamModeDetails)(nil)).Elem()
}

func (o StreamModeDetailsPtrOutput) ToStreamModeDetailsPtrOutput() StreamModeDetailsPtrOutput {
	return o
}

func (o StreamModeDetailsPtrOutput) ToStreamModeDetailsPtrOutputWithContext(ctx context.Context) StreamModeDetailsPtrOutput {
	return o
}

func (o StreamModeDetailsPtrOutput) Elem() StreamModeDetailsOutput {
	return o.ApplyT(func(v *StreamModeDetails) StreamModeDetails {
		if v != nil {
			return *v
		}
		var ret StreamModeDetails
		return ret
	}).(StreamModeDetailsOutput)
}

// The mode of the stream
func (o StreamModeDetailsPtrOutput) StreamMode() StreamModeDetailsStreamModePtrOutput {
	return o.ApplyT(func(v *StreamModeDetails) *StreamModeDetailsStreamMode {
		if v == nil {
			return nil
		}
		return &v.StreamMode
	}).(StreamModeDetailsStreamModePtrOutput)
}

// An arbitrary set of tags (key-value pairs) to associate with the Kinesis stream.
type StreamTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// StreamTagInput is an input type that accepts StreamTagArgs and StreamTagOutput values.
// You can construct a concrete instance of `StreamTagInput` via:
//
//          StreamTagArgs{...}
type StreamTagInput interface {
	pulumi.Input

	ToStreamTagOutput() StreamTagOutput
	ToStreamTagOutputWithContext(context.Context) StreamTagOutput
}

// An arbitrary set of tags (key-value pairs) to associate with the Kinesis stream.
type StreamTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value pulumi.StringInput `pulumi:"value"`
}

func (StreamTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamTag)(nil)).Elem()
}

func (i StreamTagArgs) ToStreamTagOutput() StreamTagOutput {
	return i.ToStreamTagOutputWithContext(context.Background())
}

func (i StreamTagArgs) ToStreamTagOutputWithContext(ctx context.Context) StreamTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamTagOutput)
}

// StreamTagArrayInput is an input type that accepts StreamTagArray and StreamTagArrayOutput values.
// You can construct a concrete instance of `StreamTagArrayInput` via:
//
//          StreamTagArray{ StreamTagArgs{...} }
type StreamTagArrayInput interface {
	pulumi.Input

	ToStreamTagArrayOutput() StreamTagArrayOutput
	ToStreamTagArrayOutputWithContext(context.Context) StreamTagArrayOutput
}

type StreamTagArray []StreamTagInput

func (StreamTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamTag)(nil)).Elem()
}

func (i StreamTagArray) ToStreamTagArrayOutput() StreamTagArrayOutput {
	return i.ToStreamTagArrayOutputWithContext(context.Background())
}

func (i StreamTagArray) ToStreamTagArrayOutputWithContext(ctx context.Context) StreamTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamTagArrayOutput)
}

// An arbitrary set of tags (key-value pairs) to associate with the Kinesis stream.
type StreamTagOutput struct{ *pulumi.OutputState }

func (StreamTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamTag)(nil)).Elem()
}

func (o StreamTagOutput) ToStreamTagOutput() StreamTagOutput {
	return o
}

func (o StreamTagOutput) ToStreamTagOutputWithContext(ctx context.Context) StreamTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o StreamTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StreamTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
func (o StreamTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v StreamTag) string { return v.Value }).(pulumi.StringOutput)
}

type StreamTagArrayOutput struct{ *pulumi.OutputState }

func (StreamTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamTag)(nil)).Elem()
}

func (o StreamTagArrayOutput) ToStreamTagArrayOutput() StreamTagArrayOutput {
	return o
}

func (o StreamTagArrayOutput) ToStreamTagArrayOutputWithContext(ctx context.Context) StreamTagArrayOutput {
	return o
}

func (o StreamTagArrayOutput) Index(i pulumi.IntInput) StreamTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamTag {
		return vs[0].([]StreamTag)[vs[1].(int)]
	}).(StreamTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamEncryptionInput)(nil)).Elem(), StreamEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamEncryptionPtrInput)(nil)).Elem(), StreamEncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamModeDetailsInput)(nil)).Elem(), StreamModeDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamModeDetailsPtrInput)(nil)).Elem(), StreamModeDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamTagInput)(nil)).Elem(), StreamTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamTagArrayInput)(nil)).Elem(), StreamTagArray{})
	pulumi.RegisterOutputType(StreamEncryptionOutput{})
	pulumi.RegisterOutputType(StreamEncryptionPtrOutput{})
	pulumi.RegisterOutputType(StreamModeDetailsOutput{})
	pulumi.RegisterOutputType(StreamModeDetailsPtrOutput{})
	pulumi.RegisterOutputType(StreamTagOutput{})
	pulumi.RegisterOutputType(StreamTagArrayOutput{})
}
