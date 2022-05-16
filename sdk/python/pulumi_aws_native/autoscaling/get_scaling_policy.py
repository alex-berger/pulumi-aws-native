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
    def __init__(__self__, adjustment_type=None, auto_scaling_group_name=None, cooldown=None, estimated_instance_warmup=None, id=None, metric_aggregation_type=None, min_adjustment_magnitude=None, policy_type=None, predictive_scaling_configuration=None, scaling_adjustment=None, step_adjustments=None, target_tracking_configuration=None):
        if adjustment_type and not isinstance(adjustment_type, str):
            raise TypeError("Expected argument 'adjustment_type' to be a str")
        pulumi.set(__self__, "adjustment_type", adjustment_type)
        if auto_scaling_group_name and not isinstance(auto_scaling_group_name, str):
            raise TypeError("Expected argument 'auto_scaling_group_name' to be a str")
        pulumi.set(__self__, "auto_scaling_group_name", auto_scaling_group_name)
        if cooldown and not isinstance(cooldown, str):
            raise TypeError("Expected argument 'cooldown' to be a str")
        pulumi.set(__self__, "cooldown", cooldown)
        if estimated_instance_warmup and not isinstance(estimated_instance_warmup, int):
            raise TypeError("Expected argument 'estimated_instance_warmup' to be a int")
        pulumi.set(__self__, "estimated_instance_warmup", estimated_instance_warmup)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metric_aggregation_type and not isinstance(metric_aggregation_type, str):
            raise TypeError("Expected argument 'metric_aggregation_type' to be a str")
        pulumi.set(__self__, "metric_aggregation_type", metric_aggregation_type)
        if min_adjustment_magnitude and not isinstance(min_adjustment_magnitude, int):
            raise TypeError("Expected argument 'min_adjustment_magnitude' to be a int")
        pulumi.set(__self__, "min_adjustment_magnitude", min_adjustment_magnitude)
        if policy_type and not isinstance(policy_type, str):
            raise TypeError("Expected argument 'policy_type' to be a str")
        pulumi.set(__self__, "policy_type", policy_type)
        if predictive_scaling_configuration and not isinstance(predictive_scaling_configuration, dict):
            raise TypeError("Expected argument 'predictive_scaling_configuration' to be a dict")
        pulumi.set(__self__, "predictive_scaling_configuration", predictive_scaling_configuration)
        if scaling_adjustment and not isinstance(scaling_adjustment, int):
            raise TypeError("Expected argument 'scaling_adjustment' to be a int")
        pulumi.set(__self__, "scaling_adjustment", scaling_adjustment)
        if step_adjustments and not isinstance(step_adjustments, list):
            raise TypeError("Expected argument 'step_adjustments' to be a list")
        pulumi.set(__self__, "step_adjustments", step_adjustments)
        if target_tracking_configuration and not isinstance(target_tracking_configuration, dict):
            raise TypeError("Expected argument 'target_tracking_configuration' to be a dict")
        pulumi.set(__self__, "target_tracking_configuration", target_tracking_configuration)

    @property
    @pulumi.getter(name="adjustmentType")
    def adjustment_type(self) -> Optional[str]:
        return pulumi.get(self, "adjustment_type")

    @property
    @pulumi.getter(name="autoScalingGroupName")
    def auto_scaling_group_name(self) -> Optional[str]:
        return pulumi.get(self, "auto_scaling_group_name")

    @property
    @pulumi.getter
    def cooldown(self) -> Optional[str]:
        return pulumi.get(self, "cooldown")

    @property
    @pulumi.getter(name="estimatedInstanceWarmup")
    def estimated_instance_warmup(self) -> Optional[int]:
        return pulumi.get(self, "estimated_instance_warmup")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="metricAggregationType")
    def metric_aggregation_type(self) -> Optional[str]:
        return pulumi.get(self, "metric_aggregation_type")

    @property
    @pulumi.getter(name="minAdjustmentMagnitude")
    def min_adjustment_magnitude(self) -> Optional[int]:
        return pulumi.get(self, "min_adjustment_magnitude")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[str]:
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="predictiveScalingConfiguration")
    def predictive_scaling_configuration(self) -> Optional['outputs.ScalingPolicyPredictiveScalingConfiguration']:
        return pulumi.get(self, "predictive_scaling_configuration")

    @property
    @pulumi.getter(name="scalingAdjustment")
    def scaling_adjustment(self) -> Optional[int]:
        return pulumi.get(self, "scaling_adjustment")

    @property
    @pulumi.getter(name="stepAdjustments")
    def step_adjustments(self) -> Optional[Sequence['outputs.ScalingPolicyStepAdjustment']]:
        return pulumi.get(self, "step_adjustments")

    @property
    @pulumi.getter(name="targetTrackingConfiguration")
    def target_tracking_configuration(self) -> Optional['outputs.ScalingPolicyTargetTrackingConfiguration']:
        return pulumi.get(self, "target_tracking_configuration")


class AwaitableGetScalingPolicyResult(GetScalingPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScalingPolicyResult(
            adjustment_type=self.adjustment_type,
            auto_scaling_group_name=self.auto_scaling_group_name,
            cooldown=self.cooldown,
            estimated_instance_warmup=self.estimated_instance_warmup,
            id=self.id,
            metric_aggregation_type=self.metric_aggregation_type,
            min_adjustment_magnitude=self.min_adjustment_magnitude,
            policy_type=self.policy_type,
            predictive_scaling_configuration=self.predictive_scaling_configuration,
            scaling_adjustment=self.scaling_adjustment,
            step_adjustments=self.step_adjustments,
            target_tracking_configuration=self.target_tracking_configuration)


def get_scaling_policy(id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScalingPolicyResult:
    """
    Resource Type definition for AWS::AutoScaling::ScalingPolicy
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:autoscaling:getScalingPolicy', __args__, opts=opts, typ=GetScalingPolicyResult).value

    return AwaitableGetScalingPolicyResult(
        adjustment_type=__ret__.adjustment_type,
        auto_scaling_group_name=__ret__.auto_scaling_group_name,
        cooldown=__ret__.cooldown,
        estimated_instance_warmup=__ret__.estimated_instance_warmup,
        id=__ret__.id,
        metric_aggregation_type=__ret__.metric_aggregation_type,
        min_adjustment_magnitude=__ret__.min_adjustment_magnitude,
        policy_type=__ret__.policy_type,
        predictive_scaling_configuration=__ret__.predictive_scaling_configuration,
        scaling_adjustment=__ret__.scaling_adjustment,
        step_adjustments=__ret__.step_adjustments,
        target_tracking_configuration=__ret__.target_tracking_configuration)


@_utilities.lift_output_func(get_scaling_policy)
def get_scaling_policy_output(id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScalingPolicyResult]:
    """
    Resource Type definition for AWS::AutoScaling::ScalingPolicy
    """
    ...