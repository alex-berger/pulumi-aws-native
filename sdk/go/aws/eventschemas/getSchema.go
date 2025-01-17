// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventschemas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EventSchemas::Schema
func LookupSchema(ctx *pulumi.Context, args *LookupSchemaArgs, opts ...pulumi.InvokeOption) (*LookupSchemaResult, error) {
	var rv LookupSchemaResult
	err := ctx.Invoke("aws-native:eventschemas:getSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaArgs struct {
	Id string `pulumi:"id"`
}

type LookupSchemaResult struct {
	Content       *string           `pulumi:"content"`
	Description   *string           `pulumi:"description"`
	Id            *string           `pulumi:"id"`
	SchemaArn     *string           `pulumi:"schemaArn"`
	SchemaVersion *string           `pulumi:"schemaVersion"`
	Tags          []SchemaTagsEntry `pulumi:"tags"`
	Type          *string           `pulumi:"type"`
}

func LookupSchemaOutput(ctx *pulumi.Context, args LookupSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSchemaResult, error) {
			args := v.(LookupSchemaArgs)
			r, err := LookupSchema(ctx, &args, opts...)
			var s LookupSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSchemaResultOutput)
}

type LookupSchemaOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaArgs)(nil)).Elem()
}

type LookupSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaResult)(nil)).Elem()
}

func (o LookupSchemaResultOutput) ToLookupSchemaResultOutput() LookupSchemaResultOutput {
	return o
}

func (o LookupSchemaResultOutput) ToLookupSchemaResultOutputWithContext(ctx context.Context) LookupSchemaResultOutput {
	return o
}

func (o LookupSchemaResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) SchemaArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.SchemaArn }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) SchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.SchemaVersion }).(pulumi.StringPtrOutput)
}

func (o LookupSchemaResultOutput) Tags() SchemaTagsEntryArrayOutput {
	return o.ApplyT(func(v LookupSchemaResult) []SchemaTagsEntry { return v.Tags }).(SchemaTagsEntryArrayOutput)
}

func (o LookupSchemaResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchemaResultOutput{})
}
