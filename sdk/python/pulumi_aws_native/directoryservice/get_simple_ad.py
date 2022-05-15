# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSimpleADResult',
    'AwaitableGetSimpleADResult',
    'get_simple_ad',
    'get_simple_ad_output',
]

@pulumi.output_type
class GetSimpleADResult:
    def __init__(__self__, alias=None, dns_ip_addresses=None, enable_sso=None, id=None):
        if alias and not isinstance(alias, str):
            raise TypeError("Expected argument 'alias' to be a str")
        pulumi.set(__self__, "alias", alias)
        if dns_ip_addresses and not isinstance(dns_ip_addresses, list):
            raise TypeError("Expected argument 'dns_ip_addresses' to be a list")
        pulumi.set(__self__, "dns_ip_addresses", dns_ip_addresses)
        if enable_sso and not isinstance(enable_sso, bool):
            raise TypeError("Expected argument 'enable_sso' to be a bool")
        pulumi.set(__self__, "enable_sso", enable_sso)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def alias(self) -> Optional[str]:
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="dnsIpAddresses")
    def dns_ip_addresses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "dns_ip_addresses")

    @property
    @pulumi.getter(name="enableSso")
    def enable_sso(self) -> Optional[bool]:
        return pulumi.get(self, "enable_sso")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetSimpleADResult(GetSimpleADResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSimpleADResult(
            alias=self.alias,
            dns_ip_addresses=self.dns_ip_addresses,
            enable_sso=self.enable_sso,
            id=self.id)


def get_simple_ad(id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSimpleADResult:
    """
    Resource Type definition for AWS::DirectoryService::SimpleAD
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:directoryservice:getSimpleAD', __args__, opts=opts, typ=GetSimpleADResult).value

    return AwaitableGetSimpleADResult(
        alias=__ret__.alias,
        dns_ip_addresses=__ret__.dns_ip_addresses,
        enable_sso=__ret__.enable_sso,
        id=__ret__.id)


@_utilities.lift_output_func(get_simple_ad)
def get_simple_ad_output(id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSimpleADResult]:
    """
    Resource Type definition for AWS::DirectoryService::SimpleAD
    """
    ...
