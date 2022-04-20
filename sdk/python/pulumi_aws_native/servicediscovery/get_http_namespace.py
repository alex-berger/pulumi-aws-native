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
    'GetHttpNamespaceResult',
    'AwaitableGetHttpNamespaceResult',
    'get_http_namespace',
    'get_http_namespace_output',
]

@pulumi.output_type
class GetHttpNamespaceResult:
    def __init__(__self__, arn=None, description=None, id=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.HttpNamespaceTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetHttpNamespaceResult(GetHttpNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHttpNamespaceResult(
            arn=self.arn,
            description=self.description,
            id=self.id,
            tags=self.tags)


def get_http_namespace(id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHttpNamespaceResult:
    """
    Resource Type definition for AWS::ServiceDiscovery::HttpNamespace
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:servicediscovery:getHttpNamespace', __args__, opts=opts, typ=GetHttpNamespaceResult).value

    return AwaitableGetHttpNamespaceResult(
        arn=__ret__.arn,
        description=__ret__.description,
        id=__ret__.id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_http_namespace)
def get_http_namespace_output(id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHttpNamespaceResult]:
    """
    Resource Type definition for AWS::ServiceDiscovery::HttpNamespace
    """
    ...
