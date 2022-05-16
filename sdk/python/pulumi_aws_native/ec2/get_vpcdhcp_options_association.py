# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetVPCDHCPOptionsAssociationResult',
    'AwaitableGetVPCDHCPOptionsAssociationResult',
    'get_vpcdhcp_options_association',
    'get_vpcdhcp_options_association_output',
]

@pulumi.output_type
class GetVPCDHCPOptionsAssociationResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the VPC DHCPOptions Association.
        """
        return pulumi.get(self, "id")


class AwaitableGetVPCDHCPOptionsAssociationResult(GetVPCDHCPOptionsAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVPCDHCPOptionsAssociationResult(
            id=self.id)


def get_vpcdhcp_options_association(dhcp_options_id: Optional[str] = None,
                                    vpc_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVPCDHCPOptionsAssociationResult:
    """
    Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.


    :param str dhcp_options_id: The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
    :param str vpc_id: The ID of the VPC.
    """
    __args__ = dict()
    __args__['dhcpOptionsId'] = dhcp_options_id
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getVPCDHCPOptionsAssociation', __args__, opts=opts, typ=GetVPCDHCPOptionsAssociationResult).value

    return AwaitableGetVPCDHCPOptionsAssociationResult(
        id=__ret__.id)


@_utilities.lift_output_func(get_vpcdhcp_options_association)
def get_vpcdhcp_options_association_output(dhcp_options_id: Optional[pulumi.Input[str]] = None,
                                           vpc_id: Optional[pulumi.Input[str]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVPCDHCPOptionsAssociationResult]:
    """
    Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.


    :param str dhcp_options_id: The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
    :param str vpc_id: The ID of the VPC.
    """
    ...