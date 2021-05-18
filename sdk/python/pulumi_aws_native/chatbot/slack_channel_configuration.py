# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SlackChannelConfigurationArgs', 'SlackChannelConfiguration']

@pulumi.input_type
class SlackChannelConfigurationArgs:
    def __init__(__self__, *,
                 configuration_name: pulumi.Input[str],
                 iam_role_arn: pulumi.Input[str],
                 slack_channel_id: pulumi.Input[str],
                 slack_workspace_id: pulumi.Input[str],
                 logging_level: Optional[pulumi.Input[str]] = None,
                 sns_topic_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SlackChannelConfiguration resource.
        :param pulumi.Input[str] configuration_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-configurationname
        :param pulumi.Input[str] iam_role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-iamrolearn
        :param pulumi.Input[str] slack_channel_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-slackchannelid
        :param pulumi.Input[str] slack_workspace_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-slackworkspaceid
        :param pulumi.Input[str] logging_level: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-logginglevel
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sns_topic_arns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-snstopicarns
        """
        pulumi.set(__self__, "configuration_name", configuration_name)
        pulumi.set(__self__, "iam_role_arn", iam_role_arn)
        pulumi.set(__self__, "slack_channel_id", slack_channel_id)
        pulumi.set(__self__, "slack_workspace_id", slack_workspace_id)
        if logging_level is not None:
            pulumi.set(__self__, "logging_level", logging_level)
        if sns_topic_arns is not None:
            pulumi.set(__self__, "sns_topic_arns", sns_topic_arns)

    @property
    @pulumi.getter(name="configurationName")
    def configuration_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-configurationname
        """
        return pulumi.get(self, "configuration_name")

    @configuration_name.setter
    def configuration_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "configuration_name", value)

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-iamrolearn
        """
        return pulumi.get(self, "iam_role_arn")

    @iam_role_arn.setter
    def iam_role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "iam_role_arn", value)

    @property
    @pulumi.getter(name="slackChannelId")
    def slack_channel_id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-slackchannelid
        """
        return pulumi.get(self, "slack_channel_id")

    @slack_channel_id.setter
    def slack_channel_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "slack_channel_id", value)

    @property
    @pulumi.getter(name="slackWorkspaceId")
    def slack_workspace_id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-slackworkspaceid
        """
        return pulumi.get(self, "slack_workspace_id")

    @slack_workspace_id.setter
    def slack_workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "slack_workspace_id", value)

    @property
    @pulumi.getter(name="loggingLevel")
    def logging_level(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-logginglevel
        """
        return pulumi.get(self, "logging_level")

    @logging_level.setter
    def logging_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logging_level", value)

    @property
    @pulumi.getter(name="snsTopicArns")
    def sns_topic_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-snstopicarns
        """
        return pulumi.get(self, "sns_topic_arns")

    @sns_topic_arns.setter
    def sns_topic_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "sns_topic_arns", value)


class SlackChannelConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_name: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 logging_level: Optional[pulumi.Input[str]] = None,
                 slack_channel_id: Optional[pulumi.Input[str]] = None,
                 slack_workspace_id: Optional[pulumi.Input[str]] = None,
                 sns_topic_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] configuration_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-configurationname
        :param pulumi.Input[str] iam_role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-iamrolearn
        :param pulumi.Input[str] logging_level: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-logginglevel
        :param pulumi.Input[str] slack_channel_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-slackchannelid
        :param pulumi.Input[str] slack_workspace_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-slackworkspaceid
        :param pulumi.Input[Sequence[pulumi.Input[str]]] sns_topic_arns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-snstopicarns
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SlackChannelConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html

        :param str resource_name: The name of the resource.
        :param SlackChannelConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SlackChannelConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configuration_name: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 logging_level: Optional[pulumi.Input[str]] = None,
                 slack_channel_id: Optional[pulumi.Input[str]] = None,
                 slack_workspace_id: Optional[pulumi.Input[str]] = None,
                 sns_topic_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = SlackChannelConfigurationArgs.__new__(SlackChannelConfigurationArgs)

            if configuration_name is None and not opts.urn:
                raise TypeError("Missing required property 'configuration_name'")
            __props__.__dict__["configuration_name"] = configuration_name
            if iam_role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'iam_role_arn'")
            __props__.__dict__["iam_role_arn"] = iam_role_arn
            __props__.__dict__["logging_level"] = logging_level
            if slack_channel_id is None and not opts.urn:
                raise TypeError("Missing required property 'slack_channel_id'")
            __props__.__dict__["slack_channel_id"] = slack_channel_id
            if slack_workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'slack_workspace_id'")
            __props__.__dict__["slack_workspace_id"] = slack_workspace_id
            __props__.__dict__["sns_topic_arns"] = sns_topic_arns
            __props__.__dict__["arn"] = None
        super(SlackChannelConfiguration, __self__).__init__(
            'aws-native:Chatbot:SlackChannelConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SlackChannelConfiguration':
        """
        Get an existing SlackChannelConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SlackChannelConfigurationArgs.__new__(SlackChannelConfigurationArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["configuration_name"] = None
        __props__.__dict__["iam_role_arn"] = None
        __props__.__dict__["logging_level"] = None
        __props__.__dict__["slack_channel_id"] = None
        __props__.__dict__["slack_workspace_id"] = None
        __props__.__dict__["sns_topic_arns"] = None
        return SlackChannelConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="configurationName")
    def configuration_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-configurationname
        """
        return pulumi.get(self, "configuration_name")

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-iamrolearn
        """
        return pulumi.get(self, "iam_role_arn")

    @property
    @pulumi.getter(name="loggingLevel")
    def logging_level(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-logginglevel
        """
        return pulumi.get(self, "logging_level")

    @property
    @pulumi.getter(name="slackChannelId")
    def slack_channel_id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-slackchannelid
        """
        return pulumi.get(self, "slack_channel_id")

    @property
    @pulumi.getter(name="slackWorkspaceId")
    def slack_workspace_id(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-slackworkspaceid
        """
        return pulumi.get(self, "slack_workspace_id")

    @property
    @pulumi.getter(name="snsTopicArns")
    def sns_topic_arns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chatbot-slackchannelconfiguration.html#cfn-chatbot-slackchannelconfiguration-snstopicarns
        """
        return pulumi.get(self, "sns_topic_arns")

