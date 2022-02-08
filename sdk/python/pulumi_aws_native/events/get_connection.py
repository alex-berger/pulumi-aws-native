# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetConnectionResult',
    'AwaitableGetConnectionResult',
    'get_connection',
    'get_connection_output',
]

@pulumi.output_type
class GetConnectionResult:
    def __init__(__self__, arn=None, auth_parameters=None, authorization_type=None, description=None, secret_arn=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if auth_parameters and not isinstance(auth_parameters, dict):
            raise TypeError("Expected argument 'auth_parameters' to be a dict")
        pulumi.set(__self__, "auth_parameters", auth_parameters)
        if authorization_type and not isinstance(authorization_type, str):
            raise TypeError("Expected argument 'authorization_type' to be a str")
        pulumi.set(__self__, "authorization_type", authorization_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if secret_arn and not isinstance(secret_arn, str):
            raise TypeError("Expected argument 'secret_arn' to be a str")
        pulumi.set(__self__, "secret_arn", secret_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The arn of the connection resource.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authParameters")
    def auth_parameters(self) -> Optional['outputs.AuthParametersProperties']:
        return pulumi.get(self, "auth_parameters")

    @property
    @pulumi.getter(name="authorizationType")
    def authorization_type(self) -> Optional['ConnectionAuthorizationType']:
        return pulumi.get(self, "authorization_type")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the connection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="secretArn")
    def secret_arn(self) -> Optional[str]:
        """
        The arn of the secrets manager secret created in the customer account.
        """
        return pulumi.get(self, "secret_arn")


class AwaitableGetConnectionResult(GetConnectionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionResult(
            arn=self.arn,
            auth_parameters=self.auth_parameters,
            authorization_type=self.authorization_type,
            description=self.description,
            secret_arn=self.secret_arn)


def get_connection(name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectionResult:
    """
    Resource Type definition for AWS::Events::Connection.


    :param str name: Name of the connection.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:events:getConnection', __args__, opts=opts, typ=GetConnectionResult).value

    return AwaitableGetConnectionResult(
        arn=__ret__.arn,
        auth_parameters=__ret__.auth_parameters,
        authorization_type=__ret__.authorization_type,
        description=__ret__.description,
        secret_arn=__ret__.secret_arn)


@_utilities.lift_output_func(get_connection)
def get_connection_output(name: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConnectionResult]:
    """
    Resource Type definition for AWS::Events::Connection.


    :param str name: Name of the connection.
    """
    ...