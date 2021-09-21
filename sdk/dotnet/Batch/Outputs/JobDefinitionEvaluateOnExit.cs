// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Outputs
{

    [OutputType]
    public sealed class JobDefinitionEvaluateOnExit
    {
        public readonly string Action;
        public readonly string? OnExitCode;
        public readonly string? OnReason;
        public readonly string? OnStatusReason;

        [OutputConstructor]
        private JobDefinitionEvaluateOnExit(
            string action,

            string? onExitCode,

            string? onReason,

            string? onStatusReason)
        {
            Action = action;
            OnExitCode = onExitCode;
            OnReason = onReason;
            OnStatusReason = onStatusReason;
        }
    }
}