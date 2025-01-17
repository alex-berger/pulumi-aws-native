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
    'GetStreamResult',
    'AwaitableGetStreamResult',
    'get_stream',
    'get_stream_output',
]

@pulumi.output_type
class GetStreamResult:
    def __init__(__self__, arn=None, retention_period_hours=None, shard_count=None, stream_encryption=None, stream_mode_details=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if retention_period_hours and not isinstance(retention_period_hours, int):
            raise TypeError("Expected argument 'retention_period_hours' to be a int")
        pulumi.set(__self__, "retention_period_hours", retention_period_hours)
        if shard_count and not isinstance(shard_count, int):
            raise TypeError("Expected argument 'shard_count' to be a int")
        pulumi.set(__self__, "shard_count", shard_count)
        if stream_encryption and not isinstance(stream_encryption, dict):
            raise TypeError("Expected argument 'stream_encryption' to be a dict")
        pulumi.set(__self__, "stream_encryption", stream_encryption)
        if stream_mode_details and not isinstance(stream_mode_details, dict):
            raise TypeError("Expected argument 'stream_mode_details' to be a dict")
        pulumi.set(__self__, "stream_mode_details", stream_mode_details)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon resource name (ARN) of the Kinesis stream
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="retentionPeriodHours")
    def retention_period_hours(self) -> Optional[int]:
        """
        The number of hours for the data records that are stored in shards to remain accessible.
        """
        return pulumi.get(self, "retention_period_hours")

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> Optional[int]:
        """
        The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
        """
        return pulumi.get(self, "shard_count")

    @property
    @pulumi.getter(name="streamEncryption")
    def stream_encryption(self) -> Optional['outputs.StreamEncryption']:
        """
        When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
        """
        return pulumi.get(self, "stream_encryption")

    @property
    @pulumi.getter(name="streamModeDetails")
    def stream_mode_details(self) -> Optional['outputs.StreamModeDetails']:
        """
        The mode in which the stream is running.
        """
        return pulumi.get(self, "stream_mode_details")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.StreamTag']]:
        """
        An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
        """
        return pulumi.get(self, "tags")


class AwaitableGetStreamResult(GetStreamResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStreamResult(
            arn=self.arn,
            retention_period_hours=self.retention_period_hours,
            shard_count=self.shard_count,
            stream_encryption=self.stream_encryption,
            stream_mode_details=self.stream_mode_details,
            tags=self.tags)


def get_stream(name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStreamResult:
    """
    Resource Type definition for AWS::Kinesis::Stream


    :param str name: The name of the Kinesis stream.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:kinesis:getStream', __args__, opts=opts, typ=GetStreamResult).value

    return AwaitableGetStreamResult(
        arn=__ret__.arn,
        retention_period_hours=__ret__.retention_period_hours,
        shard_count=__ret__.shard_count,
        stream_encryption=__ret__.stream_encryption,
        stream_mode_details=__ret__.stream_mode_details,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_stream)
def get_stream_output(name: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStreamResult]:
    """
    Resource Type definition for AWS::Kinesis::Stream


    :param str name: The name of the Kinesis stream.
    """
    ...
