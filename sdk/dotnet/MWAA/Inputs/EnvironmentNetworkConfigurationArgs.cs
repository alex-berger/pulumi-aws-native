// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MWAA.Inputs
{

    /// <summary>
    /// Configures the network resources of the environment.
    /// </summary>
    public sealed class EnvironmentNetworkConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of security groups to use for the environment.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of subnets to use for the environment. These must be private subnets, in the same VPC, in two different availability zones.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public EnvironmentNetworkConfigurationArgs()
        {
        }
    }
}
