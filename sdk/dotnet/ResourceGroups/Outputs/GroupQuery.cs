// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResourceGroups.Outputs
{

    [OutputType]
    public sealed class GroupQuery
    {
        public readonly ImmutableArray<string> ResourceTypeFilters;
        public readonly string? StackIdentifier;
        public readonly ImmutableArray<Outputs.GroupTagFilter> TagFilters;

        [OutputConstructor]
        private GroupQuery(
            ImmutableArray<string> resourceTypeFilters,

            string? stackIdentifier,

            ImmutableArray<Outputs.GroupTagFilter> tagFilters)
        {
            ResourceTypeFilters = resourceTypeFilters;
            StackIdentifier = stackIdentifier;
            TagFilters = tagFilters;
        }
    }
}