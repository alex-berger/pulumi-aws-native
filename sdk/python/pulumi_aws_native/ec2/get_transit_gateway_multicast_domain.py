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
    'GetTransitGatewayMulticastDomainResult',
    'AwaitableGetTransitGatewayMulticastDomainResult',
    'get_transit_gateway_multicast_domain',
    'get_transit_gateway_multicast_domain_output',
]

@pulumi.output_type
class GetTransitGatewayMulticastDomainResult:
    def __init__(__self__, creation_time=None, options=None, state=None, tags=None, transit_gateway_multicast_domain_arn=None, transit_gateway_multicast_domain_id=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        pulumi.set(__self__, "options", options)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if transit_gateway_multicast_domain_arn and not isinstance(transit_gateway_multicast_domain_arn, str):
            raise TypeError("Expected argument 'transit_gateway_multicast_domain_arn' to be a str")
        pulumi.set(__self__, "transit_gateway_multicast_domain_arn", transit_gateway_multicast_domain_arn)
        if transit_gateway_multicast_domain_id and not isinstance(transit_gateway_multicast_domain_id, str):
            raise TypeError("Expected argument 'transit_gateway_multicast_domain_id' to be a str")
        pulumi.set(__self__, "transit_gateway_multicast_domain_id", transit_gateway_multicast_domain_id)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        """
        The time the transit gateway multicast domain was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def options(self) -> Optional['outputs.OptionsProperties']:
        """
        The options for the transit gateway multicast domain.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        The state of the transit gateway multicast domain.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.TransitGatewayMulticastDomainTag']]:
        """
        The tags for the transit gateway multicast domain.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitGatewayMulticastDomainArn")
    def transit_gateway_multicast_domain_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the transit gateway multicast domain.
        """
        return pulumi.get(self, "transit_gateway_multicast_domain_arn")

    @property
    @pulumi.getter(name="transitGatewayMulticastDomainId")
    def transit_gateway_multicast_domain_id(self) -> Optional[str]:
        """
        The ID of the transit gateway multicast domain.
        """
        return pulumi.get(self, "transit_gateway_multicast_domain_id")


class AwaitableGetTransitGatewayMulticastDomainResult(GetTransitGatewayMulticastDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitGatewayMulticastDomainResult(
            creation_time=self.creation_time,
            options=self.options,
            state=self.state,
            tags=self.tags,
            transit_gateway_multicast_domain_arn=self.transit_gateway_multicast_domain_arn,
            transit_gateway_multicast_domain_id=self.transit_gateway_multicast_domain_id)


def get_transit_gateway_multicast_domain(transit_gateway_multicast_domain_id: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitGatewayMulticastDomainResult:
    """
    The AWS::EC2::TransitGatewayMulticastDomain type


    :param str transit_gateway_multicast_domain_id: The ID of the transit gateway multicast domain.
    """
    __args__ = dict()
    __args__['transitGatewayMulticastDomainId'] = transit_gateway_multicast_domain_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getTransitGatewayMulticastDomain', __args__, opts=opts, typ=GetTransitGatewayMulticastDomainResult).value

    return AwaitableGetTransitGatewayMulticastDomainResult(
        creation_time=__ret__.creation_time,
        options=__ret__.options,
        state=__ret__.state,
        tags=__ret__.tags,
        transit_gateway_multicast_domain_arn=__ret__.transit_gateway_multicast_domain_arn,
        transit_gateway_multicast_domain_id=__ret__.transit_gateway_multicast_domain_id)


@_utilities.lift_output_func(get_transit_gateway_multicast_domain)
def get_transit_gateway_multicast_domain_output(transit_gateway_multicast_domain_id: Optional[pulumi.Input[str]] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitGatewayMulticastDomainResult]:
    """
    The AWS::EC2::TransitGatewayMulticastDomain type


    :param str transit_gateway_multicast_domain_id: The ID of the transit gateway multicast domain.
    """
    ...
