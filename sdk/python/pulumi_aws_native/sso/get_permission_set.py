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
    'GetPermissionSetResult',
    'AwaitableGetPermissionSetResult',
    'get_permission_set',
    'get_permission_set_output',
]

@pulumi.output_type
class GetPermissionSetResult:
    def __init__(__self__, description=None, inline_policy=None, managed_policies=None, permission_set_arn=None, relay_state_type=None, session_duration=None, tags=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if inline_policy and not isinstance(inline_policy, dict):
            raise TypeError("Expected argument 'inline_policy' to be a dict")
        pulumi.set(__self__, "inline_policy", inline_policy)
        if managed_policies and not isinstance(managed_policies, list):
            raise TypeError("Expected argument 'managed_policies' to be a list")
        pulumi.set(__self__, "managed_policies", managed_policies)
        if permission_set_arn and not isinstance(permission_set_arn, str):
            raise TypeError("Expected argument 'permission_set_arn' to be a str")
        pulumi.set(__self__, "permission_set_arn", permission_set_arn)
        if relay_state_type and not isinstance(relay_state_type, str):
            raise TypeError("Expected argument 'relay_state_type' to be a str")
        pulumi.set(__self__, "relay_state_type", relay_state_type)
        if session_duration and not isinstance(session_duration, str):
            raise TypeError("Expected argument 'session_duration' to be a str")
        pulumi.set(__self__, "session_duration", session_duration)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The permission set description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="inlinePolicy")
    def inline_policy(self) -> Optional[Any]:
        """
        The inline policy to put in permission set.
        """
        return pulumi.get(self, "inline_policy")

    @property
    @pulumi.getter(name="managedPolicies")
    def managed_policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "managed_policies")

    @property
    @pulumi.getter(name="permissionSetArn")
    def permission_set_arn(self) -> Optional[str]:
        """
        The permission set that the policy will be attached to
        """
        return pulumi.get(self, "permission_set_arn")

    @property
    @pulumi.getter(name="relayStateType")
    def relay_state_type(self) -> Optional[str]:
        """
        The relay state URL that redirect links to any service in the AWS Management Console.
        """
        return pulumi.get(self, "relay_state_type")

    @property
    @pulumi.getter(name="sessionDuration")
    def session_duration(self) -> Optional[str]:
        """
        The length of time that a user can be signed in to an AWS account.
        """
        return pulumi.get(self, "session_duration")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.PermissionSetTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetPermissionSetResult(GetPermissionSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPermissionSetResult(
            description=self.description,
            inline_policy=self.inline_policy,
            managed_policies=self.managed_policies,
            permission_set_arn=self.permission_set_arn,
            relay_state_type=self.relay_state_type,
            session_duration=self.session_duration,
            tags=self.tags)


def get_permission_set(instance_arn: Optional[str] = None,
                       permission_set_arn: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPermissionSetResult:
    """
    Resource Type definition for SSO PermissionSet


    :param str instance_arn: The sso instance arn that the permission set is owned.
    :param str permission_set_arn: The permission set that the policy will be attached to
    """
    __args__ = dict()
    __args__['instanceArn'] = instance_arn
    __args__['permissionSetArn'] = permission_set_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:sso:getPermissionSet', __args__, opts=opts, typ=GetPermissionSetResult).value

    return AwaitableGetPermissionSetResult(
        description=__ret__.description,
        inline_policy=__ret__.inline_policy,
        managed_policies=__ret__.managed_policies,
        permission_set_arn=__ret__.permission_set_arn,
        relay_state_type=__ret__.relay_state_type,
        session_duration=__ret__.session_duration,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_permission_set)
def get_permission_set_output(instance_arn: Optional[pulumi.Input[str]] = None,
                              permission_set_arn: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPermissionSetResult]:
    """
    Resource Type definition for SSO PermissionSet


    :param str instance_arn: The sso instance arn that the permission set is owned.
    :param str permission_set_arn: The permission set that the policy will be attached to
    """
    ...
