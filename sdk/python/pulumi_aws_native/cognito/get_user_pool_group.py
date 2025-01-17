# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetUserPoolGroupResult',
    'AwaitableGetUserPoolGroupResult',
    'get_user_pool_group',
    'get_user_pool_group_output',
]

@pulumi.output_type
class GetUserPoolGroupResult:
    def __init__(__self__, description=None, id=None, precedence=None, role_arn=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if precedence and not isinstance(precedence, float):
            raise TypeError("Expected argument 'precedence' to be a float")
        pulumi.set(__self__, "precedence", precedence)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def precedence(self) -> Optional[float]:
        return pulumi.get(self, "precedence")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        return pulumi.get(self, "role_arn")


class AwaitableGetUserPoolGroupResult(GetUserPoolGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserPoolGroupResult(
            description=self.description,
            id=self.id,
            precedence=self.precedence,
            role_arn=self.role_arn)


def get_user_pool_group(id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserPoolGroupResult:
    """
    Resource Type definition for AWS::Cognito::UserPoolGroup
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:cognito:getUserPoolGroup', __args__, opts=opts, typ=GetUserPoolGroupResult).value

    return AwaitableGetUserPoolGroupResult(
        description=__ret__.description,
        id=__ret__.id,
        precedence=__ret__.precedence,
        role_arn=__ret__.role_arn)


@_utilities.lift_output_func(get_user_pool_group)
def get_user_pool_group_output(id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserPoolGroupResult]:
    """
    Resource Type definition for AWS::Cognito::UserPoolGroup
    """
    ...
