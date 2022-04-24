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
    'GetGeoMatchSetResult',
    'AwaitableGetGeoMatchSetResult',
    'get_geo_match_set',
    'get_geo_match_set_output',
]

@pulumi.output_type
class GetGeoMatchSetResult:
    def __init__(__self__, geo_match_constraints=None, id=None):
        if geo_match_constraints and not isinstance(geo_match_constraints, list):
            raise TypeError("Expected argument 'geo_match_constraints' to be a list")
        pulumi.set(__self__, "geo_match_constraints", geo_match_constraints)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="geoMatchConstraints")
    def geo_match_constraints(self) -> Optional[Sequence['outputs.GeoMatchSetGeoMatchConstraint']]:
        return pulumi.get(self, "geo_match_constraints")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetGeoMatchSetResult(GetGeoMatchSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGeoMatchSetResult(
            geo_match_constraints=self.geo_match_constraints,
            id=self.id)


def get_geo_match_set(id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGeoMatchSetResult:
    """
    Resource Type definition for AWS::WAFRegional::GeoMatchSet
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:wafregional:getGeoMatchSet', __args__, opts=opts, typ=GetGeoMatchSetResult).value

    return AwaitableGetGeoMatchSetResult(
        geo_match_constraints=__ret__.geo_match_constraints,
        id=__ret__.id)


@_utilities.lift_output_func(get_geo_match_set)
def get_geo_match_set_output(id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGeoMatchSetResult]:
    """
    Resource Type definition for AWS::WAFRegional::GeoMatchSet
    """
    ...
