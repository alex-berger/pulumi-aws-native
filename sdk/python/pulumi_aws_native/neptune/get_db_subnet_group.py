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
    'GetDBSubnetGroupResult',
    'AwaitableGetDBSubnetGroupResult',
    'get_db_subnet_group',
    'get_db_subnet_group_output',
]

@pulumi.output_type
class GetDBSubnetGroupResult:
    def __init__(__self__, d_b_subnet_group_description=None, id=None, subnet_ids=None, tags=None):
        if d_b_subnet_group_description and not isinstance(d_b_subnet_group_description, str):
            raise TypeError("Expected argument 'd_b_subnet_group_description' to be a str")
        pulumi.set(__self__, "d_b_subnet_group_description", d_b_subnet_group_description)
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
    @pulumi.getter(name="dBSubnetGroupDescription")
    def d_b_subnet_group_description(self) -> Optional[str]:
        return pulumi.get(self, "d_b_subnet_group_description")

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
    def tags(self) -> Optional[Sequence['outputs.DBSubnetGroupTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetDBSubnetGroupResult(GetDBSubnetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDBSubnetGroupResult(
            d_b_subnet_group_description=self.d_b_subnet_group_description,
            id=self.id,
            subnet_ids=self.subnet_ids,
            tags=self.tags)


def get_db_subnet_group(id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDBSubnetGroupResult:
    """
    Resource Type definition for AWS::Neptune::DBSubnetGroup
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:neptune:getDBSubnetGroup', __args__, opts=opts, typ=GetDBSubnetGroupResult).value

    return AwaitableGetDBSubnetGroupResult(
        d_b_subnet_group_description=__ret__.d_b_subnet_group_description,
        id=__ret__.id,
        subnet_ids=__ret__.subnet_ids,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_db_subnet_group)
def get_db_subnet_group_output(id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDBSubnetGroupResult]:
    """
    Resource Type definition for AWS::Neptune::DBSubnetGroup
    """
    ...
