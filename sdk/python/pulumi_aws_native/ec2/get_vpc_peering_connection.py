# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetVPCPeeringConnectionResult',
    'AwaitableGetVPCPeeringConnectionResult',
    'get_vpc_peering_connection',
    'get_vpc_peering_connection_output',
]

@pulumi.output_type
class GetVPCPeeringConnectionResult:
    def __init__(__self__, id=None, tags=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.VPCPeeringConnectionTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetVPCPeeringConnectionResult(GetVPCPeeringConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVPCPeeringConnectionResult(
            id=self.id,
            tags=self.tags)


def get_vpc_peering_connection(id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVPCPeeringConnectionResult:
    """
    Resource Type definition for AWS::EC2::VPCPeeringConnection
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getVPCPeeringConnection', __args__, opts=opts, typ=GetVPCPeeringConnectionResult).value

    return AwaitableGetVPCPeeringConnectionResult(
        id=__ret__.id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_vpc_peering_connection)
def get_vpc_peering_connection_output(id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVPCPeeringConnectionResult]:
    """
    Resource Type definition for AWS::EC2::VPCPeeringConnection
    """
    ...
