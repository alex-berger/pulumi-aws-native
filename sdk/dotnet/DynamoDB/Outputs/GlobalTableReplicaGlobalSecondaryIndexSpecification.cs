// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB.Outputs
{

    [OutputType]
    public sealed class GlobalTableReplicaGlobalSecondaryIndexSpecification
    {
        public readonly Outputs.GlobalTableContributorInsightsSpecification? ContributorInsightsSpecification;
        public readonly string IndexName;
        public readonly Outputs.GlobalTableReadProvisionedThroughputSettings? ReadProvisionedThroughputSettings;

        [OutputConstructor]
        private GlobalTableReplicaGlobalSecondaryIndexSpecification(
            Outputs.GlobalTableContributorInsightsSpecification? contributorInsightsSpecification,

            string indexName,

            Outputs.GlobalTableReadProvisionedThroughputSettings? readProvisionedThroughputSettings)
        {
            ContributorInsightsSpecification = contributorInsightsSpecification;
            IndexName = indexName;
            ReadProvisionedThroughputSettings = readProvisionedThroughputSettings;
        }
    }
}
