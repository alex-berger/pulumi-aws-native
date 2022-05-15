// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class DatasetFilesLimit
    {
        /// <summary>
        /// Maximum number of files
        /// </summary>
        public readonly int MaxFiles;
        /// <summary>
        /// Order
        /// </summary>
        public readonly Pulumi.AwsNative.DataBrew.DatasetFilesLimitOrder? Order;
        /// <summary>
        /// Ordered by
        /// </summary>
        public readonly Pulumi.AwsNative.DataBrew.DatasetFilesLimitOrderedBy? OrderedBy;

        [OutputConstructor]
        private DatasetFilesLimit(
            int maxFiles,

            Pulumi.AwsNative.DataBrew.DatasetFilesLimitOrder? order,

            Pulumi.AwsNative.DataBrew.DatasetFilesLimitOrderedBy? orderedBy)
        {
            MaxFiles = maxFiles;
            Order = order;
            OrderedBy = orderedBy;
        }
    }
}
