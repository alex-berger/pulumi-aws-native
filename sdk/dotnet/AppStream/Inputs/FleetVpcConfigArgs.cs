// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-vpcconfig.html
    /// </summary>
    public sealed class FleetVpcConfigArgs : Pulumi.ResourceArgs
    {
        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-vpcconfig.html#cfn-appstream-fleet-vpcconfig-securitygroupids
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-vpcconfig.html#cfn-appstream-fleet-vpcconfig-subnetids
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public FleetVpcConfigArgs()
        {
        }
    }
}
