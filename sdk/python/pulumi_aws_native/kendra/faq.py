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
from ._inputs import *

__all__ = ['FaqArgs', 'Faq']

@pulumi.input_type
class FaqArgs:
    def __init__(__self__, *,
                 index_id: pulumi.Input[str],
                 role_arn: pulumi.Input[str],
                 s3_path: pulumi.Input['FaqS3PathArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 file_format: Optional[pulumi.Input['FaqFileFormat']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FaqTagArgs']]]] = None):
        """
        The set of arguments for constructing a Faq resource.
        :param pulumi.Input[str] index_id: Index ID
        :param pulumi.Input[str] role_arn: FAQ role ARN
        :param pulumi.Input['FaqS3PathArgs'] s3_path: FAQ S3 path
        :param pulumi.Input[str] description: FAQ description
        :param pulumi.Input['FaqFileFormat'] file_format: FAQ file format
        :param pulumi.Input[str] name: FAQ name
        :param pulumi.Input[Sequence[pulumi.Input['FaqTagArgs']]] tags: Tags for labeling the FAQ
        """
        pulumi.set(__self__, "index_id", index_id)
        pulumi.set(__self__, "role_arn", role_arn)
        pulumi.set(__self__, "s3_path", s3_path)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if file_format is not None:
            pulumi.set(__self__, "file_format", file_format)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> pulumi.Input[str]:
        """
        Index ID
        """
        return pulumi.get(self, "index_id")

    @index_id.setter
    def index_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "index_id", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        """
        FAQ role ARN
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="s3Path")
    def s3_path(self) -> pulumi.Input['FaqS3PathArgs']:
        """
        FAQ S3 path
        """
        return pulumi.get(self, "s3_path")

    @s3_path.setter
    def s3_path(self, value: pulumi.Input['FaqS3PathArgs']):
        pulumi.set(self, "s3_path", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        FAQ description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="fileFormat")
    def file_format(self) -> Optional[pulumi.Input['FaqFileFormat']]:
        """
        FAQ file format
        """
        return pulumi.get(self, "file_format")

    @file_format.setter
    def file_format(self, value: Optional[pulumi.Input['FaqFileFormat']]):
        pulumi.set(self, "file_format", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        FAQ name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FaqTagArgs']]]]:
        """
        Tags for labeling the FAQ
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FaqTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Faq(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 file_format: Optional[pulumi.Input['FaqFileFormat']] = None,
                 index_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 s3_path: Optional[pulumi.Input[pulumi.InputType['FaqS3PathArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FaqTagArgs']]]]] = None,
                 __props__=None):
        """
        A Kendra FAQ resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: FAQ description
        :param pulumi.Input['FaqFileFormat'] file_format: FAQ file format
        :param pulumi.Input[str] index_id: Index ID
        :param pulumi.Input[str] name: FAQ name
        :param pulumi.Input[str] role_arn: FAQ role ARN
        :param pulumi.Input[pulumi.InputType['FaqS3PathArgs']] s3_path: FAQ S3 path
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FaqTagArgs']]]] tags: Tags for labeling the FAQ
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FaqArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Kendra FAQ resource

        :param str resource_name: The name of the resource.
        :param FaqArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FaqArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 file_format: Optional[pulumi.Input['FaqFileFormat']] = None,
                 index_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 s3_path: Optional[pulumi.Input[pulumi.InputType['FaqS3PathArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FaqTagArgs']]]]] = None,
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
            __props__ = FaqArgs.__new__(FaqArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["file_format"] = file_format
            if index_id is None and not opts.urn:
                raise TypeError("Missing required property 'index_id'")
            __props__.__dict__["index_id"] = index_id
            __props__.__dict__["name"] = name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__.__dict__["role_arn"] = role_arn
            if s3_path is None and not opts.urn:
                raise TypeError("Missing required property 's3_path'")
            __props__.__dict__["s3_path"] = s3_path
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(Faq, __self__).__init__(
            'aws-native:kendra:Faq',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Faq':
        """
        Get an existing Faq resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FaqArgs.__new__(FaqArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["file_format"] = None
        __props__.__dict__["index_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["role_arn"] = None
        __props__.__dict__["s3_path"] = None
        __props__.__dict__["tags"] = None
        return Faq(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        FAQ description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="fileFormat")
    def file_format(self) -> pulumi.Output[Optional['FaqFileFormat']]:
        """
        FAQ file format
        """
        return pulumi.get(self, "file_format")

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> pulumi.Output[str]:
        """
        Index ID
        """
        return pulumi.get(self, "index_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        FAQ name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        FAQ role ARN
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="s3Path")
    def s3_path(self) -> pulumi.Output['outputs.FaqS3Path']:
        """
        FAQ S3 path
        """
        return pulumi.get(self, "s3_path")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.FaqTag']]]:
        """
        Tags for labeling the FAQ
        """
        return pulumi.get(self, "tags")

