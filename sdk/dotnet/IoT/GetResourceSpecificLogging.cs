// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    public static class GetResourceSpecificLogging
    {
        /// <summary>
        /// Resource-specific logging allows you to specify a logging level for a specific thing group.
        /// </summary>
        public static Task<GetResourceSpecificLoggingResult> InvokeAsync(GetResourceSpecificLoggingArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourceSpecificLoggingResult>("aws-native:iot:getResourceSpecificLogging", args ?? new GetResourceSpecificLoggingArgs(), options.WithDefaults());

        /// <summary>
        /// Resource-specific logging allows you to specify a logging level for a specific thing group.
        /// </summary>
        public static Output<GetResourceSpecificLoggingResult> Invoke(GetResourceSpecificLoggingInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourceSpecificLoggingResult>("aws-native:iot:getResourceSpecificLogging", args ?? new GetResourceSpecificLoggingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourceSpecificLoggingArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.
        /// </summary>
        [Input("targetId", required: true)]
        public string TargetId { get; set; } = null!;

        public GetResourceSpecificLoggingArgs()
        {
        }
    }

    public sealed class GetResourceSpecificLoggingInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.
        /// </summary>
        [Input("targetId", required: true)]
        public Input<string> TargetId { get; set; } = null!;

        public GetResourceSpecificLoggingInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourceSpecificLoggingResult
    {
        /// <summary>
        /// The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.
        /// </summary>
        public readonly Pulumi.AwsNative.IoT.ResourceSpecificLoggingLogLevel? LogLevel;
        /// <summary>
        /// Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.
        /// </summary>
        public readonly string? TargetId;

        [OutputConstructor]
        private GetResourceSpecificLoggingResult(
            Pulumi.AwsNative.IoT.ResourceSpecificLoggingLogLevel? logLevel,

            string? targetId)
        {
            LogLevel = logLevel;
            TargetId = targetId;
        }
    }
}
