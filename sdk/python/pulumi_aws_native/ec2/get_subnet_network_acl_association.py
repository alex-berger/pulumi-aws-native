# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSubnetNetworkAclAssociationResult',
    'AwaitableGetSubnetNetworkAclAssociationResult',
    'get_subnet_network_acl_association',
    'get_subnet_network_acl_association_output',
]

@pulumi.output_type
class GetSubnetNetworkAclAssociationResult:
    def __init__(__self__, association_id=None):
        if association_id and not isinstance(association_id, str):
            raise TypeError("Expected argument 'association_id' to be a str")
        pulumi.set(__self__, "association_id", association_id)

    @property
    @pulumi.getter(name="associationId")
    def association_id(self) -> Optional[str]:
        return pulumi.get(self, "association_id")


class AwaitableGetSubnetNetworkAclAssociationResult(GetSubnetNetworkAclAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetNetworkAclAssociationResult(
            association_id=self.association_id)


def get_subnet_network_acl_association(association_id: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetNetworkAclAssociationResult:
    """
    Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation
    """
    __args__ = dict()
    __args__['associationId'] = association_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getSubnetNetworkAclAssociation', __args__, opts=opts, typ=GetSubnetNetworkAclAssociationResult).value

    return AwaitableGetSubnetNetworkAclAssociationResult(
        association_id=__ret__.association_id)


@_utilities.lift_output_func(get_subnet_network_acl_association)
def get_subnet_network_acl_association_output(association_id: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubnetNetworkAclAssociationResult]:
    """
    Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation
    """
    ...
