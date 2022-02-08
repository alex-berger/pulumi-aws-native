# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetScheduledActionResult',
    'AwaitableGetScheduledActionResult',
    'get_scheduled_action',
    'get_scheduled_action_output',
]

@pulumi.output_type
class GetScheduledActionResult:
    def __init__(__self__, desired_capacity=None, end_time=None, id=None, max_size=None, min_size=None, recurrence=None, start_time=None, time_zone=None):
        if desired_capacity and not isinstance(desired_capacity, int):
            raise TypeError("Expected argument 'desired_capacity' to be a int")
        pulumi.set(__self__, "desired_capacity", desired_capacity)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if max_size and not isinstance(max_size, int):
            raise TypeError("Expected argument 'max_size' to be a int")
        pulumi.set(__self__, "max_size", max_size)
        if min_size and not isinstance(min_size, int):
            raise TypeError("Expected argument 'min_size' to be a int")
        pulumi.set(__self__, "min_size", min_size)
        if recurrence and not isinstance(recurrence, str):
            raise TypeError("Expected argument 'recurrence' to be a str")
        pulumi.set(__self__, "recurrence", recurrence)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if time_zone and not isinstance(time_zone, str):
            raise TypeError("Expected argument 'time_zone' to be a str")
        pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> Optional[int]:
        return pulumi.get(self, "desired_capacity")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[str]:
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[int]:
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> Optional[int]:
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter
    def recurrence(self) -> Optional[str]:
        return pulumi.get(self, "recurrence")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[str]:
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[str]:
        return pulumi.get(self, "time_zone")


class AwaitableGetScheduledActionResult(GetScheduledActionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScheduledActionResult(
            desired_capacity=self.desired_capacity,
            end_time=self.end_time,
            id=self.id,
            max_size=self.max_size,
            min_size=self.min_size,
            recurrence=self.recurrence,
            start_time=self.start_time,
            time_zone=self.time_zone)


def get_scheduled_action(id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScheduledActionResult:
    """
    Resource Type definition for AWS::AutoScaling::ScheduledAction
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:autoscaling:getScheduledAction', __args__, opts=opts, typ=GetScheduledActionResult).value

    return AwaitableGetScheduledActionResult(
        desired_capacity=__ret__.desired_capacity,
        end_time=__ret__.end_time,
        id=__ret__.id,
        max_size=__ret__.max_size,
        min_size=__ret__.min_size,
        recurrence=__ret__.recurrence,
        start_time=__ret__.start_time,
        time_zone=__ret__.time_zone)


@_utilities.lift_output_func(get_scheduled_action)
def get_scheduled_action_output(id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScheduledActionResult]:
    """
    Resource Type definition for AWS::AutoScaling::ScheduledAction
    """
    ...