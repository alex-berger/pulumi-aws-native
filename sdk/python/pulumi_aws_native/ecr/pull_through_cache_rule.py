# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PullThroughCacheRuleArgs', 'PullThroughCacheRule']

@pulumi.input_type
class PullThroughCacheRuleArgs:
    def __init__(__self__, *,
                 ecr_repository_prefix: Optional[pulumi.Input[str]] = None,
                 upstream_registry_url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PullThroughCacheRule resource.
        :param pulumi.Input[str] ecr_repository_prefix: The ECRRepositoryPrefix is a custom alias for upstream registry url.
        :param pulumi.Input[str] upstream_registry_url: The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached
        """
        if ecr_repository_prefix is not None:
            pulumi.set(__self__, "ecr_repository_prefix", ecr_repository_prefix)
        if upstream_registry_url is not None:
            pulumi.set(__self__, "upstream_registry_url", upstream_registry_url)

    @property
    @pulumi.getter(name="ecrRepositoryPrefix")
    def ecr_repository_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The ECRRepositoryPrefix is a custom alias for upstream registry url.
        """
        return pulumi.get(self, "ecr_repository_prefix")

    @ecr_repository_prefix.setter
    def ecr_repository_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ecr_repository_prefix", value)

    @property
    @pulumi.getter(name="upstreamRegistryUrl")
    def upstream_registry_url(self) -> Optional[pulumi.Input[str]]:
        """
        The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached
        """
        return pulumi.get(self, "upstream_registry_url")

    @upstream_registry_url.setter
    def upstream_registry_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "upstream_registry_url", value)


class PullThroughCacheRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ecr_repository_prefix: Optional[pulumi.Input[str]] = None,
                 upstream_registry_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The AWS::ECR::PullThroughCacheRule resource configures the upstream registry configuration details for an Amazon Elastic Container Registry (Amazon Private ECR) pull-through cache.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ecr_repository_prefix: The ECRRepositoryPrefix is a custom alias for upstream registry url.
        :param pulumi.Input[str] upstream_registry_url: The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[PullThroughCacheRuleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::ECR::PullThroughCacheRule resource configures the upstream registry configuration details for an Amazon Elastic Container Registry (Amazon Private ECR) pull-through cache.

        :param str resource_name: The name of the resource.
        :param PullThroughCacheRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PullThroughCacheRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ecr_repository_prefix: Optional[pulumi.Input[str]] = None,
                 upstream_registry_url: Optional[pulumi.Input[str]] = None,
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
            __props__ = PullThroughCacheRuleArgs.__new__(PullThroughCacheRuleArgs)

            __props__.__dict__["ecr_repository_prefix"] = ecr_repository_prefix
            __props__.__dict__["upstream_registry_url"] = upstream_registry_url
        super(PullThroughCacheRule, __self__).__init__(
            'aws-native:ecr:PullThroughCacheRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PullThroughCacheRule':
        """
        Get an existing PullThroughCacheRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PullThroughCacheRuleArgs.__new__(PullThroughCacheRuleArgs)

        __props__.__dict__["ecr_repository_prefix"] = None
        __props__.__dict__["upstream_registry_url"] = None
        return PullThroughCacheRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ecrRepositoryPrefix")
    def ecr_repository_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The ECRRepositoryPrefix is a custom alias for upstream registry url.
        """
        return pulumi.get(self, "ecr_repository_prefix")

    @property
    @pulumi.getter(name="upstreamRegistryUrl")
    def upstream_registry_url(self) -> pulumi.Output[Optional[str]]:
        """
        The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached
        """
        return pulumi.get(self, "upstream_registry_url")

