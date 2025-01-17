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
from ._inputs import *

__all__ = ['TopicRuleDestinationArgs', 'TopicRuleDestination']

@pulumi.input_type
class TopicRuleDestinationArgs:
    def __init__(__self__, *,
                 http_url_properties: Optional[pulumi.Input['TopicRuleDestinationHttpUrlDestinationSummaryArgs']] = None,
                 status: Optional[pulumi.Input['TopicRuleDestinationStatus']] = None,
                 vpc_properties: Optional[pulumi.Input['TopicRuleDestinationVpcDestinationPropertiesArgs']] = None):
        """
        The set of arguments for constructing a TopicRuleDestination resource.
        :param pulumi.Input['TopicRuleDestinationHttpUrlDestinationSummaryArgs'] http_url_properties: HTTP URL destination properties.
        :param pulumi.Input['TopicRuleDestinationStatus'] status: The status of the TopicRuleDestination.
        :param pulumi.Input['TopicRuleDestinationVpcDestinationPropertiesArgs'] vpc_properties: VPC destination properties.
        """
        if http_url_properties is not None:
            pulumi.set(__self__, "http_url_properties", http_url_properties)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if vpc_properties is not None:
            pulumi.set(__self__, "vpc_properties", vpc_properties)

    @property
    @pulumi.getter(name="httpUrlProperties")
    def http_url_properties(self) -> Optional[pulumi.Input['TopicRuleDestinationHttpUrlDestinationSummaryArgs']]:
        """
        HTTP URL destination properties.
        """
        return pulumi.get(self, "http_url_properties")

    @http_url_properties.setter
    def http_url_properties(self, value: Optional[pulumi.Input['TopicRuleDestinationHttpUrlDestinationSummaryArgs']]):
        pulumi.set(self, "http_url_properties", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['TopicRuleDestinationStatus']]:
        """
        The status of the TopicRuleDestination.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['TopicRuleDestinationStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="vpcProperties")
    def vpc_properties(self) -> Optional[pulumi.Input['TopicRuleDestinationVpcDestinationPropertiesArgs']]:
        """
        VPC destination properties.
        """
        return pulumi.get(self, "vpc_properties")

    @vpc_properties.setter
    def vpc_properties(self, value: Optional[pulumi.Input['TopicRuleDestinationVpcDestinationPropertiesArgs']]):
        pulumi.set(self, "vpc_properties", value)


class TopicRuleDestination(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_url_properties: Optional[pulumi.Input[pulumi.InputType['TopicRuleDestinationHttpUrlDestinationSummaryArgs']]] = None,
                 status: Optional[pulumi.Input['TopicRuleDestinationStatus']] = None,
                 vpc_properties: Optional[pulumi.Input[pulumi.InputType['TopicRuleDestinationVpcDestinationPropertiesArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::IoT::TopicRuleDestination

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TopicRuleDestinationHttpUrlDestinationSummaryArgs']] http_url_properties: HTTP URL destination properties.
        :param pulumi.Input['TopicRuleDestinationStatus'] status: The status of the TopicRuleDestination.
        :param pulumi.Input[pulumi.InputType['TopicRuleDestinationVpcDestinationPropertiesArgs']] vpc_properties: VPC destination properties.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TopicRuleDestinationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::IoT::TopicRuleDestination

        :param str resource_name: The name of the resource.
        :param TopicRuleDestinationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TopicRuleDestinationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 http_url_properties: Optional[pulumi.Input[pulumi.InputType['TopicRuleDestinationHttpUrlDestinationSummaryArgs']]] = None,
                 status: Optional[pulumi.Input['TopicRuleDestinationStatus']] = None,
                 vpc_properties: Optional[pulumi.Input[pulumi.InputType['TopicRuleDestinationVpcDestinationPropertiesArgs']]] = None,
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
            __props__ = TopicRuleDestinationArgs.__new__(TopicRuleDestinationArgs)

            __props__.__dict__["http_url_properties"] = http_url_properties
            __props__.__dict__["status"] = status
            __props__.__dict__["vpc_properties"] = vpc_properties
            __props__.__dict__["arn"] = None
            __props__.__dict__["status_reason"] = None
        super(TopicRuleDestination, __self__).__init__(
            'aws-native:iot:TopicRuleDestination',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TopicRuleDestination':
        """
        Get an existing TopicRuleDestination resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TopicRuleDestinationArgs.__new__(TopicRuleDestinationArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["http_url_properties"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_reason"] = None
        __props__.__dict__["vpc_properties"] = None
        return TopicRuleDestination(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN).
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="httpUrlProperties")
    def http_url_properties(self) -> pulumi.Output[Optional['outputs.TopicRuleDestinationHttpUrlDestinationSummary']]:
        """
        HTTP URL destination properties.
        """
        return pulumi.get(self, "http_url_properties")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['TopicRuleDestinationStatus']]:
        """
        The status of the TopicRuleDestination.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusReason")
    def status_reason(self) -> pulumi.Output[str]:
        """
        The reasoning for the current status of the TopicRuleDestination.
        """
        return pulumi.get(self, "status_reason")

    @property
    @pulumi.getter(name="vpcProperties")
    def vpc_properties(self) -> pulumi.Output[Optional['outputs.TopicRuleDestinationVpcDestinationProperties']]:
        """
        VPC destination properties.
        """
        return pulumi.get(self, "vpc_properties")

