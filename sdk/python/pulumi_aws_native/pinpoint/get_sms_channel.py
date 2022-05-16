# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSMSChannelResult',
    'AwaitableGetSMSChannelResult',
    'get_sms_channel',
    'get_sms_channel_output',
]

@pulumi.output_type
class GetSMSChannelResult:
    def __init__(__self__, enabled=None, id=None, sender_id=None, short_code=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if sender_id and not isinstance(sender_id, str):
            raise TypeError("Expected argument 'sender_id' to be a str")
        pulumi.set(__self__, "sender_id", sender_id)
        if short_code and not isinstance(short_code, str):
            raise TypeError("Expected argument 'short_code' to be a str")
        pulumi.set(__self__, "short_code", short_code)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="senderId")
    def sender_id(self) -> Optional[str]:
        return pulumi.get(self, "sender_id")

    @property
    @pulumi.getter(name="shortCode")
    def short_code(self) -> Optional[str]:
        return pulumi.get(self, "short_code")


class AwaitableGetSMSChannelResult(GetSMSChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSMSChannelResult(
            enabled=self.enabled,
            id=self.id,
            sender_id=self.sender_id,
            short_code=self.short_code)


def get_sms_channel(id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSMSChannelResult:
    """
    Resource Type definition for AWS::Pinpoint::SMSChannel
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:pinpoint:getSMSChannel', __args__, opts=opts, typ=GetSMSChannelResult).value

    return AwaitableGetSMSChannelResult(
        enabled=__ret__.enabled,
        id=__ret__.id,
        sender_id=__ret__.sender_id,
        short_code=__ret__.short_code)


@_utilities.lift_output_func(get_sms_channel)
def get_sms_channel_output(id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSMSChannelResult]:
    """
    Resource Type definition for AWS::Pinpoint::SMSChannel
    """
    ...