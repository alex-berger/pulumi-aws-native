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
    'GetBotResult',
    'AwaitableGetBotResult',
    'get_bot',
    'get_bot_output',
]

@pulumi.output_type
class GetBotResult:
    def __init__(__self__, arn=None, data_privacy=None, description=None, id=None, idle_session_ttl_in_seconds=None, name=None, role_arn=None, test_bot_alias_settings=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if data_privacy and not isinstance(data_privacy, dict):
            raise TypeError("Expected argument 'data_privacy' to be a dict")
        pulumi.set(__self__, "data_privacy", data_privacy)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idle_session_ttl_in_seconds and not isinstance(idle_session_ttl_in_seconds, int):
            raise TypeError("Expected argument 'idle_session_ttl_in_seconds' to be a int")
        pulumi.set(__self__, "idle_session_ttl_in_seconds", idle_session_ttl_in_seconds)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if test_bot_alias_settings and not isinstance(test_bot_alias_settings, dict):
            raise TypeError("Expected argument 'test_bot_alias_settings' to be a dict")
        pulumi.set(__self__, "test_bot_alias_settings", test_bot_alias_settings)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dataPrivacy")
    def data_privacy(self) -> Optional['outputs.DataPrivacyProperties']:
        """
        Data privacy setting of the Bot.
        """
        return pulumi.get(self, "data_privacy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idleSessionTTLInSeconds")
    def idle_session_ttl_in_seconds(self) -> Optional[int]:
        """
        IdleSessionTTLInSeconds of the resource
        """
        return pulumi.get(self, "idle_session_ttl_in_seconds")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="testBotAliasSettings")
    def test_bot_alias_settings(self) -> Optional['outputs.BotTestBotAliasSettings']:
        return pulumi.get(self, "test_bot_alias_settings")


class AwaitableGetBotResult(GetBotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBotResult(
            arn=self.arn,
            data_privacy=self.data_privacy,
            description=self.description,
            id=self.id,
            idle_session_ttl_in_seconds=self.idle_session_ttl_in_seconds,
            name=self.name,
            role_arn=self.role_arn,
            test_bot_alias_settings=self.test_bot_alias_settings)


def get_bot(id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBotResult:
    """
    Amazon Lex conversational bot performing automated tasks such as ordering a pizza, booking a hotel, and so on.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:lex:getBot', __args__, opts=opts, typ=GetBotResult).value

    return AwaitableGetBotResult(
        arn=__ret__.arn,
        data_privacy=__ret__.data_privacy,
        description=__ret__.description,
        id=__ret__.id,
        idle_session_ttl_in_seconds=__ret__.idle_session_ttl_in_seconds,
        name=__ret__.name,
        role_arn=__ret__.role_arn,
        test_bot_alias_settings=__ret__.test_bot_alias_settings)


@_utilities.lift_output_func(get_bot)
def get_bot_output(id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBotResult]:
    """
    Amazon Lex conversational bot performing automated tasks such as ordering a pizza, booking a hotel, and so on.
    """
    ...
