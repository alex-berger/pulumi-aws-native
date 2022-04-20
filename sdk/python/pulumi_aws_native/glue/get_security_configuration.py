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
    'GetSecurityConfigurationResult',
    'AwaitableGetSecurityConfigurationResult',
    'get_security_configuration',
    'get_security_configuration_output',
]

@pulumi.output_type
class GetSecurityConfigurationResult:
    def __init__(__self__, encryption_configuration=None, id=None):
        if encryption_configuration and not isinstance(encryption_configuration, dict):
            raise TypeError("Expected argument 'encryption_configuration' to be a dict")
        pulumi.set(__self__, "encryption_configuration", encryption_configuration)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="encryptionConfiguration")
    def encryption_configuration(self) -> Optional['outputs.SecurityConfigurationEncryptionConfiguration']:
        return pulumi.get(self, "encryption_configuration")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetSecurityConfigurationResult(GetSecurityConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityConfigurationResult(
            encryption_configuration=self.encryption_configuration,
            id=self.id)


def get_security_configuration(id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityConfigurationResult:
    """
    Resource Type definition for AWS::Glue::SecurityConfiguration
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:glue:getSecurityConfiguration', __args__, opts=opts, typ=GetSecurityConfigurationResult).value

    return AwaitableGetSecurityConfigurationResult(
        encryption_configuration=__ret__.encryption_configuration,
        id=__ret__.id)


@_utilities.lift_output_func(get_security_configuration)
def get_security_configuration_output(id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityConfigurationResult]:
    """
    Resource Type definition for AWS::Glue::SecurityConfiguration
    """
    ...
