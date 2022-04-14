# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetObservabilityConfigurationResult',
    'AwaitableGetObservabilityConfigurationResult',
    'get_observability_configuration',
    'get_observability_configuration_output',
]

@pulumi.output_type
class GetObservabilityConfigurationResult:
    def __init__(__self__, latest=None, observability_configuration_arn=None, observability_configuration_revision=None):
        if latest and not isinstance(latest, bool):
            raise TypeError("Expected argument 'latest' to be a bool")
        pulumi.set(__self__, "latest", latest)
        if observability_configuration_arn and not isinstance(observability_configuration_arn, str):
            raise TypeError("Expected argument 'observability_configuration_arn' to be a str")
        pulumi.set(__self__, "observability_configuration_arn", observability_configuration_arn)
        if observability_configuration_revision and not isinstance(observability_configuration_revision, int):
            raise TypeError("Expected argument 'observability_configuration_revision' to be a int")
        pulumi.set(__self__, "observability_configuration_revision", observability_configuration_revision)

    @property
    @pulumi.getter
    def latest(self) -> Optional[bool]:
        """
        It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.
        """
        return pulumi.get(self, "latest")

    @property
    @pulumi.getter(name="observabilityConfigurationArn")
    def observability_configuration_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of this ObservabilityConfiguration
        """
        return pulumi.get(self, "observability_configuration_arn")

    @property
    @pulumi.getter(name="observabilityConfigurationRevision")
    def observability_configuration_revision(self) -> Optional[int]:
        """
        The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.
        """
        return pulumi.get(self, "observability_configuration_revision")


class AwaitableGetObservabilityConfigurationResult(GetObservabilityConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetObservabilityConfigurationResult(
            latest=self.latest,
            observability_configuration_arn=self.observability_configuration_arn,
            observability_configuration_revision=self.observability_configuration_revision)


def get_observability_configuration(observability_configuration_arn: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetObservabilityConfigurationResult:
    """
    The AWS::AppRunner::ObservabilityConfiguration resource  is an AWS App Runner resource type that specifies an App Runner observability configuration


    :param str observability_configuration_arn: The Amazon Resource Name (ARN) of this ObservabilityConfiguration
    """
    __args__ = dict()
    __args__['observabilityConfigurationArn'] = observability_configuration_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apprunner:getObservabilityConfiguration', __args__, opts=opts, typ=GetObservabilityConfigurationResult).value

    return AwaitableGetObservabilityConfigurationResult(
        latest=__ret__.latest,
        observability_configuration_arn=__ret__.observability_configuration_arn,
        observability_configuration_revision=__ret__.observability_configuration_revision)


@_utilities.lift_output_func(get_observability_configuration)
def get_observability_configuration_output(observability_configuration_arn: Optional[pulumi.Input[str]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetObservabilityConfigurationResult]:
    """
    The AWS::AppRunner::ObservabilityConfiguration resource  is an AWS App Runner resource type that specifies an App Runner observability configuration


    :param str observability_configuration_arn: The Amazon Resource Name (ARN) of this ObservabilityConfiguration
    """
    ...