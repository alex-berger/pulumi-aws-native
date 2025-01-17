// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager.Inputs
{

    /// <summary>
    /// The bandwidth for the link.
    /// </summary>
    public sealed class LinkBandwidthArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Download speed in Mbps.
        /// </summary>
        [Input("downloadSpeed")]
        public Input<int>? DownloadSpeed { get; set; }

        /// <summary>
        /// Upload speed in Mbps.
        /// </summary>
        [Input("uploadSpeed")]
        public Input<int>? UploadSpeed { get; set; }

        public LinkBandwidthArgs()
        {
        }
    }
}
