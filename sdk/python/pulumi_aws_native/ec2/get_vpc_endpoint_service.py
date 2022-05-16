# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetVPCEndpointServiceResult',
    'AwaitableGetVPCEndpointServiceResult',
    'get_vpc_endpoint_service',
    'get_vpc_endpoint_service_output',
]

@pulumi.output_type
class GetVPCEndpointServiceResult:
    def __init__(__self__, acceptance_required=None, gateway_load_balancer_arns=None, id=None, network_load_balancer_arns=None, payer_responsibility=None):
        if acceptance_required and not isinstance(acceptance_required, bool):
            raise TypeError("Expected argument 'acceptance_required' to be a bool")
        pulumi.set(__self__, "acceptance_required", acceptance_required)
        if gateway_load_balancer_arns and not isinstance(gateway_load_balancer_arns, list):
            raise TypeError("Expected argument 'gateway_load_balancer_arns' to be a list")
        pulumi.set(__self__, "gateway_load_balancer_arns", gateway_load_balancer_arns)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if network_load_balancer_arns and not isinstance(network_load_balancer_arns, list):
            raise TypeError("Expected argument 'network_load_balancer_arns' to be a list")
        pulumi.set(__self__, "network_load_balancer_arns", network_load_balancer_arns)
        if payer_responsibility and not isinstance(payer_responsibility, str):
            raise TypeError("Expected argument 'payer_responsibility' to be a str")
        pulumi.set(__self__, "payer_responsibility", payer_responsibility)

    @property
    @pulumi.getter(name="acceptanceRequired")
    def acceptance_required(self) -> Optional[bool]:
        return pulumi.get(self, "acceptance_required")

    @property
    @pulumi.getter(name="gatewayLoadBalancerArns")
    def gateway_load_balancer_arns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "gateway_load_balancer_arns")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="networkLoadBalancerArns")
    def network_load_balancer_arns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "network_load_balancer_arns")

    @property
    @pulumi.getter(name="payerResponsibility")
    def payer_responsibility(self) -> Optional[str]:
        return pulumi.get(self, "payer_responsibility")


class AwaitableGetVPCEndpointServiceResult(GetVPCEndpointServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVPCEndpointServiceResult(
            acceptance_required=self.acceptance_required,
            gateway_load_balancer_arns=self.gateway_load_balancer_arns,
            id=self.id,
            network_load_balancer_arns=self.network_load_balancer_arns,
            payer_responsibility=self.payer_responsibility)


def get_vpc_endpoint_service(id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVPCEndpointServiceResult:
    """
    Resource Type definition for AWS::EC2::VPCEndpointService
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getVPCEndpointService', __args__, opts=opts, typ=GetVPCEndpointServiceResult).value

    return AwaitableGetVPCEndpointServiceResult(
        acceptance_required=__ret__.acceptance_required,
        gateway_load_balancer_arns=__ret__.gateway_load_balancer_arns,
        id=__ret__.id,
        network_load_balancer_arns=__ret__.network_load_balancer_arns,
        payer_responsibility=__ret__.payer_responsibility)


@_utilities.lift_output_func(get_vpc_endpoint_service)
def get_vpc_endpoint_service_output(id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVPCEndpointServiceResult]:
    """
    Resource Type definition for AWS::EC2::VPCEndpointService
    """
    ...