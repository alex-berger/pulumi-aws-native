// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    public sealed class RuleTargetArgs : Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("batchParameters")]
        public Input<Inputs.RuleBatchParametersArgs>? BatchParameters { get; set; }

        [Input("deadLetterConfig")]
        public Input<Inputs.RuleDeadLetterConfigArgs>? DeadLetterConfig { get; set; }

        [Input("ecsParameters")]
        public Input<Inputs.RuleEcsParametersArgs>? EcsParameters { get; set; }

        [Input("httpParameters")]
        public Input<Inputs.RuleHttpParametersArgs>? HttpParameters { get; set; }

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("input")]
        public Input<string>? Input { get; set; }

        [Input("inputPath")]
        public Input<string>? InputPath { get; set; }

        [Input("inputTransformer")]
        public Input<Inputs.RuleInputTransformerArgs>? InputTransformer { get; set; }

        [Input("kinesisParameters")]
        public Input<Inputs.RuleKinesisParametersArgs>? KinesisParameters { get; set; }

        [Input("redshiftDataParameters")]
        public Input<Inputs.RuleRedshiftDataParametersArgs>? RedshiftDataParameters { get; set; }

        [Input("retryPolicy")]
        public Input<Inputs.RuleRetryPolicyArgs>? RetryPolicy { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("runCommandParameters")]
        public Input<Inputs.RuleRunCommandParametersArgs>? RunCommandParameters { get; set; }

        [Input("sageMakerPipelineParameters")]
        public Input<Inputs.RuleSageMakerPipelineParametersArgs>? SageMakerPipelineParameters { get; set; }

        [Input("sqsParameters")]
        public Input<Inputs.RuleSqsParametersArgs>? SqsParameters { get; set; }

        public RuleTargetArgs()
        {
        }
    }
}
