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
    'PlaybackConfigurationAdMarkerPassthrough',
    'PlaybackConfigurationAvailSuppression',
    'PlaybackConfigurationBumper',
    'PlaybackConfigurationCdnConfiguration',
    'PlaybackConfigurationConfigurationAliases',
    'PlaybackConfigurationDashConfigurationForPut',
    'PlaybackConfigurationLivePreRollConfiguration',
    'PlaybackConfigurationManifestProcessingRules',
    'PlaybackConfigurationTag',
]

@pulumi.output_type
class PlaybackConfigurationAdMarkerPassthrough(dict):
    """
    For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN, EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest to the MediaTailor personalized manifest. No logic is applied to these ad markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled for that ad break, MediaTailor will not set the value to 0.
    """
    def __init__(__self__, *,
                 enabled: Optional[bool] = None):
        """
        For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN, EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest to the MediaTailor personalized manifest. No logic is applied to these ad markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled for that ad break, MediaTailor will not set the value to 0.
        :param bool enabled: Enables ad marker passthrough for your configuration.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Enables ad marker passthrough for your configuration.
        """
        return pulumi.get(self, "enabled")


@pulumi.output_type
class PlaybackConfigurationAvailSuppression(dict):
    """
    The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
    """
    def __init__(__self__, *,
                 mode: Optional['PlaybackConfigurationAvailSuppressionMode'] = None,
                 value: Optional[str] = None):
        """
        The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        :param 'PlaybackConfigurationAvailSuppressionMode' mode: Sets the ad suppression mode. By default, ad suppression is set to OFF and all ad breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window.
        :param str value: A live edge offset time in HH:MM:SS. MediaTailor won't fill ad breaks on or behind this time in the manifest lookback window. If Value is set to 00:00:00, it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or behind the live edge. If you set a Value time, MediaTailor won't fill any ad breaks on or behind this time in the manifest lookback window. For example, if you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45 minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes behind the live edge.
        """
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional['PlaybackConfigurationAvailSuppressionMode']:
        """
        Sets the ad suppression mode. By default, ad suppression is set to OFF and all ad breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        """
        A live edge offset time in HH:MM:SS. MediaTailor won't fill ad breaks on or behind this time in the manifest lookback window. If Value is set to 00:00:00, it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or behind the live edge. If you set a Value time, MediaTailor won't fill any ad breaks on or behind this time in the manifest lookback window. For example, if you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45 minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes behind the live edge.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class PlaybackConfigurationBumper(dict):
    """
    The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "endUrl":
            suggest = "end_url"
        elif key == "startUrl":
            suggest = "start_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PlaybackConfigurationBumper. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PlaybackConfigurationBumper.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PlaybackConfigurationBumper.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 end_url: Optional[str] = None,
                 start_url: Optional[str] = None):
        """
        The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
        :param str end_url: The URL for the end bumper asset.
        :param str start_url: The URL for the start bumper asset.
        """
        if end_url is not None:
            pulumi.set(__self__, "end_url", end_url)
        if start_url is not None:
            pulumi.set(__self__, "start_url", start_url)

    @property
    @pulumi.getter(name="endUrl")
    def end_url(self) -> Optional[str]:
        """
        The URL for the end bumper asset.
        """
        return pulumi.get(self, "end_url")

    @property
    @pulumi.getter(name="startUrl")
    def start_url(self) -> Optional[str]:
        """
        The URL for the start bumper asset.
        """
        return pulumi.get(self, "start_url")


@pulumi.output_type
class PlaybackConfigurationCdnConfiguration(dict):
    """
    The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "adSegmentUrlPrefix":
            suggest = "ad_segment_url_prefix"
        elif key == "contentSegmentUrlPrefix":
            suggest = "content_segment_url_prefix"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PlaybackConfigurationCdnConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PlaybackConfigurationCdnConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PlaybackConfigurationCdnConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ad_segment_url_prefix: Optional[str] = None,
                 content_segment_url_prefix: Optional[str] = None):
        """
        The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
        :param str ad_segment_url_prefix: A non-default content delivery network (CDN) to serve ad segments. By default, AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN for the origin ads.mediatailor.&lt;region>.amazonaws.com. Then specify the rule's name in this AdSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
        :param str content_segment_url_prefix: A content delivery network (CDN) to cache content segments, so that content requests don't always have to go to the origin server. First, create a rule in your CDN for the content segment origin server. Then specify the rule's name in this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for content segments.
        """
        if ad_segment_url_prefix is not None:
            pulumi.set(__self__, "ad_segment_url_prefix", ad_segment_url_prefix)
        if content_segment_url_prefix is not None:
            pulumi.set(__self__, "content_segment_url_prefix", content_segment_url_prefix)

    @property
    @pulumi.getter(name="adSegmentUrlPrefix")
    def ad_segment_url_prefix(self) -> Optional[str]:
        """
        A non-default content delivery network (CDN) to serve ad segments. By default, AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN for the origin ads.mediatailor.&lt;region>.amazonaws.com. Then specify the rule's name in this AdSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
        """
        return pulumi.get(self, "ad_segment_url_prefix")

    @property
    @pulumi.getter(name="contentSegmentUrlPrefix")
    def content_segment_url_prefix(self) -> Optional[str]:
        """
        A content delivery network (CDN) to cache content segments, so that content requests don't always have to go to the origin server. First, create a rule in your CDN for the content segment origin server. Then specify the rule's name in this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for content segments.
        """
        return pulumi.get(self, "content_segment_url_prefix")


@pulumi.output_type
class PlaybackConfigurationConfigurationAliases(dict):
    """
    The predefined aliases for dynamic variables.
    """
    def __init__(__self__):
        """
        The predefined aliases for dynamic variables.
        """
        pass


