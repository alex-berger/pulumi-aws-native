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
    'GetAssetResult',
    'AwaitableGetAssetResult',
    'get_asset',
    'get_asset_output',
]

@pulumi.output_type
class GetAssetResult:
    def __init__(__self__, asset_arn=None, asset_hierarchies=None, asset_id=None, asset_model_id=None, asset_name=None, asset_properties=None, tags=None):
        if asset_arn and not isinstance(asset_arn, str):
            raise TypeError("Expected argument 'asset_arn' to be a str")
        pulumi.set(__self__, "asset_arn", asset_arn)
        if asset_hierarchies and not isinstance(asset_hierarchies, list):
            raise TypeError("Expected argument 'asset_hierarchies' to be a list")
        pulumi.set(__self__, "asset_hierarchies", asset_hierarchies)
        if asset_id and not isinstance(asset_id, str):
            raise TypeError("Expected argument 'asset_id' to be a str")
        pulumi.set(__self__, "asset_id", asset_id)
        if asset_model_id and not isinstance(asset_model_id, str):
            raise TypeError("Expected argument 'asset_model_id' to be a str")
        pulumi.set(__self__, "asset_model_id", asset_model_id)
        if asset_name and not isinstance(asset_name, str):
            raise TypeError("Expected argument 'asset_name' to be a str")
        pulumi.set(__self__, "asset_name", asset_name)
        if asset_properties and not isinstance(asset_properties, list):
            raise TypeError("Expected argument 'asset_properties' to be a list")
        pulumi.set(__self__, "asset_properties", asset_properties)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="assetArn")
    def asset_arn(self) -> Optional[str]:
        """
        The ARN of the asset
        """
        return pulumi.get(self, "asset_arn")

    @property
    @pulumi.getter(name="assetHierarchies")
    def asset_hierarchies(self) -> Optional[Sequence['outputs.AssetHierarchy']]:
        return pulumi.get(self, "asset_hierarchies")

    @property
    @pulumi.getter(name="assetId")
    def asset_id(self) -> Optional[str]:
        """
        The ID of the asset
        """
        return pulumi.get(self, "asset_id")

    @property
    @pulumi.getter(name="assetModelId")
    def asset_model_id(self) -> Optional[str]:
        """
        The ID of the asset model from which to create the asset.
        """
        return pulumi.get(self, "asset_model_id")

    @property
    @pulumi.getter(name="assetName")
    def asset_name(self) -> Optional[str]:
        """
        A unique, friendly name for the asset.
        """
        return pulumi.get(self, "asset_name")

    @property
    @pulumi.getter(name="assetProperties")
    def asset_properties(self) -> Optional[Sequence['outputs.AssetProperty']]:
        return pulumi.get(self, "asset_properties")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.AssetTag']]:
        """
        A list of key-value pairs that contain metadata for the asset.
        """
        return pulumi.get(self, "tags")


class AwaitableGetAssetResult(GetAssetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssetResult(
            asset_arn=self.asset_arn,
            asset_hierarchies=self.asset_hierarchies,
            asset_id=self.asset_id,
            asset_model_id=self.asset_model_id,
            asset_name=self.asset_name,
            asset_properties=self.asset_properties,
            tags=self.tags)


def get_asset(asset_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAssetResult:
    """
    Resource schema for AWS::IoTSiteWise::Asset


    :param str asset_id: The ID of the asset
    """
    __args__ = dict()
    __args__['assetId'] = asset_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iotsitewise:getAsset', __args__, opts=opts, typ=GetAssetResult).value

    return AwaitableGetAssetResult(
        asset_arn=__ret__.asset_arn,
        asset_hierarchies=__ret__.asset_hierarchies,
        asset_id=__ret__.asset_id,
        asset_model_id=__ret__.asset_model_id,
        asset_name=__ret__.asset_name,
        asset_properties=__ret__.asset_properties,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_asset)
def get_asset_output(asset_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAssetResult]:
    """
    Resource schema for AWS::IoTSiteWise::Asset


    :param str asset_id: The ID of the asset
    """
    ...
