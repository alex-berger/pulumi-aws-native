# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClientVpnTargetNetworkAssociationArgs', 'ClientVpnTargetNetworkAssociation']

@pulumi.input_type
class ClientVpnTargetNetworkAssociationArgs:
    def __init__(__self__, *,
                 client_vpn_endpoint_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ClientVpnTargetNetworkAssociation resource.
        """
        pulumi.set(__self__, "client_vpn_endpoint_id", client_vpn_endpoint_id)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="clientVpnEndpointId")
    def client_vpn_endpoint_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "client_vpn_endpoint_id")

    @client_vpn_endpoint_id.setter
    def client_vpn_endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_vpn_endpoint_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)


warnings.warn("""ClientVpnTargetNetworkAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class ClientVpnTargetNetworkAssociation(pulumi.CustomResource):
    warnings.warn("""ClientVpnTargetNetworkAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_vpn_endpoint_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EC2::ClientVpnTargetNetworkAssociation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClientVpnTargetNetworkAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EC2::ClientVpnTargetNetworkAssociation

        :param str resource_name: The name of the resource.
        :param ClientVpnTargetNetworkAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientVpnTargetNetworkAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_vpn_endpoint_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""ClientVpnTargetNetworkAssociation is deprecated: ClientVpnTargetNetworkAssociation is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientVpnTargetNetworkAssociationArgs.__new__(ClientVpnTargetNetworkAssociationArgs)

            if client_vpn_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_vpn_endpoint_id'")
            __props__.__dict__["client_vpn_endpoint_id"] = client_vpn_endpoint_id
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
        super(ClientVpnTargetNetworkAssociation, __self__).__init__(
            'aws-native:ec2:ClientVpnTargetNetworkAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ClientVpnTargetNetworkAssociation':
        """
        Get an existing ClientVpnTargetNetworkAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ClientVpnTargetNetworkAssociationArgs.__new__(ClientVpnTargetNetworkAssociationArgs)

        __props__.__dict__["client_vpn_endpoint_id"] = None
        __props__.__dict__["subnet_id"] = None
        return ClientVpnTargetNetworkAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientVpnEndpointId")
    def client_vpn_endpoint_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_vpn_endpoint_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subnet_id")

