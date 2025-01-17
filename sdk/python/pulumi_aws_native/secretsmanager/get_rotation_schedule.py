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
    'GetRotationScheduleResult',
    'AwaitableGetRotationScheduleResult',
    'get_rotation_schedule',
    'get_rotation_schedule_output',
]

@pulumi.output_type
class GetRotationScheduleResult:
    def __init__(__self__, hosted_rotation_lambda=None, id=None, rotate_immediately_on_update=None, rotation_lambda_arn=None, rotation_rules=None):
        if hosted_rotation_lambda and not isinstance(hosted_rotation_lambda, dict):
            raise TypeError("Expected argument 'hosted_rotation_lambda' to be a dict")
        pulumi.set(__self__, "hosted_rotation_lambda", hosted_rotation_lambda)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if rotate_immediately_on_update and not isinstance(rotate_immediately_on_update, bool):
            raise TypeError("Expected argument 'rotate_immediately_on_update' to be a bool")
        pulumi.set(__self__, "rotate_immediately_on_update", rotate_immediately_on_update)
        if rotation_lambda_arn and not isinstance(rotation_lambda_arn, str):
            raise TypeError("Expected argument 'rotation_lambda_arn' to be a str")
        pulumi.set(__self__, "rotation_lambda_arn", rotation_lambda_arn)
        if rotation_rules and not isinstance(rotation_rules, dict):
            raise TypeError("Expected argument 'rotation_rules' to be a dict")
        pulumi.set(__self__, "rotation_rules", rotation_rules)

    @property
    @pulumi.getter(name="hostedRotationLambda")
    def hosted_rotation_lambda(self) -> Optional['outputs.RotationScheduleHostedRotationLambda']:
        return pulumi.get(self, "hosted_rotation_lambda")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="rotateImmediatelyOnUpdate")
    def rotate_immediately_on_update(self) -> Optional[bool]:
        return pulumi.get(self, "rotate_immediately_on_update")

    @property
    @pulumi.getter(name="rotationLambdaARN")
    def rotation_lambda_arn(self) -> Optional[str]:
        return pulumi.get(self, "rotation_lambda_arn")

    @property
    @pulumi.getter(name="rotationRules")
    def rotation_rules(self) -> Optional['outputs.RotationScheduleRotationRules']:
        return pulumi.get(self, "rotation_rules")


class AwaitableGetRotationScheduleResult(GetRotationScheduleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRotationScheduleResult(
            hosted_rotation_lambda=self.hosted_rotation_lambda,
            id=self.id,
            rotate_immediately_on_update=self.rotate_immediately_on_update,
            rotation_lambda_arn=self.rotation_lambda_arn,
            rotation_rules=self.rotation_rules)


def get_rotation_schedule(id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRotationScheduleResult:
    """
    Resource Type definition for AWS::SecretsManager::RotationSchedule
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:secretsmanager:getRotationSchedule', __args__, opts=opts, typ=GetRotationScheduleResult).value

    return AwaitableGetRotationScheduleResult(
        hosted_rotation_lambda=__ret__.hosted_rotation_lambda,
        id=__ret__.id,
        rotate_immediately_on_update=__ret__.rotate_immediately_on_update,
        rotation_lambda_arn=__ret__.rotation_lambda_arn,
        rotation_rules=__ret__.rotation_rules)


@_utilities.lift_output_func(get_rotation_schedule)
def get_rotation_schedule_output(id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRotationScheduleResult]:
    """
    Resource Type definition for AWS::SecretsManager::RotationSchedule
    """
    ...
