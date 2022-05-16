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

__all__ = ['PlaybackConfigurationArgs', 'PlaybackConfiguration']

@pulumi.input_type
class PlaybackConfigurationArgs:
    def __init__(__self__, *,
                 ad_decision_server_url: pulumi.Input[str],
                 video_content_source_url: pulumi.Input[str],
                 avail_suppression: Optional[pulumi.Input['PlaybackConfigurationAvailSuppressionArgs']] = None,
                 bumper: Optional[pulumi.Input['PlaybackConfigurationBumperArgs']] = None,
                 cdn_configuration: Optional[pulumi.Input['PlaybackConfigurationCdnConfigurationArgs']] = None,
                 configuration_aliases: Optional[pulumi.Input['PlaybackConfigurationConfigurationAliasesArgs']] = None,
                 dash_configuration: Optional[pulumi.Input['PlaybackConfigurationDashConfigurationArgs']] = None,
                 hls_configuration: Optional[pulumi.Input['PlaybackConfigurationHlsConfigurationArgs']] = None,
                 live_pre_roll_configuration: Optional[pulumi.Input['PlaybackConfigurationLivePreRollConfigurationArgs']] = None,
                 manifest_processing_rules: Optional[pulumi.Input['PlaybackConfigurationManifestProcessingRulesArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 personalization_threshold_seconds: Optional[pulumi.Input[int]] = None,
                 slate_ad_url: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['PlaybackConfigurationTagArgs']]]] = None,
                 transcode_profile_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PlaybackConfiguration resource.
        :param pulumi.Input[str] ad_decision_server_url: The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
        :param pulumi.Input[str] video_content_source_url: The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
        :param pulumi.Input['PlaybackConfigurationAvailSuppressionArgs'] avail_suppression: The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        :param pulumi.Input['PlaybackConfigurationBumperArgs'] bumper: The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
        :param pulumi.Input['PlaybackConfigurationCdnConfigurationArgs'] cdn_configuration: The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
        :param pulumi.Input['PlaybackConfigurationConfigurationAliasesArgs'] configuration_aliases: The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables. 
        :param pulumi.Input['PlaybackConfigurationDashConfigurationArgs'] dash_configuration: The configuration for DASH content.
        :param pulumi.Input['PlaybackConfigurationHlsConfigurationArgs'] hls_configuration: The configuration for HLS content.
        :param pulumi.Input['PlaybackConfigurationLivePreRollConfigurationArgs'] live_pre_roll_configuration: The configuration for pre-roll ad insertion.
        :param pulumi.Input['PlaybackConfigurationManifestProcessingRulesArgs'] manifest_processing_rules: The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
        :param pulumi.Input[str] name: The identifier for the playback configuration.
        :param pulumi.Input[int] personalization_threshold_seconds: Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        :param pulumi.Input[str] slate_ad_url: The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
        :param pulumi.Input[Sequence[pulumi.Input['PlaybackConfigurationTagArgs']]] tags: The tags to assign to the playback configuration.
        :param pulumi.Input[str] transcode_profile_name: The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
        """
        pulumi.set(__self__, "ad_decision_server_url", ad_decision_server_url)
        pulumi.set(__self__, "video_content_source_url", video_content_source_url)
        if avail_suppression is not None:
            pulumi.set(__self__, "avail_suppression", avail_suppression)
        if bumper is not None:
            pulumi.set(__self__, "bumper", bumper)
        if cdn_configuration is not None:
            pulumi.set(__self__, "cdn_configuration", cdn_configuration)
        if configuration_aliases is not None:
            pulumi.set(__self__, "configuration_aliases", configuration_aliases)
        if dash_configuration is not None:
            pulumi.set(__self__, "dash_configuration", dash_configuration)
        if hls_configuration is not None:
            pulumi.set(__self__, "hls_configuration", hls_configuration)
        if live_pre_roll_configuration is not None:
            pulumi.set(__self__, "live_pre_roll_configuration", live_pre_roll_configuration)
        if manifest_processing_rules is not None:
            pulumi.set(__self__, "manifest_processing_rules", manifest_processing_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if personalization_threshold_seconds is not None:
            pulumi.set(__self__, "personalization_threshold_seconds", personalization_threshold_seconds)
        if slate_ad_url is not None:
            pulumi.set(__self__, "slate_ad_url", slate_ad_url)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if transcode_profile_name is not None:
            pulumi.set(__self__, "transcode_profile_name", transcode_profile_name)

    @property
    @pulumi.getter(name="adDecisionServerUrl")
    def ad_decision_server_url(self) -> pulumi.Input[str]:
        """
        The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
        """
        return pulumi.get(self, "ad_decision_server_url")

    @ad_decision_server_url.setter
    def ad_decision_server_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "ad_decision_server_url", value)

    @property
    @pulumi.getter(name="videoContentSourceUrl")
    def video_content_source_url(self) -> pulumi.Input[str]:
        """
        The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
        """
        return pulumi.get(self, "video_content_source_url")

    @video_content_source_url.setter
    def video_content_source_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "video_content_source_url", value)

    @property
    @pulumi.getter(name="availSuppression")
    def avail_suppression(self) -> Optional[pulumi.Input['PlaybackConfigurationAvailSuppressionArgs']]:
        """
        The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        """
        return pulumi.get(self, "avail_suppression")

    @avail_suppression.setter
    def avail_suppression(self, value: Optional[pulumi.Input['PlaybackConfigurationAvailSuppressionArgs']]):
        pulumi.set(self, "avail_suppression", value)

    @property
    @pulumi.getter
    def bumper(self) -> Optional[pulumi.Input['PlaybackConfigurationBumperArgs']]:
        """
        The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
        """
        return pulumi.get(self, "bumper")

    @bumper.setter
    def bumper(self, value: Optional[pulumi.Input['PlaybackConfigurationBumperArgs']]):
        pulumi.set(self, "bumper", value)

    @property
    @pulumi.getter(name="cdnConfiguration")
    def cdn_configuration(self) -> Optional[pulumi.Input['PlaybackConfigurationCdnConfigurationArgs']]:
        """
        The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
        """
        return pulumi.get(self, "cdn_configuration")

    @cdn_configuration.setter
    def cdn_configuration(self, value: Optional[pulumi.Input['PlaybackConfigurationCdnConfigurationArgs']]):
        pulumi.set(self, "cdn_configuration", value)

    @property
    @pulumi.getter(name="configurationAliases")
    def configuration_aliases(self) -> Optional[pulumi.Input['PlaybackConfigurationConfigurationAliasesArgs']]:
        """
        The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables. 
        """
        return pulumi.get(self, "configuration_aliases")

    @configuration_aliases.setter
    def configuration_aliases(self, value: Optional[pulumi.Input['PlaybackConfigurationConfigurationAliasesArgs']]):
        pulumi.set(self, "configuration_aliases", value)

    @property
    @pulumi.getter(name="dashConfiguration")
    def dash_configuration(self) -> Optional[pulumi.Input['PlaybackConfigurationDashConfigurationArgs']]:
        """
        The configuration for DASH content.
        """
        return pulumi.get(self, "dash_configuration")

    @dash_configuration.setter
    def dash_configuration(self, value: Optional[pulumi.Input['PlaybackConfigurationDashConfigurationArgs']]):
        pulumi.set(self, "dash_configuration", value)

    @property
    @pulumi.getter(name="hlsConfiguration")
    def hls_configuration(self) -> Optional[pulumi.Input['PlaybackConfigurationHlsConfigurationArgs']]:
        """
        The configuration for HLS content.
        """
        return pulumi.get(self, "hls_configuration")

    @hls_configuration.setter
    def hls_configuration(self, value: Optional[pulumi.Input['PlaybackConfigurationHlsConfigurationArgs']]):
        pulumi.set(self, "hls_configuration", value)

    @property
    @pulumi.getter(name="livePreRollConfiguration")
    def live_pre_roll_configuration(self) -> Optional[pulumi.Input['PlaybackConfigurationLivePreRollConfigurationArgs']]:
        """
        The configuration for pre-roll ad insertion.
        """
        return pulumi.get(self, "live_pre_roll_configuration")

    @live_pre_roll_configuration.setter
    def live_pre_roll_configuration(self, value: Optional[pulumi.Input['PlaybackConfigurationLivePreRollConfigurationArgs']]):
        pulumi.set(self, "live_pre_roll_configuration", value)

    @property
    @pulumi.getter(name="manifestProcessingRules")
    def manifest_processing_rules(self) -> Optional[pulumi.Input['PlaybackConfigurationManifestProcessingRulesArgs']]:
        """
        The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
        """
        return pulumi.get(self, "manifest_processing_rules")

    @manifest_processing_rules.setter
    def manifest_processing_rules(self, value: Optional[pulumi.Input['PlaybackConfigurationManifestProcessingRulesArgs']]):
        pulumi.set(self, "manifest_processing_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the playback configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="personalizationThresholdSeconds")
    def personalization_threshold_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        """
        return pulumi.get(self, "personalization_threshold_seconds")

    @personalization_threshold_seconds.setter
    def personalization_threshold_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "personalization_threshold_seconds", value)

    @property
    @pulumi.getter(name="slateAdUrl")
    def slate_ad_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
        """
        return pulumi.get(self, "slate_ad_url")

    @slate_ad_url.setter
    def slate_ad_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slate_ad_url", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PlaybackConfigurationTagArgs']]]]:
        """
        The tags to assign to the playback configuration.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PlaybackConfigurationTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="transcodeProfileName")
    def transcode_profile_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
        """
        return pulumi.get(self, "transcode_profile_name")

    @transcode_profile_name.setter
    def transcode_profile_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transcode_profile_name", value)


class PlaybackConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ad_decision_server_url: Optional[pulumi.Input[str]] = None,
                 avail_suppression: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationAvailSuppressionArgs']]] = None,
                 bumper: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationBumperArgs']]] = None,
                 cdn_configuration: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationCdnConfigurationArgs']]] = None,
                 configuration_aliases: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationConfigurationAliasesArgs']]] = None,
                 dash_configuration: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationDashConfigurationArgs']]] = None,
                 hls_configuration: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationHlsConfigurationArgs']]] = None,
                 live_pre_roll_configuration: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationLivePreRollConfigurationArgs']]] = None,
                 manifest_processing_rules: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationManifestProcessingRulesArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 personalization_threshold_seconds: Optional[pulumi.Input[int]] = None,
                 slate_ad_url: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PlaybackConfigurationTagArgs']]]]] = None,
                 transcode_profile_name: Optional[pulumi.Input[str]] = None,
                 video_content_source_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource schema for AWS::MediaTailor::PlaybackConfiguration

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ad_decision_server_url: The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
        :param pulumi.Input[pulumi.InputType['PlaybackConfigurationAvailSuppressionArgs']] avail_suppression: The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        :param pulumi.Input[pulumi.InputType['PlaybackConfigurationBumperArgs']] bumper: The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
        :param pulumi.Input[pulumi.InputType['PlaybackConfigurationCdnConfigurationArgs']] cdn_configuration: The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
        :param pulumi.Input[pulumi.InputType['PlaybackConfigurationConfigurationAliasesArgs']] configuration_aliases: The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables. 
        :param pulumi.Input[pulumi.InputType['PlaybackConfigurationDashConfigurationArgs']] dash_configuration: The configuration for DASH content.
        :param pulumi.Input[pulumi.InputType['PlaybackConfigurationHlsConfigurationArgs']] hls_configuration: The configuration for HLS content.
        :param pulumi.Input[pulumi.InputType['PlaybackConfigurationLivePreRollConfigurationArgs']] live_pre_roll_configuration: The configuration for pre-roll ad insertion.
        :param pulumi.Input[pulumi.InputType['PlaybackConfigurationManifestProcessingRulesArgs']] manifest_processing_rules: The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
        :param pulumi.Input[str] name: The identifier for the playback configuration.
        :param pulumi.Input[int] personalization_threshold_seconds: Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        :param pulumi.Input[str] slate_ad_url: The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PlaybackConfigurationTagArgs']]]] tags: The tags to assign to the playback configuration.
        :param pulumi.Input[str] transcode_profile_name: The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
        :param pulumi.Input[str] video_content_source_url: The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PlaybackConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::MediaTailor::PlaybackConfiguration

        :param str resource_name: The name of the resource.
        :param PlaybackConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PlaybackConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ad_decision_server_url: Optional[pulumi.Input[str]] = None,
                 avail_suppression: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationAvailSuppressionArgs']]] = None,
                 bumper: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationBumperArgs']]] = None,
                 cdn_configuration: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationCdnConfigurationArgs']]] = None,
                 configuration_aliases: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationConfigurationAliasesArgs']]] = None,
                 dash_configuration: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationDashConfigurationArgs']]] = None,
                 hls_configuration: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationHlsConfigurationArgs']]] = None,
                 live_pre_roll_configuration: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationLivePreRollConfigurationArgs']]] = None,
                 manifest_processing_rules: Optional[pulumi.Input[pulumi.InputType['PlaybackConfigurationManifestProcessingRulesArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 personalization_threshold_seconds: Optional[pulumi.Input[int]] = None,
                 slate_ad_url: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PlaybackConfigurationTagArgs']]]]] = None,
                 transcode_profile_name: Optional[pulumi.Input[str]] = None,
                 video_content_source_url: Optional[pulumi.Input[str]] = None,
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
            __props__ = PlaybackConfigurationArgs.__new__(PlaybackConfigurationArgs)

            if ad_decision_server_url is None and not opts.urn:
                raise TypeError("Missing required property 'ad_decision_server_url'")
            __props__.__dict__["ad_decision_server_url"] = ad_decision_server_url
            __props__.__dict__["avail_suppression"] = avail_suppression
            __props__.__dict__["bumper"] = bumper
            __props__.__dict__["cdn_configuration"] = cdn_configuration
            __props__.__dict__["configuration_aliases"] = configuration_aliases
            __props__.__dict__["dash_configuration"] = dash_configuration
            __props__.__dict__["hls_configuration"] = hls_configuration
            __props__.__dict__["live_pre_roll_configuration"] = live_pre_roll_configuration
            __props__.__dict__["manifest_processing_rules"] = manifest_processing_rules
            __props__.__dict__["name"] = name
            __props__.__dict__["personalization_threshold_seconds"] = personalization_threshold_seconds
            __props__.__dict__["slate_ad_url"] = slate_ad_url
            __props__.__dict__["tags"] = tags
            __props__.__dict__["transcode_profile_name"] = transcode_profile_name
            if video_content_source_url is None and not opts.urn:
                raise TypeError("Missing required property 'video_content_source_url'")
            __props__.__dict__["video_content_source_url"] = video_content_source_url
            __props__.__dict__["playback_configuration_arn"] = None
            __props__.__dict__["playback_endpoint_prefix"] = None
            __props__.__dict__["session_initialization_endpoint_prefix"] = None
        super(PlaybackConfiguration, __self__).__init__(
            'aws-native:mediatailor:PlaybackConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PlaybackConfiguration':
        """
        Get an existing PlaybackConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PlaybackConfigurationArgs.__new__(PlaybackConfigurationArgs)

        __props__.__dict__["ad_decision_server_url"] = None
        __props__.__dict__["avail_suppression"] = None
        __props__.__dict__["bumper"] = None
        __props__.__dict__["cdn_configuration"] = None
        __props__.__dict__["configuration_aliases"] = None
        __props__.__dict__["dash_configuration"] = None
        __props__.__dict__["hls_configuration"] = None
        __props__.__dict__["live_pre_roll_configuration"] = None
        __props__.__dict__["manifest_processing_rules"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["personalization_threshold_seconds"] = None
        __props__.__dict__["playback_configuration_arn"] = None
        __props__.__dict__["playback_endpoint_prefix"] = None
        __props__.__dict__["session_initialization_endpoint_prefix"] = None
        __props__.__dict__["slate_ad_url"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["transcode_profile_name"] = None
        __props__.__dict__["video_content_source_url"] = None
        return PlaybackConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adDecisionServerUrl")
    def ad_decision_server_url(self) -> pulumi.Output[str]:
        """
        The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
        """
        return pulumi.get(self, "ad_decision_server_url")

    @property
    @pulumi.getter(name="availSuppression")
    def avail_suppression(self) -> pulumi.Output[Optional['outputs.PlaybackConfigurationAvailSuppression']]:
        """
        The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        """
        return pulumi.get(self, "avail_suppression")

    @property
    @pulumi.getter
    def bumper(self) -> pulumi.Output[Optional['outputs.PlaybackConfigurationBumper']]:
        """
        The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
        """
        return pulumi.get(self, "bumper")

    @property
    @pulumi.getter(name="cdnConfiguration")
    def cdn_configuration(self) -> pulumi.Output[Optional['outputs.PlaybackConfigurationCdnConfiguration']]:
        """
        The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
        """
        return pulumi.get(self, "cdn_configuration")

    @property
    @pulumi.getter(name="configurationAliases")
    def configuration_aliases(self) -> pulumi.Output[Optional['outputs.PlaybackConfigurationConfigurationAliases']]:
        """
        The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables. 
        """
        return pulumi.get(self, "configuration_aliases")

    @property
    @pulumi.getter(name="dashConfiguration")
    def dash_configuration(self) -> pulumi.Output[Optional['outputs.PlaybackConfigurationDashConfiguration']]:
        """
        The configuration for DASH content.
        """
        return pulumi.get(self, "dash_configuration")

    @property
    @pulumi.getter(name="hlsConfiguration")
    def hls_configuration(self) -> pulumi.Output[Optional['outputs.PlaybackConfigurationHlsConfiguration']]:
        """
        The configuration for HLS content.
        """
        return pulumi.get(self, "hls_configuration")

    @property
    @pulumi.getter(name="livePreRollConfiguration")
    def live_pre_roll_configuration(self) -> pulumi.Output[Optional['outputs.PlaybackConfigurationLivePreRollConfiguration']]:
        """
        The configuration for pre-roll ad insertion.
        """
        return pulumi.get(self, "live_pre_roll_configuration")

    @property
    @pulumi.getter(name="manifestProcessingRules")
    def manifest_processing_rules(self) -> pulumi.Output[Optional['outputs.PlaybackConfigurationManifestProcessingRules']]:
        """
        The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
        """
        return pulumi.get(self, "manifest_processing_rules")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The identifier for the playback configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="personalizationThresholdSeconds")
    def personalization_threshold_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        """
        return pulumi.get(self, "personalization_threshold_seconds")

    @property
    @pulumi.getter(name="playbackConfigurationArn")
    def playback_configuration_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) for the playback configuration.
        """
        return pulumi.get(self, "playback_configuration_arn")

    @property
    @pulumi.getter(name="playbackEndpointPrefix")
    def playback_endpoint_prefix(self) -> pulumi.Output[str]:
        """
        The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.
        """
        return pulumi.get(self, "playback_endpoint_prefix")

    @property
    @pulumi.getter(name="sessionInitializationEndpointPrefix")
    def session_initialization_endpoint_prefix(self) -> pulumi.Output[str]:
        """
        The URL that the player uses to initialize a session that uses client-side reporting.
        """
        return pulumi.get(self, "session_initialization_endpoint_prefix")

    @property
    @pulumi.getter(name="slateAdUrl")
    def slate_ad_url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
        """
        return pulumi.get(self, "slate_ad_url")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.PlaybackConfigurationTag']]]:
        """
        The tags to assign to the playback configuration.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transcodeProfileName")
    def transcode_profile_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
        """
        return pulumi.get(self, "transcode_profile_name")

    @property
    @pulumi.getter(name="videoContentSourceUrl")
    def video_content_source_url(self) -> pulumi.Output[str]:
        """
        The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
        """
        return pulumi.get(self, "video_content_source_url")
