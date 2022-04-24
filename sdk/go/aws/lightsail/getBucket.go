// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lightsail::Bucket
func LookupBucket(ctx *pulumi.Context, args *LookupBucketArgs, opts ...pulumi.InvokeOption) (*LookupBucketResult, error) {
	var rv LookupBucketResult
	err := ctx.Invoke("aws-native:lightsail:getBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBucketArgs struct {
	// The name for the bucket.
	BucketName string `pulumi:"bucketName"`
}

type LookupBucketResult struct {
	// Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.
	AbleToUpdateBundle *bool              `pulumi:"ableToUpdateBundle"`
	AccessRules        *BucketAccessRules `pulumi:"accessRules"`
	BucketArn          *string            `pulumi:"bucketArn"`
	// The ID of the bundle to use for the bucket.
	BundleId *string `pulumi:"bundleId"`
	// Specifies whether to enable or disable versioning of objects in the bucket.
	ObjectVersioning *bool `pulumi:"objectVersioning"`
	// An array of strings to specify the AWS account IDs that can access the bucket.
	ReadOnlyAccessAccounts []string `pulumi:"readOnlyAccessAccounts"`
	// The names of the Lightsail resources for which to set bucket access.
	ResourcesReceivingAccess []string `pulumi:"resourcesReceivingAccess"`
	// An array of key-value pairs to apply to this resource.
	Tags []BucketTag `pulumi:"tags"`
	// The URL of the bucket.
	Url *string `pulumi:"url"`
}

func LookupBucketOutput(ctx *pulumi.Context, args LookupBucketOutputArgs, opts ...pulumi.InvokeOption) LookupBucketResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketResult, error) {
			args := v.(LookupBucketArgs)
			r, err := LookupBucket(ctx, &args, opts...)
			var s LookupBucketResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBucketResultOutput)
}

type LookupBucketOutputArgs struct {
	// The name for the bucket.
	BucketName pulumi.StringInput `pulumi:"bucketName"`
}

func (LookupBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketArgs)(nil)).Elem()
}

type LookupBucketResultOutput struct{ *pulumi.OutputState }

func (LookupBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketResult)(nil)).Elem()
}

func (o LookupBucketResultOutput) ToLookupBucketResultOutput() LookupBucketResultOutput {
	return o
}

func (o LookupBucketResultOutput) ToLookupBucketResultOutputWithContext(ctx context.Context) LookupBucketResultOutput {
	return o
}

// Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.
func (o LookupBucketResultOutput) AbleToUpdateBundle() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *bool { return v.AbleToUpdateBundle }).(pulumi.BoolPtrOutput)
}

func (o LookupBucketResultOutput) AccessRules() BucketAccessRulesPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketAccessRules { return v.AccessRules }).(BucketAccessRulesPtrOutput)
}

func (o LookupBucketResultOutput) BucketArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.BucketArn }).(pulumi.StringPtrOutput)
}

// The ID of the bundle to use for the bucket.
func (o LookupBucketResultOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.BundleId }).(pulumi.StringPtrOutput)
}

// Specifies whether to enable or disable versioning of objects in the bucket.
func (o LookupBucketResultOutput) ObjectVersioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *bool { return v.ObjectVersioning }).(pulumi.BoolPtrOutput)
}

// An array of strings to specify the AWS account IDs that can access the bucket.
func (o LookupBucketResultOutput) ReadOnlyAccessAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []string { return v.ReadOnlyAccessAccounts }).(pulumi.StringArrayOutput)
}

// The names of the Lightsail resources for which to set bucket access.
func (o LookupBucketResultOutput) ResourcesReceivingAccess() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []string { return v.ResourcesReceivingAccess }).(pulumi.StringArrayOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupBucketResultOutput) Tags() BucketTagArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketTag { return v.Tags }).(BucketTagArrayOutput)
}

// The URL of the bucket.
func (o LookupBucketResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketResultOutput{})
}
