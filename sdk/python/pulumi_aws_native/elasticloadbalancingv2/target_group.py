# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TargetGroupArgs', 'TargetGroup']

@pulumi.input_type
class TargetGroupArgs:
    def __init__(__self__, *,
                 health_check_enabled: Optional[pulumi.Input[bool]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[str]] = None,
                 health_check_protocol: Optional[pulumi.Input[str]] = None,
                 health_check_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 healthy_threshold_count: Optional[pulumi.Input[int]] = None,
                 ip_address_type: Optional[pulumi.Input[str]] = None,
                 matcher: Optional[pulumi.Input['TargetGroupMatcherArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTagArgs']]]] = None,
                 target_group_attributes: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupAttributeArgs']]]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetDescriptionArgs']]]] = None,
                 unhealthy_threshold_count: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TargetGroup resource.
        """
        if health_check_enabled is not None:
            pulumi.set(__self__, "health_check_enabled", health_check_enabled)
        if health_check_interval_seconds is not None:
            pulumi.set(__self__, "health_check_interval_seconds", health_check_interval_seconds)
        if health_check_path is not None:
            pulumi.set(__self__, "health_check_path", health_check_path)
        if health_check_port is not None:
            pulumi.set(__self__, "health_check_port", health_check_port)
        if health_check_protocol is not None:
            pulumi.set(__self__, "health_check_protocol", health_check_protocol)
        if health_check_timeout_seconds is not None:
            pulumi.set(__self__, "health_check_timeout_seconds", health_check_timeout_seconds)
        if healthy_threshold_count is not None:
            pulumi.set(__self__, "healthy_threshold_count", healthy_threshold_count)
        if ip_address_type is not None:
            pulumi.set(__self__, "ip_address_type", ip_address_type)
        if matcher is not None:
            pulumi.set(__self__, "matcher", matcher)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if protocol_version is not None:
            pulumi.set(__self__, "protocol_version", protocol_version)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_group_attributes is not None:
            pulumi.set(__self__, "target_group_attributes", target_group_attributes)
        if target_type is not None:
            pulumi.set(__self__, "target_type", target_type)
        if targets is not None:
            pulumi.set(__self__, "targets", targets)
        if unhealthy_threshold_count is not None:
            pulumi.set(__self__, "unhealthy_threshold_count", unhealthy_threshold_count)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="healthCheckEnabled")
    def health_check_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "health_check_enabled")

    @health_check_enabled.setter
    def health_check_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "health_check_enabled", value)

    @property
    @pulumi.getter(name="healthCheckIntervalSeconds")
    def health_check_interval_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_check_interval_seconds")

    @health_check_interval_seconds.setter
    def health_check_interval_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_interval_seconds", value)

    @property
    @pulumi.getter(name="healthCheckPath")
    def health_check_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "health_check_path")

    @health_check_path.setter
    def health_check_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_path", value)

    @property
    @pulumi.getter(name="healthCheckPort")
    def health_check_port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "health_check_port")

    @health_check_port.setter
    def health_check_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_port", value)

    @property
    @pulumi.getter(name="healthCheckProtocol")
    def health_check_protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "health_check_protocol")

    @health_check_protocol.setter
    def health_check_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_check_protocol", value)

    @property
    @pulumi.getter(name="healthCheckTimeoutSeconds")
    def health_check_timeout_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "health_check_timeout_seconds")

    @health_check_timeout_seconds.setter
    def health_check_timeout_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "health_check_timeout_seconds", value)

    @property
    @pulumi.getter(name="healthyThresholdCount")
    def healthy_threshold_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "healthy_threshold_count")

    @healthy_threshold_count.setter
    def healthy_threshold_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "healthy_threshold_count", value)

    @property
    @pulumi.getter(name="ipAddressType")
    def ip_address_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_address_type")

    @ip_address_type.setter
    def ip_address_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address_type", value)

    @property
    @pulumi.getter
    def matcher(self) -> Optional[pulumi.Input['TargetGroupMatcherArgs']]:
        return pulumi.get(self, "matcher")

    @matcher.setter
    def matcher(self, value: Optional[pulumi.Input['TargetGroupMatcherArgs']]):
        pulumi.set(self, "matcher", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="protocolVersion")
    def protocol_version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol_version")

    @protocol_version.setter
    def protocol_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol_version", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetGroupAttributes")
    def target_group_attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupAttributeArgs']]]]:
        return pulumi.get(self, "target_group_attributes")

    @target_group_attributes.setter
    def target_group_attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupAttributeArgs']]]]):
        pulumi.set(self, "target_group_attributes", value)

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target_type")

    @target_type.setter
    def target_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_type", value)

    @property
    @pulumi.getter
    def targets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetDescriptionArgs']]]]:
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TargetGroupTargetDescriptionArgs']]]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter(name="unhealthyThresholdCount")
    def unhealthy_threshold_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "unhealthy_threshold_count")

    @unhealthy_threshold_count.setter
    def unhealthy_threshold_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "unhealthy_threshold_count", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


warnings.warn("""TargetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class TargetGroup(pulumi.CustomResource):
    warnings.warn("""TargetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 health_check_enabled: Optional[pulumi.Input[bool]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[str]] = None,
                 health_check_protocol: Optional[pulumi.Input[str]] = None,
                 health_check_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 healthy_threshold_count: Optional[pulumi.Input[int]] = None,
                 ip_address_type: Optional[pulumi.Input[str]] = None,
                 matcher: Optional[pulumi.Input[pulumi.InputType['TargetGroupMatcherArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTagArgs']]]]] = None,
                 target_group_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupAttributeArgs']]]]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTargetDescriptionArgs']]]]] = None,
                 unhealthy_threshold_count: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TargetGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::ElasticLoadBalancingV2::TargetGroup

        :param str resource_name: The name of the resource.
        :param TargetGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TargetGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 health_check_enabled: Optional[pulumi.Input[bool]] = None,
                 health_check_interval_seconds: Optional[pulumi.Input[int]] = None,
                 health_check_path: Optional[pulumi.Input[str]] = None,
                 health_check_port: Optional[pulumi.Input[str]] = None,
                 health_check_protocol: Optional[pulumi.Input[str]] = None,
                 health_check_timeout_seconds: Optional[pulumi.Input[int]] = None,
                 healthy_threshold_count: Optional[pulumi.Input[int]] = None,
                 ip_address_type: Optional[pulumi.Input[str]] = None,
                 matcher: Optional[pulumi.Input[pulumi.InputType['TargetGroupMatcherArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 protocol_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTagArgs']]]]] = None,
                 target_group_attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupAttributeArgs']]]]] = None,
                 target_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TargetGroupTargetDescriptionArgs']]]]] = None,
                 unhealthy_threshold_count: Optional[pulumi.Input[int]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""TargetGroup is deprecated: TargetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TargetGroupArgs.__new__(TargetGroupArgs)

            __props__.__dict__["health_check_enabled"] = health_check_enabled
            __props__.__dict__["health_check_interval_seconds"] = health_check_interval_seconds
            __props__.__dict__["health_check_path"] = health_check_path
            __props__.__dict__["health_check_port"] = health_check_port
            __props__.__dict__["health_check_protocol"] = health_check_protocol
            __props__.__dict__["health_check_timeout_seconds"] = health_check_timeout_seconds
            __props__.__dict__["healthy_threshold_count"] = healthy_threshold_count
            __props__.__dict__["ip_address_type"] = ip_address_type
            __props__.__dict__["matcher"] = matcher
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["protocol_version"] = protocol_version
            __props__.__dict__["tags"] = tags
            __props__.__dict__["target_group_attributes"] = target_group_attributes
            __props__.__dict__["target_type"] = target_type
            __props__.__dict__["targets"] = targets
            __props__.__dict__["unhealthy_threshold_count"] = unhealthy_threshold_count
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["load_balancer_arns"] = None
            __props__.__dict__["target_group_full_name"] = None
            __props__.__dict__["target_group_name"] = None
        super(TargetGroup, __self__).__init__(
            'aws-native:elasticloadbalancingv2:TargetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TargetGroup':
        """
        Get an existing TargetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TargetGroupArgs.__new__(TargetGroupArgs)

        __props__.__dict__["health_check_enabled"] = None
        __props__.__dict__["health_check_interval_seconds"] = None
        __props__.__dict__["health_check_path"] = None
        __props__.__dict__["health_check_port"] = None
        __props__.__dict__["health_check_protocol"] = None
        __props__.__dict__["health_check_timeout_seconds"] = None
        __props__.__dict__["healthy_threshold_count"] = None
        __props__.__dict__["ip_address_type"] = None
        __props__.__dict__["load_balancer_arns"] = None
        __props__.__dict__["matcher"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["port"] = None
        __props__.__dict__["protocol"] = None
        __props__.__dict__["protocol_version"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["target_group_attributes"] = None
        __props__.__dict__["target_group_full_name"] = None
        __props__.__dict__["target_group_name"] = None
        __props__.__dict__["target_type"] = None
        __props__.__dict__["targets"] = None
        __props__.__dict__["unhealthy_threshold_count"] = None
        __props__.__dict__["vpc_id"] = None
        return TargetGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="healthCheckEnabled")
    def health_check_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "health_check_enabled")

    @property
    @pulumi.getter(name="healthCheckIntervalSeconds")
    def health_check_interval_seconds(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "health_check_interval_seconds")

    @property
    @pulumi.getter(name="healthCheckPath")
    def health_check_path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "health_check_path")

    @property
    @pulumi.getter(name="healthCheckPort")
    def health_check_port(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "health_check_port")

    @property
    @pulumi.getter(name="healthCheckProtocol")
    def health_check_protocol(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "health_check_protocol")

    @property
    @pulumi.getter(name="healthCheckTimeoutSeconds")
    def health_check_timeout_seconds(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "health_check_timeout_seconds")

    @property
    @pulumi.getter(name="healthyThresholdCount")
    def healthy_threshold_count(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "healthy_threshold_count")

    @property
    @pulumi.getter(name="ipAddressType")
    def ip_address_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ip_address_type")

    @property
    @pulumi.getter(name="loadBalancerArns")
    def load_balancer_arns(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "load_balancer_arns")

    @property
    @pulumi.getter
    def matcher(self) -> pulumi.Output[Optional['outputs.TargetGroupMatcher']]:
        return pulumi.get(self, "matcher")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolVersion")
    def protocol_version(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "protocol_version")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.TargetGroupTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetGroupAttributes")
    def target_group_attributes(self) -> pulumi.Output[Optional[Sequence['outputs.TargetGroupAttribute']]]:
        return pulumi.get(self, "target_group_attributes")

    @property
    @pulumi.getter(name="targetGroupFullName")
    def target_group_full_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_group_full_name")

    @property
    @pulumi.getter(name="targetGroupName")
    def target_group_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "target_group_name")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Optional[Sequence['outputs.TargetGroupTargetDescription']]]:
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="unhealthyThresholdCount")
    def unhealthy_threshold_count(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "unhealthy_threshold_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "vpc_id")

