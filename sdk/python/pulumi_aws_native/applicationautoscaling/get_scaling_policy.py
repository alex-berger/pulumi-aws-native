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
    'GetScalingPolicyResult',
    'AwaitableGetScalingPolicyResult',
    'get_scaling_policy',
    'get_scaling_policy_output',
]

@pulumi.output_type
class GetScalingPolicyResult:
    def __init__(__self__, id=None, policy_type=None, step_scaling_policy_configuration=None, target_tracking_scaling_policy_configuration=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy_type and not isinstance(policy_type, str):
            raise TypeError("Expected argument 'policy_type' to be a str")
        pulumi.set(__self__, "policy_type", policy_type)
        if step_scaling_policy_configuration and not isinstance(step_scaling_policy_configuration, dict):
            raise TypeError("Expected argument 'step_scaling_policy_configuration' to be a dict")
        pulumi.set(__self__, "step_scaling_policy_configuration", step_scaling_policy_configuration)
        if target_tracking_scaling_policy_configuration and not isinstance(target_tracking_scaling_policy_configuration, dict):
            raise TypeError("Expected argument 'target_tracking_scaling_policy_configuration' to be a dict")
        pulumi.set(__self__, "target_tracking_scaling_policy_configuration", target_tracking_scaling_policy_configuration)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[str]:
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="stepScalingPolicyConfiguration")
    def step_scaling_policy_configuration(self) -> Optional['outputs.ScalingPolicyStepScalingPolicyConfiguration']:
        return pulumi.get(self, "step_scaling_policy_configuration")

    @property
    @pulumi.getter(name="targetTrackingScalingPolicyConfiguration")
    def target_tracking_scaling_policy_configuration(self) -> Optional['outputs.ScalingPolicyTargetTrackingScalingPolicyConfiguration']:
        return pulumi.get(self, "target_tracking_scaling_policy_configuration")


class AwaitableGetScalingPolicyResult(GetScalingPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScalingPolicyResult(
            id=self.id,
            policy_type=self.policy_type,
            step_scaling_policy_configuration=self.step_scaling_policy_configuration,
            target_tracking_scaling_policy_configuration=self.target_tracking_scaling_policy_configuration)


def get_scaling_policy(id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScalingPolicyResult:
    """
    Resource Type definition for AWS::ApplicationAutoScaling::ScalingPolicy
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:applicationautoscaling:getScalingPolicy', __args__, opts=opts, typ=GetScalingPolicyResult).value

    return AwaitableGetScalingPolicyResult(
        id=__ret__.id,
        policy_type=__ret__.policy_type,
        step_scaling_policy_configuration=__ret__.step_scaling_policy_configuration,
        target_tracking_scaling_policy_configuration=__ret__.target_tracking_scaling_policy_configuration)


@_utilities.lift_output_func(get_scaling_policy)
def get_scaling_policy_output(id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScalingPolicyResult]:
    """
    Resource Type definition for AWS::ApplicationAutoScaling::ScalingPolicy
    """
    ...
