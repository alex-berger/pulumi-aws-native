// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CE
{
    public static class GetCostCategory
    {
        /// <summary>
        /// Cost Category enables you to map your cost and usage into meaningful categories. You can use Cost Category to organize your costs using a rule-based engine.
        /// </summary>
        public static Task<GetCostCategoryResult> InvokeAsync(GetCostCategoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCostCategoryResult>("aws-native:ce:getCostCategory", args ?? new GetCostCategoryArgs(), options.WithDefaults());

        /// <summary>
        /// Cost Category enables you to map your cost and usage into meaningful categories. You can use Cost Category to organize your costs using a rule-based engine.
        /// </summary>
        public static Output<GetCostCategoryResult> Invoke(GetCostCategoryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCostCategoryResult>("aws-native:ce:getCostCategory", args ?? new GetCostCategoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCostCategoryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cost category ARN
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetCostCategoryArgs()
        {
        }
    }

    public sealed class GetCostCategoryInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cost category ARN
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetCostCategoryInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCostCategoryResult
    {
        /// <summary>
        /// Cost category ARN
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The default value for the cost category
        /// </summary>
        public readonly string? DefaultValue;
        public readonly string? EffectiveStart;
        public readonly Pulumi.AwsNative.CE.CostCategoryRuleVersion? RuleVersion;
        /// <summary>
        /// JSON array format of Expression in Billing and Cost Management API
        /// </summary>
        public readonly string? Rules;
        /// <summary>
        /// Json array format of CostCategorySplitChargeRule in Billing and Cost Management API
        /// </summary>
        public readonly string? SplitChargeRules;

        [OutputConstructor]
        private GetCostCategoryResult(
            string? arn,

            string? defaultValue,

            string? effectiveStart,

            Pulumi.AwsNative.CE.CostCategoryRuleVersion? ruleVersion,

            string? rules,

            string? splitChargeRules)
        {
            Arn = arn;
            DefaultValue = defaultValue;
            EffectiveStart = effectiveStart;
            RuleVersion = ruleVersion;
            Rules = rules;
            SplitChargeRules = splitChargeRules;
        }
    }
}
