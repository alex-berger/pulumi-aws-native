# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetMemberResult',
    'AwaitableGetMemberResult',
    'get_member',
    'get_member_output',
]

@pulumi.output_type
class GetMemberResult:
    def __init__(__self__, disable_email_notification=None, message=None, status=None):
        if disable_email_notification and not isinstance(disable_email_notification, bool):
            raise TypeError("Expected argument 'disable_email_notification' to be a bool")
        pulumi.set(__self__, "disable_email_notification", disable_email_notification)
        if message and not isinstance(message, str):
            raise TypeError("Expected argument 'message' to be a str")
        pulumi.set(__self__, "message", message)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="disableEmailNotification")
    def disable_email_notification(self) -> Optional[bool]:
        return pulumi.get(self, "disable_email_notification")

    @property
    @pulumi.getter
    def message(self) -> Optional[str]:
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


class AwaitableGetMemberResult(GetMemberResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMemberResult(
            disable_email_notification=self.disable_email_notification,
            message=self.message,
            status=self.status)


def get_member(member_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMemberResult:
    """
    Resource Type definition for AWS::GuardDuty::Member
    """
    __args__ = dict()
    __args__['memberId'] = member_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:guardduty:getMember', __args__, opts=opts, typ=GetMemberResult).value

    return AwaitableGetMemberResult(
        disable_email_notification=__ret__.disable_email_notification,
        message=__ret__.message,
        status=__ret__.status)


@_utilities.lift_output_func(get_member)
def get_member_output(member_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMemberResult]:
    """
    Resource Type definition for AWS::GuardDuty::Member
    """
    ...