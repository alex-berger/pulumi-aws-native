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
    'GetGraphResult',
    'AwaitableGetGraphResult',
    'get_graph',
    'get_graph_output',
]

@pulumi.output_type
class GetGraphResult:
    def __init__(__self__, arn=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Detective graph ARN
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.GraphTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetGraphResult(GetGraphResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGraphResult(
            arn=self.arn,
            tags=self.tags)


def get_graph(arn: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGraphResult:
    """
    Resource schema for AWS::Detective::Graph


    :param str arn: The Detective graph ARN
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:detective:getGraph', __args__, opts=opts, typ=GetGraphResult).value

    return AwaitableGetGraphResult(
        arn=__ret__.arn,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_graph)
def get_graph_output(arn: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGraphResult]:
    """
    Resource schema for AWS::Detective::Graph


    :param str arn: The Detective graph ARN
    """
    ...