// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// Describes the request headers that a Lightsail distribution bases caching on.
    /// </summary>
    public sealed class DistributionHeaderObjectArgs : Pulumi.ResourceArgs
    {
        [Input("headersAllowList")]
        private InputList<string>? _headersAllowList;

        /// <summary>
        /// The specific headers to forward to your distribution's origin.
        /// </summary>
        public InputList<string> HeadersAllowList
        {
            get => _headersAllowList ?? (_headersAllowList = new InputList<string>());
            set => _headersAllowList = value;
        }

        /// <summary>
        /// The headers that you want your distribution to forward to your origin and base caching on.
        /// </summary>
        [Input("option")]
        public Input<string>? Option { get; set; }

        public DistributionHeaderObjectArgs()
        {
        }
    }
}