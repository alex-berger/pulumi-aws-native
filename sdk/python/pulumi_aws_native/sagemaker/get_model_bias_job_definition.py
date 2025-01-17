# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetModelBiasJobDefinitionResult',
    'AwaitableGetModelBiasJobDefinitionResult',
    'get_model_bias_job_definition',
    'get_model_bias_job_definition_output',
]

@pulumi.output_type
class GetModelBiasJobDefinitionResult:
    def __init__(__self__, creation_time=None, job_definition_arn=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if job_definition_arn and not isinstance(job_definition_arn, str):
            raise TypeError("Expected argument 'job_definition_arn' to be a str")
        pulumi.set(__self__, "job_definition_arn", job_definition_arn)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        """
        The time at which the job definition was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="jobDefinitionArn")
    def job_definition_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of job definition.
        """
        return pulumi.get(self, "job_definition_arn")


class AwaitableGetModelBiasJobDefinitionResult(GetModelBiasJobDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetModelBiasJobDefinitionResult(
            creation_time=self.creation_time,
            job_definition_arn=self.job_definition_arn)


def get_model_bias_job_definition(job_definition_arn: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetModelBiasJobDefinitionResult:
    """
    Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition


    :param str job_definition_arn: The Amazon Resource Name (ARN) of job definition.
    """
    __args__ = dict()
    __args__['jobDefinitionArn'] = job_definition_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:sagemaker:getModelBiasJobDefinition', __args__, opts=opts, typ=GetModelBiasJobDefinitionResult).value

    return AwaitableGetModelBiasJobDefinitionResult(
        creation_time=__ret__.creation_time,
        job_definition_arn=__ret__.job_definition_arn)


@_utilities.lift_output_func(get_model_bias_job_definition)
def get_model_bias_job_definition_output(job_definition_arn: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetModelBiasJobDefinitionResult]:
    """
    Resource Type definition for AWS::SageMaker::ModelBiasJobDefinition


    :param str job_definition_arn: The Amazon Resource Name (ARN) of job definition.
    """
    ...
