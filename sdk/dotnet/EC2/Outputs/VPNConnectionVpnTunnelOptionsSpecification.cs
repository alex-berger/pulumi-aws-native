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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html
    /// </summary>
    [OutputType]
    public sealed class VPNConnectionVpnTunnelOptionsSpecification
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-presharedkey
        /// </summary>
        public readonly string? PreSharedKey;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-vpntunneloptionsspecification.html#cfn-ec2-vpnconnection-vpntunneloptionsspecification-tunnelinsidecidr
        /// </summary>
        public readonly string? TunnelInsideCidr;

        [OutputConstructor]
        private VPNConnectionVpnTunnelOptionsSpecification(
            string? preSharedKey,

            string? tunnelInsideCidr)
        {
            PreSharedKey = preSharedKey;
            TunnelInsideCidr = tunnelInsideCidr;
        }
    }
}