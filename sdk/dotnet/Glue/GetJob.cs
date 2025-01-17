// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetJob
    {
        /// <summary>
        /// Resource Type definition for AWS::Glue::Job
        /// </summary>
        public static Task<GetJobResult> InvokeAsync(GetJobArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobResult>("aws-native:glue:getJob", args ?? new GetJobArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Glue::Job
        /// </summary>
        public static Output<GetJobResult> Invoke(GetJobInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetJobResult>("aws-native:glue:getJob", args ?? new GetJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetJobArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetJobArgs()
        {
        }
    }

    public sealed class GetJobInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetJobInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetJobResult
    {
        public readonly double? AllocatedCapacity;
        public readonly Outputs.JobCommand? Command;
        public readonly Outputs.JobConnectionsList? Connections;
        public readonly object? DefaultArguments;
        public readonly string? Description;
        public readonly Outputs.JobExecutionProperty? ExecutionProperty;
        public readonly string? GlueVersion;
        public readonly string? Id;
        public readonly string? LogUri;
        public readonly double? MaxCapacity;
        public readonly double? MaxRetries;
        public readonly Outputs.JobNotificationProperty? NotificationProperty;
        public readonly int? NumberOfWorkers;
        public readonly string? Role;
        public readonly string? SecurityConfiguration;
        public readonly object? Tags;
        public readonly int? Timeout;
        public readonly string? WorkerType;

        [OutputConstructor]
        private GetJobResult(
            double? allocatedCapacity,

            Outputs.JobCommand? command,

            Outputs.JobConnectionsList? connections,

            object? defaultArguments,

            string? description,

            Outputs.JobExecutionProperty? executionProperty,

            string? glueVersion,

            string? id,

            string? logUri,

            double? maxCapacity,

            double? maxRetries,

            Outputs.JobNotificationProperty? notificationProperty,

            int? numberOfWorkers,

            string? role,

            string? securityConfiguration,

            object? tags,

            int? timeout,

            string? workerType)
        {
            AllocatedCapacity = allocatedCapacity;
            Command = command;
            Connections = connections;
            DefaultArguments = defaultArguments;
            Description = description;
            ExecutionProperty = executionProperty;
            GlueVersion = glueVersion;
            Id = id;
            LogUri = logUri;
            MaxCapacity = maxCapacity;
            MaxRetries = maxRetries;
            NotificationProperty = notificationProperty;
            NumberOfWorkers = numberOfWorkers;
            Role = role;
            SecurityConfiguration = securityConfiguration;
            Tags = tags;
            Timeout = timeout;
            WorkerType = workerType;
        }
    }
}
