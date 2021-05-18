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

__all__ = ['HostedZoneArgs', 'HostedZone']

@pulumi.input_type
class HostedZoneArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 hosted_zone_config: Optional[pulumi.Input['HostedZoneHostedZoneConfigArgs']] = None,
                 hosted_zone_tags: Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneHostedZoneTagArgs']]]] = None,
                 query_logging_config: Optional[pulumi.Input['HostedZoneQueryLoggingConfigArgs']] = None,
                 v_pcs: Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneVPCArgs']]]] = None):
        """
        The set of arguments for constructing a HostedZone resource.
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name
        :param pulumi.Input['HostedZoneHostedZoneConfigArgs'] hosted_zone_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig
        :param pulumi.Input[Sequence[pulumi.Input['HostedZoneHostedZoneTagArgs']]] hosted_zone_tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags
        :param pulumi.Input['HostedZoneQueryLoggingConfigArgs'] query_logging_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-queryloggingconfig
        :param pulumi.Input[Sequence[pulumi.Input['HostedZoneVPCArgs']]] v_pcs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs
        """
        pulumi.set(__self__, "name", name)
        if hosted_zone_config is not None:
            pulumi.set(__self__, "hosted_zone_config", hosted_zone_config)
        if hosted_zone_tags is not None:
            pulumi.set(__self__, "hosted_zone_tags", hosted_zone_tags)
        if query_logging_config is not None:
            pulumi.set(__self__, "query_logging_config", query_logging_config)
        if v_pcs is not None:
            pulumi.set(__self__, "v_pcs", v_pcs)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="hostedZoneConfig")
    def hosted_zone_config(self) -> Optional[pulumi.Input['HostedZoneHostedZoneConfigArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig
        """
        return pulumi.get(self, "hosted_zone_config")

    @hosted_zone_config.setter
    def hosted_zone_config(self, value: Optional[pulumi.Input['HostedZoneHostedZoneConfigArgs']]):
        pulumi.set(self, "hosted_zone_config", value)

    @property
    @pulumi.getter(name="hostedZoneTags")
    def hosted_zone_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneHostedZoneTagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags
        """
        return pulumi.get(self, "hosted_zone_tags")

    @hosted_zone_tags.setter
    def hosted_zone_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneHostedZoneTagArgs']]]]):
        pulumi.set(self, "hosted_zone_tags", value)

    @property
    @pulumi.getter(name="queryLoggingConfig")
    def query_logging_config(self) -> Optional[pulumi.Input['HostedZoneQueryLoggingConfigArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-queryloggingconfig
        """
        return pulumi.get(self, "query_logging_config")

    @query_logging_config.setter
    def query_logging_config(self, value: Optional[pulumi.Input['HostedZoneQueryLoggingConfigArgs']]):
        pulumi.set(self, "query_logging_config", value)

    @property
    @pulumi.getter(name="vPCs")
    def v_pcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneVPCArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs
        """
        return pulumi.get(self, "v_pcs")

    @v_pcs.setter
    def v_pcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneVPCArgs']]]]):
        pulumi.set(self, "v_pcs", value)


class HostedZone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hosted_zone_config: Optional[pulumi.Input[pulumi.InputType['HostedZoneHostedZoneConfigArgs']]] = None,
                 hosted_zone_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HostedZoneHostedZoneTagArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_logging_config: Optional[pulumi.Input[pulumi.InputType['HostedZoneQueryLoggingConfigArgs']]] = None,
                 v_pcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HostedZoneVPCArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['HostedZoneHostedZoneConfigArgs']] hosted_zone_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HostedZoneHostedZoneTagArgs']]]] hosted_zone_tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name
        :param pulumi.Input[pulumi.InputType['HostedZoneQueryLoggingConfigArgs']] query_logging_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-queryloggingconfig
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HostedZoneVPCArgs']]]] v_pcs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostedZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html

        :param str resource_name: The name of the resource.
        :param HostedZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostedZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hosted_zone_config: Optional[pulumi.Input[pulumi.InputType['HostedZoneHostedZoneConfigArgs']]] = None,
                 hosted_zone_tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HostedZoneHostedZoneTagArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_logging_config: Optional[pulumi.Input[pulumi.InputType['HostedZoneQueryLoggingConfigArgs']]] = None,
                 v_pcs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HostedZoneVPCArgs']]]]] = None,
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
            __props__ = HostedZoneArgs.__new__(HostedZoneArgs)

            __props__.__dict__["hosted_zone_config"] = hosted_zone_config
            __props__.__dict__["hosted_zone_tags"] = hosted_zone_tags
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["query_logging_config"] = query_logging_config
            __props__.__dict__["v_pcs"] = v_pcs
            __props__.__dict__["id"] = None
            __props__.__dict__["name_servers"] = None
        super(HostedZone, __self__).__init__(
            'aws-native:Route53:HostedZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'HostedZone':
        """
        Get an existing HostedZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = HostedZoneArgs.__new__(HostedZoneArgs)

        __props__.__dict__["hosted_zone_config"] = None
        __props__.__dict__["hosted_zone_tags"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["name_servers"] = None
        __props__.__dict__["query_logging_config"] = None
        __props__.__dict__["v_pcs"] = None
        return HostedZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="hostedZoneConfig")
    def hosted_zone_config(self) -> pulumi.Output[Optional['outputs.HostedZoneHostedZoneConfig']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig
        """
        return pulumi.get(self, "hosted_zone_config")

    @property
    @pulumi.getter(name="hostedZoneTags")
    def hosted_zone_tags(self) -> pulumi.Output[Optional[Sequence['outputs.HostedZoneHostedZoneTag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags
        """
        return pulumi.get(self, "hosted_zone_tags")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter(name="queryLoggingConfig")
    def query_logging_config(self) -> pulumi.Output[Optional['outputs.HostedZoneQueryLoggingConfig']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-queryloggingconfig
        """
        return pulumi.get(self, "query_logging_config")

    @property
    @pulumi.getter(name="vPCs")
    def v_pcs(self) -> pulumi.Output[Optional[Sequence['outputs.HostedZoneVPC']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs
        """
        return pulumi.get(self, "v_pcs")

