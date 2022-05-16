// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediatailor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaTailor::PlaybackConfiguration
type PlaybackConfiguration struct {
	pulumi.CustomResourceState

	// The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl pulumi.StringOutput `pulumi:"adDecisionServerUrl"`
	// The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	AvailSuppression PlaybackConfigurationAvailSuppressionPtrOutput `pulumi:"availSuppression"`
	// The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
	Bumper PlaybackConfigurationBumperPtrOutput `pulumi:"bumper"`
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	CdnConfiguration PlaybackConfigurationCdnConfigurationPtrOutput `pulumi:"cdnConfiguration"`
	// The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables.
	ConfigurationAliases PlaybackConfigurationConfigurationAliasesPtrOutput `pulumi:"configurationAliases"`
	// The configuration for DASH content.
	DashConfiguration PlaybackConfigurationDashConfigurationPtrOutput `pulumi:"dashConfiguration"`
	// The configuration for HLS content.
	HlsConfiguration PlaybackConfigurationHlsConfigurationPtrOutput `pulumi:"hlsConfiguration"`
	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration PlaybackConfigurationLivePreRollConfigurationPtrOutput `pulumi:"livePreRollConfiguration"`
	// The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules PlaybackConfigurationManifestProcessingRulesPtrOutput `pulumi:"manifestProcessingRules"`
	// The identifier for the playback configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	PersonalizationThresholdSeconds pulumi.IntPtrOutput `pulumi:"personalizationThresholdSeconds"`
	// The Amazon Resource Name (ARN) for the playback configuration.
	PlaybackConfigurationArn pulumi.StringOutput `pulumi:"playbackConfigurationArn"`
	// The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.
	PlaybackEndpointPrefix pulumi.StringOutput `pulumi:"playbackEndpointPrefix"`
	// The URL that the player uses to initialize a session that uses client-side reporting.
	SessionInitializationEndpointPrefix pulumi.StringOutput `pulumi:"sessionInitializationEndpointPrefix"`
	// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl pulumi.StringPtrOutput `pulumi:"slateAdUrl"`
	// The tags to assign to the playback configuration.
	Tags PlaybackConfigurationTagArrayOutput `pulumi:"tags"`
	// The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	TranscodeProfileName pulumi.StringPtrOutput `pulumi:"transcodeProfileName"`
	// The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
	VideoContentSourceUrl pulumi.StringOutput `pulumi:"videoContentSourceUrl"`
}

