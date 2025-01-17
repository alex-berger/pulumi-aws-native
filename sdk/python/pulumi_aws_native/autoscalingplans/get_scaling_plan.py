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
    'GetScalingPlanResult',
    'AwaitableGetScalingPlanResult',
    'get_scaling_plan',
    'get_scaling_plan_output',
]

@pulumi.output_type
class GetScalingPlanResult:
    def __init__(__self__, application_source=None, id=None, scaling_instructions=None, scaling_plan_name=None, scaling_plan_version=None):
        if application_source and not isinstance(application_source, dict):
            raise TypeError("Expected argument 'application_source' to be a dict")
        pulumi.set(__self__, "application_source", application_source)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if scaling_instructions and not isinstance(scaling_instructions, list):
            raise TypeError("Expected argument 'scaling_instructions' to be a list")
        pulumi.set(__self__, "scaling_instructions", scaling_instructions)
        if scaling_plan_name and not isinstance(scaling_plan_name, str):
            raise TypeError("Expected argument 'scaling_plan_name' to be a str")
        pulumi.set(__self__, "scaling_plan_name", scaling_plan_name)
        if scaling_plan_version and not isinstance(scaling_plan_version, str):
            raise TypeError("Expected argument 'scaling_plan_version' to be a str")
        pulumi.set(__self__, "scaling_plan_version", scaling_plan_version)

    @property
    @pulumi.getter(name="applicationSource")
    def application_source(self) -> Optional['outputs.ScalingPlanApplicationSource']:
        return pulumi.get(self, "application_source")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="scalingInstructions")
    def scaling_instructions(self) -> Optional[Sequence['outputs.ScalingPlanScalingInstruction']]:
        return pulumi.get(self, "scaling_instructions")

    @property
    @pulumi.getter(name="scalingPlanName")
    def scaling_plan_name(self) -> Optional[str]:
        return pulumi.get(self, "scaling_plan_name")

    @property
    @pulumi.getter(name="scalingPlanVersion")
    def scaling_plan_version(self) -> Optional[str]:
        return pulumi.get(self, "scaling_plan_version")


class AwaitableGetScalingPlanResult(GetScalingPlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetScalingPlanResult(
            application_source=self.application_source,
            id=self.id,
            scaling_instructions=self.scaling_instructions,
            scaling_plan_name=self.scaling_plan_name,
            scaling_plan_version=self.scaling_plan_version)


def get_scaling_plan(id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetScalingPlanResult:
    """
    Resource Type definition for AWS::AutoScalingPlans::ScalingPlan
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:autoscalingplans:getScalingPlan', __args__, opts=opts, typ=GetScalingPlanResult).value

    return AwaitableGetScalingPlanResult(
        application_source=__ret__.application_source,
        id=__ret__.id,
        scaling_instructions=__ret__.scaling_instructions,
        scaling_plan_name=__ret__.scaling_plan_name,
        scaling_plan_version=__ret__.scaling_plan_version)


@_utilities.lift_output_func(get_scaling_plan)
def get_scaling_plan_output(id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetScalingPlanResult]:
    """
    Resource Type definition for AWS::AutoScalingPlans::ScalingPlan
    """
    ...
