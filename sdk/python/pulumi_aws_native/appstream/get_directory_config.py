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
    'GetDirectoryConfigResult',
    'AwaitableGetDirectoryConfigResult',
    'get_directory_config',
    'get_directory_config_output',
]

@pulumi.output_type
class GetDirectoryConfigResult:
    def __init__(__self__, id=None, organizational_unit_distinguished_names=None, service_account_credentials=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organizational_unit_distinguished_names and not isinstance(organizational_unit_distinguished_names, list):
            raise TypeError("Expected argument 'organizational_unit_distinguished_names' to be a list")
        pulumi.set(__self__, "organizational_unit_distinguished_names", organizational_unit_distinguished_names)
        if service_account_credentials and not isinstance(service_account_credentials, dict):
            raise TypeError("Expected argument 'service_account_credentials' to be a dict")
        pulumi.set(__self__, "service_account_credentials", service_account_credentials)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="organizationalUnitDistinguishedNames")
    def organizational_unit_distinguished_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "organizational_unit_distinguished_names")

    @property
    @pulumi.getter(name="serviceAccountCredentials")
    def service_account_credentials(self) -> Optional['outputs.DirectoryConfigServiceAccountCredentials']:
        return pulumi.get(self, "service_account_credentials")


class AwaitableGetDirectoryConfigResult(GetDirectoryConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDirectoryConfigResult(
            id=self.id,
            organizational_unit_distinguished_names=self.organizational_unit_distinguished_names,
            service_account_credentials=self.service_account_credentials)


def get_directory_config(id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDirectoryConfigResult:
    """
    Resource Type definition for AWS::AppStream::DirectoryConfig
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:appstream:getDirectoryConfig', __args__, opts=opts, typ=GetDirectoryConfigResult).value

    return AwaitableGetDirectoryConfigResult(
        id=__ret__.id,
        organizational_unit_distinguished_names=__ret__.organizational_unit_distinguished_names,
        service_account_credentials=__ret__.service_account_credentials)


@_utilities.lift_output_func(get_directory_config)
def get_directory_config_output(id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDirectoryConfigResult]:
    """
    Resource Type definition for AWS::AppStream::DirectoryConfig
    """
    ...
