# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetIdentityPoolRoleAttachmentResult',
    'AwaitableGetIdentityPoolRoleAttachmentResult',
    'get_identity_pool_role_attachment',
    'get_identity_pool_role_attachment_output',
]

@pulumi.output_type
class GetIdentityPoolRoleAttachmentResult:
    def __init__(__self__, id=None, role_mappings=None, roles=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if role_mappings and not isinstance(role_mappings, dict):
            raise TypeError("Expected argument 'role_mappings' to be a dict")
        pulumi.set(__self__, "role_mappings", role_mappings)
        if roles and not isinstance(roles, dict):
            raise TypeError("Expected argument 'roles' to be a dict")
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="roleMappings")
    def role_mappings(self) -> Optional[Any]:
        return pulumi.get(self, "role_mappings")

    @property
    @pulumi.getter
    def roles(self) -> Optional[Any]:
        return pulumi.get(self, "roles")


class AwaitableGetIdentityPoolRoleAttachmentResult(GetIdentityPoolRoleAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdentityPoolRoleAttachmentResult(
            id=self.id,
            role_mappings=self.role_mappings,
            roles=self.roles)


def get_identity_pool_role_attachment(id: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdentityPoolRoleAttachmentResult:
    """
    Resource Type definition for AWS::Cognito::IdentityPoolRoleAttachment
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:cognito:getIdentityPoolRoleAttachment', __args__, opts=opts, typ=GetIdentityPoolRoleAttachmentResult).value

    return AwaitableGetIdentityPoolRoleAttachmentResult(
        id=__ret__.id,
        role_mappings=__ret__.role_mappings,
        roles=__ret__.roles)


@_utilities.lift_output_func(get_identity_pool_role_attachment)
def get_identity_pool_role_attachment_output(id: Optional[pulumi.Input[str]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIdentityPoolRoleAttachmentResult]:
    """
    Resource Type definition for AWS::Cognito::IdentityPoolRoleAttachment
    """
    ...
