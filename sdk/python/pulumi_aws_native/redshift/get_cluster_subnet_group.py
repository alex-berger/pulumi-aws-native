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
    'GetClusterSubnetGroupResult',
    'AwaitableGetClusterSubnetGroupResult',
    'get_cluster_subnet_group',
    'get_cluster_subnet_group_output',
]

@pulumi.output_type
class GetClusterSubnetGroupResult:
    def __init__(__self__, description=None, id=None, subnet_ids=None, tags=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ClusterSubnetGroupTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetClusterSubnetGroupResult(GetClusterSubnetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterSubnetGroupResult(
            description=self.description,
            id=self.id,
            subnet_ids=self.subnet_ids,
            tags=self.tags)


def get_cluster_subnet_group(id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterSubnetGroupResult:
    """
    Resource Type definition for AWS::Redshift::ClusterSubnetGroup
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:redshift:getClusterSubnetGroup', __args__, opts=opts, typ=GetClusterSubnetGroupResult).value

    return AwaitableGetClusterSubnetGroupResult(
        description=__ret__.description,
        id=__ret__.id,
        subnet_ids=__ret__.subnet_ids,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_cluster_subnet_group)
def get_cluster_subnet_group_output(id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClusterSubnetGroupResult]:
    """
    Resource Type definition for AWS::Redshift::ClusterSubnetGroup
    """
    ...
