# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VPCDHCPOptionsAssociationArgs', 'VPCDHCPOptionsAssociation']

@pulumi.input_type
class VPCDHCPOptionsAssociationArgs:
    def __init__(__self__, *,
                 dhcp_options_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a VPCDHCPOptionsAssociation resource.
        :param pulumi.Input[str] dhcp_options_id: The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        pulumi.set(__self__, "dhcp_options_id", dhcp_options_id)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="dhcpOptionsId")
    def dhcp_options_id(self) -> pulumi.Input[str]:
        """
        The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
        """
        return pulumi.get(self, "dhcp_options_id")

    @dhcp_options_id.setter
    def dhcp_options_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dhcp_options_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


class VPCDHCPOptionsAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dhcp_options_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dhcp_options_id: The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VPCDHCPOptionsAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.

        :param str resource_name: The name of the resource.
        :param VPCDHCPOptionsAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VPCDHCPOptionsAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dhcp_options_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VPCDHCPOptionsAssociationArgs.__new__(VPCDHCPOptionsAssociationArgs)

            if dhcp_options_id is None and not opts.urn:
                raise TypeError("Missing required property 'dhcp_options_id'")
            __props__.__dict__["dhcp_options_id"] = dhcp_options_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(VPCDHCPOptionsAssociation, __self__).__init__(
            'aws-native:ec2:VPCDHCPOptionsAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VPCDHCPOptionsAssociation':
        """
        Get an existing VPCDHCPOptionsAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VPCDHCPOptionsAssociationArgs.__new__(VPCDHCPOptionsAssociationArgs)

        __props__.__dict__["dhcp_options_id"] = None
        __props__.__dict__["vpc_id"] = None
        return VPCDHCPOptionsAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dhcpOptionsId")
    def dhcp_options_id(self) -> pulumi.Output[str]:
        """
        The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
        """
        return pulumi.get(self, "dhcp_options_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

