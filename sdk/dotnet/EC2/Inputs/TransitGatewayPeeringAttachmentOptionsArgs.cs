// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class TransitGatewayPeeringAttachmentOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable dynamic routing. (enable/disable)
        /// </summary>
        [Input("dynamicRouting")]
        public Input<string>? DynamicRouting { get; set; }

        public TransitGatewayPeeringAttachmentOptionsArgs()
        {
        }
    }
}