// NewPlaybackConfiguration registers a new resource with the given unique name, arguments, and options.
func NewPlaybackConfiguration(ctx *pulumi.Context,
	name string, args *PlaybackConfigurationArgs, opts ...pulumi.ResourceOption) (*PlaybackConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdDecisionServerUrl == nil {
		return nil, errors.New("invalid value for required argument 'AdDecisionServerUrl'")
	}
	if args.VideoContentSourceUrl == nil {
		return nil, errors.New("invalid value for required argument 'VideoContentSourceUrl'")
	}
	var resource PlaybackConfiguration
	err := ctx.RegisterResource("aws-native:mediatailor:PlaybackConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlaybackConfiguration gets an existing PlaybackConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlaybackConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlaybackConfigurationState, opts ...pulumi.ResourceOption) (*PlaybackConfiguration, error) {
	var resource PlaybackConfiguration
	err := ctx.ReadResource("aws-native:mediatailor:PlaybackConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlaybackConfiguration resources.
type playbackConfigurationState struct {
}

type PlaybackConfigurationState struct {
}

func (PlaybackConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*playbackConfigurationState)(nil)).Elem()
}

type playbackConfigurationArgs struct {
	// The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl string `pulumi:"adDecisionServerUrl"`
	// The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	AvailSuppression *PlaybackConfigurationAvailSuppression `pulumi:"availSuppression"`
	// The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
	Bumper *PlaybackConfigurationBumper `pulumi:"bumper"`
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	CdnConfiguration *PlaybackConfigurationCdnConfiguration `pulumi:"cdnConfiguration"`
	// The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables.
	ConfigurationAliases *PlaybackConfigurationConfigurationAliases `pulumi:"configurationAliases"`
	// The configuration for DASH content.
	DashConfiguration *PlaybackConfigurationDashConfiguration `pulumi:"dashConfiguration"`
	// The configuration for HLS content.
	HlsConfiguration *PlaybackConfigurationHlsConfiguration `pulumi:"hlsConfiguration"`
	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *PlaybackConfigurationLivePreRollConfiguration `pulumi:"livePreRollConfiguration"`
	// The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *PlaybackConfigurationManifestProcessingRules `pulumi:"manifestProcessingRules"`
	// The identifier for the playback configuration.
	Name *string `pulumi:"name"`
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	PersonalizationThresholdSeconds *int `pulumi:"personalizationThresholdSeconds"`
	// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `pulumi:"slateAdUrl"`
	// The tags to assign to the playback configuration.
	Tags []PlaybackConfigurationTag `pulumi:"tags"`
	// The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	TranscodeProfileName *string `pulumi:"transcodeProfileName"`
	// The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
	VideoContentSourceUrl string `pulumi:"videoContentSourceUrl"`
}

// The set of arguments for constructing a PlaybackConfiguration resource.
type PlaybackConfigurationArgs struct {
	// The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl pulumi.StringInput
	// The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	AvailSuppression PlaybackConfigurationAvailSuppressionPtrInput
	// The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
	Bumper PlaybackConfigurationBumperPtrInput
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	CdnConfiguration PlaybackConfigurationCdnConfigurationPtrInput
	// The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables.
	ConfigurationAliases PlaybackConfigurationConfigurationAliasesPtrInput
	// The configuration for DASH content.
	DashConfiguration PlaybackConfigurationDashConfigurationPtrInput
	// The configuration for HLS content.
	HlsConfiguration PlaybackConfigurationHlsConfigurationPtrInput
	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration PlaybackConfigurationLivePreRollConfigurationPtrInput
	// The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules PlaybackConfigurationManifestProcessingRulesPtrInput
	// The identifier for the playback configuration.
	Name pulumi.StringPtrInput
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	PersonalizationThresholdSeconds pulumi.IntPtrInput
	// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl pulumi.StringPtrInput
	// The tags to assign to the playback configuration.
	Tags PlaybackConfigurationTagArrayInput
	// The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	TranscodeProfileName pulumi.StringPtrInput
	// The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
	VideoContentSourceUrl pulumi.StringInput
}

func (PlaybackConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*playbackConfigurationArgs)(nil)).Elem()
}

type PlaybackConfigurationInput interface {
	pulumi.Input

	ToPlaybackConfigurationOutput() PlaybackConfigurationOutput
	ToPlaybackConfigurationOutputWithContext(ctx context.Context) PlaybackConfigurationOutput
}

func (*PlaybackConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**PlaybackConfiguration)(nil)).Elem()
}

func (i *PlaybackConfiguration) ToPlaybackConfigurationOutput() PlaybackConfigurationOutput {
	return i.ToPlaybackConfigurationOutputWithContext(context.Background())
}

func (i *PlaybackConfiguration) ToPlaybackConfigurationOutputWithContext(ctx context.Context) PlaybackConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlaybackConfigurationOutput)
}

type PlaybackConfigurationOutput struct{ *pulumi.OutputState }

func (PlaybackConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlaybackConfiguration)(nil)).Elem()
}

func (o PlaybackConfigurationOutput) ToPlaybackConfigurationOutput() PlaybackConfigurationOutput {
	return o
}

func (o PlaybackConfigurationOutput) ToPlaybackConfigurationOutputWithContext(ctx context.Context) PlaybackConfigurationOutput {
	return o
}

// The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
func (o PlaybackConfigurationOutput) AdDecisionServerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.StringOutput { return v.AdDecisionServerUrl }).(pulumi.StringOutput)
}

// The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
func (o PlaybackConfigurationOutput) AvailSuppression() PlaybackConfigurationAvailSuppressionPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationAvailSuppressionPtrOutput {
		return v.AvailSuppression
	}).(PlaybackConfigurationAvailSuppressionPtrOutput)
}

