// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaPackage::Asset
func LookupAsset(ctx *pulumi.Context, args *LookupAssetArgs, opts ...pulumi.InvokeOption) (*LookupAssetResult, error) {
	var rv LookupAssetResult
	err := ctx.Invoke("aws-native:mediapackage:getAsset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssetArgs struct {
	// The unique identifier for the Asset.
	Id string `pulumi:"id"`
}

type LookupAssetResult struct {
	// The ARN of the Asset.
	Arn *string `pulumi:"arn"`
	// The time the Asset was initially submitted for Ingest.
	CreatedAt *string `pulumi:"createdAt"`
	// The list of egress endpoints available for the Asset.
	EgressEndpoints []AssetEgressEndpoint `pulumi:"egressEndpoints"`
	// The unique identifier for the Asset.
	Id *string `pulumi:"id"`
	// The ID of the PackagingGroup for the Asset.
	PackagingGroupId *string `pulumi:"packagingGroupId"`
	// The resource ID to include in SPEKE key requests.
	ResourceId *string `pulumi:"resourceId"`
	// ARN of the source object in S3.
	SourceArn *string `pulumi:"sourceArn"`
	// The IAM role_arn used to access the source S3 bucket.
	SourceRoleArn *string `pulumi:"sourceRoleArn"`
	// A collection of tags associated with a resource
	Tags []AssetTag `pulumi:"tags"`
}

func LookupAssetOutput(ctx *pulumi.Context, args LookupAssetOutputArgs, opts ...pulumi.InvokeOption) LookupAssetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssetResult, error) {
			args := v.(LookupAssetArgs)
			r, err := LookupAsset(ctx, &args, opts...)
			var s LookupAssetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssetResultOutput)
}

type LookupAssetOutputArgs struct {
	// The unique identifier for the Asset.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAssetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetArgs)(nil)).Elem()
}

type LookupAssetResultOutput struct{ *pulumi.OutputState }

func (LookupAssetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssetResult)(nil)).Elem()
}

func (o LookupAssetResultOutput) ToLookupAssetResultOutput() LookupAssetResultOutput {
	return o
}

func (o LookupAssetResultOutput) ToLookupAssetResultOutputWithContext(ctx context.Context) LookupAssetResultOutput {
	return o
}

// The ARN of the Asset.
func (o LookupAssetResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The time the Asset was initially submitted for Ingest.
func (o LookupAssetResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The list of egress endpoints available for the Asset.
func (o LookupAssetResultOutput) EgressEndpoints() AssetEgressEndpointArrayOutput {
	return o.ApplyT(func(v LookupAssetResult) []AssetEgressEndpoint { return v.EgressEndpoints }).(AssetEgressEndpointArrayOutput)
}

// The unique identifier for the Asset.
func (o LookupAssetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ID of the PackagingGroup for the Asset.
func (o LookupAssetResultOutput) PackagingGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.PackagingGroupId }).(pulumi.StringPtrOutput)
}

// The resource ID to include in SPEKE key requests.
func (o LookupAssetResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// ARN of the source object in S3.
func (o LookupAssetResultOutput) SourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.SourceArn }).(pulumi.StringPtrOutput)
}

// The IAM role_arn used to access the source S3 bucket.
func (o LookupAssetResultOutput) SourceRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssetResult) *string { return v.SourceRoleArn }).(pulumi.StringPtrOutput)
}

// A collection of tags associated with a resource
func (o LookupAssetResultOutput) Tags() AssetTagArrayOutput {
	return o.ApplyT(func(v LookupAssetResult) []AssetTag { return v.Tags }).(AssetTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssetResultOutput{})
}
