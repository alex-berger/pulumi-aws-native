// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Outputs
{

    /// <summary>
    /// The settings for attaching a VPC interface to an output.
    /// </summary>
    [OutputType]
    public sealed class FlowOutputVpcInterfaceAttachment
    {
        /// <summary>
        /// The name of the VPC interface to use for this output.
        /// </summary>
        public readonly string? VpcInterfaceName;

        [OutputConstructor]
        private FlowOutputVpcInterfaceAttachment(string? vpcInterfaceName)
        {
            VpcInterfaceName = vpcInterfaceName;
        }
    }
}
