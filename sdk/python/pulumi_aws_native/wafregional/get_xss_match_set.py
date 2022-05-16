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
    'GetXssMatchSetResult',
    'AwaitableGetXssMatchSetResult',
    'get_xss_match_set',
    'get_xss_match_set_output',
]

@pulumi.output_type
class GetXssMatchSetResult:
    def __init__(__self__, id=None, xss_match_tuples=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if xss_match_tuples and not isinstance(xss_match_tuples, list):
            raise TypeError("Expected argument 'xss_match_tuples' to be a list")
        pulumi.set(__self__, "xss_match_tuples", xss_match_tuples)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="xssMatchTuples")
    def xss_match_tuples(self) -> Optional[Sequence['outputs.XssMatchSetXssMatchTuple']]:
        return pulumi.get(self, "xss_match_tuples")


class AwaitableGetXssMatchSetResult(GetXssMatchSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetXssMatchSetResult(
            id=self.id,
            xss_match_tuples=self.xss_match_tuples)


def get_xss_match_set(id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetXssMatchSetResult:
    """
    Resource Type definition for AWS::WAFRegional::XssMatchSet
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:wafregional:getXssMatchSet', __args__, opts=opts, typ=GetXssMatchSetResult).value

    return AwaitableGetXssMatchSetResult(
        id=__ret__.id,
        xss_match_tuples=__ret__.xss_match_tuples)


@_utilities.lift_output_func(get_xss_match_set)
def get_xss_match_set_output(id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetXssMatchSetResult]:
    """
    Resource Type definition for AWS::WAFRegional::XssMatchSet
    """
    ...