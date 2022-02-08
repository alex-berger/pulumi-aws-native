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
    'GetResolverQueryLoggingConfigResult',
    'AwaitableGetResolverQueryLoggingConfigResult',
    'get_resolver_query_logging_config',
    'get_resolver_query_logging_config_output',
]

@pulumi.output_type
class GetResolverQueryLoggingConfigResult:
    def __init__(__self__, arn=None, association_count=None, creation_time=None, creator_request_id=None, id=None, owner_id=None, share_status=None, status=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if association_count and not isinstance(association_count, int):
            raise TypeError("Expected argument 'association_count' to be a int")
        pulumi.set(__self__, "association_count", association_count)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if creator_request_id and not isinstance(creator_request_id, str):
            raise TypeError("Expected argument 'creator_request_id' to be a str")
        pulumi.set(__self__, "creator_request_id", creator_request_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if share_status and not isinstance(share_status, str):
            raise TypeError("Expected argument 'share_status' to be a str")
        pulumi.set(__self__, "share_status", share_status)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Arn
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="associationCount")
    def association_count(self) -> Optional[int]:
        """
        Count
        """
        return pulumi.get(self, "association_count")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[str]:
        """
        Rfc3339TimeString
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="creatorRequestId")
    def creator_request_id(self) -> Optional[str]:
        """
        The id of the creator request.
        """
        return pulumi.get(self, "creator_request_id")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        ResourceId
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
    @pulumi.getter(name="shareStatus")
    def share_status(self) -> Optional['ResolverQueryLoggingConfigShareStatus']:
        """
        ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.
        """
        return pulumi.get(self, "share_status")

    @property
    @pulumi.getter
    def status(self) -> Optional['ResolverQueryLoggingConfigStatus']:
        """
        ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.
        """
        return pulumi.get(self, "status")


class AwaitableGetResolverQueryLoggingConfigResult(GetResolverQueryLoggingConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverQueryLoggingConfigResult(
            arn=self.arn,
            association_count=self.association_count,
            creation_time=self.creation_time,
            creator_request_id=self.creator_request_id,
            id=self.id,
            owner_id=self.owner_id,
            share_status=self.share_status,
            status=self.status)


def get_resolver_query_logging_config(id: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResolverQueryLoggingConfigResult:
    """
    Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.


    :param str id: ResourceId
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:route53resolver:getResolverQueryLoggingConfig', __args__, opts=opts, typ=GetResolverQueryLoggingConfigResult).value

    return AwaitableGetResolverQueryLoggingConfigResult(
        arn=__ret__.arn,
        association_count=__ret__.association_count,
        creation_time=__ret__.creation_time,
        creator_request_id=__ret__.creator_request_id,
        id=__ret__.id,
        owner_id=__ret__.owner_id,
        share_status=__ret__.share_status,
        status=__ret__.status)


@_utilities.lift_output_func(get_resolver_query_logging_config)
def get_resolver_query_logging_config_output(id: Optional[pulumi.Input[str]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResolverQueryLoggingConfigResult]:
    """
    Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.


    :param str id: ResourceId
    """
    ...