@pulumi.output_type
class PlaybackConfigurationDashConfigurationForPut(dict):
    """
    The configuration for DASH PUT operations.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "mpdLocation":
            suggest = "mpd_location"
        elif key == "originManifestType":
            suggest = "origin_manifest_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PlaybackConfigurationDashConfigurationForPut. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PlaybackConfigurationDashConfigurationForPut.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PlaybackConfigurationDashConfigurationForPut.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 mpd_location: Optional[str] = None,
                 origin_manifest_type: Optional['PlaybackConfigurationDashConfigurationForPutOriginManifestType'] = None):
        """
        The configuration for DASH PUT operations.
        :param str mpd_location: The setting that controls whether MediaTailor includes the Location tag in DASH manifests. MediaTailor populates the Location tag with the URL for manifest update requests, to be used by players that don't support sticky redirects. Disable this if you have CDN routing rules set up for accessing MediaTailor manifests, and you are either using client-side reporting or your players support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
        :param 'PlaybackConfigurationDashConfigurationForPutOriginManifestType' origin_manifest_type: The setting that controls whether MediaTailor handles manifests from the origin server as multi-period manifests or single-period manifests. If your origin server produces single-period manifests, set this to SINGLE_PERIOD. The default setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it to MULTI_PERIOD.
        """
        if mpd_location is not None:
            pulumi.set(__self__, "mpd_location", mpd_location)
        if origin_manifest_type is not None:
            pulumi.set(__self__, "origin_manifest_type", origin_manifest_type)

    @property
    @pulumi.getter(name="mpdLocation")
    def mpd_location(self) -> Optional[str]:
        """
        The setting that controls whether MediaTailor includes the Location tag in DASH manifests. MediaTailor populates the Location tag with the URL for manifest update requests, to be used by players that don't support sticky redirects. Disable this if you have CDN routing rules set up for accessing MediaTailor manifests, and you are either using client-side reporting or your players support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
        """
        return pulumi.get(self, "mpd_location")

    @property
    @pulumi.getter(name="originManifestType")
    def origin_manifest_type(self) -> Optional['PlaybackConfigurationDashConfigurationForPutOriginManifestType']:
        """
        The setting that controls whether MediaTailor handles manifests from the origin server as multi-period manifests or single-period manifests. If your origin server produces single-period manifests, set this to SINGLE_PERIOD. The default setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it to MULTI_PERIOD.
        """
        return pulumi.get(self, "origin_manifest_type")


@pulumi.output_type
class PlaybackConfigurationLivePreRollConfiguration(dict):
    """
    The configuration for pre-roll ad insertion.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "adDecisionServerUrl":
            suggest = "ad_decision_server_url"
        elif key == "maxDurationSeconds":
            suggest = "max_duration_seconds"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PlaybackConfigurationLivePreRollConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PlaybackConfigurationLivePreRollConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PlaybackConfigurationLivePreRollConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ad_decision_server_url: Optional[str] = None,
                 max_duration_seconds: Optional[int] = None):
        """
        The configuration for pre-roll ad insertion.
        :param str ad_decision_server_url: The URL for the ad decision server (ADS) for pre-roll ads. This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing, you can provide a static VAST URL. The maximum length is 25,000 characters.
        :param int max_duration_seconds: The maximum allowed duration for the pre-roll ad avail. AWS Elemental MediaTailor won't play pre-roll ads to exceed this duration, regardless of the total duration of ads that the ADS returns.
        """
        if ad_decision_server_url is not None:
            pulumi.set(__self__, "ad_decision_server_url", ad_decision_server_url)
        if max_duration_seconds is not None:
            pulumi.set(__self__, "max_duration_seconds", max_duration_seconds)

    @property
    @pulumi.getter(name="adDecisionServerUrl")
    def ad_decision_server_url(self) -> Optional[str]:
        """
        The URL for the ad decision server (ADS) for pre-roll ads. This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing, you can provide a static VAST URL. The maximum length is 25,000 characters.
        """
        return pulumi.get(self, "ad_decision_server_url")

    @property
    @pulumi.getter(name="maxDurationSeconds")
    def max_duration_seconds(self) -> Optional[int]:
        """
        The maximum allowed duration for the pre-roll ad avail. AWS Elemental MediaTailor won't play pre-roll ads to exceed this duration, regardless of the total duration of ads that the ADS returns.
        """
        return pulumi.get(self, "max_duration_seconds")


@pulumi.output_type
class PlaybackConfigurationManifestProcessingRules(dict):
    """
    The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "adMarkerPassthrough":
            suggest = "ad_marker_passthrough"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PlaybackConfigurationManifestProcessingRules. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PlaybackConfigurationManifestProcessingRules.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PlaybackConfigurationManifestProcessingRules.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ad_marker_passthrough: Optional['outputs.PlaybackConfigurationAdMarkerPassthrough'] = None):
        """
        The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
        :param 'PlaybackConfigurationAdMarkerPassthrough' ad_marker_passthrough: For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN, EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest to the MediaTailor personalized manifest. No logic is applied to these ad markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled for that ad break, MediaTailor will not set the value to 0.
        """
        if ad_marker_passthrough is not None:
            pulumi.set(__self__, "ad_marker_passthrough", ad_marker_passthrough)

    @property
    @pulumi.getter(name="adMarkerPassthrough")
    def ad_marker_passthrough(self) -> Optional['outputs.PlaybackConfigurationAdMarkerPassthrough']:
        """
        For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN, EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest to the MediaTailor personalized manifest. No logic is applied to these ad markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled for that ad break, MediaTailor will not set the value to 0.
        """
        return pulumi.get(self, "ad_marker_passthrough")


@pulumi.output_type
class PlaybackConfigurationTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")


