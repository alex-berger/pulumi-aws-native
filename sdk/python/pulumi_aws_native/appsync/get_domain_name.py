# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetDomainNameResult',
    'AwaitableGetDomainNameResult',
    'get_domain_name',
    'get_domain_name_output',
]

@pulumi.output_type
class GetDomainNameResult:
    def __init__(__self__, app_sync_domain_name=None, description=None, hosted_zone_id=None):
        if app_sync_domain_name and not isinstance(app_sync_domain_name, str):
            raise TypeError("Expected argument 'app_sync_domain_name' to be a str")
        pulumi.set(__self__, "app_sync_domain_name", app_sync_domain_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if hosted_zone_id and not isinstance(hosted_zone_id, str):
            raise TypeError("Expected argument 'hosted_zone_id' to be a str")
        pulumi.set(__self__, "hosted_zone_id", hosted_zone_id)

    @property
    @pulumi.getter(name="appSyncDomainName")
    def app_sync_domain_name(self) -> Optional[str]:
        return pulumi.get(self, "app_sync_domain_name")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> Optional[str]:
        return pulumi.get(self, "hosted_zone_id")


class AwaitableGetDomainNameResult(GetDomainNameResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainNameResult(
            app_sync_domain_name=self.app_sync_domain_name,
            description=self.description,
            hosted_zone_id=self.hosted_zone_id)


def get_domain_name(domain_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainNameResult:
    """
    Resource Type definition for AWS::AppSync::DomainName
    """
    __args__ = dict()
    __args__['domainName'] = domain_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:appsync:getDomainName', __args__, opts=opts, typ=GetDomainNameResult).value

    return AwaitableGetDomainNameResult(
        app_sync_domain_name=__ret__.app_sync_domain_name,
        description=__ret__.description,
        hosted_zone_id=__ret__.hosted_zone_id)


@_utilities.lift_output_func(get_domain_name)
def get_domain_name_output(domain_name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDomainNameResult]:
    """
    Resource Type definition for AWS::AppSync::DomainName
    """
    ...
