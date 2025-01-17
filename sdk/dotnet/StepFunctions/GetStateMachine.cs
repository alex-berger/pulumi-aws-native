// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.StepFunctions
{
    public static class GetStateMachine
    {
        /// <summary>
        /// Resource schema for StateMachine
        /// </summary>
        public static Task<GetStateMachineResult> InvokeAsync(GetStateMachineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStateMachineResult>("aws-native:stepfunctions:getStateMachine", args ?? new GetStateMachineArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for StateMachine
        /// </summary>
        public static Output<GetStateMachineResult> Invoke(GetStateMachineInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStateMachineResult>("aws-native:stepfunctions:getStateMachine", args ?? new GetStateMachineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStateMachineArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetStateMachineArgs()
        {
        }
    }

    public sealed class GetStateMachineInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetStateMachineInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStateMachineResult
    {
        public readonly string? Arn;
        public readonly string? DefinitionString;
        public readonly Outputs.StateMachineLoggingConfiguration? LoggingConfiguration;
        public readonly string? Name;
        public readonly string? RoleArn;
        public readonly ImmutableArray<Outputs.StateMachineTagsEntry> Tags;
        public readonly Outputs.StateMachineTracingConfiguration? TracingConfiguration;

        [OutputConstructor]
        private GetStateMachineResult(
            string? arn,

            string? definitionString,

            Outputs.StateMachineLoggingConfiguration? loggingConfiguration,

            string? name,

            string? roleArn,

            ImmutableArray<Outputs.StateMachineTagsEntry> tags,

            Outputs.StateMachineTracingConfiguration? tracingConfiguration)
        {
            Arn = arn;
            DefinitionString = definitionString;
            LoggingConfiguration = loggingConfiguration;
            Name = name;
            RoleArn = roleArn;
            Tags = tags;
            TracingConfiguration = tracingConfiguration;
        }
    }
}
