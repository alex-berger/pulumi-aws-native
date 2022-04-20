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
    public sealed class EC2FleetTagSpecification
    {
        public readonly Pulumi.AwsNative.EC2.EC2FleetTagSpecificationResourceType? ResourceType;
        public readonly ImmutableArray<Outputs.EC2FleetTag> Tags;

        [OutputConstructor]
        private EC2FleetTagSpecification(
            Pulumi.AwsNative.EC2.EC2FleetTagSpecificationResourceType? resourceType,

            ImmutableArray<Outputs.EC2FleetTag> tags)
        {
            ResourceType = resourceType;
            Tags = tags;
        }
    }
}
