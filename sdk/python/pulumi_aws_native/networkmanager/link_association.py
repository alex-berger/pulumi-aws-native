# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LinkAssociationArgs', 'LinkAssociation']

@pulumi.input_type
class LinkAssociationArgs:
    def __init__(__self__, *,
                 device_id: pulumi.Input[str],
                 global_network_id: pulumi.Input[str],
                 link_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a LinkAssociation resource.
        :param pulumi.Input[str] device_id: The ID of the device
        :param pulumi.Input[str] global_network_id: The ID of the global network.
        :param pulumi.Input[str] link_id: The ID of the link
        """
        pulumi.set(__self__, "device_id", device_id)
        pulumi.set(__self__, "global_network_id", global_network_id)
        pulumi.set(__self__, "link_id", link_id)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Input[str]:
        """
        The ID of the device
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="globalNetworkId")
    def global_network_id(self) -> pulumi.Input[str]:
        """
        The ID of the global network.
        """
        return pulumi.get(self, "global_network_id")

    @global_network_id.setter
    def global_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "global_network_id", value)

    @property
    @pulumi.getter(name="linkId")
    def link_id(self) -> pulumi.Input[str]:
        """
        The ID of the link
        """
        return pulumi.get(self, "link_id")

    @link_id.setter
    def link_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "link_id", value)


class LinkAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 global_network_id: Optional[pulumi.Input[str]] = None,
                 link_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The AWS::NetworkManager::LinkAssociation type associates a link to a device. The device and link must be in the same global network and the same site.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: The ID of the device
        :param pulumi.Input[str] global_network_id: The ID of the global network.
        :param pulumi.Input[str] link_id: The ID of the link
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LinkAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::NetworkManager::LinkAssociation type associates a link to a device. The device and link must be in the same global network and the same site.

        :param str resource_name: The name of the resource.
        :param LinkAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LinkAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 global_network_id: Optional[pulumi.Input[str]] = None,
                 link_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = LinkAssociationArgs.__new__(LinkAssociationArgs)

            if device_id is None and not opts.urn:
                raise TypeError("Missing required property 'device_id'")
            __props__.__dict__["device_id"] = device_id
            if global_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'global_network_id'")
            __props__.__dict__["global_network_id"] = global_network_id
            if link_id is None and not opts.urn:
                raise TypeError("Missing required property 'link_id'")
            __props__.__dict__["link_id"] = link_id
        super(LinkAssociation, __self__).__init__(
            'aws-native:networkmanager:LinkAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'LinkAssociation':
        """
        Get an existing LinkAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = LinkAssociationArgs.__new__(LinkAssociationArgs)

        __props__.__dict__["device_id"] = None
        __props__.__dict__["global_network_id"] = None
        __props__.__dict__["link_id"] = None
        return LinkAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        The ID of the device
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="globalNetworkId")
    def global_network_id(self) -> pulumi.Output[str]:
        """
        The ID of the global network.
        """
        return pulumi.get(self, "global_network_id")

    @property
    @pulumi.getter(name="linkId")
    def link_id(self) -> pulumi.Output[str]:
        """
        The ID of the link
        """
        return pulumi.get(self, "link_id")
