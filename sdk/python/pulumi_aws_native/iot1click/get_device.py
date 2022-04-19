# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDeviceResult',
    'AwaitableGetDeviceResult',
    'get_device',
    'get_device_output',
]

@pulumi.output_type
class GetDeviceResult:
    def __init__(__self__, arn=None, enabled=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")


class AwaitableGetDeviceResult(GetDeviceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceResult(
            arn=self.arn,
            enabled=self.enabled)


def get_device(device_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceResult:
    """
    Resource Type definition for AWS::IoT1Click::Device
    """
    __args__ = dict()
    __args__['deviceId'] = device_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iot1click:getDevice', __args__, opts=opts, typ=GetDeviceResult).value

    return AwaitableGetDeviceResult(
        arn=__ret__.arn,
        enabled=__ret__.enabled)


@_utilities.lift_output_func(get_device)
def get_device_output(device_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeviceResult]:
    """
    Resource Type definition for AWS::IoT1Click::Device
    """
    ...
