# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetApplicationVersionResult',
    'AwaitableGetApplicationVersionResult',
    'get_application_version',
    'get_application_version_output',
]

@pulumi.output_type
class GetApplicationVersionResult:
    def __init__(__self__, description=None, id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetApplicationVersionResult(GetApplicationVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationVersionResult(
            description=self.description,
            id=self.id)


def get_application_version(id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationVersionResult:
    """
    Resource Type definition for AWS::ElasticBeanstalk::ApplicationVersion
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:elasticbeanstalk:getApplicationVersion', __args__, opts=opts, typ=GetApplicationVersionResult).value

    return AwaitableGetApplicationVersionResult(
        description=__ret__.description,
        id=__ret__.id)


@_utilities.lift_output_func(get_application_version)
def get_application_version_output(id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationVersionResult]:
    """
    Resource Type definition for AWS::ElasticBeanstalk::ApplicationVersion
    """
    ...