# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetTrackerConsumerResult',
    'AwaitableGetTrackerConsumerResult',
    'get_tracker_consumer',
    'get_tracker_consumer_output',
]

@pulumi.output_type
class GetTrackerConsumerResult:
    def __init__(__self__):


class AwaitableGetTrackerConsumerResult(GetTrackerConsumerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrackerConsumerResult(
)


def get_tracker_consumer(consumer_arn: Optional[str] = None,
                         tracker_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrackerConsumerResult:
    """
    Definition of AWS::Location::TrackerConsumer Resource Type
    """
    __args__ = dict()
    __args__['consumerArn'] = consumer_arn
    __args__['trackerName'] = tracker_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:location:getTrackerConsumer', __args__, opts=opts, typ=GetTrackerConsumerResult).value

    return AwaitableGetTrackerConsumerResult(


@_utilities.lift_output_func(get_tracker_consumer)
def get_tracker_consumer_output(consumer_arn: Optional[pulumi.Input[str]] = None,
                                tracker_name: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTrackerConsumerResult]:
    """
    Definition of AWS::Location::TrackerConsumer Resource Type
    """
    ...