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
    'GetFargateProfileResult',
    'AwaitableGetFargateProfileResult',
    'get_fargate_profile',
    'get_fargate_profile_output',
]

@pulumi.output_type
class GetFargateProfileResult:
    def __init__(__self__, arn=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.FargateProfileTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetFargateProfileResult(GetFargateProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFargateProfileResult(
            arn=self.arn,
            tags=self.tags)


def get_fargate_profile(cluster_name: Optional[str] = None,
                        fargate_profile_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFargateProfileResult:
    """
    Resource Schema for AWS::EKS::FargateProfile


    :param str cluster_name: Name of the Cluster
    :param str fargate_profile_name: Name of FargateProfile
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['fargateProfileName'] = fargate_profile_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:eks:getFargateProfile', __args__, opts=opts, typ=GetFargateProfileResult).value

    return AwaitableGetFargateProfileResult(
        arn=__ret__.arn,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_fargate_profile)
def get_fargate_profile_output(cluster_name: Optional[pulumi.Input[str]] = None,
                               fargate_profile_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFargateProfileResult]:
    """
    Resource Schema for AWS::EKS::FargateProfile


    :param str cluster_name: Name of the Cluster
    :param str fargate_profile_name: Name of FargateProfile
    """
    ...
