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
    /// An HTTP Live Streaming (HLS) encryption configuration.
    /// </summary>
    public sealed class OriginEndpointHlsEncryptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A constant initialization vector for encryption (optional). When not specified the initialization vector will be periodically rotated.
        /// </summary>
        [Input("constantInitializationVector")]
        public Input<string>? ConstantInitializationVector { get; set; }

        /// <summary>
        /// The encryption method to use.
        /// </summary>
        [Input("encryptionMethod")]
        public Input<Pulumi.AwsNative.MediaPackage.OriginEndpointHlsEncryptionEncryptionMethod>? EncryptionMethod { get; set; }

        /// <summary>
        /// Interval (in seconds) between each encryption key rotation.
        /// </summary>
        [Input("keyRotationIntervalSeconds")]
        public Input<int>? KeyRotationIntervalSeconds { get; set; }

        /// <summary>
        /// When enabled, the EXT-X-KEY tag will be repeated in output manifests.
        /// </summary>
        [Input("repeatExtXKey")]
        public Input<bool>? RepeatExtXKey { get; set; }

        [Input("spekeKeyProvider", required: true)]
        public Input<Inputs.OriginEndpointSpekeKeyProviderArgs> SpekeKeyProvider { get; set; } = null!;

        public OriginEndpointHlsEncryptionArgs()
        {
        }
    }
}