// The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
func (o PlaybackConfigurationOutput) Bumper() PlaybackConfigurationBumperPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationBumperPtrOutput { return v.Bumper }).(PlaybackConfigurationBumperPtrOutput)
}

// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
func (o PlaybackConfigurationOutput) CdnConfiguration() PlaybackConfigurationCdnConfigurationPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationCdnConfigurationPtrOutput {
		return v.CdnConfiguration
	}).(PlaybackConfigurationCdnConfigurationPtrOutput)
}

// The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables.
func (o PlaybackConfigurationOutput) ConfigurationAliases() PlaybackConfigurationConfigurationAliasesPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationConfigurationAliasesPtrOutput {
		return v.ConfigurationAliases
	}).(PlaybackConfigurationConfigurationAliasesPtrOutput)
}

// The configuration for DASH content.
func (o PlaybackConfigurationOutput) DashConfiguration() PlaybackConfigurationDashConfigurationPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationDashConfigurationPtrOutput {
		return v.DashConfiguration
	}).(PlaybackConfigurationDashConfigurationPtrOutput)
}

// The configuration for HLS content.
func (o PlaybackConfigurationOutput) HlsConfiguration() PlaybackConfigurationHlsConfigurationPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationHlsConfigurationPtrOutput {
		return v.HlsConfiguration
	}).(PlaybackConfigurationHlsConfigurationPtrOutput)
}

// The configuration for pre-roll ad insertion.
func (o PlaybackConfigurationOutput) LivePreRollConfiguration() PlaybackConfigurationLivePreRollConfigurationPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationLivePreRollConfigurationPtrOutput {
		return v.LivePreRollConfiguration
	}).(PlaybackConfigurationLivePreRollConfigurationPtrOutput)
}

// The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
func (o PlaybackConfigurationOutput) ManifestProcessingRules() PlaybackConfigurationManifestProcessingRulesPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationManifestProcessingRulesPtrOutput {
		return v.ManifestProcessingRules
	}).(PlaybackConfigurationManifestProcessingRulesPtrOutput)
}

// The identifier for the playback configuration.
func (o PlaybackConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
func (o PlaybackConfigurationOutput) PersonalizationThresholdSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.IntPtrOutput { return v.PersonalizationThresholdSeconds }).(pulumi.IntPtrOutput)
}

// The Amazon Resource Name (ARN) for the playback configuration.
func (o PlaybackConfigurationOutput) PlaybackConfigurationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.StringOutput { return v.PlaybackConfigurationArn }).(pulumi.StringOutput)
}

// The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.
func (o PlaybackConfigurationOutput) PlaybackEndpointPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.StringOutput { return v.PlaybackEndpointPrefix }).(pulumi.StringOutput)
}

// The URL that the player uses to initialize a session that uses client-side reporting.
func (o PlaybackConfigurationOutput) SessionInitializationEndpointPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.StringOutput { return v.SessionInitializationEndpointPrefix }).(pulumi.StringOutput)
}

// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
func (o PlaybackConfigurationOutput) SlateAdUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.StringPtrOutput { return v.SlateAdUrl }).(pulumi.StringPtrOutput)
}

// The tags to assign to the playback configuration.
func (o PlaybackConfigurationOutput) Tags() PlaybackConfigurationTagArrayOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) PlaybackConfigurationTagArrayOutput { return v.Tags }).(PlaybackConfigurationTagArrayOutput)
}

// The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
func (o PlaybackConfigurationOutput) TranscodeProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.StringPtrOutput { return v.TranscodeProfileName }).(pulumi.StringPtrOutput)
}

// The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
func (o PlaybackConfigurationOutput) VideoContentSourceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackConfiguration) pulumi.StringOutput { return v.VideoContentSourceUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlaybackConfigurationInput)(nil)).Elem(), &PlaybackConfiguration{})
	pulumi.RegisterOutputType(PlaybackConfigurationOutput{})
}