// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetTrigger
    {
        /// <summary>
        /// Resource Type definition for AWS::Glue::Trigger
        /// </summary>
        public static Task<GetTriggerResult> InvokeAsync(GetTriggerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTriggerResult>("aws-native:glue:getTrigger", args ?? new GetTriggerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Glue::Trigger
        /// </summary>
        public static Output<GetTriggerResult> Invoke(GetTriggerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTriggerResult>("aws-native:glue:getTrigger", args ?? new GetTriggerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTriggerArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTriggerArgs()
        {
        }
    }

    public sealed class GetTriggerInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTriggerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTriggerResult
    {
        public readonly ImmutableArray<Outputs.TriggerAction> Actions;
        public readonly string? Description;
        public readonly string? Id;
        public readonly Outputs.TriggerPredicate? Predicate;
        public readonly string? Schedule;
        public readonly bool? StartOnCreation;
        public readonly object? Tags;
        public readonly string? Type;

        [OutputConstructor]
        private GetTriggerResult(
            ImmutableArray<Outputs.TriggerAction> actions,

            string? description,

            string? id,

            Outputs.TriggerPredicate? predicate,

            string? schedule,

            bool? startOnCreation,

            object? tags,

            string? type)
        {
            Actions = actions;
            Description = description;
            Id = id;
            Predicate = predicate;
            Schedule = schedule;
            StartOnCreation = startOnCreation;
            Tags = tags;
            Type = type;
        }
    }
}