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
    'GetPlaybackKeyPairResult',
    'AwaitableGetPlaybackKeyPairResult',
    'get_playback_key_pair',
    'get_playback_key_pair_output',
]

@pulumi.output_type
class GetPlaybackKeyPairResult:
    def __init__(__self__, arn=None, fingerprint=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Key-pair identifier.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[str]:
        """
        Key-pair identifier.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.PlaybackKeyPairTag']]:
        """
        A list of key-value pairs that contain metadata for the asset model.
        """
        return pulumi.get(self, "tags")


class AwaitableGetPlaybackKeyPairResult(GetPlaybackKeyPairResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPlaybackKeyPairResult(
            arn=self.arn,
            fingerprint=self.fingerprint,
            tags=self.tags)


def get_playback_key_pair(arn: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPlaybackKeyPairResult:
    """
    Resource Type definition for AWS::IVS::PlaybackKeyPair


    :param str arn: Key-pair identifier.
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:ivs:getPlaybackKeyPair', __args__, opts=opts, typ=GetPlaybackKeyPairResult).value

    return AwaitableGetPlaybackKeyPairResult(
        arn=__ret__.arn,
        fingerprint=__ret__.fingerprint,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_playback_key_pair)
def get_playback_key_pair_output(arn: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPlaybackKeyPairResult]:
    """
    Resource Type definition for AWS::IVS::PlaybackKeyPair


    :param str arn: Key-pair identifier.
    """
    ...
