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
    'GetResolverQueryLoggingConfigAssociationResult',
    'AwaitableGetResolverQueryLoggingConfigAssociationResult',
    'get_resolver_query_logging_config_association',
    'get_resolver_query_logging_config_association_output',
]

@pulumi.output_type
class GetResolverQueryLoggingConfigAssociationResult:
    def __init__(__self__, creation_time=None, error=None, error_message=None, id=None, status=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if error and not isinstance(error, str):
            raise TypeError("Expected argument 'error' to be a str")
        pulumi.set(__self__, "error", error)
        if error_message and not isinstance(error_message, str):
            raise TypeError("Expected argument 'error_message' to be a str")
        pulumi.set(__self__, "error_message", error_message)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        """
        Rfc3339TimeString
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def error(self) -> Optional['ResolverQueryLoggingConfigAssociationError']:
        """
        ResolverQueryLogConfigAssociationError
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter(name="errorMessage")
    def error_message(self) -> Optional[str]:
        """
        ResolverQueryLogConfigAssociationErrorMessage
        """
        return pulumi.get(self, "error_message")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> Optional['ResolverQueryLoggingConfigAssociationStatus']:
        """
        ResolverQueryLogConfigAssociationStatus
        """
        return pulumi.get(self, "status")


class AwaitableGetResolverQueryLoggingConfigAssociationResult(GetResolverQueryLoggingConfigAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverQueryLoggingConfigAssociationResult(
            creation_time=self.creation_time,
            error=self.error,
            error_message=self.error_message,
            id=self.id,
            status=self.status)


def get_resolver_query_logging_config_association(id: Optional[str] = None,
                                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResolverQueryLoggingConfigAssociationResult:
    """
    Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.


    :param str id: Id
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:route53resolver:getResolverQueryLoggingConfigAssociation', __args__, opts=opts, typ=GetResolverQueryLoggingConfigAssociationResult).value

    return AwaitableGetResolverQueryLoggingConfigAssociationResult(
        creation_time=__ret__.creation_time,
        error=__ret__.error,
        error_message=__ret__.error_message,
        id=__ret__.id,
        status=__ret__.status)


@_utilities.lift_output_func(get_resolver_query_logging_config_association)
def get_resolver_query_logging_config_association_output(id: Optional[pulumi.Input[str]] = None,
                                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResolverQueryLoggingConfigAssociationResult]:
    """
    Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.


    :param str id: Id
    """
    ...
