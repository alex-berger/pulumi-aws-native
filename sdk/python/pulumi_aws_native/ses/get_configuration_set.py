# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetConfigurationSetResult',
    'AwaitableGetConfigurationSetResult',
    'get_configuration_set',
    'get_configuration_set_output',
]

@pulumi.output_type
class GetConfigurationSetResult:
    def __init__(__self__):


class AwaitableGetConfigurationSetResult(GetConfigurationSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigurationSetResult(
)


def get_configuration_set(name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigurationSetResult:
    """
    Resource schema for AWS::SES::ConfigurationSet.


    :param str name: The name of the configuration set.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ses:getConfigurationSet', __args__, opts=opts, typ=GetConfigurationSetResult).value

    return AwaitableGetConfigurationSetResult(


@_utilities.lift_output_func(get_configuration_set)
def get_configuration_set_output(name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConfigurationSetResult]:
    """
    Resource schema for AWS::SES::ConfigurationSet.


    :param str name: The name of the configuration set.
    """
    ...
