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
    'GetACLResult',
    'AwaitableGetACLResult',
    'get_acl',
    'get_acl_output',
]

@pulumi.output_type
class GetACLResult:
    def __init__(__self__, arn=None, status=None, tags=None, user_names=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if user_names and not isinstance(user_names, list):
            raise TypeError("Expected argument 'user_names' to be a list")
        pulumi.set(__self__, "user_names", user_names)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the acl.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Indicates acl status. Can be "creating", "active", "modifying", "deleting".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ACLTag']]:
        """
        An array of key-value pairs to apply to this cluster.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="userNames")
    def user_names(self) -> Optional[Sequence[str]]:
        """
        List of users associated to this acl.
        """
        return pulumi.get(self, "user_names")


class AwaitableGetACLResult(GetACLResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetACLResult(
            arn=self.arn,
            status=self.status,
            tags=self.tags,
            user_names=self.user_names)


def get_acl(a_cl_name: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetACLResult:
    """
    Resource Type definition for AWS::MemoryDB::ACL


    :param str a_cl_name: The name of the acl.
    """
    __args__ = dict()
    __args__['aCLName'] = a_cl_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:memorydb:getACL', __args__, opts=opts, typ=GetACLResult).value

    return AwaitableGetACLResult(
        arn=__ret__.arn,
        status=__ret__.status,
        tags=__ret__.tags,
        user_names=__ret__.user_names)


@_utilities.lift_output_func(get_acl)
def get_acl_output(a_cl_name: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetACLResult]:
    """
    Resource Type definition for AWS::MemoryDB::ACL


    :param str a_cl_name: The name of the acl.
    """
    ...
