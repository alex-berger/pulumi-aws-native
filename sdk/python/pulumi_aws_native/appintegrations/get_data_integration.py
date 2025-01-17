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
    'GetDataIntegrationResult',
    'AwaitableGetDataIntegrationResult',
    'get_data_integration',
    'get_data_integration_output',
]

@pulumi.output_type
class GetDataIntegrationResult:
    def __init__(__self__, data_integration_arn=None, description=None, id=None, name=None, tags=None):
        if data_integration_arn and not isinstance(data_integration_arn, str):
            raise TypeError("Expected argument 'data_integration_arn' to be a str")
        pulumi.set(__self__, "data_integration_arn", data_integration_arn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="dataIntegrationArn")
    def data_integration_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the data integration.
        """
        return pulumi.get(self, "data_integration_arn")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The data integration description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The unique identifer of the data integration.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the data integration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DataIntegrationTag']]:
        """
        The tags (keys and values) associated with the data integration.
        """
        return pulumi.get(self, "tags")


class AwaitableGetDataIntegrationResult(GetDataIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataIntegrationResult(
            data_integration_arn=self.data_integration_arn,
            description=self.description,
            id=self.id,
            name=self.name,
            tags=self.tags)


def get_data_integration(id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataIntegrationResult:
    """
    Resource Type definition for AWS::AppIntegrations::DataIntegration


    :param str id: The unique identifer of the data integration.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:appintegrations:getDataIntegration', __args__, opts=opts, typ=GetDataIntegrationResult).value

    return AwaitableGetDataIntegrationResult(
        data_integration_arn=__ret__.data_integration_arn,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_data_integration)
def get_data_integration_output(id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataIntegrationResult]:
    """
    Resource Type definition for AWS::AppIntegrations::DataIntegration


    :param str id: The unique identifer of the data integration.
    """
    ...
