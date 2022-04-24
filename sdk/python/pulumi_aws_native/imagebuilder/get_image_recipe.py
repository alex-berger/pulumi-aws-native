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
    'GetImageRecipeResult',
    'AwaitableGetImageRecipeResult',
    'get_image_recipe',
    'get_image_recipe_output',
]

@pulumi.output_type
class GetImageRecipeResult:
    def __init__(__self__, additional_instance_configuration=None, arn=None):
        if additional_instance_configuration and not isinstance(additional_instance_configuration, dict):
            raise TypeError("Expected argument 'additional_instance_configuration' to be a dict")
        pulumi.set(__self__, "additional_instance_configuration", additional_instance_configuration)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter(name="additionalInstanceConfiguration")
    def additional_instance_configuration(self) -> Optional['outputs.ImageRecipeAdditionalInstanceConfiguration']:
        """
        Specify additional settings and launch scripts for your build instances.
        """
        return pulumi.get(self, "additional_instance_configuration")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the image recipe.
        """
        return pulumi.get(self, "arn")


class AwaitableGetImageRecipeResult(GetImageRecipeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageRecipeResult(
            additional_instance_configuration=self.additional_instance_configuration,
            arn=self.arn)


def get_image_recipe(arn: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageRecipeResult:
    """
    Resource schema for AWS::ImageBuilder::ImageRecipe


    :param str arn: The Amazon Resource Name (ARN) of the image recipe.
    """
    __args__ = dict()
    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:imagebuilder:getImageRecipe', __args__, opts=opts, typ=GetImageRecipeResult).value

    return AwaitableGetImageRecipeResult(
        additional_instance_configuration=__ret__.additional_instance_configuration,
        arn=__ret__.arn)


@_utilities.lift_output_func(get_image_recipe)
def get_image_recipe_output(arn: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImageRecipeResult]:
    """
    Resource schema for AWS::ImageBuilder::ImageRecipe


    :param str arn: The Amazon Resource Name (ARN) of the image recipe.
    """
    ...
