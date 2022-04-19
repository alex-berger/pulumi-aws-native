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

__all__ = ['VPCPeeringConnectionArgs', 'VPCPeeringConnection']

@pulumi.input_type
class VPCPeeringConnectionArgs:
    def __init__(__self__, *,
                 peer_vpc_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 peer_owner_id: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 peer_role_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['VPCPeeringConnectionTagArgs']]]] = None):
        """
        The set of arguments for constructing a VPCPeeringConnection resource.
        """
        pulumi.set(__self__, "peer_vpc_id", peer_vpc_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if peer_owner_id is not None:
            pulumi.set(__self__, "peer_owner_id", peer_owner_id)
        if peer_region is not None:
            pulumi.set(__self__, "peer_region", peer_region)
        if peer_role_arn is not None:
            pulumi.set(__self__, "peer_role_arn", peer_role_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "peer_vpc_id")

    @peer_vpc_id.setter
    def peer_vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_vpc_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="peerOwnerId")
    def peer_owner_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peer_owner_id")

    @peer_owner_id.setter
    def peer_owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_owner_id", value)

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peer_region")

    @peer_region.setter
    def peer_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_region", value)

    @property
    @pulumi.getter(name="peerRoleArn")
    def peer_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peer_role_arn")

    @peer_role_arn.setter
    def peer_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_role_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VPCPeeringConnectionTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VPCPeeringConnectionTagArgs']]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""VPCPeeringConnection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class VPCPeeringConnection(pulumi.CustomResource):
    warnings.warn("""VPCPeeringConnection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peer_owner_id: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 peer_role_arn: Optional[pulumi.Input[str]] = None,
                 peer_vpc_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VPCPeeringConnectionTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EC2::VPCPeeringConnection

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VPCPeeringConnectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EC2::VPCPeeringConnection

        :param str resource_name: The name of the resource.
        :param VPCPeeringConnectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VPCPeeringConnectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peer_owner_id: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 peer_role_arn: Optional[pulumi.Input[str]] = None,
                 peer_vpc_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VPCPeeringConnectionTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""VPCPeeringConnection is deprecated: VPCPeeringConnection is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VPCPeeringConnectionArgs.__new__(VPCPeeringConnectionArgs)

            __props__.__dict__["peer_owner_id"] = peer_owner_id
            __props__.__dict__["peer_region"] = peer_region
            __props__.__dict__["peer_role_arn"] = peer_role_arn
            if peer_vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'peer_vpc_id'")
            __props__.__dict__["peer_vpc_id"] = peer_vpc_id
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(VPCPeeringConnection, __self__).__init__(
            'aws-native:ec2:VPCPeeringConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VPCPeeringConnection':
        """
        Get an existing VPCPeeringConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = VPCPeeringConnectionArgs.__new__(VPCPeeringConnectionArgs)

        __props__.__dict__["peer_owner_id"] = None
        __props__.__dict__["peer_region"] = None
        __props__.__dict__["peer_role_arn"] = None
        __props__.__dict__["peer_vpc_id"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["vpc_id"] = None
        return VPCPeeringConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="peerOwnerId")
    def peer_owner_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "peer_owner_id")

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "peer_region")

    @property
    @pulumi.getter(name="peerRoleArn")
    def peer_role_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "peer_role_arn")

    @property
    @pulumi.getter(name="peerVpcId")
    def peer_vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "peer_vpc_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.VPCPeeringConnectionTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

