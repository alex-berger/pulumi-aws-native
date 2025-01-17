# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetEC2FleetResult',
    'AwaitableGetEC2FleetResult',
    'get_ec2_fleet',
    'get_ec2_fleet_output',
]

@pulumi.output_type
class GetEC2FleetResult:
    def __init__(__self__, context=None, excess_capacity_termination_policy=None, fleet_id=None, target_capacity_specification=None):
        if context and not isinstance(context, str):
            raise TypeError("Expected argument 'context' to be a str")
        pulumi.set(__self__, "context", context)
        if excess_capacity_termination_policy and not isinstance(excess_capacity_termination_policy, str):
            raise TypeError("Expected argument 'excess_capacity_termination_policy' to be a str")
        pulumi.set(__self__, "excess_capacity_termination_policy", excess_capacity_termination_policy)
        if fleet_id and not isinstance(fleet_id, str):
            raise TypeError("Expected argument 'fleet_id' to be a str")
        pulumi.set(__self__, "fleet_id", fleet_id)
        if target_capacity_specification and not isinstance(target_capacity_specification, dict):
            raise TypeError("Expected argument 'target_capacity_specification' to be a dict")
        pulumi.set(__self__, "target_capacity_specification", target_capacity_specification)

    @property
    @pulumi.getter
    def context(self) -> Optional[str]:
        return pulumi.get(self, "context")

    @property
    @pulumi.getter(name="excessCapacityTerminationPolicy")
    def excess_capacity_termination_policy(self) -> Optional['EC2FleetExcessCapacityTerminationPolicy']:
        return pulumi.get(self, "excess_capacity_termination_policy")

    @property
    @pulumi.getter(name="fleetId")
    def fleet_id(self) -> Optional[str]:
        return pulumi.get(self, "fleet_id")

    @property
    @pulumi.getter(name="targetCapacitySpecification")
    def target_capacity_specification(self) -> Optional['outputs.EC2FleetTargetCapacitySpecificationRequest']:
        return pulumi.get(self, "target_capacity_specification")


class AwaitableGetEC2FleetResult(GetEC2FleetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEC2FleetResult(
            context=self.context,
            excess_capacity_termination_policy=self.excess_capacity_termination_policy,
            fleet_id=self.fleet_id,
            target_capacity_specification=self.target_capacity_specification)


def get_ec2_fleet(fleet_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEC2FleetResult:
    """
    Resource Type definition for AWS::EC2::EC2Fleet
    """
    __args__ = dict()
    __args__['fleetId'] = fleet_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getEC2Fleet', __args__, opts=opts, typ=GetEC2FleetResult).value

    return AwaitableGetEC2FleetResult(
        context=__ret__.context,
        excess_capacity_termination_policy=__ret__.excess_capacity_termination_policy,
        fleet_id=__ret__.fleet_id,
        target_capacity_specification=__ret__.target_capacity_specification)


@_utilities.lift_output_func(get_ec2_fleet)
def get_ec2_fleet_output(fleet_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEC2FleetResult]:
    """
    Resource Type definition for AWS::EC2::EC2Fleet
    """
    ...
