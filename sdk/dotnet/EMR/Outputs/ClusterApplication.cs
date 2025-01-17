// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Outputs
{

    [OutputType]
    public sealed class ClusterApplication
    {
        public readonly object? AdditionalInfo;
        public readonly ImmutableArray<string> Args;
        public readonly string? Name;
        public readonly string? Version;

        [OutputConstructor]
        private ClusterApplication(
            object? additionalInfo,

            ImmutableArray<string> args,

            string? name,

            string? version)
        {
            AdditionalInfo = additionalInfo;
            Args = args;
            Name = name;
            Version = version;
        }
    }
}
