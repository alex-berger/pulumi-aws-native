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
    'GetFunctionResult',
    'AwaitableGetFunctionResult',
    'get_function',
    'get_function_output',
]

@pulumi.output_type
class GetFunctionResult:
    def __init__(__self__, function_arn=None, function_config=None, function_metadata=None, name=None, stage=None):
        if function_arn and not isinstance(function_arn, str):
            raise TypeError("Expected argument 'function_arn' to be a str")
        pulumi.set(__self__, "function_arn", function_arn)
        if function_config and not isinstance(function_config, dict):
            raise TypeError("Expected argument 'function_config' to be a dict")
        pulumi.set(__self__, "function_config", function_config)
        if function_metadata and not isinstance(function_metadata, dict):
            raise TypeError("Expected argument 'function_metadata' to be a dict")
        pulumi.set(__self__, "function_metadata", function_metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if stage and not isinstance(stage, str):
            raise TypeError("Expected argument 'stage' to be a str")
        pulumi.set(__self__, "stage", stage)

    @property
    @pulumi.getter(name="functionARN")
    def function_arn(self) -> Optional[str]:
        return pulumi.get(self, "function_arn")

    @property
    @pulumi.getter(name="functionConfig")
    def function_config(self) -> Optional['outputs.FunctionConfig']:
        return pulumi.get(self, "function_config")

    @property
    @pulumi.getter(name="functionMetadata")
    def function_metadata(self) -> Optional['outputs.FunctionMetadata']:
        return pulumi.get(self, "function_metadata")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def stage(self) -> Optional[str]:
        return pulumi.get(self, "stage")


class AwaitableGetFunctionResult(GetFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionResult(
            function_arn=self.function_arn,
            function_config=self.function_config,
            function_metadata=self.function_metadata,
            name=self.name,
            stage=self.stage)


def get_function(function_arn: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionResult:
    """
    Resource Type definition for AWS::CloudFront::Function
    """
    __args__ = dict()
    __args__['functionARN'] = function_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:cloudfront:getFunction', __args__, opts=opts, typ=GetFunctionResult).value

    return AwaitableGetFunctionResult(
        function_arn=__ret__.function_arn,
        function_config=__ret__.function_config,
        function_metadata=__ret__.function_metadata,
        name=__ret__.name,
        stage=__ret__.stage)


@_utilities.lift_output_func(get_function)
def get_function_output(function_arn: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFunctionResult]:
    """
    Resource Type definition for AWS::CloudFront::Function
    """
    ...
