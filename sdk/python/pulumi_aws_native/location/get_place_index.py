# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetPlaceIndexResult',
    'AwaitableGetPlaceIndexResult',
    'get_place_index',
    'get_place_index_output',
]

@pulumi.output_type
class GetPlaceIndexResult:
    def __init__(__self__, arn=None, create_time=None, index_arn=None, update_time=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if index_arn and not isinstance(index_arn, str):
            raise TypeError("Expected argument 'index_arn' to be a str")
        pulumi.set(__self__, "index_arn", index_arn)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="indexArn")
    def index_arn(self) -> Optional[str]:
        return pulumi.get(self, "index_arn")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[str]:
        return pulumi.get(self, "update_time")


class AwaitableGetPlaceIndexResult(GetPlaceIndexResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlaceIndexResult(
            arn=self.arn,
            create_time=self.create_time,
            index_arn=self.index_arn,
            update_time=self.update_time)


def get_place_index(index_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlaceIndexResult:
    """
    Definition of AWS::Location::PlaceIndex Resource Type
    """
    __args__ = dict()
    __args__['indexName'] = index_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:location:getPlaceIndex', __args__, opts=opts, typ=GetPlaceIndexResult).value

    return AwaitableGetPlaceIndexResult(
        arn=__ret__.arn,
        create_time=__ret__.create_time,
        index_arn=__ret__.index_arn,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_place_index)
def get_place_index_output(index_name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPlaceIndexResult]:
    """
    Definition of AWS::Location::PlaceIndex Resource Type
    """
    ...
