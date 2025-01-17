// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class TransitGatewayPeeringAttachmentTagArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public TransitGatewayPeeringAttachmentTagArgs()
        {
        }
    }
}
