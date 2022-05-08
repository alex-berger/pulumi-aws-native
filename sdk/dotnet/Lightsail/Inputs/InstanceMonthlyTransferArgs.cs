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
    /// Monthly Transfer of the Instance.
    /// </summary>
    public sealed class InstanceMonthlyTransferArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// GbPerMonthAllocated of the Instance.
        /// </summary>
        [Input("gbPerMonthAllocated")]
        public Input<string>? GbPerMonthAllocated { get; set; }

        public InstanceMonthlyTransferArgs()
        {
        }
    }
}
