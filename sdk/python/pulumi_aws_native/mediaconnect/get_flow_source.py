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
    'GetFlowSourceResult',
    'AwaitableGetFlowSourceResult',
    'get_flow_source',
    'get_flow_source_output',
]

@pulumi.output_type
class GetFlowSourceResult:
    def __init__(__self__, decryption=None, description=None, entitlement_arn=None, flow_arn=None, ingest_ip=None, ingest_port=None, max_bitrate=None, max_latency=None, protocol=None, source_arn=None, source_ingest_port=None, stream_id=None, vpc_interface_name=None, whitelist_cidr=None):
        if decryption and not isinstance(decryption, dict):
            raise TypeError("Expected argument 'decryption' to be a dict")
        pulumi.set(__self__, "decryption", decryption)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if entitlement_arn and not isinstance(entitlement_arn, str):
            raise TypeError("Expected argument 'entitlement_arn' to be a str")
        pulumi.set(__self__, "entitlement_arn", entitlement_arn)
        if flow_arn and not isinstance(flow_arn, str):
            raise TypeError("Expected argument 'flow_arn' to be a str")
        pulumi.set(__self__, "flow_arn", flow_arn)
        if ingest_ip and not isinstance(ingest_ip, str):
            raise TypeError("Expected argument 'ingest_ip' to be a str")
        pulumi.set(__self__, "ingest_ip", ingest_ip)
        if ingest_port and not isinstance(ingest_port, int):
            raise TypeError("Expected argument 'ingest_port' to be a int")
        pulumi.set(__self__, "ingest_port", ingest_port)
        if max_bitrate and not isinstance(max_bitrate, int):
            raise TypeError("Expected argument 'max_bitrate' to be a int")
        pulumi.set(__self__, "max_bitrate", max_bitrate)
        if max_latency and not isinstance(max_latency, int):
            raise TypeError("Expected argument 'max_latency' to be a int")
        pulumi.set(__self__, "max_latency", max_latency)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if source_arn and not isinstance(source_arn, str):
            raise TypeError("Expected argument 'source_arn' to be a str")
        pulumi.set(__self__, "source_arn", source_arn)
        if source_ingest_port and not isinstance(source_ingest_port, str):
            raise TypeError("Expected argument 'source_ingest_port' to be a str")
        pulumi.set(__self__, "source_ingest_port", source_ingest_port)
        if stream_id and not isinstance(stream_id, str):
            raise TypeError("Expected argument 'stream_id' to be a str")
        pulumi.set(__self__, "stream_id", stream_id)
        if vpc_interface_name and not isinstance(vpc_interface_name, str):
            raise TypeError("Expected argument 'vpc_interface_name' to be a str")
        pulumi.set(__self__, "vpc_interface_name", vpc_interface_name)
        if whitelist_cidr and not isinstance(whitelist_cidr, str):
            raise TypeError("Expected argument 'whitelist_cidr' to be a str")
        pulumi.set(__self__, "whitelist_cidr", whitelist_cidr)

    @property
    @pulumi.getter
    def decryption(self) -> Optional['outputs.FlowSourceEncryption']:
        """
        The type of encryption that is used on the content ingested from this source.
        """
        return pulumi.get(self, "decryption")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="entitlementArn")
    def entitlement_arn(self) -> Optional[str]:
        """
        The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
        """
        return pulumi.get(self, "entitlement_arn")

    @property
    @pulumi.getter(name="flowArn")
    def flow_arn(self) -> Optional[str]:
        """
        The ARN of the flow.
        """
        return pulumi.get(self, "flow_arn")

    @property
    @pulumi.getter(name="ingestIp")
    def ingest_ip(self) -> Optional[str]:
        """
        The IP address that the flow will be listening on for incoming content.
        """
        return pulumi.get(self, "ingest_ip")

    @property
    @pulumi.getter(name="ingestPort")
    def ingest_port(self) -> Optional[int]:
        """
        The port that the flow will be listening on for incoming content.
        """
        return pulumi.get(self, "ingest_port")

    @property
    @pulumi.getter(name="maxBitrate")
    def max_bitrate(self) -> Optional[int]:
        """
        The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
        """
        return pulumi.get(self, "max_bitrate")

    @property
    @pulumi.getter(name="maxLatency")
    def max_latency(self) -> Optional[int]:
        """
        The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
        """
        return pulumi.get(self, "max_latency")

    @property
    @pulumi.getter
    def protocol(self) -> Optional['FlowSourceProtocol']:
        """
        The protocol that is used by the source.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="sourceArn")
    def source_arn(self) -> Optional[str]:
        """
        The ARN of the source.
        """
        return pulumi.get(self, "source_arn")

    @property
    @pulumi.getter(name="sourceIngestPort")
    def source_ingest_port(self) -> Optional[str]:
        """
        The port that the flow will be listening on for incoming content.(ReadOnly)
        """
        return pulumi.get(self, "source_ingest_port")

    @property
    @pulumi.getter(name="streamId")
    def stream_id(self) -> Optional[str]:
        """
        The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
        """
        return pulumi.get(self, "stream_id")

    @property
    @pulumi.getter(name="vpcInterfaceName")
    def vpc_interface_name(self) -> Optional[str]:
        """
        The name of the VPC Interface this Source is configured with.
        """
        return pulumi.get(self, "vpc_interface_name")

    @property
    @pulumi.getter(name="whitelistCidr")
    def whitelist_cidr(self) -> Optional[str]:
        """
        The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
        """
        return pulumi.get(self, "whitelist_cidr")


class AwaitableGetFlowSourceResult(GetFlowSourceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlowSourceResult(
            decryption=self.decryption,
            description=self.description,
            entitlement_arn=self.entitlement_arn,
            flow_arn=self.flow_arn,
            ingest_ip=self.ingest_ip,
            ingest_port=self.ingest_port,
            max_bitrate=self.max_bitrate,
            max_latency=self.max_latency,
            protocol=self.protocol,
            source_arn=self.source_arn,
            source_ingest_port=self.source_ingest_port,
            stream_id=self.stream_id,
            vpc_interface_name=self.vpc_interface_name,
            whitelist_cidr=self.whitelist_cidr)


def get_flow_source(source_arn: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlowSourceResult:
    """
    Resource schema for AWS::MediaConnect::FlowSource


    :param str source_arn: The ARN of the source.
    """
    __args__ = dict()
    __args__['sourceArn'] = source_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:mediaconnect:getFlowSource', __args__, opts=opts, typ=GetFlowSourceResult).value

    return AwaitableGetFlowSourceResult(
        decryption=__ret__.decryption,
        description=__ret__.description,
        entitlement_arn=__ret__.entitlement_arn,
        flow_arn=__ret__.flow_arn,
        ingest_ip=__ret__.ingest_ip,
        ingest_port=__ret__.ingest_port,
        max_bitrate=__ret__.max_bitrate,
        max_latency=__ret__.max_latency,
        protocol=__ret__.protocol,
        source_arn=__ret__.source_arn,
        source_ingest_port=__ret__.source_ingest_port,
        stream_id=__ret__.stream_id,
        vpc_interface_name=__ret__.vpc_interface_name,
        whitelist_cidr=__ret__.whitelist_cidr)


@_utilities.lift_output_func(get_flow_source)
def get_flow_source_output(source_arn: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlowSourceResult]:
    """
    Resource schema for AWS::MediaConnect::FlowSource


    :param str source_arn: The ARN of the source.
    """
    ...
