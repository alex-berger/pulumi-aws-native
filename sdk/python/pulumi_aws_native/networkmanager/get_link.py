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
    'GetLinkResult',
    'AwaitableGetLinkResult',
    'get_link',
    'get_link_output',
]

@pulumi.output_type
class GetLinkResult:
    def __init__(__self__, bandwidth=None, description=None, link_arn=None, link_id=None, provider=None, tags=None, type=None):
        if bandwidth and not isinstance(bandwidth, dict):
            raise TypeError("Expected argument 'bandwidth' to be a dict")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if link_arn and not isinstance(link_arn, str):
            raise TypeError("Expected argument 'link_arn' to be a str")
        pulumi.set(__self__, "link_arn", link_arn)
        if link_id and not isinstance(link_id, str):
            raise TypeError("Expected argument 'link_id' to be a str")
        pulumi.set(__self__, "link_id", link_id)
        if provider and not isinstance(provider, str):
            raise TypeError("Expected argument 'provider' to be a str")
        pulumi.set(__self__, "provider", provider)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional['outputs.LinkBandwidth']:
        """
        The Bandwidth for the link.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the link.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="linkArn")
    def link_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the link.
        """
        return pulumi.get(self, "link_arn")

    @property
    @pulumi.getter(name="linkId")
    def link_id(self) -> Optional[str]:
        """
        The ID of the link.
        """
        return pulumi.get(self, "link_id")

    @property
    @pulumi.getter
    def provider(self) -> Optional[str]:
        """
        The provider of the link.
        """
        return pulumi.get(self, "provider")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.LinkTag']]:
        """
        The tags for the link.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the link.
        """
        return pulumi.get(self, "type")


class AwaitableGetLinkResult(GetLinkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLinkResult(
            bandwidth=self.bandwidth,
            description=self.description,
            link_arn=self.link_arn,
            link_id=self.link_id,
            provider=self.provider,
            tags=self.tags,
            type=self.type)


def get_link(global_network_id: Optional[str] = None,
             link_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLinkResult:
    """
    The AWS::NetworkManager::Link type describes a link.


    :param str global_network_id: The ID of the global network.
    :param str link_id: The ID of the link.
    """
    __args__ = dict()
    __args__['globalNetworkId'] = global_network_id
    __args__['linkId'] = link_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:networkmanager:getLink', __args__, opts=opts, typ=GetLinkResult).value

    return AwaitableGetLinkResult(
        bandwidth=__ret__.bandwidth,
        description=__ret__.description,
        link_arn=__ret__.link_arn,
        link_id=__ret__.link_id,
        provider=__ret__.provider,
        tags=__ret__.tags,
        type=__ret__.type)


@_utilities.lift_output_func(get_link)
def get_link_output(global_network_id: Optional[pulumi.Input[str]] = None,
                    link_id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLinkResult]:
    """
    The AWS::NetworkManager::Link type describes a link.


    :param str global_network_id: The ID of the global network.
    :param str link_id: The ID of the link.
    """
    ...