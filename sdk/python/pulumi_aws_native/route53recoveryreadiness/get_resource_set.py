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
    'GetResourceSetResult',
    'AwaitableGetResourceSetResult',
    'get_resource_set',
    'get_resource_set_output',
]

@pulumi.output_type
class GetResourceSetResult:
    def __init__(__self__, resource_set_arn=None, resources=None, tags=None):
        if resource_set_arn and not isinstance(resource_set_arn, str):
            raise TypeError("Expected argument 'resource_set_arn' to be a str")
        pulumi.set(__self__, "resource_set_arn", resource_set_arn)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="resourceSetArn")
    def resource_set_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the resource set.
        """
        return pulumi.get(self, "resource_set_arn")

    @property
    @pulumi.getter
    def resources(self) -> Optional[Sequence['outputs.ResourceSetResource']]:
        """
        A list of resource objects in the resource set.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ResourceSetTag']]:
        """
        A tag to associate with the parameters for a resource set.
        """
        return pulumi.get(self, "tags")


class AwaitableGetResourceSetResult(GetResourceSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceSetResult(
            resource_set_arn=self.resource_set_arn,
            resources=self.resources,
            tags=self.tags)


def get_resource_set(resource_set_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceSetResult:
    """
    Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.


    :param str resource_set_name: The name of the resource set to create.
    """
    __args__ = dict()
    __args__['resourceSetName'] = resource_set_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:route53recoveryreadiness:getResourceSet', __args__, opts=opts, typ=GetResourceSetResult).value

    return AwaitableGetResourceSetResult(
        resource_set_arn=__ret__.resource_set_arn,
        resources=__ret__.resources,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_resource_set)
def get_resource_set_output(resource_set_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourceSetResult]:
    """
    Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.


    :param str resource_set_name: The name of the resource set to create.
    """
    ...
