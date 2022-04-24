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
    'GetDistributionConfigurationResult',
    'AwaitableGetDistributionConfigurationResult',
    'get_distribution_configuration',
    'get_distribution_configuration_output',
]

@pulumi.output_type
class GetDistributionConfigurationResult:
    def __init__(__self__, arn=None, description=None, distributions=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if distributions and not isinstance(distributions, list):
            raise TypeError("Expected argument 'distributions' to be a list")
        pulumi.set(__self__, "distributions", distributions)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the distribution configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the distribution configuration.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def distributions(self) -> Optional[Sequence['outputs.DistributionConfigurationDistribution']]:
        """
        The distributions of the distribution configuration.
        """
        return pulumi.get(self, "distributions")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        """
        The tags associated with the component.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDistributionConfigurationResult(GetDistributionConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDistributionConfigurationResult(
            arn=self.arn,
            description=self.description,
            distributions=self.distributions,
            tags=self.tags)


def get_distribution_configuration(arn: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDistributionConfigurationResult:
    """
    Resource schema for AWS::ImageBuilder::DistributionConfiguration


    :param str arn: The Amazon Resource Name (ARN) of the distribution configuration.
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:imagebuilder:getDistributionConfiguration', __args__, opts=opts, typ=GetDistributionConfigurationResult).value

    return AwaitableGetDistributionConfigurationResult(
        arn=__ret__.arn,
        description=__ret__.description,
        distributions=__ret__.distributions,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_distribution_configuration)
def get_distribution_configuration_output(arn: Optional[pulumi.Input[str]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDistributionConfigurationResult]:
    """
    Resource schema for AWS::ImageBuilder::DistributionConfiguration


    :param str arn: The Amazon Resource Name (ARN) of the distribution configuration.
    """
    ...
