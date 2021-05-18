# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['AssetModelArgs', 'AssetModel']

@pulumi.input_type
class AssetModelArgs:
    def __init__(__self__, *,
                 asset_model_name: pulumi.Input[str],
                 asset_model_description: Optional[pulumi.Input[str]] = None,
                 asset_model_hierarchies: Optional[pulumi.Input[Sequence[pulumi.Input['AssetModelAssetModelHierarchyArgs']]]] = None,
                 asset_model_properties: Optional[pulumi.Input[Sequence[pulumi.Input['AssetModelAssetModelPropertyArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a AssetModel resource.
        :param pulumi.Input[str] asset_model_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
        :param pulumi.Input[str] asset_model_description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
        :param pulumi.Input[Sequence[pulumi.Input['AssetModelAssetModelHierarchyArgs']]] asset_model_hierarchies: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
        :param pulumi.Input[Sequence[pulumi.Input['AssetModelAssetModelPropertyArgs']]] asset_model_properties: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
        """
        pulumi.set(__self__, "asset_model_name", asset_model_name)
        if asset_model_description is not None:
            pulumi.set(__self__, "asset_model_description", asset_model_description)
        if asset_model_hierarchies is not None:
            pulumi.set(__self__, "asset_model_hierarchies", asset_model_hierarchies)
        if asset_model_properties is not None:
            pulumi.set(__self__, "asset_model_properties", asset_model_properties)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="assetModelName")
    def asset_model_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
        """
        return pulumi.get(self, "asset_model_name")

    @asset_model_name.setter
    def asset_model_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "asset_model_name", value)

    @property
    @pulumi.getter(name="assetModelDescription")
    def asset_model_description(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
        """
        return pulumi.get(self, "asset_model_description")

    @asset_model_description.setter
    def asset_model_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asset_model_description", value)

    @property
    @pulumi.getter(name="assetModelHierarchies")
    def asset_model_hierarchies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AssetModelAssetModelHierarchyArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
        """
        return pulumi.get(self, "asset_model_hierarchies")

    @asset_model_hierarchies.setter
    def asset_model_hierarchies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AssetModelAssetModelHierarchyArgs']]]]):
        pulumi.set(self, "asset_model_hierarchies", value)

    @property
    @pulumi.getter(name="assetModelProperties")
    def asset_model_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AssetModelAssetModelPropertyArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
        """
        return pulumi.get(self, "asset_model_properties")

    @asset_model_properties.setter
    def asset_model_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AssetModelAssetModelPropertyArgs']]]]):
        pulumi.set(self, "asset_model_properties", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class AssetModel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_model_description: Optional[pulumi.Input[str]] = None,
                 asset_model_hierarchies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssetModelAssetModelHierarchyArgs']]]]] = None,
                 asset_model_name: Optional[pulumi.Input[str]] = None,
                 asset_model_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssetModelAssetModelPropertyArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asset_model_description: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssetModelAssetModelHierarchyArgs']]]] asset_model_hierarchies: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
        :param pulumi.Input[str] asset_model_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssetModelAssetModelPropertyArgs']]]] asset_model_properties: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AssetModelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html

        :param str resource_name: The name of the resource.
        :param AssetModelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssetModelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_model_description: Optional[pulumi.Input[str]] = None,
                 asset_model_hierarchies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssetModelAssetModelHierarchyArgs']]]]] = None,
                 asset_model_name: Optional[pulumi.Input[str]] = None,
                 asset_model_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AssetModelAssetModelPropertyArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AssetModelArgs.__new__(AssetModelArgs)

            __props__.__dict__["asset_model_description"] = asset_model_description
            __props__.__dict__["asset_model_hierarchies"] = asset_model_hierarchies
            if asset_model_name is None and not opts.urn:
                raise TypeError("Missing required property 'asset_model_name'")
            __props__.__dict__["asset_model_name"] = asset_model_name
            __props__.__dict__["asset_model_properties"] = asset_model_properties
            __props__.__dict__["tags"] = tags
            __props__.__dict__["asset_model_arn"] = None
            __props__.__dict__["asset_model_id"] = None
        super(AssetModel, __self__).__init__(
            'aws-native:IoTSiteWise:AssetModel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AssetModel':
        """
        Get an existing AssetModel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AssetModelArgs.__new__(AssetModelArgs)

        __props__.__dict__["asset_model_arn"] = None
        __props__.__dict__["asset_model_description"] = None
        __props__.__dict__["asset_model_hierarchies"] = None
        __props__.__dict__["asset_model_id"] = None
        __props__.__dict__["asset_model_name"] = None
        __props__.__dict__["asset_model_properties"] = None
        __props__.__dict__["tags"] = None
        return AssetModel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="assetModelArn")
    def asset_model_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "asset_model_arn")

    @property
    @pulumi.getter(name="assetModelDescription")
    def asset_model_description(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodeldescription
        """
        return pulumi.get(self, "asset_model_description")

    @property
    @pulumi.getter(name="assetModelHierarchies")
    def asset_model_hierarchies(self) -> pulumi.Output[Optional[Sequence['outputs.AssetModelAssetModelHierarchy']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelhierarchies
        """
        return pulumi.get(self, "asset_model_hierarchies")

    @property
    @pulumi.getter(name="assetModelId")
    def asset_model_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "asset_model_id")

    @property
    @pulumi.getter(name="assetModelName")
    def asset_model_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelname
        """
        return pulumi.get(self, "asset_model_name")

    @property
    @pulumi.getter(name="assetModelProperties")
    def asset_model_properties(self) -> pulumi.Output[Optional[Sequence['outputs.AssetModelAssetModelProperty']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-assetmodelproperties
        """
        return pulumi.get(self, "asset_model_properties")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html#cfn-iotsitewise-assetmodel-tags
        """
        return pulumi.get(self, "tags")

