// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionHostEntry
    {
        public readonly string? Hostname;
        public readonly string? IpAddress;

        [OutputConstructor]
        private TaskDefinitionHostEntry(
            string? hostname,

            string? ipAddress)
        {
            Hostname = hostname;
            IpAddress = ipAddress;
        }
    }
}
