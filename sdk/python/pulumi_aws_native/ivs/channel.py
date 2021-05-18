# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs

__all__ = ['ChannelArgs', 'Channel']

@pulumi.input_type
class ChannelArgs:
    def __init__(__self__, *,
                 authorized: Optional[pulumi.Input[bool]] = None,
                 latency_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Channel resource.
        :param pulumi.Input[bool] authorized: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-authorized
        :param pulumi.Input[str] latency_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-latencymode
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-name
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-tags
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-type
        """
        if authorized is not None:
            pulumi.set(__self__, "authorized", authorized)
        if latency_mode is not None:
            pulumi.set(__self__, "latency_mode", latency_mode)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def authorized(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-authorized
        """
        return pulumi.get(self, "authorized")

    @authorized.setter
    def authorized(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "authorized", value)

    @property
    @pulumi.getter(name="latencyMode")
    def latency_mode(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-latencymode
        """
        return pulumi.get(self, "latency_mode")

    @latency_mode.setter
    def latency_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "latency_mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Channel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[bool]] = None,
                 latency_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] authorized: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-authorized
        :param pulumi.Input[str] latency_mode: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-latencymode
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-name
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-tags
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-type
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ChannelArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html

        :param str resource_name: The name of the resource.
        :param ChannelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ChannelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized: Optional[pulumi.Input[bool]] = None,
                 latency_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            __props__ = ChannelArgs.__new__(ChannelArgs)

            __props__.__dict__["authorized"] = authorized
            __props__.__dict__["latency_mode"] = latency_mode
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["arn"] = None
            __props__.__dict__["ingest_endpoint"] = None
            __props__.__dict__["playback_url"] = None
        super(Channel, __self__).__init__(
            'aws-native:IVS:Channel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Channel':
        """
        Get an existing Channel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ChannelArgs.__new__(ChannelArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["authorized"] = None
        __props__.__dict__["ingest_endpoint"] = None
        __props__.__dict__["latency_mode"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["playback_url"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        return Channel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def authorized(self) -> pulumi.Output[Optional[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-authorized
        """
        return pulumi.get(self, "authorized")

    @property
    @pulumi.getter(name="ingestEndpoint")
    def ingest_endpoint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ingest_endpoint")

    @property
    @pulumi.getter(name="latencyMode")
    def latency_mode(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-latencymode
        """
        return pulumi.get(self, "latency_mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="playbackUrl")
    def playback_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "playback_url")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-channel.html#cfn-ivs-channel-type
        """
        return pulumi.get(self, "type")

