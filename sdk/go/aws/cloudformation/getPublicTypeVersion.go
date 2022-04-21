// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Test and Publish a resource that has been registered in the CloudFormation Registry.
func LookupPublicTypeVersion(ctx *pulumi.Context, args *LookupPublicTypeVersionArgs, opts ...pulumi.InvokeOption) (*LookupPublicTypeVersionResult, error) {
	var rv LookupPublicTypeVersionResult
	err := ctx.Invoke("aws-native:cloudformation:getPublicTypeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicTypeVersionArgs struct {
	// The Amazon Resource Number (ARN) assigned to the public extension upon publication
	PublicTypeArn string `pulumi:"publicTypeArn"`
}

type LookupPublicTypeVersionResult struct {
	// The Amazon Resource Number (ARN) assigned to the public extension upon publication
	PublicTypeArn *string `pulumi:"publicTypeArn"`
	// The publisher id assigned by CloudFormation for publishing in this region.
	PublisherId *string `pulumi:"publisherId"`
	// The Amazon Resource Number (ARN) of the extension with the versionId.
	TypeVersionArn *string `pulumi:"typeVersionArn"`
}

func LookupPublicTypeVersionOutput(ctx *pulumi.Context, args LookupPublicTypeVersionOutputArgs, opts ...pulumi.InvokeOption) LookupPublicTypeVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPublicTypeVersionResult, error) {
			args := v.(LookupPublicTypeVersionArgs)
			r, err := LookupPublicTypeVersion(ctx, &args, opts...)
			var s LookupPublicTypeVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPublicTypeVersionResultOutput)
}

type LookupPublicTypeVersionOutputArgs struct {
	// The Amazon Resource Number (ARN) assigned to the public extension upon publication
	PublicTypeArn pulumi.StringInput `pulumi:"publicTypeArn"`
}

func (LookupPublicTypeVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicTypeVersionArgs)(nil)).Elem()
}

type LookupPublicTypeVersionResultOutput struct{ *pulumi.OutputState }

func (LookupPublicTypeVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPublicTypeVersionResult)(nil)).Elem()
}

func (o LookupPublicTypeVersionResultOutput) ToLookupPublicTypeVersionResultOutput() LookupPublicTypeVersionResultOutput {
	return o
}

func (o LookupPublicTypeVersionResultOutput) ToLookupPublicTypeVersionResultOutputWithContext(ctx context.Context) LookupPublicTypeVersionResultOutput {
	return o
}

// The Amazon Resource Number (ARN) assigned to the public extension upon publication
func (o LookupPublicTypeVersionResultOutput) PublicTypeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicTypeVersionResult) *string { return v.PublicTypeArn }).(pulumi.StringPtrOutput)
}

// The publisher id assigned by CloudFormation for publishing in this region.
func (o LookupPublicTypeVersionResultOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicTypeVersionResult) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Number (ARN) of the extension with the versionId.
func (o LookupPublicTypeVersionResultOutput) TypeVersionArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPublicTypeVersionResult) *string { return v.TypeVersionArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPublicTypeVersionResultOutput{})
}
