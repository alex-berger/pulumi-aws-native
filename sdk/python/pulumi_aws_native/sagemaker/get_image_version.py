# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetImageVersionResult',
    'AwaitableGetImageVersionResult',
    'get_image_version',
    'get_image_version_output',
]

@pulumi.output_type
class GetImageVersionResult:
    def __init__(__self__, container_image=None, image_arn=None, image_version_arn=None, version=None):
        if container_image and not isinstance(container_image, str):
            raise TypeError("Expected argument 'container_image' to be a str")
        pulumi.set(__self__, "container_image", container_image)
        if image_arn and not isinstance(image_arn, str):
            raise TypeError("Expected argument 'image_arn' to be a str")
        pulumi.set(__self__, "image_arn", image_arn)
        if image_version_arn and not isinstance(image_version_arn, str):
            raise TypeError("Expected argument 'image_version_arn' to be a str")
        pulumi.set(__self__, "image_version_arn", image_version_arn)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="containerImage")
    def container_image(self) -> Optional[str]:
        return pulumi.get(self, "container_image")

    @property
    @pulumi.getter(name="imageArn")
    def image_arn(self) -> Optional[str]:
        return pulumi.get(self, "image_arn")

    @property
    @pulumi.getter(name="imageVersionArn")
    def image_version_arn(self) -> Optional[str]:
        return pulumi.get(self, "image_version_arn")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        return pulumi.get(self, "version")


class AwaitableGetImageVersionResult(GetImageVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageVersionResult(
            container_image=self.container_image,
            image_arn=self.image_arn,
            image_version_arn=self.image_version_arn,
            version=self.version)


def get_image_version(image_version_arn: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageVersionResult:
    """
    Resource Type definition for AWS::SageMaker::ImageVersion
    """
    __args__ = dict()
    __args__['imageVersionArn'] = image_version_arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:sagemaker:getImageVersion', __args__, opts=opts, typ=GetImageVersionResult).value

    return AwaitableGetImageVersionResult(
        container_image=__ret__.container_image,
        image_arn=__ret__.image_arn,
        image_version_arn=__ret__.image_version_arn,
        version=__ret__.version)


@_utilities.lift_output_func(get_image_version)
def get_image_version_output(image_version_arn: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImageVersionResult]:
    """
    Resource Type definition for AWS::SageMaker::ImageVersion
    """
    ...
