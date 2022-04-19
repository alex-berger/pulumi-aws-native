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
    'GetDimensionResult',
    'AwaitableGetDimensionResult',
    'get_dimension',
    'get_dimension_output',
]

@pulumi.output_type
class GetDimensionResult:
    def __init__(__self__, arn=None, string_values=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if string_values and not isinstance(string_values, list):
            raise TypeError("Expected argument 'string_values' to be a list")
        pulumi.set(__self__, "string_values", string_values)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The ARN (Amazon resource name) of the created dimension.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="stringValues")
    def string_values(self) -> Optional[Sequence[str]]:
        """
        Specifies the value or list of values for the dimension.
        """
        return pulumi.get(self, "string_values")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DimensionTag']]:
        """
        Metadata that can be used to manage the dimension.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDimensionResult(GetDimensionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDimensionResult(
            arn=self.arn,
            string_values=self.string_values,
            tags=self.tags)


def get_dimension(name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDimensionResult:
    """
    A dimension can be used to limit the scope of a metric used in a security profile for AWS IoT Device Defender.


    :param str name: A unique identifier for the dimension.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iot:getDimension', __args__, opts=opts, typ=GetDimensionResult).value

    return AwaitableGetDimensionResult(
        arn=__ret__.arn,
        string_values=__ret__.string_values,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_dimension)
def get_dimension_output(name: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDimensionResult]:
    """
    A dimension can be used to limit the scope of a metric used in a security profile for AWS IoT Device Defender.


    :param str name: A unique identifier for the dimension.
    """
    ...
