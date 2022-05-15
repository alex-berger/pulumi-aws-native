// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelInputSettings
    {
        public readonly ImmutableArray<Outputs.ChannelAudioSelector> AudioSelectors;
        public readonly ImmutableArray<Outputs.ChannelCaptionSelector> CaptionSelectors;
        public readonly string? DeblockFilter;
        public readonly string? DenoiseFilter;
        public readonly int? FilterStrength;
        public readonly string? InputFilter;
        public readonly Outputs.ChannelNetworkInputSettings? NetworkInputSettings;
        public readonly int? Scte35Pid;
        public readonly string? Smpte2038DataPreference;
        public readonly string? SourceEndBehavior;
        public readonly Outputs.ChannelVideoSelector? VideoSelector;

        [OutputConstructor]
        private ChannelInputSettings(
            ImmutableArray<Outputs.ChannelAudioSelector> audioSelectors,

            ImmutableArray<Outputs.ChannelCaptionSelector> captionSelectors,

            string? deblockFilter,

            string? denoiseFilter,

            int? filterStrength,

            string? inputFilter,

            Outputs.ChannelNetworkInputSettings? networkInputSettings,

            int? scte35Pid,

            string? smpte2038DataPreference,

            string? sourceEndBehavior,

            Outputs.ChannelVideoSelector? videoSelector)
        {
            AudioSelectors = audioSelectors;
            CaptionSelectors = captionSelectors;
            DeblockFilter = deblockFilter;
            DenoiseFilter = denoiseFilter;
            FilterStrength = filterStrength;
            InputFilter = inputFilter;
            NetworkInputSettings = networkInputSettings;
            Scte35Pid = scte35Pid;
            Smpte2038DataPreference = smpte2038DataPreference;
            SourceEndBehavior = sourceEndBehavior;
            VideoSelector = videoSelector;
        }
    }
}
