// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html
    /// </summary>
    public sealed class ChannelH264SettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-adaptivequantization
        /// </summary>
        [Input("adaptiveQuantization")]
        public Input<string>? AdaptiveQuantization { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-afdsignaling
        /// </summary>
        [Input("afdSignaling")]
        public Input<string>? AfdSignaling { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-bitrate
        /// </summary>
        [Input("bitrate")]
        public Input<int>? Bitrate { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-buffillpct
        /// </summary>
        [Input("bufFillPct")]
        public Input<int>? BufFillPct { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-bufsize
        /// </summary>
        [Input("bufSize")]
        public Input<int>? BufSize { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-colormetadata
        /// </summary>
        [Input("colorMetadata")]
        public Input<string>? ColorMetadata { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-colorspacesettings
        /// </summary>
        [Input("colorSpaceSettings")]
        public Input<Inputs.ChannelH264ColorSpaceSettingsArgs>? ColorSpaceSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-entropyencoding
        /// </summary>
        [Input("entropyEncoding")]
        public Input<string>? EntropyEncoding { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-filtersettings
        /// </summary>
        [Input("filterSettings")]
        public Input<Inputs.ChannelH264FilterSettingsArgs>? FilterSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-fixedafd
        /// </summary>
        [Input("fixedAfd")]
        public Input<string>? FixedAfd { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-flickeraq
        /// </summary>
        [Input("flickerAq")]
        public Input<string>? FlickerAq { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-forcefieldpictures
        /// </summary>
        [Input("forceFieldPictures")]
        public Input<string>? ForceFieldPictures { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-frameratecontrol
        /// </summary>
        [Input("framerateControl")]
        public Input<string>? FramerateControl { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-frameratedenominator
        /// </summary>
        [Input("framerateDenominator")]
        public Input<int>? FramerateDenominator { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-frameratenumerator
        /// </summary>
        [Input("framerateNumerator")]
        public Input<int>? FramerateNumerator { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopbreference
        /// </summary>
        [Input("gopBReference")]
        public Input<string>? GopBReference { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopclosedcadence
        /// </summary>
        [Input("gopClosedCadence")]
        public Input<int>? GopClosedCadence { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopnumbframes
        /// </summary>
        [Input("gopNumBFrames")]
        public Input<int>? GopNumBFrames { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopsize
        /// </summary>
        [Input("gopSize")]
        public Input<double>? GopSize { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-gopsizeunits
        /// </summary>
        [Input("gopSizeUnits")]
        public Input<string>? GopSizeUnits { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-level
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-lookaheadratecontrol
        /// </summary>
        [Input("lookAheadRateControl")]
        public Input<string>? LookAheadRateControl { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-maxbitrate
        /// </summary>
        [Input("maxBitrate")]
        public Input<int>? MaxBitrate { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-miniinterval
        /// </summary>
        [Input("minIInterval")]
        public Input<int>? MinIInterval { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-numrefframes
        /// </summary>
        [Input("numRefFrames")]
        public Input<int>? NumRefFrames { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-parcontrol
        /// </summary>
        [Input("parControl")]
        public Input<string>? ParControl { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-pardenominator
        /// </summary>
        [Input("parDenominator")]
        public Input<int>? ParDenominator { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-parnumerator
        /// </summary>
        [Input("parNumerator")]
        public Input<int>? ParNumerator { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-profile
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-qualitylevel
        /// </summary>
        [Input("qualityLevel")]
        public Input<string>? QualityLevel { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-qvbrqualitylevel
        /// </summary>
        [Input("qvbrQualityLevel")]
        public Input<int>? QvbrQualityLevel { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-ratecontrolmode
        /// </summary>
        [Input("rateControlMode")]
        public Input<string>? RateControlMode { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-scantype
        /// </summary>
        [Input("scanType")]
        public Input<string>? ScanType { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-scenechangedetect
        /// </summary>
        [Input("sceneChangeDetect")]
        public Input<string>? SceneChangeDetect { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-slices
        /// </summary>
        [Input("slices")]
        public Input<int>? Slices { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-softness
        /// </summary>
        [Input("softness")]
        public Input<int>? Softness { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-spatialaq
        /// </summary>
        [Input("spatialAq")]
        public Input<string>? SpatialAq { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-subgoplength
        /// </summary>
        [Input("subgopLength")]
        public Input<string>? SubgopLength { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-syntax
        /// </summary>
        [Input("syntax")]
        public Input<string>? Syntax { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-temporalaq
        /// </summary>
        [Input("temporalAq")]
        public Input<string>? TemporalAq { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264settings.html#cfn-medialive-channel-h264settings-timecodeinsertion
        /// </summary>
        [Input("timecodeInsertion")]
        public Input<string>? TimecodeInsertion { get; set; }

        public ChannelH264SettingsArgs()
        {
        }
    }
}