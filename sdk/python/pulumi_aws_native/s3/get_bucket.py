# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetBucketResult',
    'AwaitableGetBucketResult',
    'get_bucket',
    'get_bucket_output',
]

@pulumi.output_type
class GetBucketResult:
    def __init__(__self__, accelerate_configuration=None, access_control=None, analytics_configurations=None, arn=None, bucket_encryption=None, cors_configuration=None, domain_name=None, dual_stack_domain_name=None, intelligent_tiering_configurations=None, inventory_configurations=None, lifecycle_configuration=None, logging_configuration=None, metrics_configurations=None, notification_configuration=None, object_lock_configuration=None, ownership_controls=None, public_access_block_configuration=None, regional_domain_name=None, replication_configuration=None, tags=None, versioning_configuration=None, website_configuration=None, website_url=None):
        if accelerate_configuration and not isinstance(accelerate_configuration, dict):
            raise TypeError("Expected argument 'accelerate_configuration' to be a dict")
        pulumi.set(__self__, "accelerate_configuration", accelerate_configuration)
        if access_control and not isinstance(access_control, str):
            raise TypeError("Expected argument 'access_control' to be a str")
        pulumi.set(__self__, "access_control", access_control)
        if analytics_configurations and not isinstance(analytics_configurations, list):
            raise TypeError("Expected argument 'analytics_configurations' to be a list")
        pulumi.set(__self__, "analytics_configurations", analytics_configurations)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if bucket_encryption and not isinstance(bucket_encryption, dict):
            raise TypeError("Expected argument 'bucket_encryption' to be a dict")
        pulumi.set(__self__, "bucket_encryption", bucket_encryption)
        if cors_configuration and not isinstance(cors_configuration, dict):
            raise TypeError("Expected argument 'cors_configuration' to be a dict")
        pulumi.set(__self__, "cors_configuration", cors_configuration)
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        pulumi.set(__self__, "domain_name", domain_name)
        if dual_stack_domain_name and not isinstance(dual_stack_domain_name, str):
            raise TypeError("Expected argument 'dual_stack_domain_name' to be a str")
        pulumi.set(__self__, "dual_stack_domain_name", dual_stack_domain_name)
        if intelligent_tiering_configurations and not isinstance(intelligent_tiering_configurations, list):
            raise TypeError("Expected argument 'intelligent_tiering_configurations' to be a list")
        pulumi.set(__self__, "intelligent_tiering_configurations", intelligent_tiering_configurations)
        if inventory_configurations and not isinstance(inventory_configurations, list):
            raise TypeError("Expected argument 'inventory_configurations' to be a list")
        pulumi.set(__self__, "inventory_configurations", inventory_configurations)
        if lifecycle_configuration and not isinstance(lifecycle_configuration, dict):
            raise TypeError("Expected argument 'lifecycle_configuration' to be a dict")
        pulumi.set(__self__, "lifecycle_configuration", lifecycle_configuration)
        if logging_configuration and not isinstance(logging_configuration, dict):
            raise TypeError("Expected argument 'logging_configuration' to be a dict")
        pulumi.set(__self__, "logging_configuration", logging_configuration)
        if metrics_configurations and not isinstance(metrics_configurations, list):
            raise TypeError("Expected argument 'metrics_configurations' to be a list")
        pulumi.set(__self__, "metrics_configurations", metrics_configurations)
        if notification_configuration and not isinstance(notification_configuration, dict):
            raise TypeError("Expected argument 'notification_configuration' to be a dict")
        pulumi.set(__self__, "notification_configuration", notification_configuration)
        if object_lock_configuration and not isinstance(object_lock_configuration, dict):
            raise TypeError("Expected argument 'object_lock_configuration' to be a dict")
        pulumi.set(__self__, "object_lock_configuration", object_lock_configuration)
        if ownership_controls and not isinstance(ownership_controls, dict):
            raise TypeError("Expected argument 'ownership_controls' to be a dict")
        pulumi.set(__self__, "ownership_controls", ownership_controls)
        if public_access_block_configuration and not isinstance(public_access_block_configuration, dict):
            raise TypeError("Expected argument 'public_access_block_configuration' to be a dict")
        pulumi.set(__self__, "public_access_block_configuration", public_access_block_configuration)
        if regional_domain_name and not isinstance(regional_domain_name, str):
            raise TypeError("Expected argument 'regional_domain_name' to be a str")
        pulumi.set(__self__, "regional_domain_name", regional_domain_name)
        if replication_configuration and not isinstance(replication_configuration, dict):
            raise TypeError("Expected argument 'replication_configuration' to be a dict")
        pulumi.set(__self__, "replication_configuration", replication_configuration)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if versioning_configuration and not isinstance(versioning_configuration, dict):
            raise TypeError("Expected argument 'versioning_configuration' to be a dict")
        pulumi.set(__self__, "versioning_configuration", versioning_configuration)
        if website_configuration and not isinstance(website_configuration, dict):
            raise TypeError("Expected argument 'website_configuration' to be a dict")
        pulumi.set(__self__, "website_configuration", website_configuration)
        if website_url and not isinstance(website_url, str):
            raise TypeError("Expected argument 'website_url' to be a str")
        pulumi.set(__self__, "website_url", website_url)

    @property
    @pulumi.getter(name="accelerateConfiguration")
    def accelerate_configuration(self) -> Optional['outputs.BucketAccelerateConfiguration']:
        """
        Configuration for the transfer acceleration state.
        """
        return pulumi.get(self, "accelerate_configuration")

    @property
    @pulumi.getter(name="accessControl")
    def access_control(self) -> Optional['BucketAccessControl']:
        """
        A canned access control list (ACL) that grants predefined permissions to the bucket.
        """
        return pulumi.get(self, "access_control")

    @property
    @pulumi.getter(name="analyticsConfigurations")
    def analytics_configurations(self) -> Optional[Sequence['outputs.BucketAnalyticsConfiguration']]:
        """
        The configuration and any analyses for the analytics filter of an Amazon S3 bucket.
        """
        return pulumi.get(self, "analytics_configurations")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the specified bucket.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="bucketEncryption")
    def bucket_encryption(self) -> Optional['outputs.BucketEncryption']:
        return pulumi.get(self, "bucket_encryption")

    @property
    @pulumi.getter(name="corsConfiguration")
    def cors_configuration(self) -> Optional['outputs.BucketCorsConfiguration']:
        """
        Rules that define cross-origin resource sharing of objects in this bucket.
        """
        return pulumi.get(self, "cors_configuration")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[str]:
        """
        The IPv4 DNS name of the specified bucket.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="dualStackDomainName")
    def dual_stack_domain_name(self) -> Optional[str]:
        """
        The IPv6 DNS name of the specified bucket. For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html).
        """
        return pulumi.get(self, "dual_stack_domain_name")

    @property
    @pulumi.getter(name="intelligentTieringConfigurations")
    def intelligent_tiering_configurations(self) -> Optional[Sequence['outputs.BucketIntelligentTieringConfiguration']]:
        """
        Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
        """
        return pulumi.get(self, "intelligent_tiering_configurations")

    @property
    @pulumi.getter(name="inventoryConfigurations")
    def inventory_configurations(self) -> Optional[Sequence['outputs.BucketInventoryConfiguration']]:
        """
        The inventory configuration for an Amazon S3 bucket.
        """
        return pulumi.get(self, "inventory_configurations")

    @property
    @pulumi.getter(name="lifecycleConfiguration")
    def lifecycle_configuration(self) -> Optional['outputs.BucketLifecycleConfiguration']:
        """
        Rules that define how Amazon S3 manages objects during their lifetime.
        """
        return pulumi.get(self, "lifecycle_configuration")

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> Optional['outputs.BucketLoggingConfiguration']:
        """
        Settings that define where logs are stored.
        """
        return pulumi.get(self, "logging_configuration")

    @property
    @pulumi.getter(name="metricsConfigurations")
    def metrics_configurations(self) -> Optional[Sequence['outputs.BucketMetricsConfiguration']]:
        """
        Settings that define a metrics configuration for the CloudWatch request metrics from the bucket.
        """
        return pulumi.get(self, "metrics_configurations")

    @property
    @pulumi.getter(name="notificationConfiguration")
    def notification_configuration(self) -> Optional['outputs.BucketNotificationConfiguration']:
        """
        Configuration that defines how Amazon S3 handles bucket notifications.
        """
        return pulumi.get(self, "notification_configuration")

    @property
    @pulumi.getter(name="objectLockConfiguration")
    def object_lock_configuration(self) -> Optional['outputs.BucketObjectLockConfiguration']:
        """
        Places an Object Lock configuration on the specified bucket.
        """
        return pulumi.get(self, "object_lock_configuration")

    @property
    @pulumi.getter(name="ownershipControls")
    def ownership_controls(self) -> Optional['outputs.BucketOwnershipControls']:
        """
        Specifies the container element for object ownership rules.
        """
        return pulumi.get(self, "ownership_controls")

    @property
    @pulumi.getter(name="publicAccessBlockConfiguration")
    def public_access_block_configuration(self) -> Optional['outputs.BucketPublicAccessBlockConfiguration']:
        return pulumi.get(self, "public_access_block_configuration")

    @property
    @pulumi.getter(name="regionalDomainName")
    def regional_domain_name(self) -> Optional[str]:
        """
        Returns the regional domain name of the specified bucket.
        """
        return pulumi.get(self, "regional_domain_name")

    @property
    @pulumi.getter(name="replicationConfiguration")
    def replication_configuration(self) -> Optional['outputs.BucketReplicationConfiguration']:
        """
        Configuration for replicating objects in an S3 bucket.
        """
        return pulumi.get(self, "replication_configuration")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.BucketTag']]:
        """
        An arbitrary set of tags (key-value pairs) for this S3 bucket.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="versioningConfiguration")
    def versioning_configuration(self) -> Optional['outputs.BucketVersioningConfiguration']:
        return pulumi.get(self, "versioning_configuration")

    @property
    @pulumi.getter(name="websiteConfiguration")
    def website_configuration(self) -> Optional['outputs.BucketWebsiteConfiguration']:
        return pulumi.get(self, "website_configuration")

    @property
    @pulumi.getter(name="websiteURL")
    def website_url(self) -> Optional[str]:
        """
        The Amazon S3 website endpoint for the specified bucket.
        """
        return pulumi.get(self, "website_url")


