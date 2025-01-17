# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetContactChannelResult',
    'AwaitableGetContactChannelResult',
    'get_contact_channel',
    'get_contact_channel_output',
]

@pulumi.output_type
class GetContactChannelResult:
    def __init__(__self__, arn=None, channel_address=None, channel_name=None, defer_activation=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if channel_address and not isinstance(channel_address, str):
            raise TypeError("Expected argument 'channel_address' to be a str")
        pulumi.set(__self__, "channel_address", channel_address)
        if channel_name and not isinstance(channel_name, str):
            raise TypeError("Expected argument 'channel_name' to be a str")
        pulumi.set(__self__, "channel_name", channel_name)
        if defer_activation and not isinstance(defer_activation, bool):
            raise TypeError("Expected argument 'defer_activation' to be a bool")
        pulumi.set(__self__, "defer_activation", defer_activation)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the engagement to a contact channel.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="channelAddress")
    def channel_address(self) -> Optional[str]:
        """
        The details that SSM Incident Manager uses when trying to engage the contact channel.
        """
        return pulumi.get(self, "channel_address")

    @property
    @pulumi.getter(name="channelName")
    def channel_name(self) -> Optional[str]:
        """
        The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
        """
        return pulumi.get(self, "channel_name")

    @property
    @pulumi.getter(name="deferActivation")
    def defer_activation(self) -> Optional[bool]:
        """
        If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.
        """
        return pulumi.get(self, "defer_activation")


class AwaitableGetContactChannelResult(GetContactChannelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContactChannelResult(
            arn=self.arn,
            channel_address=self.channel_address,
            channel_name=self.channel_name,
            defer_activation=self.defer_activation)


def get_contact_channel(arn: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContactChannelResult:
    """
    Resource Type definition for AWS::SSMContacts::ContactChannel


    :param str arn: The Amazon Resource Name (ARN) of the engagement to a contact channel.
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ssmcontacts:getContactChannel', __args__, opts=opts, typ=GetContactChannelResult).value

    return AwaitableGetContactChannelResult(
        arn=__ret__.arn,
        channel_address=__ret__.channel_address,
        channel_name=__ret__.channel_name,
        defer_activation=__ret__.defer_activation)


@_utilities.lift_output_func(get_contact_channel)
def get_contact_channel_output(arn: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContactChannelResult]:
    """
    Resource Type definition for AWS::SSMContacts::ContactChannel


    :param str arn: The Amazon Resource Name (ARN) of the engagement to a contact channel.
    """
    ...
