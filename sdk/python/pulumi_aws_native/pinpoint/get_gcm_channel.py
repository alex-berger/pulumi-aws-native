# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetGCMChannelResult',
    'AwaitableGetGCMChannelResult',
    'get_gcm_channel',
    'get_gcm_channel_output',
]

@pulumi.output_type
class GetGCMChannelResult:
    def __init__(__self__, api_key=None, enabled=None, id=None):
        if api_key and not isinstance(api_key, str):
            raise TypeError("Expected argument 'api_key' to be a str")
        pulumi.set(__self__, "api_key", api_key)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> Optional[str]:
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetGCMChannelResult(GetGCMChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGCMChannelResult(
            api_key=self.api_key,
            enabled=self.enabled,
            id=self.id)


def get_gcm_channel(id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGCMChannelResult:
    """
    Resource Type definition for AWS::Pinpoint::GCMChannel
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:pinpoint:getGCMChannel', __args__, opts=opts, typ=GetGCMChannelResult).value

    return AwaitableGetGCMChannelResult(
        api_key=__ret__.api_key,
        enabled=__ret__.enabled,
        id=__ret__.id)


@_utilities.lift_output_func(get_gcm_channel)
def get_gcm_channel_output(id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGCMChannelResult]:
    """
    Resource Type definition for AWS::Pinpoint::GCMChannel
    """
    ...