# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetParameterResult',
    'AwaitableGetParameterResult',
    'get_parameter',
    'get_parameter_output',
]

@pulumi.output_type
class GetParameterResult:
    def __init__(__self__, allowed_pattern=None, data_type=None, description=None, id=None, policies=None, tags=None, tier=None, type=None, value=None):
        if allowed_pattern and not isinstance(allowed_pattern, str):
            raise TypeError("Expected argument 'allowed_pattern' to be a str")
        pulumi.set(__self__, "allowed_pattern", allowed_pattern)
        if data_type and not isinstance(data_type, str):
            raise TypeError("Expected argument 'data_type' to be a str")
        pulumi.set(__self__, "data_type", data_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policies and not isinstance(policies, str):
            raise TypeError("Expected argument 'policies' to be a str")
        pulumi.set(__self__, "policies", policies)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if tier and not isinstance(tier, str):
            raise TypeError("Expected argument 'tier' to be a str")
        pulumi.set(__self__, "tier", tier)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="allowedPattern")
    def allowed_pattern(self) -> Optional[str]:
        return pulumi.get(self, "allowed_pattern")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional[str]:
        return pulumi.get(self, "data_type")

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
    def policies(self) -> Optional[str]:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def tier(self) -> Optional[str]:
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


class AwaitableGetParameterResult(GetParameterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetParameterResult(
            allowed_pattern=self.allowed_pattern,
            data_type=self.data_type,
            description=self.description,
            id=self.id,
            policies=self.policies,
            tags=self.tags,
            tier=self.tier,
            type=self.type,
            value=self.value)


def get_parameter(id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetParameterResult:
    """
    Resource Type definition for AWS::SSM::Parameter
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ssm:getParameter', __args__, opts=opts, typ=GetParameterResult).value

    return AwaitableGetParameterResult(
        allowed_pattern=__ret__.allowed_pattern,
        data_type=__ret__.data_type,
        description=__ret__.description,
        id=__ret__.id,
        policies=__ret__.policies,
        tags=__ret__.tags,
        tier=__ret__.tier,
        type=__ret__.type,
        value=__ret__.value)


@_utilities.lift_output_func(get_parameter)
def get_parameter_output(id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetParameterResult]:
    """
    Resource Type definition for AWS::SSM::Parameter
    """
    ...
