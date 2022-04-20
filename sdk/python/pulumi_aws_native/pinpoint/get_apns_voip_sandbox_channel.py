# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAPNSVoipSandboxChannelResult',
    'AwaitableGetAPNSVoipSandboxChannelResult',
    'get_apns_voip_sandbox_channel',
    'get_apns_voip_sandbox_channel_output',
]

@pulumi.output_type
class GetAPNSVoipSandboxChannelResult:
    def __init__(__self__, bundle_id=None, certificate=None, default_authentication_method=None, enabled=None, id=None, private_key=None, team_id=None, token_key=None, token_key_id=None):
        if bundle_id and not isinstance(bundle_id, str):
            raise TypeError("Expected argument 'bundle_id' to be a str")
        pulumi.set(__self__, "bundle_id", bundle_id)
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        pulumi.set(__self__, "certificate", certificate)
        if default_authentication_method and not isinstance(default_authentication_method, str):
            raise TypeError("Expected argument 'default_authentication_method' to be a str")
        pulumi.set(__self__, "default_authentication_method", default_authentication_method)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if private_key and not isinstance(private_key, str):
            raise TypeError("Expected argument 'private_key' to be a str")
        pulumi.set(__self__, "private_key", private_key)
        if team_id and not isinstance(team_id, str):
            raise TypeError("Expected argument 'team_id' to be a str")
        pulumi.set(__self__, "team_id", team_id)
        if token_key and not isinstance(token_key, str):
            raise TypeError("Expected argument 'token_key' to be a str")
        pulumi.set(__self__, "token_key", token_key)
        if token_key_id and not isinstance(token_key_id, str):
            raise TypeError("Expected argument 'token_key_id' to be a str")
        pulumi.set(__self__, "token_key_id", token_key_id)

    @property
    @pulumi.getter(name="bundleId")
    def bundle_id(self) -> Optional[str]:
        return pulumi.get(self, "bundle_id")

    @property
    @pulumi.getter
    def certificate(self) -> Optional[str]:
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="defaultAuthenticationMethod")
    def default_authentication_method(self) -> Optional[str]:
        return pulumi.get(self, "default_authentication_method")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[str]:
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[str]:
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter(name="tokenKey")
    def token_key(self) -> Optional[str]:
        return pulumi.get(self, "token_key")

    @property
    @pulumi.getter(name="tokenKeyId")
    def token_key_id(self) -> Optional[str]:
        return pulumi.get(self, "token_key_id")


class AwaitableGetAPNSVoipSandboxChannelResult(GetAPNSVoipSandboxChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAPNSVoipSandboxChannelResult(
            bundle_id=self.bundle_id,
            certificate=self.certificate,
            default_authentication_method=self.default_authentication_method,
            enabled=self.enabled,
            id=self.id,
            private_key=self.private_key,
            team_id=self.team_id,
            token_key=self.token_key,
            token_key_id=self.token_key_id)


def get_apns_voip_sandbox_channel(id: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAPNSVoipSandboxChannelResult:
    """
    Resource Type definition for AWS::Pinpoint::APNSVoipSandboxChannel
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:pinpoint:getAPNSVoipSandboxChannel', __args__, opts=opts, typ=GetAPNSVoipSandboxChannelResult).value

    return AwaitableGetAPNSVoipSandboxChannelResult(
        bundle_id=__ret__.bundle_id,
        certificate=__ret__.certificate,
        default_authentication_method=__ret__.default_authentication_method,
        enabled=__ret__.enabled,
        id=__ret__.id,
        private_key=__ret__.private_key,
        team_id=__ret__.team_id,
        token_key=__ret__.token_key,
        token_key_id=__ret__.token_key_id)


@_utilities.lift_output_func(get_apns_voip_sandbox_channel)
def get_apns_voip_sandbox_channel_output(id: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAPNSVoipSandboxChannelResult]:
    """
    Resource Type definition for AWS::Pinpoint::APNSVoipSandboxChannel
    """
    ...
