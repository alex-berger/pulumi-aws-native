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
    'GetResponseHeadersPolicyResult',
    'AwaitableGetResponseHeadersPolicyResult',
    'get_response_headers_policy',
    'get_response_headers_policy_output',
]

@pulumi.output_type
class GetResponseHeadersPolicyResult:
    def __init__(__self__, id=None, last_modified_time=None, response_headers_policy_config=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if response_headers_policy_config and not isinstance(response_headers_policy_config, dict):
            raise TypeError("Expected argument 'response_headers_policy_config' to be a dict")
        pulumi.set(__self__, "response_headers_policy_config", response_headers_policy_config)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> Optional[str]:
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="responseHeadersPolicyConfig")
    def response_headers_policy_config(self) -> Optional['outputs.ResponseHeadersPolicyConfig']:
        return pulumi.get(self, "response_headers_policy_config")


class AwaitableGetResponseHeadersPolicyResult(GetResponseHeadersPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResponseHeadersPolicyResult(
            id=self.id,
            last_modified_time=self.last_modified_time,
            response_headers_policy_config=self.response_headers_policy_config)


def get_response_headers_policy(id: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResponseHeadersPolicyResult:
    """
    Resource Type definition for AWS::CloudFront::ResponseHeadersPolicy
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:cloudfront:getResponseHeadersPolicy', __args__, opts=opts, typ=GetResponseHeadersPolicyResult).value

    return AwaitableGetResponseHeadersPolicyResult(
        id=__ret__.id,
        last_modified_time=__ret__.last_modified_time,
        response_headers_policy_config=__ret__.response_headers_policy_config)


@_utilities.lift_output_func(get_response_headers_policy)
def get_response_headers_policy_output(id: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResponseHeadersPolicyResult]:
    """
    Resource Type definition for AWS::CloudFront::ResponseHeadersPolicy
    """
    ...
