# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetTrackerResult',
    'AwaitableGetTrackerResult',
    'get_tracker',
    'get_tracker_output',
]

@pulumi.output_type
class GetTrackerResult:
    def __init__(__self__, arn=None, create_time=None, tracker_arn=None, update_time=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if tracker_arn and not isinstance(tracker_arn, str):
            raise TypeError("Expected argument 'tracker_arn' to be a str")
        pulumi.set(__self__, "tracker_arn", tracker_arn)
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
    @pulumi.getter(name="trackerArn")
    def tracker_arn(self) -> Optional[str]:
        return pulumi.get(self, "tracker_arn")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[str]:
        return pulumi.get(self, "update_time")


class AwaitableGetTrackerResult(GetTrackerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrackerResult(
            arn=self.arn,
            create_time=self.create_time,
            tracker_arn=self.tracker_arn,
            update_time=self.update_time)


def get_tracker(tracker_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrackerResult:
    """
    Definition of AWS::Location::Tracker Resource Type
    """
    __args__ = dict()
    __args__['trackerName'] = tracker_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:location:getTracker', __args__, opts=opts, typ=GetTrackerResult).value

    return AwaitableGetTrackerResult(
        arn=__ret__.arn,
        create_time=__ret__.create_time,
        tracker_arn=__ret__.tracker_arn,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_tracker)
def get_tracker_output(tracker_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTrackerResult]:
    """
    Definition of AWS::Location::Tracker Resource Type
    """
    ...
