// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Inputs
{

    /// <summary>
    /// An HTTP Live Streaming (HLS) packaging configuration.
    /// </summary>
    public sealed class PackagingConfigurationHlsPackageArgs : Pulumi.ResourceArgs
    {
        [Input("encryption")]
        public Input<Inputs.PackagingConfigurationHlsEncryptionArgs>? Encryption { get; set; }

        [Input("hlsManifests", required: true)]
        private InputList<Inputs.PackagingConfigurationHlsManifestArgs>? _hlsManifests;

        /// <summary>
        /// A list of HLS manifest configurations.
        /// </summary>
        public InputList<Inputs.PackagingConfigurationHlsManifestArgs> HlsManifests
        {
            get => _hlsManifests ?? (_hlsManifests = new InputList<Inputs.PackagingConfigurationHlsManifestArgs>());
            set => _hlsManifests = value;
        }

        [Input("segmentDurationSeconds")]
        public Input<int>? SegmentDurationSeconds { get; set; }

        /// <summary>
        /// When enabled, audio streams will be placed in rendition groups in the output.
        /// </summary>
        [Input("useAudioRenditionGroup")]
        public Input<bool>? UseAudioRenditionGroup { get; set; }

        public PackagingConfigurationHlsPackageArgs()
        {
        }
    }
}
