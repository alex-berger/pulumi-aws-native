# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetAssessmentTargetResult',
    'AwaitableGetAssessmentTargetResult',
    'get_assessment_target',
    'get_assessment_target_output',
]

@pulumi.output_type
class GetAssessmentTargetResult:
    def __init__(__self__, arn=None, id=None, resource_group_arn=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if resource_group_arn and not isinstance(resource_group_arn, str):
            raise TypeError("Expected argument 'resource_group_arn' to be a str")
        pulumi.set(__self__, "resource_group_arn", resource_group_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceGroupArn")
    def resource_group_arn(self) -> Optional[str]:
        return pulumi.get(self, "resource_group_arn")


class AwaitableGetAssessmentTargetResult(GetAssessmentTargetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssessmentTargetResult(
            arn=self.arn,
            id=self.id,
            resource_group_arn=self.resource_group_arn)


def get_assessment_target(id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAssessmentTargetResult:
    """
    Resource Type definition for AWS::Inspector::AssessmentTarget
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:inspector:getAssessmentTarget', __args__, opts=opts, typ=GetAssessmentTargetResult).value

    return AwaitableGetAssessmentTargetResult(
        arn=__ret__.arn,
        id=__ret__.id,
        resource_group_arn=__ret__.resource_group_arn)


@_utilities.lift_output_func(get_assessment_target)
def get_assessment_target_output(id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAssessmentTargetResult]:
    """
    Resource Type definition for AWS::Inspector::AssessmentTarget
    """
    ...