# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDNSSECResult',
    'AwaitableGetDNSSECResult',
    'get_dnssec',
    'get_dnssec_output',
]

@pulumi.output_type
class GetDNSSECResult:
    def __init__(__self__):


class AwaitableGetDNSSECResult(GetDNSSECResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDNSSECResult(
)


def get_dnssec(hosted_zone_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDNSSECResult:
    """
    Resource used to control (enable/disable) DNSSEC in a specific hosted zone.


    :param str hosted_zone_id: The unique string (ID) used to identify a hosted zone.
    """
    __args__ = dict()
    __args__['hostedZoneId'] = hosted_zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:route53:getDNSSEC', __args__, opts=opts, typ=GetDNSSECResult).value

    return AwaitableGetDNSSECResult(


@_utilities.lift_output_func(get_dnssec)
def get_dnssec_output(hosted_zone_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDNSSECResult]:
    """
    Resource used to control (enable/disable) DNSSEC in a specific hosted zone.


    :param str hosted_zone_id: The unique string (ID) used to identify a hosted zone.
    """
    ...