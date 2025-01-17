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
    'GetTrafficMirrorTargetResult',
    'AwaitableGetTrafficMirrorTargetResult',
    'get_traffic_mirror_target',
    'get_traffic_mirror_target_output',
]

@pulumi.output_type
class GetTrafficMirrorTargetResult:
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
    def tags(self) -> Optional[Sequence['outputs.TrafficMirrorTargetTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetTrafficMirrorTargetResult(GetTrafficMirrorTargetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrafficMirrorTargetResult(
            id=self.id,
            tags=self.tags)


def get_traffic_mirror_target(id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrafficMirrorTargetResult:
    """
    Resource Type definition for AWS::EC2::TrafficMirrorTarget
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getTrafficMirrorTarget', __args__, opts=opts, typ=GetTrafficMirrorTargetResult).value

    return AwaitableGetTrafficMirrorTargetResult(
        id=__ret__.id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_traffic_mirror_target)
def get_traffic_mirror_target_output(id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTrafficMirrorTargetResult]:
    """
    Resource Type definition for AWS::EC2::TrafficMirrorTarget
    """
    ...
