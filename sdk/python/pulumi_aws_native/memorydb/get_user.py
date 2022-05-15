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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    def __init__(__self__, arn=None, status=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the user account.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Indicates the user status. Can be "active", "modifying" or "deleting".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.UserTag']]:
        """
        An array of key-value pairs to apply to this user.
        """
        return pulumi.get(self, "tags")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            arn=self.arn,
            status=self.status,
            tags=self.tags)


def get_user(user_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Resource Type definition for AWS::MemoryDB::User


    :param str user_name: The name of the user.
    """
    __args__ = dict()
    __args__['userName'] = user_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:memorydb:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        arn=__ret__.arn,
        status=__ret__.status,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_user)
def get_user_output(user_name: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Resource Type definition for AWS::MemoryDB::User


    :param str user_name: The name of the user.
    """
    ...
