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
    'GetGameSessionQueueResult',
    'AwaitableGetGameSessionQueueResult',
    'get_game_session_queue',
    'get_game_session_queue_output',
]

@pulumi.output_type
class GetGameSessionQueueResult:
    def __init__(__self__, arn=None, custom_event_data=None, destinations=None, filter_configuration=None, id=None, notification_target=None, player_latency_policies=None, priority_configuration=None, tags=None, timeout_in_seconds=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if custom_event_data and not isinstance(custom_event_data, str):
            raise TypeError("Expected argument 'custom_event_data' to be a str")
        pulumi.set(__self__, "custom_event_data", custom_event_data)
        if destinations and not isinstance(destinations, list):
            raise TypeError("Expected argument 'destinations' to be a list")
        pulumi.set(__self__, "destinations", destinations)
        if filter_configuration and not isinstance(filter_configuration, dict):
            raise TypeError("Expected argument 'filter_configuration' to be a dict")
        pulumi.set(__self__, "filter_configuration", filter_configuration)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if notification_target and not isinstance(notification_target, str):
            raise TypeError("Expected argument 'notification_target' to be a str")
        pulumi.set(__self__, "notification_target", notification_target)
        if player_latency_policies and not isinstance(player_latency_policies, list):
            raise TypeError("Expected argument 'player_latency_policies' to be a list")
        pulumi.set(__self__, "player_latency_policies", player_latency_policies)
        if priority_configuration and not isinstance(priority_configuration, dict):
            raise TypeError("Expected argument 'priority_configuration' to be a dict")
        pulumi.set(__self__, "priority_configuration", priority_configuration)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if timeout_in_seconds and not isinstance(timeout_in_seconds, int):
            raise TypeError("Expected argument 'timeout_in_seconds' to be a int")
        pulumi.set(__self__, "timeout_in_seconds", timeout_in_seconds)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="customEventData")
    def custom_event_data(self) -> Optional[str]:
        return pulumi.get(self, "custom_event_data")

    @property
    @pulumi.getter
    def destinations(self) -> Optional[Sequence['outputs.GameSessionQueueDestination']]:
        return pulumi.get(self, "destinations")

    @property
    @pulumi.getter(name="filterConfiguration")
    def filter_configuration(self) -> Optional['outputs.GameSessionQueueFilterConfiguration']:
        return pulumi.get(self, "filter_configuration")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="notificationTarget")
    def notification_target(self) -> Optional[str]:
        return pulumi.get(self, "notification_target")

    @property
    @pulumi.getter(name="playerLatencyPolicies")
    def player_latency_policies(self) -> Optional[Sequence['outputs.GameSessionQueuePlayerLatencyPolicy']]:
        return pulumi.get(self, "player_latency_policies")

    @property
    @pulumi.getter(name="priorityConfiguration")
    def priority_configuration(self) -> Optional['outputs.GameSessionQueuePriorityConfiguration']:
        return pulumi.get(self, "priority_configuration")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.GameSessionQueueTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeoutInSeconds")
    def timeout_in_seconds(self) -> Optional[int]:
        return pulumi.get(self, "timeout_in_seconds")


class AwaitableGetGameSessionQueueResult(GetGameSessionQueueResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGameSessionQueueResult(
            arn=self.arn,
            custom_event_data=self.custom_event_data,
            destinations=self.destinations,
            filter_configuration=self.filter_configuration,
            id=self.id,
            notification_target=self.notification_target,
            player_latency_policies=self.player_latency_policies,
            priority_configuration=self.priority_configuration,
            tags=self.tags,
            timeout_in_seconds=self.timeout_in_seconds)


def get_game_session_queue(id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGameSessionQueueResult:
    """
    Resource Type definition for AWS::GameLift::GameSessionQueue
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:gamelift:getGameSessionQueue', __args__, opts=opts, typ=GetGameSessionQueueResult).value

    return AwaitableGetGameSessionQueueResult(
        arn=__ret__.arn,
        custom_event_data=__ret__.custom_event_data,
        destinations=__ret__.destinations,
        filter_configuration=__ret__.filter_configuration,
        id=__ret__.id,
        notification_target=__ret__.notification_target,
        player_latency_policies=__ret__.player_latency_policies,
        priority_configuration=__ret__.priority_configuration,
        tags=__ret__.tags,
        timeout_in_seconds=__ret__.timeout_in_seconds)


@_utilities.lift_output_func(get_game_session_queue)
def get_game_session_queue_output(id: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGameSessionQueueResult]:
    """
    Resource Type definition for AWS::GameLift::GameSessionQueue
    """
    ...
