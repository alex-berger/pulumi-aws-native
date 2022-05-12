# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetUsagePlanKeyResult',
    'AwaitableGetUsagePlanKeyResult',
    'get_usage_plan_key',
    'get_usage_plan_key_output',
]

@pulumi.output_type
class GetUsagePlanKeyResult:
    def __init__(__self__, id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
        """
        return pulumi.get(self, "id")


class AwaitableGetUsagePlanKeyResult(GetUsagePlanKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsagePlanKeyResult(
            id=self.id)


def get_usage_plan_key(id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsagePlanKeyResult:
    """
    Resource Type definition for AWS::ApiGateway::UsagePlanKey


    :param str id: An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:apigateway:getUsagePlanKey', __args__, opts=opts, typ=GetUsagePlanKeyResult).value

    return AwaitableGetUsagePlanKeyResult(
        id=__ret__.id)


@_utilities.lift_output_func(get_usage_plan_key)
def get_usage_plan_key_output(id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUsagePlanKeyResult]:
    """
    Resource Type definition for AWS::ApiGateway::UsagePlanKey


    :param str id: An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
    """
    ...
