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
    'GetStoredQueryResult',
    'AwaitableGetStoredQueryResult',
    'get_stored_query',
    'get_stored_query_output',
]

@pulumi.output_type
class GetStoredQueryResult:
    def __init__(__self__, query_arn=None, query_description=None, query_expression=None, query_id=None, tags=None):
        if query_arn and not isinstance(query_arn, str):
            raise TypeError("Expected argument 'query_arn' to be a str")
        pulumi.set(__self__, "query_arn", query_arn)
        if query_description and not isinstance(query_description, str):
            raise TypeError("Expected argument 'query_description' to be a str")
        pulumi.set(__self__, "query_description", query_description)
        if query_expression and not isinstance(query_expression, str):
            raise TypeError("Expected argument 'query_expression' to be a str")
        pulumi.set(__self__, "query_expression", query_expression)
        if query_id and not isinstance(query_id, str):
            raise TypeError("Expected argument 'query_id' to be a str")
        pulumi.set(__self__, "query_id", query_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="queryArn")
    def query_arn(self) -> Optional[str]:
        return pulumi.get(self, "query_arn")

    @property
    @pulumi.getter(name="queryDescription")
    def query_description(self) -> Optional[str]:
        return pulumi.get(self, "query_description")

    @property
    @pulumi.getter(name="queryExpression")
    def query_expression(self) -> Optional[str]:
        return pulumi.get(self, "query_expression")

    @property
    @pulumi.getter(name="queryId")
    def query_id(self) -> Optional[str]:
        return pulumi.get(self, "query_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.StoredQueryTag']]:
        """
        The tags for the stored query.
        """
        return pulumi.get(self, "tags")


class AwaitableGetStoredQueryResult(GetStoredQueryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStoredQueryResult(
            query_arn=self.query_arn,
            query_description=self.query_description,
            query_expression=self.query_expression,
            query_id=self.query_id,
            tags=self.tags)


def get_stored_query(query_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStoredQueryResult:
    """
    Resource Type definition for AWS::Config::StoredQuery
    """
    __args__ = dict()
    __args__['queryName'] = query_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:configuration:getStoredQuery', __args__, opts=opts, typ=GetStoredQueryResult).value

    return AwaitableGetStoredQueryResult(
        query_arn=__ret__.query_arn,
        query_description=__ret__.query_description,
        query_expression=__ret__.query_expression,
        query_id=__ret__.query_id,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_stored_query)
def get_stored_query_output(query_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStoredQueryResult]:
    """
    Resource Type definition for AWS::Config::StoredQuery
    """
    ...