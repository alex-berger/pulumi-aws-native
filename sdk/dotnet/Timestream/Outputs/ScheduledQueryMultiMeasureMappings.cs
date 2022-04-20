// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Timestream.Outputs
{

    /// <summary>
    /// Only one of MixedMeasureMappings or MultiMeasureMappings is to be provided. MultiMeasureMappings can be used to ingest data as multi measures in the derived table.
    /// </summary>
    [OutputType]
    public sealed class ScheduledQueryMultiMeasureMappings
    {
        public readonly ImmutableArray<Outputs.ScheduledQueryMultiMeasureAttributeMapping> MultiMeasureAttributeMappings;
        public readonly string? TargetMultiMeasureName;

        [OutputConstructor]
        private ScheduledQueryMultiMeasureMappings(
            ImmutableArray<Outputs.ScheduledQueryMultiMeasureAttributeMapping> multiMeasureAttributeMappings,

            string? targetMultiMeasureName)
        {
            MultiMeasureAttributeMappings = multiMeasureAttributeMappings;
            TargetMultiMeasureName = targetMultiMeasureName;
        }
    }
}
