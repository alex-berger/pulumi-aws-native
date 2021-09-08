// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html
    /// </summary>
    [OutputType]
    public sealed class NetworkInterfaceInstanceIpv6Address
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html#cfn-ec2-networkinterface-instanceipv6address-ipv6address
        /// </summary>
        public readonly string Ipv6Address;

        [OutputConstructor]
        private NetworkInterfaceInstanceIpv6Address(string ipv6Address)
        {
            Ipv6Address = ipv6Address;
        }
    }
}
