# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetTransitGatewayRegistrationResult',
    'AwaitableGetTransitGatewayRegistrationResult',
    'get_transit_gateway_registration',
    'get_transit_gateway_registration_output',
]

@pulumi.output_type
class GetTransitGatewayRegistrationResult:
    def __init__(__self__):


class AwaitableGetTransitGatewayRegistrationResult(GetTransitGatewayRegistrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransitGatewayRegistrationResult(
)


def get_transit_gateway_registration(global_network_id: Optional[str] = None,
                                     transit_gateway_arn: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransitGatewayRegistrationResult:
    """
    The AWS::NetworkManager::TransitGatewayRegistration type registers a transit gateway in your global network. The transit gateway can be in any AWS Region, but it must be owned by the same AWS account that owns the global network. You cannot register a transit gateway in more than one global network.


    :param str global_network_id: The ID of the global network.
    :param str transit_gateway_arn: The Amazon Resource Name (ARN) of the transit gateway.
    """
    __args__ = dict()
    __args__['globalNetworkId'] = global_network_id
    __args__['transitGatewayArn'] = transit_gateway_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:networkmanager:getTransitGatewayRegistration', __args__, opts=opts, typ=GetTransitGatewayRegistrationResult).value

    return AwaitableGetTransitGatewayRegistrationResult(


@_utilities.lift_output_func(get_transit_gateway_registration)
def get_transit_gateway_registration_output(global_network_id: Optional[pulumi.Input[str]] = None,
                                            transit_gateway_arn: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransitGatewayRegistrationResult]:
    """
    The AWS::NetworkManager::TransitGatewayRegistration type registers a transit gateway in your global network. The transit gateway can be in any AWS Region, but it must be owned by the same AWS account that owns the global network. You cannot register a transit gateway in more than one global network.


    :param str global_network_id: The ID of the global network.
    :param str transit_gateway_arn: The Amazon Resource Name (ARN) of the transit gateway.
    """
    ...
