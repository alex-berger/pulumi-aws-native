// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch
{
    public static class GetInsightRule
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudWatch::InsightRule
        /// </summary>
        public static Task<GetInsightRuleResult> InvokeAsync(GetInsightRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInsightRuleResult>("aws-native:cloudwatch:getInsightRule", args ?? new GetInsightRuleArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudWatch::InsightRule
        /// </summary>
        public static Output<GetInsightRuleResult> Invoke(GetInsightRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInsightRuleResult>("aws-native:cloudwatch:getInsightRule", args ?? new GetInsightRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInsightRuleArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetInsightRuleArgs()
        {
        }
    }

    public sealed class GetInsightRuleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetInsightRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInsightRuleResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly string? RuleBody;
        public readonly string? RuleState;
        public readonly Outputs.InsightRuleTags? Tags;

        [OutputConstructor]
        private GetInsightRuleResult(
            string? arn,

            string? id,

            string? ruleBody,

            string? ruleState,

            Outputs.InsightRuleTags? tags)
        {
            Arn = arn;
            Id = id;
            RuleBody = ruleBody;
            RuleState = ruleState;
            Tags = tags;
        }
    }
}
