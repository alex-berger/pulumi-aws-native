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
    /// A DASH manifest configuration.
    /// </summary>
    public sealed class PackagingConfigurationDashManifestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines the position of some tags in the Media Presentation Description (MPD). When set to FULL, elements like SegmentTemplate and ContentProtection are included in each Representation. When set to COMPACT, duplicate elements are combined and presented at the AdaptationSet level.
        /// </summary>
        [Input("manifestLayout")]
        public Input<Pulumi.AwsNative.MediaPackage.PackagingConfigurationDashManifestManifestLayout>? ManifestLayout { get; set; }

        [Input("manifestName")]
        public Input<string>? ManifestName { get; set; }

        /// <summary>
        /// Minimum duration (in seconds) that a player will buffer media before starting the presentation.
        /// </summary>
        [Input("minBufferTimeSeconds")]
        public Input<int>? MinBufferTimeSeconds { get; set; }

        /// <summary>
        /// The Dynamic Adaptive Streaming over HTTP (DASH) profile type. When set to "HBBTV_1_5", HbbTV 1.5 compliant output is enabled.
        /// </summary>
        [Input("profile")]
        public Input<Pulumi.AwsNative.MediaPackage.PackagingConfigurationDashManifestProfile>? Profile { get; set; }

        /// <summary>
        /// The source of scte markers used. When set to SEGMENTS, the scte markers are sourced from the segments of the ingested content. When set to MANIFEST, the scte markers are sourced from the manifest of the ingested content.
        /// </summary>
        [Input("scteMarkersSource")]
        public Input<Pulumi.AwsNative.MediaPackage.PackagingConfigurationDashManifestScteMarkersSource>? ScteMarkersSource { get; set; }

        [Input("streamSelection")]
        public Input<Inputs.PackagingConfigurationStreamSelectionArgs>? StreamSelection { get; set; }

        public PackagingConfigurationDashManifestArgs()
        {
        }
    }
}
