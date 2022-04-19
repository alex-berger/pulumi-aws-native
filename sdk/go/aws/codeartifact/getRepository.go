// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codeartifact

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource schema to create a CodeArtifact repository.
func LookupRepository(ctx *pulumi.Context, args *LookupRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryResult, error) {
	var rv LookupRepositoryResult
	err := ctx.Invoke("aws-native:codeartifact:getRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRepositoryArgs struct {
	// The ARN of the repository.
	Arn string `pulumi:"arn"`
}

type LookupRepositoryResult struct {
	// The ARN of the repository.
	Arn *string `pulumi:"arn"`
	// A text description of the repository.
	Description *string `pulumi:"description"`
	// A list of external connections associated with the repository.
	ExternalConnections []string `pulumi:"externalConnections"`
	// The name of the repository. This is used for GetAtt
	Name *string `pulumi:"name"`
	// The access control resource policy on the provided repository.
	PermissionsPolicyDocument interface{} `pulumi:"permissionsPolicyDocument"`
	// An array of key-value pairs to apply to this resource.
	Tags []RepositoryTag `pulumi:"tags"`
	// A list of upstream repositories associated with the repository.
	Upstreams []string `pulumi:"upstreams"`
}

func LookupRepositoryOutput(ctx *pulumi.Context, args LookupRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRepositoryResult, error) {
			args := v.(LookupRepositoryArgs)
			r, err := LookupRepository(ctx, &args, opts...)
			var s LookupRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRepositoryResultOutput)
}

type LookupRepositoryOutputArgs struct {
	// The ARN of the repository.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryArgs)(nil)).Elem()
}

type LookupRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryResult)(nil)).Elem()
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutput() LookupRepositoryResultOutput {
	return o
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutputWithContext(ctx context.Context) LookupRepositoryResultOutput {
	return o
}

// The ARN of the repository.
func (o LookupRepositoryResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A text description of the repository.
func (o LookupRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A list of external connections associated with the repository.
func (o LookupRepositoryResultOutput) ExternalConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRepositoryResult) []string { return v.ExternalConnections }).(pulumi.StringArrayOutput)
}

// The name of the repository. This is used for GetAtt
func (o LookupRepositoryResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRepositoryResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The access control resource policy on the provided repository.
func (o LookupRepositoryResultOutput) PermissionsPolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupRepositoryResult) interface{} { return v.PermissionsPolicyDocument }).(pulumi.AnyOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupRepositoryResultOutput) Tags() RepositoryTagArrayOutput {
	return o.ApplyT(func(v LookupRepositoryResult) []RepositoryTag { return v.Tags }).(RepositoryTagArrayOutput)
}

// A list of upstream repositories associated with the repository.
func (o LookupRepositoryResultOutput) Upstreams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRepositoryResult) []string { return v.Upstreams }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryResultOutput{})
}
