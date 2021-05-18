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

__all__ = ['PackagingConfigurationArgs', 'PackagingConfiguration']

@pulumi.input_type
class PackagingConfigurationArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 packaging_group_id: pulumi.Input[str],
                 cmaf_package: Optional[pulumi.Input['PackagingConfigurationCmafPackageArgs']] = None,
                 dash_package: Optional[pulumi.Input['PackagingConfigurationDashPackageArgs']] = None,
                 hls_package: Optional[pulumi.Input['PackagingConfigurationHlsPackageArgs']] = None,
                 mss_package: Optional[pulumi.Input['PackagingConfigurationMssPackageArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a PackagingConfiguration resource.
        :param pulumi.Input[str] id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-id
        :param pulumi.Input[str] packaging_group_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-packaginggroupid
        :param pulumi.Input['PackagingConfigurationCmafPackageArgs'] cmaf_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-cmafpackage
        :param pulumi.Input['PackagingConfigurationDashPackageArgs'] dash_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-dashpackage
        :param pulumi.Input['PackagingConfigurationHlsPackageArgs'] hls_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-hlspackage
        :param pulumi.Input['PackagingConfigurationMssPackageArgs'] mss_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-msspackage
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-tags
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "packaging_group_id", packaging_group_id)
        if cmaf_package is not None:
            pulumi.set(__self__, "cmaf_package", cmaf_package)
        if dash_package is not None:
            pulumi.set(__self__, "dash_package", dash_package)
        if hls_package is not None:
            pulumi.set(__self__, "hls_package", hls_package)
        if mss_package is not None:
            pulumi.set(__self__, "mss_package", mss_package)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-id
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="packagingGroupId")
    def packaging_group_id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-packaginggroupid
        """
        return pulumi.get(self, "packaging_group_id")

    @packaging_group_id.setter
    def packaging_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "packaging_group_id", value)

    @property
    @pulumi.getter(name="cmafPackage")
    def cmaf_package(self) -> Optional[pulumi.Input['PackagingConfigurationCmafPackageArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-cmafpackage
        """
        return pulumi.get(self, "cmaf_package")

    @cmaf_package.setter
    def cmaf_package(self, value: Optional[pulumi.Input['PackagingConfigurationCmafPackageArgs']]):
        pulumi.set(self, "cmaf_package", value)

    @property
    @pulumi.getter(name="dashPackage")
    def dash_package(self) -> Optional[pulumi.Input['PackagingConfigurationDashPackageArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-dashpackage
        """
        return pulumi.get(self, "dash_package")

    @dash_package.setter
    def dash_package(self, value: Optional[pulumi.Input['PackagingConfigurationDashPackageArgs']]):
        pulumi.set(self, "dash_package", value)

    @property
    @pulumi.getter(name="hlsPackage")
    def hls_package(self) -> Optional[pulumi.Input['PackagingConfigurationHlsPackageArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-hlspackage
        """
        return pulumi.get(self, "hls_package")

    @hls_package.setter
    def hls_package(self, value: Optional[pulumi.Input['PackagingConfigurationHlsPackageArgs']]):
        pulumi.set(self, "hls_package", value)

    @property
    @pulumi.getter(name="mssPackage")
    def mss_package(self) -> Optional[pulumi.Input['PackagingConfigurationMssPackageArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-msspackage
        """
        return pulumi.get(self, "mss_package")

    @mss_package.setter
    def mss_package(self, value: Optional[pulumi.Input['PackagingConfigurationMssPackageArgs']]):
        pulumi.set(self, "mss_package", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class PackagingConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cmaf_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationCmafPackageArgs']]] = None,
                 dash_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationDashPackageArgs']]] = None,
                 hls_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationHlsPackageArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 mss_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationMssPackageArgs']]] = None,
                 packaging_group_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PackagingConfigurationCmafPackageArgs']] cmaf_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-cmafpackage
        :param pulumi.Input[pulumi.InputType['PackagingConfigurationDashPackageArgs']] dash_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-dashpackage
        :param pulumi.Input[pulumi.InputType['PackagingConfigurationHlsPackageArgs']] hls_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-hlspackage
        :param pulumi.Input[str] id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-id
        :param pulumi.Input[pulumi.InputType['PackagingConfigurationMssPackageArgs']] mss_package: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-msspackage
        :param pulumi.Input[str] packaging_group_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-packaginggroupid
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PackagingConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html

        :param str resource_name: The name of the resource.
        :param PackagingConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PackagingConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cmaf_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationCmafPackageArgs']]] = None,
                 dash_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationDashPackageArgs']]] = None,
                 hls_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationHlsPackageArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 mss_package: Optional[pulumi.Input[pulumi.InputType['PackagingConfigurationMssPackageArgs']]] = None,
                 packaging_group_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = PackagingConfigurationArgs.__new__(PackagingConfigurationArgs)

            __props__.__dict__["cmaf_package"] = cmaf_package
            __props__.__dict__["dash_package"] = dash_package
            __props__.__dict__["hls_package"] = hls_package
            if id is None and not opts.urn:
                raise TypeError("Missing required property 'id'")
            __props__.__dict__["id"] = id
            __props__.__dict__["mss_package"] = mss_package
            if packaging_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'packaging_group_id'")
            __props__.__dict__["packaging_group_id"] = packaging_group_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(PackagingConfiguration, __self__).__init__(
            'aws-native:MediaPackage:PackagingConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PackagingConfiguration':
        """
        Get an existing PackagingConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PackagingConfigurationArgs.__new__(PackagingConfigurationArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["cmaf_package"] = None
        __props__.__dict__["dash_package"] = None
        __props__.__dict__["hls_package"] = None
        __props__.__dict__["id"] = None
        __props__.__dict__["mss_package"] = None
        __props__.__dict__["packaging_group_id"] = None
        __props__.__dict__["tags"] = None
        return PackagingConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cmafPackage")
    def cmaf_package(self) -> pulumi.Output[Optional['outputs.PackagingConfigurationCmafPackage']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-cmafpackage
        """
        return pulumi.get(self, "cmaf_package")

    @property
    @pulumi.getter(name="dashPackage")
    def dash_package(self) -> pulumi.Output[Optional['outputs.PackagingConfigurationDashPackage']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-dashpackage
        """
        return pulumi.get(self, "dash_package")

    @property
    @pulumi.getter(name="hlsPackage")
    def hls_package(self) -> pulumi.Output[Optional['outputs.PackagingConfigurationHlsPackage']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-hlspackage
        """
        return pulumi.get(self, "hls_package")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mssPackage")
    def mss_package(self) -> pulumi.Output[Optional['outputs.PackagingConfigurationMssPackage']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-msspackage
        """
        return pulumi.get(self, "mss_package")

    @property
    @pulumi.getter(name="packagingGroupId")
    def packaging_group_id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-packaginggroupid
        """
        return pulumi.get(self, "packaging_group_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packagingconfiguration.html#cfn-mediapackage-packagingconfiguration-tags
        """
        return pulumi.get(self, "tags")

