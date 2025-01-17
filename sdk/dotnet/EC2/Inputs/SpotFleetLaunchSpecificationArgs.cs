// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class SpotFleetLaunchSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("blockDeviceMappings")]
        private InputList<Inputs.SpotFleetBlockDeviceMappingArgs>? _blockDeviceMappings;
        public InputList<Inputs.SpotFleetBlockDeviceMappingArgs> BlockDeviceMappings
        {
            get => _blockDeviceMappings ?? (_blockDeviceMappings = new InputList<Inputs.SpotFleetBlockDeviceMappingArgs>());
            set => _blockDeviceMappings = value;
        }

        [Input("ebsOptimized")]
        public Input<bool>? EbsOptimized { get; set; }

        [Input("iamInstanceProfile")]
        public Input<Inputs.SpotFleetIamInstanceProfileSpecificationArgs>? IamInstanceProfile { get; set; }

        [Input("imageId", required: true)]
        public Input<string> ImageId { get; set; } = null!;

        [Input("instanceRequirements")]
        public Input<Inputs.SpotFleetInstanceRequirementsRequestArgs>? InstanceRequirements { get; set; }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("kernelId")]
        public Input<string>? KernelId { get; set; }

        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        [Input("monitoring")]
        public Input<Inputs.SpotFleetMonitoringArgs>? Monitoring { get; set; }

        [Input("networkInterfaces")]
        private InputList<Inputs.SpotFleetInstanceNetworkInterfaceSpecificationArgs>? _networkInterfaces;
        public InputList<Inputs.SpotFleetInstanceNetworkInterfaceSpecificationArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.SpotFleetInstanceNetworkInterfaceSpecificationArgs>());
            set => _networkInterfaces = value;
        }

        [Input("placement")]
        public Input<Inputs.SpotFleetSpotPlacementArgs>? Placement { get; set; }

        [Input("ramdiskId")]
        public Input<string>? RamdiskId { get; set; }

        [Input("securityGroups")]
        private InputList<Inputs.SpotFleetGroupIdentifierArgs>? _securityGroups;
        public InputList<Inputs.SpotFleetGroupIdentifierArgs> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<Inputs.SpotFleetGroupIdentifierArgs>());
            set => _securityGroups = value;
        }

        [Input("spotPrice")]
        public Input<string>? SpotPrice { get; set; }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tagSpecifications")]
        private InputList<Inputs.SpotFleetTagSpecificationArgs>? _tagSpecifications;
        public InputList<Inputs.SpotFleetTagSpecificationArgs> TagSpecifications
        {
            get => _tagSpecifications ?? (_tagSpecifications = new InputList<Inputs.SpotFleetTagSpecificationArgs>());
            set => _tagSpecifications = value;
        }

        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("weightedCapacity")]
        public Input<double>? WeightedCapacity { get; set; }

        public SpotFleetLaunchSpecificationArgs()
        {
        }
    }
}
