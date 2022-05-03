// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Inputs
{

    /// <summary>
    /// Contains the ARN of AWS IoT Greengrass Group V1 that the gateway runs on.
    /// </summary>
    public sealed class GatewayGreengrassArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Greengrass group.
        /// </summary>
        [Input("groupArn", required: true)]
        public Input<string> GroupArn { get; set; } = null!;

        public GatewayGreengrassArgs()
        {
        }
    }
}
