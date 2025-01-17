# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'GetResolverDNSSECConfigResult',
    'AwaitableGetResolverDNSSECConfigResult',
    'get_resolver_dnssec_config',
    'get_resolver_dnssec_config_output',
]

@pulumi.output_type
class GetResolverDNSSECConfigResult:
    def __init__(__self__, id=None, owner_id=None, validation_status=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if validation_status and not isinstance(validation_status, str):
            raise TypeError("Expected argument 'validation_status' to be a str")
        pulumi.set(__self__, "validation_status", validation_status)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[str]:
        """
        AccountId
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="validationStatus")
    def validation_status(self) -> Optional['ResolverDNSSECConfigValidationStatus']:
        """
        ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
        """
        return pulumi.get(self, "validation_status")


class AwaitableGetResolverDNSSECConfigResult(GetResolverDNSSECConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverDNSSECConfigResult(
            id=self.id,
            owner_id=self.owner_id,
            validation_status=self.validation_status)


def get_resolver_dnssec_config(id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResolverDNSSECConfigResult:
    """
    Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.


    :param str id: Id
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:route53resolver:getResolverDNSSECConfig', __args__, opts=opts, typ=GetResolverDNSSECConfigResult).value

    return AwaitableGetResolverDNSSECConfigResult(
        id=__ret__.id,
        owner_id=__ret__.owner_id,
        validation_status=__ret__.validation_status)


@_utilities.lift_output_func(get_resolver_dnssec_config)
def get_resolver_dnssec_config_output(id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResolverDNSSECConfigResult]:
    """
    Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.


    :param str id: Id
    """
    ...
