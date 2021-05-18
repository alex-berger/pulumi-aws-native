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

__all__ = ['HealthCheckArgs', 'HealthCheck']

@pulumi.input_type
class HealthCheckArgs:
    def __init__(__self__, *,
                 health_check_config: pulumi.Input[Union[Any, str]],
                 health_check_tags: Optional[pulumi.Input[Sequence[pulumi.Input['HealthCheckHealthCheckTagArgs']]]] = None):
        """
        The set of arguments for constructing a HealthCheck resource.
        :param pulumi.Input[Union[Any, str]] health_check_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
        :param pulumi.Input[Sequence[pulumi.Input['HealthCheckHealthCheckTagArgs']]] health_check_tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
        """
        pulumi.set(__self__, "health_check_config", health_check_config)
        if health_check_tags is not None:
            pulumi.set(__self__, "health_check_tags", health_check_tags)

    @property
    @pulumi.getter(name="healthCheckConfig")
    def health_check_config(self) -> pulumi.Input[Union[Any, str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
        """
        return pulumi.get(self, "health_check_config")

    @health_check_config.setter
    def health_check_config(self, value: pulumi.Input[Union[Any, str]]):
        pulumi.set(self, "health_check_config", value)

    @property
    @pulumi.getter(name="healthCheckTags")
    def health_check_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['HealthCheckHealthCheckTagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
        """
        return pulumi.get(self, "health_check_tags")

    @health_check_tags.setter
    def health_check_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['HealthCheckHealthCheckTagArgs']]]]):
        pulumi.set(self, "health_check_tags", value)


class HealthCheck(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 health_check_config: Optional[pulumi.Input[Union[Any, str]]] = None,
                 health_check_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HealthCheckHealthCheckTagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[Any, str]] health_check_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HealthCheckHealthCheckTagArgs']]]] health_check_tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HealthCheckArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html

        :param str resource_name: The name of the resource.
        :param HealthCheckArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HealthCheckArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 health_check_config: Optional[pulumi.Input[Union[Any, str]]] = None,
                 health_check_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HealthCheckHealthCheckTagArgs']]]]] = None,
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
            __props__ = HealthCheckArgs.__new__(HealthCheckArgs)

            if health_check_config is None and not opts.urn:
                raise TypeError("Missing required property 'health_check_config'")
            __props__.__dict__["health_check_config"] = health_check_config
            __props__.__dict__["health_check_tags"] = health_check_tags
            __props__.__dict__["health_check_id"] = None
        super(HealthCheck, __self__).__init__(
            'aws-native:Route53:HealthCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HealthCheck':
        """
        Get an existing HealthCheck resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = HealthCheckArgs.__new__(HealthCheckArgs)

        __props__.__dict__["health_check_config"] = None
        __props__.__dict__["health_check_id"] = None
        __props__.__dict__["health_check_tags"] = None
        return HealthCheck(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="healthCheckConfig")
    def health_check_config(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
        """
        return pulumi.get(self, "health_check_config")

    @property
    @pulumi.getter(name="healthCheckId")
    def health_check_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "health_check_id")

    @property
    @pulumi.getter(name="healthCheckTags")
    def health_check_tags(self) -> pulumi.Output[Optional[Sequence['outputs.HealthCheckHealthCheckTag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
        """
        return pulumi.get(self, "health_check_tags")

