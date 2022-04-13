// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class EC2FleetSpotOptionsRequestArgs : Pulumi.ResourceArgs
    {
        [Input("allocationStrategy")]
        public Input<Pulumi.AwsNative.EC2.EC2FleetSpotOptionsRequestAllocationStrategy>? AllocationStrategy { get; set; }

        [Input("instanceInterruptionBehavior")]
        public Input<Pulumi.AwsNative.EC2.EC2FleetSpotOptionsRequestInstanceInterruptionBehavior>? InstanceInterruptionBehavior { get; set; }

        [Input("instancePoolsToUseCount")]
        public Input<int>? InstancePoolsToUseCount { get; set; }

        [Input("maintenanceStrategies")]
        public Input<Inputs.EC2FleetMaintenanceStrategiesArgs>? MaintenanceStrategies { get; set; }

        [Input("maxTotalPrice")]
        public Input<string>? MaxTotalPrice { get; set; }

        [Input("minTargetCapacity")]
        public Input<int>? MinTargetCapacity { get; set; }

        [Input("singleAvailabilityZone")]
        public Input<bool>? SingleAvailabilityZone { get; set; }

        [Input("singleInstanceType")]
        public Input<bool>? SingleInstanceType { get; set; }

        public EC2FleetSpotOptionsRequestArgs()
        {
        }
    }
}