class AwaitableGetBucketResult(GetBucketResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBucketResult(
            accelerate_configuration=self.accelerate_configuration,
            access_control=self.access_control,
            analytics_configurations=self.analytics_configurations,
            arn=self.arn,
            bucket_encryption=self.bucket_encryption,
            cors_configuration=self.cors_configuration,
            domain_name=self.domain_name,
            dual_stack_domain_name=self.dual_stack_domain_name,
            intelligent_tiering_configurations=self.intelligent_tiering_configurations,
            inventory_configurations=self.inventory_configurations,
            lifecycle_configuration=self.lifecycle_configuration,
            logging_configuration=self.logging_configuration,
            metrics_configurations=self.metrics_configurations,
            notification_configuration=self.notification_configuration,
            object_lock_configuration=self.object_lock_configuration,
            ownership_controls=self.ownership_controls,
            public_access_block_configuration=self.public_access_block_configuration,
            regional_domain_name=self.regional_domain_name,
            replication_configuration=self.replication_configuration,
            tags=self.tags,
            versioning_configuration=self.versioning_configuration,
            website_configuration=self.website_configuration,
            website_url=self.website_url)


def get_bucket(bucket_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBucketResult:
    """
    Resource Type definition for AWS::S3::Bucket


    :param str bucket_name: A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
    """
    __args__ = dict()
    __args__['bucketName'] = bucket_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:s3:getBucket', __args__, opts=opts, typ=GetBucketResult).value

    return AwaitableGetBucketResult(
        accelerate_configuration=__ret__.accelerate_configuration,
        access_control=__ret__.access_control,
        analytics_configurations=__ret__.analytics_configurations,
        arn=__ret__.arn,
        bucket_encryption=__ret__.bucket_encryption,
        cors_configuration=__ret__.cors_configuration,
        domain_name=__ret__.domain_name,
        dual_stack_domain_name=__ret__.dual_stack_domain_name,
        intelligent_tiering_configurations=__ret__.intelligent_tiering_configurations,
        inventory_configurations=__ret__.inventory_configurations,
        lifecycle_configuration=__ret__.lifecycle_configuration,
        logging_configuration=__ret__.logging_configuration,
        metrics_configurations=__ret__.metrics_configurations,
        notification_configuration=__ret__.notification_configuration,
        object_lock_configuration=__ret__.object_lock_configuration,
        ownership_controls=__ret__.ownership_controls,
        public_access_block_configuration=__ret__.public_access_block_configuration,
        regional_domain_name=__ret__.regional_domain_name,
        replication_configuration=__ret__.replication_configuration,
        tags=__ret__.tags,
        versioning_configuration=__ret__.versioning_configuration,
        website_configuration=__ret__.website_configuration,
        website_url=__ret__.website_url)


@_utilities.lift_output_func(get_bucket)
def get_bucket_output(bucket_name: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBucketResult]:
    """
    Resource Type definition for AWS::S3::Bucket


    :param str bucket_name: A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
    """
    ...
