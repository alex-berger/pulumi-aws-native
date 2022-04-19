// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class DomainElasticsearchClusterConfig
    {
        public readonly Outputs.DomainColdStorageOptions? ColdStorageOptions;
        public readonly int? DedicatedMasterCount;
        public readonly bool? DedicatedMasterEnabled;
        public readonly string? DedicatedMasterType;
        public readonly int? InstanceCount;
        public readonly string? InstanceType;
        public readonly int? WarmCount;
        public readonly bool? WarmEnabled;
        public readonly string? WarmType;
        public readonly Outputs.DomainZoneAwarenessConfig? ZoneAwarenessConfig;
        public readonly bool? ZoneAwarenessEnabled;

        [OutputConstructor]
        private DomainElasticsearchClusterConfig(
            Outputs.DomainColdStorageOptions? coldStorageOptions,

            int? dedicatedMasterCount,

            bool? dedicatedMasterEnabled,

            string? dedicatedMasterType,

            int? instanceCount,

            string? instanceType,

            int? warmCount,

            bool? warmEnabled,

            string? warmType,

            Outputs.DomainZoneAwarenessConfig? zoneAwarenessConfig,

            bool? zoneAwarenessEnabled)
        {
            ColdStorageOptions = coldStorageOptions;
            DedicatedMasterCount = dedicatedMasterCount;
            DedicatedMasterEnabled = dedicatedMasterEnabled;
            DedicatedMasterType = dedicatedMasterType;
            InstanceCount = instanceCount;
            InstanceType = instanceType;
            WarmCount = warmCount;
            WarmEnabled = warmEnabled;
            WarmType = warmType;
            ZoneAwarenessConfig = zoneAwarenessConfig;
            ZoneAwarenessEnabled = zoneAwarenessEnabled;
        }
    }
}
