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
    def __init__(__self__, arn=None, groups=None, id=None, login_profile=None, managed_policy_arns=None, path=None, permissions_boundary=None, policies=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if login_profile and not isinstance(login_profile, dict):
            raise TypeError("Expected argument 'login_profile' to be a dict")
        pulumi.set(__self__, "login_profile", login_profile)
        if managed_policy_arns and not isinstance(managed_policy_arns, list):
            raise TypeError("Expected argument 'managed_policy_arns' to be a list")
        pulumi.set(__self__, "managed_policy_arns", managed_policy_arns)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if permissions_boundary and not isinstance(permissions_boundary, str):
            raise TypeError("Expected argument 'permissions_boundary' to be a str")
        pulumi.set(__self__, "permissions_boundary", permissions_boundary)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loginProfile")
    def login_profile(self) -> Optional['outputs.UserLoginProfile']:
        return pulumi.get(self, "login_profile")

    @property
    @pulumi.getter(name="managedPolicyArns")
    def managed_policy_arns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "managed_policy_arns")

    @property
    @pulumi.getter
    def path(self) -> Optional[str]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="permissionsBoundary")
    def permissions_boundary(self) -> Optional[str]:
        return pulumi.get(self, "permissions_boundary")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence['outputs.UserPolicy']]:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.UserTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            arn=self.arn,
            groups=self.groups,
            id=self.id,
            login_profile=self.login_profile,
            managed_policy_arns=self.managed_policy_arns,
            path=self.path,
            permissions_boundary=self.permissions_boundary,
            policies=self.policies,
            tags=self.tags)


def get_user(id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Resource Type definition for AWS::IAM::User
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iam:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        arn=__ret__.arn,
        groups=__ret__.groups,
        id=__ret__.id,
        login_profile=__ret__.login_profile,
        managed_policy_arns=__ret__.managed_policy_arns,
        path=__ret__.path,
        permissions_boundary=__ret__.permissions_boundary,
        policies=__ret__.policies,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_user)
def get_user_output(id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Resource Type definition for AWS::IAM::User
    """
    ...
