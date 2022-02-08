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
    'GetTrafficMirrorSessionResult',
    'AwaitableGetTrafficMirrorSessionResult',
    'get_traffic_mirror_session',
    'get_traffic_mirror_session_output',
]

@pulumi.output_type
class GetTrafficMirrorSessionResult:
    def __init__(__self__, description=None, id=None, packet_length=None, session_number=None, tags=None, traffic_mirror_filter_id=None, traffic_mirror_target_id=None, virtual_network_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if packet_length and not isinstance(packet_length, int):
            raise TypeError("Expected argument 'packet_length' to be a int")
        pulumi.set(__self__, "packet_length", packet_length)
        if session_number and not isinstance(session_number, int):
            raise TypeError("Expected argument 'session_number' to be a int")
        pulumi.set(__self__, "session_number", session_number)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if traffic_mirror_filter_id and not isinstance(traffic_mirror_filter_id, str):
            raise TypeError("Expected argument 'traffic_mirror_filter_id' to be a str")
        pulumi.set(__self__, "traffic_mirror_filter_id", traffic_mirror_filter_id)
        if traffic_mirror_target_id and not isinstance(traffic_mirror_target_id, str):
            raise TypeError("Expected argument 'traffic_mirror_target_id' to be a str")
        pulumi.set(__self__, "traffic_mirror_target_id", traffic_mirror_target_id)
        if virtual_network_id and not isinstance(virtual_network_id, int):
            raise TypeError("Expected argument 'virtual_network_id' to be a int")
        pulumi.set(__self__, "virtual_network_id", virtual_network_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="packetLength")
    def packet_length(self) -> Optional[int]:
        return pulumi.get(self, "packet_length")

    @property
    @pulumi.getter(name="sessionNumber")
    def session_number(self) -> Optional[int]:
        return pulumi.get(self, "session_number")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.TrafficMirrorSessionTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> Optional[str]:
        return pulumi.get(self, "traffic_mirror_filter_id")

    @property
    @pulumi.getter(name="trafficMirrorTargetId")
    def traffic_mirror_target_id(self) -> Optional[str]:
        return pulumi.get(self, "traffic_mirror_target_id")

    @property
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> Optional[int]:
        return pulumi.get(self, "virtual_network_id")


class AwaitableGetTrafficMirrorSessionResult(GetTrafficMirrorSessionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTrafficMirrorSessionResult(
            description=self.description,
            id=self.id,
            packet_length=self.packet_length,
            session_number=self.session_number,
            tags=self.tags,
            traffic_mirror_filter_id=self.traffic_mirror_filter_id,
            traffic_mirror_target_id=self.traffic_mirror_target_id,
            virtual_network_id=self.virtual_network_id)


def get_traffic_mirror_session(id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTrafficMirrorSessionResult:
    """
    Resource Type definition for AWS::EC2::TrafficMirrorSession
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ec2:getTrafficMirrorSession', __args__, opts=opts, typ=GetTrafficMirrorSessionResult).value

    return AwaitableGetTrafficMirrorSessionResult(
        description=__ret__.description,
        id=__ret__.id,
        packet_length=__ret__.packet_length,
        session_number=__ret__.session_number,
        tags=__ret__.tags,
        traffic_mirror_filter_id=__ret__.traffic_mirror_filter_id,
        traffic_mirror_target_id=__ret__.traffic_mirror_target_id,
        virtual_network_id=__ret__.virtual_network_id)


@_utilities.lift_output_func(get_traffic_mirror_session)
def get_traffic_mirror_session_output(id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTrafficMirrorSessionResult]:
    """
    Resource Type definition for AWS::EC2::TrafficMirrorSession
    """
    ...