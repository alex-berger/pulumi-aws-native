# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FlowVpcInterfaceArgs', 'FlowVpcInterface']

@pulumi.input_type
class FlowVpcInterfaceArgs:
    def __init__(__self__, *,
                 flow_arn: pulumi.Input[str],
                 role_arn: pulumi.Input[str],
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnet_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FlowVpcInterface resource.
        :param pulumi.Input[str] flow_arn: The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        :param pulumi.Input[str] role_arn: Role Arn MediaConnect can assumes to create ENIs in customer's account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Security Group IDs to be used on ENI.
        :param pulumi.Input[str] subnet_id: Subnet must be in the AZ of the Flow
        :param pulumi.Input[str] name: Immutable and has to be a unique against other VpcInterfaces in this Flow.
        """
        pulumi.set(__self__, "flow_arn", flow_arn)
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        """
        return pulumi.get(self, "flow_arn")

    @flow_arn.setter
    def flow_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "flow_arn", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        Role Arn MediaConnect can assumes to create ENIs in customer's account.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Security Group IDs to be used on ENI.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        Subnet must be in the AZ of the Flow
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable and has to be a unique against other VpcInterfaces in this Flow.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class FlowVpcInterface(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flow_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource schema for AWS::MediaConnect::FlowVpcInterface

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] flow_arn: The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        :param pulumi.Input[str] name: Immutable and has to be a unique against other VpcInterfaces in this Flow.
        :param pulumi.Input[str] role_arn: Role Arn MediaConnect can assumes to create ENIs in customer's account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Security Group IDs to be used on ENI.
        :param pulumi.Input[str] subnet_id: Subnet must be in the AZ of the Flow
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowVpcInterfaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::MediaConnect::FlowVpcInterface

        :param str resource_name: The name of the resource.
        :param FlowVpcInterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowVpcInterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flow_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = FlowVpcInterfaceArgs.__new__(FlowVpcInterfaceArgs)

            if flow_arn is None and not opts.urn:
                raise TypeError("Missing required property 'flow_arn'")
            __props__.__dict__["flow_arn"] = flow_arn
            __props__.__dict__["name"] = name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            if security_group_ids is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_ids'")
            __props__.__dict__["security_group_ids"] = security_group_ids
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["network_interface_ids"] = None
        super(FlowVpcInterface, __self__).__init__(
            'aws-native:mediaconnect:FlowVpcInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FlowVpcInterface':
        """
        Get an existing FlowVpcInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FlowVpcInterfaceArgs.__new__(FlowVpcInterfaceArgs)

        __props__.__dict__["flow_arn"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_interface_ids"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["security_group_ids"] = None
        __props__.__dict__["subnet_id"] = None
        return FlowVpcInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
        """
        return pulumi.get(self, "flow_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable and has to be a unique against other VpcInterfaces in this Flow.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaceIds")
    def network_interface_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        IDs of the network interfaces created in customer's account by MediaConnect.
        """
        return pulumi.get(self, "network_interface_ids")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        Role Arn MediaConnect can assumes to create ENIs in customer's account.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Security Group IDs to be used on ENI.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        Subnet must be in the AZ of the Flow
        """
        return pulumi.get(self, "subnet_id")
