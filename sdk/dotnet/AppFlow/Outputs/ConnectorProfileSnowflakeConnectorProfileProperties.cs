// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class ConnectorProfileSnowflakeConnectorProfileProperties
    {
        /// <summary>
        /// The name of the account.
        /// </summary>
        public readonly string? AccountName;
        /// <summary>
        /// The name of the Amazon S3 bucket associated with Snowﬂake.
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// The bucket prefix that refers to the Amazon S3 bucket associated with Snowﬂake.
        /// </summary>
        public readonly string? BucketPrefix;
        /// <summary>
        /// The Snowﬂake Private Link service name to be used for private data transfers.
        /// </summary>
        public readonly string? PrivateLinkServiceName;
        /// <summary>
        /// The region of the Snowﬂake account.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the
        /// Snowﬂake account. This is written in the following format: &lt; Database&gt;&lt; Schema&gt;&lt;Stage Name&gt;.
        /// </summary>
        public readonly string Stage;
        /// <summary>
        /// The name of the Snowﬂake warehouse.
        /// </summary>
        public readonly string Warehouse;

        [OutputConstructor]
        private ConnectorProfileSnowflakeConnectorProfileProperties(
            string? accountName,

            string bucketName,

            string? bucketPrefix,

            string? privateLinkServiceName,

            string? region,

            string stage,

            string warehouse)
        {
            AccountName = accountName;
            BucketName = bucketName;
            BucketPrefix = bucketPrefix;
            PrivateLinkServiceName = privateLinkServiceName;
            Region = region;
            Stage = stage;
            Warehouse = warehouse;
        }
    }
}
