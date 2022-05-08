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
    public sealed class TaskDefinitionPortMapping
    {
        public readonly int? ContainerPort;
        public readonly int? HostPort;
        public readonly string? Protocol;

        [OutputConstructor]
        private TaskDefinitionPortMapping(
            int? containerPort,

            int? hostPort,

            string? protocol)
        {
            ContainerPort = containerPort;
            HostPort = hostPort;
            Protocol = protocol;
        }
    }
}
