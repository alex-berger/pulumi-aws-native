# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'PlaybackConfigurationAdMarkerPassthroughArgs',
    'PlaybackConfigurationAvailSuppressionArgs',
    'PlaybackConfigurationBumperArgs',
    'PlaybackConfigurationCdnConfigurationArgs',
    'PlaybackConfigurationConfigurationAliasesArgs',
    'PlaybackConfigurationDashConfigurationForPutArgs',
    'PlaybackConfigurationLivePreRollConfigurationArgs',
    'PlaybackConfigurationManifestProcessingRulesArgs',
    'PlaybackConfigurationTagArgs',
]

@pulumi.input_type
class PlaybackConfigurationAdMarkerPassthroughArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        """
        For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN, EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest to the MediaTailor personalized manifest. No logic is applied to these ad markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled for that ad break, MediaTailor will not set the value to 0.
        :param pulumi.Input[bool] enabled: Enables ad marker passthrough for your configuration.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables ad marker passthrough for your configuration.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class PlaybackConfigurationAvailSuppressionArgs:
    def __init__(__self__, *,
                 mode: Optional[pulumi.Input['PlaybackConfigurationAvailSuppressionMode']] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        :param pulumi.Input['PlaybackConfigurationAvailSuppressionMode'] mode: Sets the ad suppression mode. By default, ad suppression is set to OFF and all ad breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window.
        :param pulumi.Input[str] value: A live edge offset time in HH:MM:SS. MediaTailor won't fill ad breaks on or behind this time in the manifest lookback window. If Value is set to 00:00:00, it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or behind the live edge. If you set a Value time, MediaTailor won't fill any ad breaks on or behind this time in the manifest lookback window. For example, if you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45 minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes behind the live edge.
        """
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input['PlaybackConfigurationAvailSuppressionMode']]:
        """
        Sets the ad suppression mode. By default, ad suppression is set to OFF and all ad breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input['PlaybackConfigurationAvailSuppressionMode']]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        A live edge offset time in HH:MM:SS. MediaTailor won't fill ad breaks on or behind this time in the manifest lookback window. If Value is set to 00:00:00, it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or behind the live edge. If you set a Value time, MediaTailor won't fill any ad breaks on or behind this time in the manifest lookback window. For example, if you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45 minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes behind the live edge.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class PlaybackConfigurationBumperArgs:
    def __init__(__self__, *,
                 end_url: Optional[pulumi.Input[str]] = None,
                 start_url: Optional[pulumi.Input[str]] = None):
        """
        The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
        :param pulumi.Input[str] end_url: The URL for the end bumper asset.
        :param pulumi.Input[str] start_url: The URL for the start bumper asset.
        """
        if end_url is not None:
            pulumi.set(__self__, "end_url", end_url)
        if start_url is not None:
            pulumi.set(__self__, "start_url", start_url)

    @property
    @pulumi.getter(name="endUrl")
    def end_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for the end bumper asset.
        """
        return pulumi.get(self, "end_url")

    @end_url.setter
    def end_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_url", value)

    @property
    @pulumi.getter(name="startUrl")
    def start_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for the start bumper asset.
        """
        return pulumi.get(self, "start_url")

    @start_url.setter
    def start_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_url", value)


@pulumi.input_type
class PlaybackConfigurationCdnConfigurationArgs:
    def __init__(__self__, *,
                 ad_segment_url_prefix: Optional[pulumi.Input[str]] = None,
                 content_segment_url_prefix: Optional[pulumi.Input[str]] = None):
        """
        The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
        :param pulumi.Input[str] ad_segment_url_prefix: A non-default content delivery network (CDN) to serve ad segments. By default, AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN for the origin ads.mediatailor.&lt;region>.amazonaws.com. Then specify the rule's name in this AdSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
        :param pulumi.Input[str] content_segment_url_prefix: A content delivery network (CDN) to cache content segments, so that content requests don't always have to go to the origin server. First, create a rule in your CDN for the content segment origin server. Then specify the rule's name in this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for content segments.
        """
        if ad_segment_url_prefix is not None:
            pulumi.set(__self__, "ad_segment_url_prefix", ad_segment_url_prefix)
        if content_segment_url_prefix is not None:
            pulumi.set(__self__, "content_segment_url_prefix", content_segment_url_prefix)

    @property
    @pulumi.getter(name="adSegmentUrlPrefix")
    def ad_segment_url_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        A non-default content delivery network (CDN) to serve ad segments. By default, AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN for the origin ads.mediatailor.&lt;region>.amazonaws.com. Then specify the rule's name in this AdSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
        """
        return pulumi.get(self, "ad_segment_url_prefix")

    @ad_segment_url_prefix.setter
    def ad_segment_url_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ad_segment_url_prefix", value)

    @property
    @pulumi.getter(name="contentSegmentUrlPrefix")
    def content_segment_url_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        A content delivery network (CDN) to cache content segments, so that content requests don't always have to go to the origin server. First, create a rule in your CDN for the content segment origin server. Then specify the rule's name in this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for content segments.
        """
        return pulumi.get(self, "content_segment_url_prefix")

    @content_segment_url_prefix.setter
    def content_segment_url_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_segment_url_prefix", value)


@pulumi.input_type
class PlaybackConfigurationConfigurationAliasesArgs:
    def __init__(__self__):
        """
        The predefined aliases for dynamic variables.
        """
        pass


@pulumi.input_type
class PlaybackConfigurationDashConfigurationForPutArgs:
    def __init__(__self__, *,
                 mpd_location: Optional[pulumi.Input[str]] = None,
                 origin_manifest_type: Optional[pulumi.Input['PlaybackConfigurationDashConfigurationForPutOriginManifestType']] = None):
        """
        The configuration for DASH PUT operations.
        :param pulumi.Input[str] mpd_location: The setting that controls whether MediaTailor includes the Location tag in DASH manifests. MediaTailor populates the Location tag with the URL for manifest update requests, to be used by players that don't support sticky redirects. Disable this if you have CDN routing rules set up for accessing MediaTailor manifests, and you are either using client-side reporting or your players support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
        :param pulumi.Input['PlaybackConfigurationDashConfigurationForPutOriginManifestType'] origin_manifest_type: The setting that controls whether MediaTailor handles manifests from the origin server as multi-period manifests or single-period manifests. If your origin server produces single-period manifests, set this to SINGLE_PERIOD. The default setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it to MULTI_PERIOD.
        """
        if mpd_location is not None:
            pulumi.set(__self__, "mpd_location", mpd_location)
        if origin_manifest_type is not None:
            pulumi.set(__self__, "origin_manifest_type", origin_manifest_type)

    @property
    @pulumi.getter(name="mpdLocation")
    def mpd_location(self) -> Optional[pulumi.Input[str]]:
        """
        The setting that controls whether MediaTailor includes the Location tag in DASH manifests. MediaTailor populates the Location tag with the URL for manifest update requests, to be used by players that don't support sticky redirects. Disable this if you have CDN routing rules set up for accessing MediaTailor manifests, and you are either using client-side reporting or your players support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
        """
        return pulumi.get(self, "mpd_location")

    @mpd_location.setter
    def mpd_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mpd_location", value)

    @property
    @pulumi.getter(name="originManifestType")
    def origin_manifest_type(self) -> Optional[pulumi.Input['PlaybackConfigurationDashConfigurationForPutOriginManifestType']]:
        """
        The setting that controls whether MediaTailor handles manifests from the origin server as multi-period manifests or single-period manifests. If your origin server produces single-period manifests, set this to SINGLE_PERIOD. The default setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it to MULTI_PERIOD.
        """
        return pulumi.get(self, "origin_manifest_type")

    @origin_manifest_type.setter
    def origin_manifest_type(self, value: Optional[pulumi.Input['PlaybackConfigurationDashConfigurationForPutOriginManifestType']]):
        pulumi.set(self, "origin_manifest_type", value)


@pulumi.input_type
class PlaybackConfigurationLivePreRollConfigurationArgs:
    def __init__(__self__, *,
                 ad_decision_server_url: Optional[pulumi.Input[str]] = None,
                 max_duration_seconds: Optional[pulumi.Input[int]] = None):
        """
        The configuration for pre-roll ad insertion.
        :param pulumi.Input[str] ad_decision_server_url: The URL for the ad decision server (ADS) for pre-roll ads. This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing, you can provide a static VAST URL. The maximum length is 25,000 characters.
        :param pulumi.Input[int] max_duration_seconds: The maximum allowed duration for the pre-roll ad avail. AWS Elemental MediaTailor won't play pre-roll ads to exceed this duration, regardless of the total duration of ads that the ADS returns.
        """
        if ad_decision_server_url is not None:
            pulumi.set(__self__, "ad_decision_server_url", ad_decision_server_url)
        if max_duration_seconds is not None:
            pulumi.set(__self__, "max_duration_seconds", max_duration_seconds)

    @property
    @pulumi.getter(name="adDecisionServerUrl")
    def ad_decision_server_url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL for the ad decision server (ADS) for pre-roll ads. This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing, you can provide a static VAST URL. The maximum length is 25,000 characters.
        """
        return pulumi.get(self, "ad_decision_server_url")

    @ad_decision_server_url.setter
    def ad_decision_server_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ad_decision_server_url", value)

    @property
    @pulumi.getter(name="maxDurationSeconds")
    def max_duration_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum allowed duration for the pre-roll ad avail. AWS Elemental MediaTailor won't play pre-roll ads to exceed this duration, regardless of the total duration of ads that the ADS returns.
        """
        return pulumi.get(self, "max_duration_seconds")

    @max_duration_seconds.setter
    def max_duration_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_duration_seconds", value)


@pulumi.input_type
class PlaybackConfigurationManifestProcessingRulesArgs:
    def __init__(__self__, *,
                 ad_marker_passthrough: Optional[pulumi.Input['PlaybackConfigurationAdMarkerPassthroughArgs']] = None):
        """
        The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
        :param pulumi.Input['PlaybackConfigurationAdMarkerPassthroughArgs'] ad_marker_passthrough: For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN, EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest to the MediaTailor personalized manifest. No logic is applied to these ad markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled for that ad break, MediaTailor will not set the value to 0.
        """
        if ad_marker_passthrough is not None:
            pulumi.set(__self__, "ad_marker_passthrough", ad_marker_passthrough)

    @property
    @pulumi.getter(name="adMarkerPassthrough")
    def ad_marker_passthrough(self) -> Optional[pulumi.Input['PlaybackConfigurationAdMarkerPassthroughArgs']]:
        """
        For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN, EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest to the MediaTailor personalized manifest. No logic is applied to these ad markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled for that ad break, MediaTailor will not set the value to 0.
        """
        return pulumi.get(self, "ad_marker_passthrough")

    @ad_marker_passthrough.setter
    def ad_marker_passthrough(self, value: Optional[pulumi.Input['PlaybackConfigurationAdMarkerPassthroughArgs']]):
        pulumi.set(self, "ad_marker_passthrough", value)


@pulumi.input_type
class PlaybackConfigurationTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


