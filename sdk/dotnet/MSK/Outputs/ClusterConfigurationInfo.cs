// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MSK.Outputs
{

    [OutputType]
    public sealed class ClusterConfigurationInfo
    {
        public readonly string Arn;
        public readonly int Revision;

        [OutputConstructor]
        private ClusterConfigurationInfo(
            string arn,

            int revision)
        {
            Arn = arn;
            Revision = revision;
        }
    }
}