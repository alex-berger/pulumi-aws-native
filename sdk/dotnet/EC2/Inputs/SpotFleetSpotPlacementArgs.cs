// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class SpotFleetSpotPlacementArgs : Pulumi.ResourceArgs
    {
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("tenancy")]
        public Input<Pulumi.AwsNative.EC2.SpotFleetSpotPlacementTenancy>? Tenancy { get; set; }

        public SpotFleetSpotPlacementArgs()
        {
        }
    }
}
