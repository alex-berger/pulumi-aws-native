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
    'GetDatasetResult',
    'AwaitableGetDatasetResult',
    'get_dataset',
    'get_dataset_output',
]

@pulumi.output_type
class GetDatasetResult:
    def __init__(__self__, format=None, format_options=None, input=None, path_options=None):
        if format and not isinstance(format, str):
            raise TypeError("Expected argument 'format' to be a str")
        pulumi.set(__self__, "format", format)
        if format_options and not isinstance(format_options, dict):
            raise TypeError("Expected argument 'format_options' to be a dict")
        pulumi.set(__self__, "format_options", format_options)
        if input and not isinstance(input, dict):
            raise TypeError("Expected argument 'input' to be a dict")
        pulumi.set(__self__, "input", input)
        if path_options and not isinstance(path_options, dict):
            raise TypeError("Expected argument 'path_options' to be a dict")
        pulumi.set(__self__, "path_options", path_options)

    @property
    @pulumi.getter
    def format(self) -> Optional['DatasetFormat']:
        """
        Dataset format
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="formatOptions")
    def format_options(self) -> Optional['outputs.DatasetFormatOptions']:
        """
        Format options for dataset
        """
        return pulumi.get(self, "format_options")

    @property
    @pulumi.getter
    def input(self) -> Optional['outputs.DatasetInput']:
        """
        Input
        """
        return pulumi.get(self, "input")

    @property
    @pulumi.getter(name="pathOptions")
    def path_options(self) -> Optional['outputs.DatasetPathOptions']:
        """
        PathOptions
        """
        return pulumi.get(self, "path_options")


class AwaitableGetDatasetResult(GetDatasetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatasetResult(
            format=self.format,
            format_options=self.format_options,
            input=self.input,
            path_options=self.path_options)


def get_dataset(name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatasetResult:
    """
    Resource schema for AWS::DataBrew::Dataset.


    :param str name: Dataset name
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:databrew:getDataset', __args__, opts=opts, typ=GetDatasetResult).value

    return AwaitableGetDatasetResult(
        format=__ret__.format,
        format_options=__ret__.format_options,
        input=__ret__.input,
        path_options=__ret__.path_options)


@_utilities.lift_output_func(get_dataset)
def get_dataset_output(name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatasetResult]:
    """
    Resource schema for AWS::DataBrew::Dataset.


    :param str name: Dataset name
    """
    ...
