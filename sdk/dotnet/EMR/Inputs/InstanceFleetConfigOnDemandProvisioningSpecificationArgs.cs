// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Inputs
{

    public sealed class InstanceFleetConfigOnDemandProvisioningSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("allocationStrategy", required: true)]
        public Input<string> AllocationStrategy { get; set; } = null!;

        public InstanceFleetConfigOnDemandProvisioningSpecificationArgs()
        {
        }
    }
}
