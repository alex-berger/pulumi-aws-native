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
    'GetRobotResult',
    'AwaitableGetRobotResult',
    'get_robot',
    'get_robot_output',
]

@pulumi.output_type
class GetRobotResult:
    def __init__(__self__, arn=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional['outputs.RobotTags']:
        return pulumi.get(self, "tags")


class AwaitableGetRobotResult(GetRobotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRobotResult(
            arn=self.arn,
            tags=self.tags)


def get_robot(arn: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRobotResult:
    """
    AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:robomaker:getRobot', __args__, opts=opts, typ=GetRobotResult).value

    return AwaitableGetRobotResult(
        arn=__ret__.arn,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_robot)
def get_robot_output(arn: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRobotResult]:
    """
    AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.
    """
    ...
