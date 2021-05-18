# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'CidrResult',
    'AwaitableCidrResult',
    'cidr',
]

@pulumi.output_type
class CidrResult:
    def __init__(__self__, subnets=None):
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        pulumi.set(__self__, "subnets", subnets)

    @property
    @pulumi.getter
    def subnets(self) -> Sequence[str]:
        return pulumi.get(self, "subnets")


class AwaitableCidrResult(CidrResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return CidrResult(
            subnets=self.subnets)


def cidr(cidr_bits: Optional[int] = None,
         count: Optional[int] = None,
         ip_block: Optional[str] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableCidrResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['cidrBits'] = cidr_bits
    __args__['count'] = count
    __args__['ipBlock'] = ip_block
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:index:cidr', __args__, opts=opts, typ=CidrResult).value

    return AwaitableCidrResult(
        subnets=__ret__.subnets)
