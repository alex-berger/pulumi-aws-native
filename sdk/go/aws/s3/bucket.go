// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::S3::Bucket
type Bucket struct {
	pulumi.CustomResourceState

	// Configuration for the transfer acceleration state.
	AccelerateConfiguration BucketAccelerateConfigurationPtrOutput `pulumi:"accelerateConfiguration"`
	// A canned access control list (ACL) that grants predefined permissions to the bucket.
	AccessControl BucketAccessControlPtrOutput `pulumi:"accessControl"`
	// The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	AnalyticsConfigurations BucketAnalyticsConfigurationArrayOutput `pulumi:"analyticsConfigurations"`
	// The Amazon Resource Name (ARN) of the specified bucket.
	Arn              pulumi.StringOutput       `pulumi:"arn"`
	BucketEncryption BucketEncryptionPtrOutput `pulumi:"bucketEncryption"`
	// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	BucketName pulumi.StringPtrOutput `pulumi:"bucketName"`
	// Rules that define cross-origin resource sharing of objects in this bucket.
	CorsConfiguration BucketCorsConfigurationPtrOutput `pulumi:"corsConfiguration"`
	// The IPv4 DNS name of the specified bucket.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
	DualStackDomainName pulumi.StringOutput `pulumi:"dualStackDomainName"`
	// Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
	IntelligentTieringConfigurations BucketIntelligentTieringConfigurationArrayOutput `pulumi:"intelligentTieringConfigurations"`
	// The inventory configuration for an Amazon S3 bucket.
	InventoryConfigurations BucketInventoryConfigurationArrayOutput `pulumi:"inventoryConfigurations"`
	// Rules that define how Amazon S3 manages objects during their lifetime.
	LifecycleConfiguration BucketLifecycleConfigurationPtrOutput `pulumi:"lifecycleConfiguration"`
	// Settings that define where logs are stored.
	LoggingConfiguration BucketLoggingConfigurationPtrOutput `pulumi:"loggingConfiguration"`
	// Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
	MetricsConfigurations BucketMetricsConfigurationArrayOutput `pulumi:"metricsConfigurations"`
	// Configuration that defines how Amazon S3 handles bucket notifications.
	NotificationConfiguration BucketNotificationConfigurationPtrOutput `pulumi:"notificationConfiguration"`
	// Places an Object Lock configuration on the specified bucket.
	ObjectLockConfiguration BucketObjectLockConfigurationPtrOutput `pulumi:"objectLockConfiguration"`
	// Indicates whether this bucket has an Object Lock configuration enabled.
	ObjectLockEnabled pulumi.BoolPtrOutput `pulumi:"objectLockEnabled"`
	// Specifies the container element for object ownership rules.
	OwnershipControls              BucketOwnershipControlsPtrOutput              `pulumi:"ownershipControls"`
	PublicAccessBlockConfiguration BucketPublicAccessBlockConfigurationPtrOutput `pulumi:"publicAccessBlockConfiguration"`
	// Returns the regional domain name of the specified bucket.
	RegionalDomainName pulumi.StringOutput `pulumi:"regionalDomainName"`
	// Configuration for replicating objects in an S3 bucket.
	ReplicationConfiguration BucketReplicationConfigurationPtrOutput `pulumi:"replicationConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	Tags                    BucketTagArrayOutput                   `pulumi:"tags"`
	VersioningConfiguration BucketVersioningConfigurationPtrOutput `pulumi:"versioningConfiguration"`
	WebsiteConfiguration    BucketWebsiteConfigurationPtrOutput    `pulumi:"websiteConfiguration"`
	// The Amazon S3 website endpoint for the specified bucket.
	WebsiteURL pulumi.StringOutput `pulumi:"websiteURL"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}

	var resource Bucket
	err := ctx.RegisterResource("aws-native:s3:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("aws-native:s3:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
}

type BucketState struct {
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// Configuration for the transfer acceleration state.
	AccelerateConfiguration *BucketAccelerateConfiguration `pulumi:"accelerateConfiguration"`
	// A canned access control list (ACL) that grants predefined permissions to the bucket.
	AccessControl *BucketAccessControl `pulumi:"accessControl"`
	// The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	AnalyticsConfigurations []BucketAnalyticsConfiguration `pulumi:"analyticsConfigurations"`
	BucketEncryption        *BucketEncryption              `pulumi:"bucketEncryption"`
	// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	BucketName *string `pulumi:"bucketName"`
	// Rules that define cross-origin resource sharing of objects in this bucket.
	CorsConfiguration *BucketCorsConfiguration `pulumi:"corsConfiguration"`
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
	// Indicates whether this bucket has an Object Lock configuration enabled.
	ObjectLockEnabled *bool `pulumi:"objectLockEnabled"`
	// Specifies the container element for object ownership rules.
	OwnershipControls              *BucketOwnershipControls              `pulumi:"ownershipControls"`
	PublicAccessBlockConfiguration *BucketPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// Configuration for replicating objects in an S3 bucket.
	ReplicationConfiguration *BucketReplicationConfiguration `pulumi:"replicationConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	Tags                    []BucketTag                    `pulumi:"tags"`
	VersioningConfiguration *BucketVersioningConfiguration `pulumi:"versioningConfiguration"`
	WebsiteConfiguration    *BucketWebsiteConfiguration    `pulumi:"websiteConfiguration"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// Configuration for the transfer acceleration state.
	AccelerateConfiguration BucketAccelerateConfigurationPtrInput
	// A canned access control list (ACL) that grants predefined permissions to the bucket.
	AccessControl BucketAccessControlPtrInput
	// The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	AnalyticsConfigurations BucketAnalyticsConfigurationArrayInput
	BucketEncryption        BucketEncryptionPtrInput
	// A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	BucketName pulumi.StringPtrInput
	// Rules that define cross-origin resource sharing of objects in this bucket.
	CorsConfiguration BucketCorsConfigurationPtrInput
	// Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
	IntelligentTieringConfigurations BucketIntelligentTieringConfigurationArrayInput
	// The inventory configuration for an Amazon S3 bucket.
	InventoryConfigurations BucketInventoryConfigurationArrayInput
	// Rules that define how Amazon S3 manages objects during their lifetime.
	LifecycleConfiguration BucketLifecycleConfigurationPtrInput
	// Settings that define where logs are stored.
	LoggingConfiguration BucketLoggingConfigurationPtrInput
	// Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
	MetricsConfigurations BucketMetricsConfigurationArrayInput
	// Configuration that defines how Amazon S3 handles bucket notifications.
	NotificationConfiguration BucketNotificationConfigurationPtrInput
	// Places an Object Lock configuration on the specified bucket.
	ObjectLockConfiguration BucketObjectLockConfigurationPtrInput
	// Indicates whether this bucket has an Object Lock configuration enabled.
	ObjectLockEnabled pulumi.BoolPtrInput
	// Specifies the container element for object ownership rules.
	OwnershipControls              BucketOwnershipControlsPtrInput
	PublicAccessBlockConfiguration BucketPublicAccessBlockConfigurationPtrInput
	// Configuration for replicating objects in an S3 bucket.
	ReplicationConfiguration BucketReplicationConfigurationPtrInput
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	Tags                    BucketTagArrayInput
	VersioningConfiguration BucketVersioningConfigurationPtrInput
	WebsiteConfiguration    BucketWebsiteConfigurationPtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterOutputType(BucketOutput{})
}
