// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Outputs
{

    /// <summary>
    /// A portal resource.
    /// </summary>
    [OutputType]
    public sealed class AccessPolicyPortal
    {
        /// <summary>
        /// The ID of the portal.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private AccessPolicyPortal(string? id)
        {
            Id = id;
        }
    }
}
