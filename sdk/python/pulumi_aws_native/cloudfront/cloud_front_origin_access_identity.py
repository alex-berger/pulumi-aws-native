# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CloudFrontOriginAccessIdentityArgs', 'CloudFrontOriginAccessIdentity']

@pulumi.input_type
class CloudFrontOriginAccessIdentityArgs:
    def __init__(__self__, *,
                 cloud_front_origin_access_identity_config: pulumi.Input['CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigArgs']):
        """
        The set of arguments for constructing a CloudFrontOriginAccessIdentity resource.
        :param pulumi.Input['CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigArgs'] cloud_front_origin_access_identity_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig
        """
        pulumi.set(__self__, "cloud_front_origin_access_identity_config", cloud_front_origin_access_identity_config)

    @property
    @pulumi.getter(name="cloudFrontOriginAccessIdentityConfig")
    def cloud_front_origin_access_identity_config(self) -> pulumi.Input['CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig
        """
        return pulumi.get(self, "cloud_front_origin_access_identity_config")

    @cloud_front_origin_access_identity_config.setter
    def cloud_front_origin_access_identity_config(self, value: pulumi.Input['CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigArgs']):
        pulumi.set(self, "cloud_front_origin_access_identity_config", value)


class CloudFrontOriginAccessIdentity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_front_origin_access_identity_config: Optional[pulumi.Input[pulumi.InputType['CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigArgs']]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigArgs']] cloud_front_origin_access_identity_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudFrontOriginAccessIdentityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html

        :param str resource_name: The name of the resource.
        :param CloudFrontOriginAccessIdentityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudFrontOriginAccessIdentityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_front_origin_access_identity_config: Optional[pulumi.Input[pulumi.InputType['CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfigArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudFrontOriginAccessIdentityArgs.__new__(CloudFrontOriginAccessIdentityArgs)

            if cloud_front_origin_access_identity_config is None and not opts.urn:
                raise TypeError("Missing required property 'cloud_front_origin_access_identity_config'")
            __props__.__dict__["cloud_front_origin_access_identity_config"] = cloud_front_origin_access_identity_config
            __props__.__dict__["s3_canonical_user_id"] = None
        super(CloudFrontOriginAccessIdentity, __self__).__init__(
            'aws-native:CloudFront:CloudFrontOriginAccessIdentity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CloudFrontOriginAccessIdentity':
        """
        Get an existing CloudFrontOriginAccessIdentity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CloudFrontOriginAccessIdentityArgs.__new__(CloudFrontOriginAccessIdentityArgs)

        __props__.__dict__["cloud_front_origin_access_identity_config"] = None
        __props__.__dict__["s3_canonical_user_id"] = None
        return CloudFrontOriginAccessIdentity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cloudFrontOriginAccessIdentityConfig")
    def cloud_front_origin_access_identity_config(self) -> pulumi.Output['outputs.CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfig']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig
        """
        return pulumi.get(self, "cloud_front_origin_access_identity_config")

    @property
    @pulumi.getter(name="s3CanonicalUserId")
    def s3_canonical_user_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "s3_canonical_user_id")

