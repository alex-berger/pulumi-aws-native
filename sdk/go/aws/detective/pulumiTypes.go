// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package detective

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A key-value pair to associate with a resource.
type GraphTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @
	Key *string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @
	Value *string `pulumi:"value"`
}

// GraphTagInput is an input type that accepts GraphTagArgs and GraphTagOutput values.
// You can construct a concrete instance of `GraphTagInput` via:
//
//          GraphTagArgs{...}
type GraphTagInput interface {
	pulumi.Input

	ToGraphTagOutput() GraphTagOutput
	ToGraphTagOutputWithContext(context.Context) GraphTagOutput
}

// A key-value pair to associate with a resource.
type GraphTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @
	Key pulumi.StringPtrInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (GraphTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphTag)(nil)).Elem()
}

func (i GraphTagArgs) ToGraphTagOutput() GraphTagOutput {
	return i.ToGraphTagOutputWithContext(context.Background())
}

func (i GraphTagArgs) ToGraphTagOutputWithContext(ctx context.Context) GraphTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphTagOutput)
}

// GraphTagArrayInput is an input type that accepts GraphTagArray and GraphTagArrayOutput values.
// You can construct a concrete instance of `GraphTagArrayInput` via:
//
//          GraphTagArray{ GraphTagArgs{...} }
type GraphTagArrayInput interface {
	pulumi.Input

	ToGraphTagArrayOutput() GraphTagArrayOutput
	ToGraphTagArrayOutputWithContext(context.Context) GraphTagArrayOutput
}

type GraphTagArray []GraphTagInput

func (GraphTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphTag)(nil)).Elem()
}

func (i GraphTagArray) ToGraphTagArrayOutput() GraphTagArrayOutput {
	return i.ToGraphTagArrayOutputWithContext(context.Background())
}

func (i GraphTagArray) ToGraphTagArrayOutputWithContext(ctx context.Context) GraphTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphTagArrayOutput)
}

// A key-value pair to associate with a resource.
type GraphTagOutput struct{ *pulumi.OutputState }

func (GraphTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphTag)(nil)).Elem()
}

func (o GraphTagOutput) ToGraphTagOutput() GraphTagOutput {
	return o
}

func (o GraphTagOutput) ToGraphTagOutputWithContext(ctx context.Context) GraphTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @
func (o GraphTagOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphTag) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. Valid characters are Unicode letters, digits, white space, and any of the following symbols: _ . : / = + - @
func (o GraphTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GraphTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type GraphTagArrayOutput struct{ *pulumi.OutputState }

func (GraphTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GraphTag)(nil)).Elem()
}

func (o GraphTagArrayOutput) ToGraphTagArrayOutput() GraphTagArrayOutput {
	return o
}

func (o GraphTagArrayOutput) ToGraphTagArrayOutputWithContext(ctx context.Context) GraphTagArrayOutput {
	return o
}

func (o GraphTagArrayOutput) Index(i pulumi.IntInput) GraphTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GraphTag {
		return vs[0].([]GraphTag)[vs[1].(int)]
	}).(GraphTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GraphTagInput)(nil)).Elem(), GraphTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GraphTagArrayInput)(nil)).Elem(), GraphTagArray{})
	pulumi.RegisterOutputType(GraphTagOutput{})
	pulumi.RegisterOutputType(GraphTagArrayOutput{})
}
