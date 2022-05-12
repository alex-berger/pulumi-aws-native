// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class ComputeEnvironmentComputeResourcesArgs : Pulumi.ResourceArgs
    {
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        [Input("bidPercentage")]
        public Input<int>? BidPercentage { get; set; }

        [Input("desiredvCpus")]
        public Input<int>? DesiredvCpus { get; set; }

        [Input("ec2Configuration")]
        private InputList<Inputs.ComputeEnvironmentEc2ConfigurationObjectArgs>? _ec2Configuration;
        public InputList<Inputs.ComputeEnvironmentEc2ConfigurationObjectArgs> Ec2Configuration
        {
            get => _ec2Configuration ?? (_ec2Configuration = new InputList<Inputs.ComputeEnvironmentEc2ConfigurationObjectArgs>());
            set => _ec2Configuration = value;
        }

        [Input("ec2KeyPair")]
        public Input<string>? Ec2KeyPair { get; set; }

        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        [Input("instanceRole")]
        public Input<string>? InstanceRole { get; set; }

        [Input("instanceTypes")]
        private InputList<string>? _instanceTypes;
        public InputList<string> InstanceTypes
        {
            get => _instanceTypes ?? (_instanceTypes = new InputList<string>());
            set => _instanceTypes = value;
        }

        [Input("launchTemplate")]
        public Input<Inputs.ComputeEnvironmentLaunchTemplateSpecificationArgs>? LaunchTemplate { get; set; }

        [Input("maxvCpus", required: true)]
        public Input<int> MaxvCpus { get; set; } = null!;

        [Input("minvCpus")]
        public Input<int>? MinvCpus { get; set; }

        [Input("placementGroup")]
        public Input<string>? PlacementGroup { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("spotIamFleetRole")]
        public Input<string>? SpotIamFleetRole { get; set; }

        [Input("subnets", required: true)]
        private InputList<string>? _subnets;
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        /// <summary>
        /// A key-value pair to associate with a resource.
        /// </summary>
        [Input("tags")]
        public Input<object>? Tags { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("updateToLatestImageVersion")]
        public Input<bool>? UpdateToLatestImageVersion { get; set; }

        public ComputeEnvironmentComputeResourcesArgs()
        {
        }
    }
}
