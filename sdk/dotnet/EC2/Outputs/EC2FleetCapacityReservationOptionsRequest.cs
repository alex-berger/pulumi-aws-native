// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    [OutputType]
    public sealed class EC2FleetCapacityReservationOptionsRequest
    {
        public readonly Pulumi.AwsNative.EC2.EC2FleetCapacityReservationOptionsRequestUsageStrategy? UsageStrategy;

        [OutputConstructor]
        private EC2FleetCapacityReservationOptionsRequest(Pulumi.AwsNative.EC2.EC2FleetCapacityReservationOptionsRequestUsageStrategy? usageStrategy)
        {
            UsageStrategy = usageStrategy;
        }
    }
}
