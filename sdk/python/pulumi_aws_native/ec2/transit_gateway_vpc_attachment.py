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

__all__ = ['TransitGatewayVpcAttachmentArgs', 'TransitGatewayVpcAttachment']

@pulumi.input_type
class TransitGatewayVpcAttachmentArgs:
    def __init__(__self__, *,
                 subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 transit_gateway_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 add_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 options: Optional[pulumi.Input['OptionsPropertiesArgs']] = None,
                 remove_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TransitGatewayVpcAttachmentTagArgs']]]] = None):
        """
        The set of arguments for constructing a TransitGatewayVpcAttachment resource.
        :param pulumi.Input['OptionsPropertiesArgs'] options: The options for the transit gateway vpc attachment.
        """
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "transit_gateway_id", transit_gateway_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if add_subnet_ids is not None:
            pulumi.set(__self__, "add_subnet_ids", add_subnet_ids)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if remove_subnet_ids is not None:
            pulumi.set(__self__, "remove_subnet_ids", remove_subnet_ids)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "transit_gateway_id")

    @transit_gateway_id.setter
    def transit_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_gateway_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="addSubnetIds")
    def add_subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "add_subnet_ids")

    @add_subnet_ids.setter
    def add_subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "add_subnet_ids", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input['OptionsPropertiesArgs']]:
        """
        The options for the transit gateway vpc attachment.
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input['OptionsPropertiesArgs']]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter(name="removeSubnetIds")
    def remove_subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "remove_subnet_ids")

    @remove_subnet_ids.setter
    def remove_subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "remove_subnet_ids", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TransitGatewayVpcAttachmentTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TransitGatewayVpcAttachmentTagArgs']]]]):
        pulumi.set(self, "tags", value)


class TransitGatewayVpcAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 options: Optional[pulumi.Input[pulumi.InputType['OptionsPropertiesArgs']]] = None,
                 remove_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransitGatewayVpcAttachmentTagArgs']]]]] = None,
                 transit_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['OptionsPropertiesArgs']] options: The options for the transit gateway vpc attachment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransitGatewayVpcAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment

        :param str resource_name: The name of the resource.
        :param TransitGatewayVpcAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransitGatewayVpcAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 options: Optional[pulumi.Input[pulumi.InputType['OptionsPropertiesArgs']]] = None,
                 remove_subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransitGatewayVpcAttachmentTagArgs']]]]] = None,
                 transit_gateway_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TransitGatewayVpcAttachmentArgs.__new__(TransitGatewayVpcAttachmentArgs)

            __props__.__dict__["add_subnet_ids"] = add_subnet_ids
            __props__.__dict__["options"] = options
            __props__.__dict__["remove_subnet_ids"] = remove_subnet_ids
            if subnet_ids is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__.__dict__["subnet_ids"] = subnet_ids
            __props__.__dict__["tags"] = tags
            if transit_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_gateway_id'")
            __props__.__dict__["transit_gateway_id"] = transit_gateway_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(TransitGatewayVpcAttachment, __self__).__init__(
            'aws-native:ec2:TransitGatewayVpcAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TransitGatewayVpcAttachment':
        """
        Get an existing TransitGatewayVpcAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TransitGatewayVpcAttachmentArgs.__new__(TransitGatewayVpcAttachmentArgs)

        __props__.__dict__["add_subnet_ids"] = None
        __props__.__dict__["options"] = None
        __props__.__dict__["remove_subnet_ids"] = None
        __props__.__dict__["subnet_ids"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["transit_gateway_id"] = None
        __props__.__dict__["vpc_id"] = None
        return TransitGatewayVpcAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addSubnetIds")
    def add_subnet_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "add_subnet_ids")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional['outputs.OptionsProperties']]:
        """
        The options for the transit gateway vpc attachment.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="removeSubnetIds")
    def remove_subnet_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "remove_subnet_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.TransitGatewayVpcAttachmentTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "transit_gateway_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

