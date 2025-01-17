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
    'GetAttributeGroupResult',
    'AwaitableGetAttributeGroupResult',
    'get_attribute_group',
    'get_attribute_group_output',
]

@pulumi.output_type
class GetAttributeGroupResult:
    def __init__(__self__, arn=None, attributes=None, description=None, id=None, name=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def attributes(self) -> Optional[Any]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the attribute group. 
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the attribute group. 
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Optional['outputs.AttributeGroupTags']:
        return pulumi.get(self, "tags")


class AwaitableGetAttributeGroupResult(GetAttributeGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAttributeGroupResult(
            arn=self.arn,
            attributes=self.attributes,
            description=self.description,
            id=self.id,
            name=self.name,
            tags=self.tags)


def get_attribute_group(id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAttributeGroupResult:
    """
    Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:servicecatalogappregistry:getAttributeGroup', __args__, opts=opts, typ=GetAttributeGroupResult).value

    return AwaitableGetAttributeGroupResult(
        arn=__ret__.arn,
        attributes=__ret__.attributes,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_attribute_group)
def get_attribute_group_output(id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAttributeGroupResult]:
    """
    Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.
    """
    ...
