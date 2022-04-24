// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::S3::Bucket
func LookupBucket(ctx *pulumi.Context, args *LookupBucketArgs, opts ...pulumi.InvokeOption) (*LookupBucketResult, error) {
	var rv LookupBucketResult
	err := ctx.Invoke("aws-native:s3:getBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBucketArgs struct {
	// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	BucketName string `pulumi:"bucketName"`
}

type LookupBucketResult struct {
	// Configuration for the transfer acceleration state.
	AccelerateConfiguration *BucketAccelerateConfiguration `pulumi:"accelerateConfiguration"`
	// A canned access control list (ACL) that grants predefined permissions to the bucket.
	AccessControl *BucketAccessControl `pulumi:"accessControl"`
	// The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	AnalyticsConfigurations []BucketAnalyticsConfiguration `pulumi:"analyticsConfigurations"`
	// The Amazon Resource Name (ARN) of the specified bucket.
	Arn              *string           `pulumi:"arn"`
	BucketEncryption *BucketEncryption `pulumi:"bucketEncryption"`
	// Rules that define cross-origin resource sharing of objects in this bucket.
	CorsConfiguration *BucketCorsConfiguration `pulumi:"corsConfiguration"`
	// The IPv4 DNS name of the specified bucket.
	DomainName *string `pulumi:"domainName"`
	// The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
	DualStackDomainName *string `pulumi:"dualStackDomainName"`
	// Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
	IntelligentTieringConfigurations []BucketIntelligentTieringConfiguration `pulumi:"intelligentTieringConfigurations"`
	// The inventory configuration for an Amazon S3 bucket.
	InventoryConfigurations []BucketInventoryConfiguration `pulumi:"inventoryConfigurations"`
	// Rules that define how Amazon S3 manages objects during their lifetime.
	LifecycleConfiguration *BucketLifecycleConfiguration `pulumi:"lifecycleConfiguration"`
	// Settings that define where logs are stored.
	LoggingConfiguration *BucketLoggingConfiguration `pulumi:"loggingConfiguration"`
	// Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
	MetricsConfigurations []BucketMetricsConfiguration `pulumi:"metricsConfigurations"`
	// Configuration that defines how Amazon S3 handles bucket notifications.
	NotificationConfiguration *BucketNotificationConfiguration `pulumi:"notificationConfiguration"`
	// Places an Object Lock configuration on the specified bucket.
	ObjectLockConfiguration *BucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	// Specifies the container element for object ownership rules.
	OwnershipControls              *BucketOwnershipControls              `pulumi:"ownershipControls"`
	PublicAccessBlockConfiguration *BucketPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// Returns the regional domain name of the specified bucket.
	RegionalDomainName *string `pulumi:"regionalDomainName"`
	// Configuration for replicating objects in an S3 bucket.
	ReplicationConfiguration *BucketReplicationConfiguration `pulumi:"replicationConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	Tags                    []BucketTag                    `pulumi:"tags"`
	VersioningConfiguration *BucketVersioningConfiguration `pulumi:"versioningConfiguration"`
	WebsiteConfiguration    *BucketWebsiteConfiguration    `pulumi:"websiteConfiguration"`
	// The Amazon S3 website endpoint for the specified bucket.
	WebsiteURL *string `pulumi:"websiteURL"`
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
	// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
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

// Configuration for the transfer acceleration state.
func (o LookupBucketResultOutput) AccelerateConfiguration() BucketAccelerateConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketAccelerateConfiguration { return v.AccelerateConfiguration }).(BucketAccelerateConfigurationPtrOutput)
}

// A canned access control list (ACL) that grants predefined permissions to the bucket.
func (o LookupBucketResultOutput) AccessControl() BucketAccessControlPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketAccessControl { return v.AccessControl }).(BucketAccessControlPtrOutput)
}

// The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
func (o LookupBucketResultOutput) AnalyticsConfigurations() BucketAnalyticsConfigurationArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketAnalyticsConfiguration { return v.AnalyticsConfigurations }).(BucketAnalyticsConfigurationArrayOutput)
}

// The Amazon Resource Name (ARN) of the specified bucket.
func (o LookupBucketResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupBucketResultOutput) BucketEncryption() BucketEncryptionPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketEncryption { return v.BucketEncryption }).(BucketEncryptionPtrOutput)
}

// Rules that define cross-origin resource sharing of objects in this bucket.
func (o LookupBucketResultOutput) CorsConfiguration() BucketCorsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketCorsConfiguration { return v.CorsConfiguration }).(BucketCorsConfigurationPtrOutput)
}

// The IPv4 DNS name of the specified bucket.
func (o LookupBucketResultOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.DomainName }).(pulumi.StringPtrOutput)
}

// The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
func (o LookupBucketResultOutput) DualStackDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.DualStackDomainName }).(pulumi.StringPtrOutput)
}

// Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
func (o LookupBucketResultOutput) IntelligentTieringConfigurations() BucketIntelligentTieringConfigurationArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketIntelligentTieringConfiguration {
		return v.IntelligentTieringConfigurations
	}).(BucketIntelligentTieringConfigurationArrayOutput)
}

// The inventory configuration for an Amazon S3 bucket.
func (o LookupBucketResultOutput) InventoryConfigurations() BucketInventoryConfigurationArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketInventoryConfiguration { return v.InventoryConfigurations }).(BucketInventoryConfigurationArrayOutput)
}

// Rules that define how Amazon S3 manages objects during their lifetime.
func (o LookupBucketResultOutput) LifecycleConfiguration() BucketLifecycleConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketLifecycleConfiguration { return v.LifecycleConfiguration }).(BucketLifecycleConfigurationPtrOutput)
}

// Settings that define where logs are stored.
func (o LookupBucketResultOutput) LoggingConfiguration() BucketLoggingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketLoggingConfiguration { return v.LoggingConfiguration }).(BucketLoggingConfigurationPtrOutput)
}

// Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
func (o LookupBucketResultOutput) MetricsConfigurations() BucketMetricsConfigurationArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketMetricsConfiguration { return v.MetricsConfigurations }).(BucketMetricsConfigurationArrayOutput)
}

// Configuration that defines how Amazon S3 handles bucket notifications.
func (o LookupBucketResultOutput) NotificationConfiguration() BucketNotificationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketNotificationConfiguration { return v.NotificationConfiguration }).(BucketNotificationConfigurationPtrOutput)
}

// Places an Object Lock configuration on the specified bucket.
func (o LookupBucketResultOutput) ObjectLockConfiguration() BucketObjectLockConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketObjectLockConfiguration { return v.ObjectLockConfiguration }).(BucketObjectLockConfigurationPtrOutput)
}

// Specifies the container element for object ownership rules.
func (o LookupBucketResultOutput) OwnershipControls() BucketOwnershipControlsPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketOwnershipControls { return v.OwnershipControls }).(BucketOwnershipControlsPtrOutput)
}

func (o LookupBucketResultOutput) PublicAccessBlockConfiguration() BucketPublicAccessBlockConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketPublicAccessBlockConfiguration {
		return v.PublicAccessBlockConfiguration
	}).(BucketPublicAccessBlockConfigurationPtrOutput)
}

// Returns the regional domain name of the specified bucket.
func (o LookupBucketResultOutput) RegionalDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.RegionalDomainName }).(pulumi.StringPtrOutput)
}

// Configuration for replicating objects in an S3 bucket.
func (o LookupBucketResultOutput) ReplicationConfiguration() BucketReplicationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketReplicationConfiguration { return v.ReplicationConfiguration }).(BucketReplicationConfigurationPtrOutput)
}

// An arbitrary set of tags (key-value pairs) for this S3 bucket.
func (o LookupBucketResultOutput) Tags() BucketTagArrayOutput {
	return o.ApplyT(func(v LookupBucketResult) []BucketTag { return v.Tags }).(BucketTagArrayOutput)
}

func (o LookupBucketResultOutput) VersioningConfiguration() BucketVersioningConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketVersioningConfiguration { return v.VersioningConfiguration }).(BucketVersioningConfigurationPtrOutput)
}

func (o LookupBucketResultOutput) WebsiteConfiguration() BucketWebsiteConfigurationPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *BucketWebsiteConfiguration { return v.WebsiteConfiguration }).(BucketWebsiteConfigurationPtrOutput)
}

// The Amazon S3 website endpoint for the specified bucket.
func (o LookupBucketResultOutput) WebsiteURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBucketResult) *string { return v.WebsiteURL }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketResultOutput{})
}
