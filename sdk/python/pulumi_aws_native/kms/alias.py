# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AliasArgs', 'Alias']

@pulumi.input_type
class AliasArgs:
    def __init__(__self__, *,
                 alias_name: pulumi.Input[str],
                 target_key_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a Alias resource.
        :param pulumi.Input[str] alias_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
        :param pulumi.Input[str] target_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
        """
        pulumi.set(__self__, "alias_name", alias_name)
        pulumi.set(__self__, "target_key_id", target_key_id)

    @property
    @pulumi.getter(name="aliasName")
    def alias_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
        """
        return pulumi.get(self, "alias_name")

    @alias_name.setter
    def alias_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias_name", value)

    @property
    @pulumi.getter(name="targetKeyId")
    def target_key_id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
        """
        return pulumi.get(self, "target_key_id")

    @target_key_id.setter
    def target_key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_key_id", value)


class Alias(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias_name: Optional[pulumi.Input[str]] = None,
                 target_key_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
        :param pulumi.Input[str] target_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AliasArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html

        :param str resource_name: The name of the resource.
        :param AliasArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AliasArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias_name: Optional[pulumi.Input[str]] = None,
                 target_key_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = AliasArgs.__new__(AliasArgs)

            if alias_name is None and not opts.urn:
                raise TypeError("Missing required property 'alias_name'")
            __props__.__dict__["alias_name"] = alias_name
            if target_key_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_key_id'")
            __props__.__dict__["target_key_id"] = target_key_id
        super(Alias, __self__).__init__(
            'aws-native:KMS:Alias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Alias':
        """
        Get an existing Alias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AliasArgs.__new__(AliasArgs)

        __props__.__dict__["alias_name"] = None
        __props__.__dict__["target_key_id"] = None
        return Alias(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aliasName")
    def alias_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
        """
        return pulumi.get(self, "alias_name")

    @property
    @pulumi.getter(name="targetKeyId")
    def target_key_id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
        """
        return pulumi.get(self, "target_key_id")

