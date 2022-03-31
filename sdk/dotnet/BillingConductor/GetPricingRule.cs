// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.BillingConductor
{
    public static class GetPricingRule
    {
        /// <summary>
        /// A markup/discount that is defined for a specific set of services that can later be associated with a pricing plan.
        /// </summary>
        public static Task<GetPricingRuleResult> InvokeAsync(GetPricingRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPricingRuleResult>("aws-native:billingconductor:getPricingRule", args ?? new GetPricingRuleArgs(), options.WithDefaults());

        /// <summary>
        /// A markup/discount that is defined for a specific set of services that can later be associated with a pricing plan.
        /// </summary>
        public static Output<GetPricingRuleResult> Invoke(GetPricingRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetPricingRuleResult>("aws-native:billingconductor:getPricingRule", args ?? new GetPricingRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPricingRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Pricing rule ARN
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetPricingRuleArgs()
        {
        }
    }

    public sealed class GetPricingRuleInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Pricing rule ARN
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetPricingRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPricingRuleResult
    {
        /// <summary>
        /// Pricing rule ARN
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The number of pricing plans associated with pricing rule
        /// </summary>
        public readonly int? AssociatedPricingPlanCount;
        /// <summary>
        /// Creation timestamp in UNIX epoch time format
        /// </summary>
        public readonly int? CreationTime;
        /// <summary>
        /// Pricing rule description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Latest modified timestamp in UNIX epoch time format
        /// </summary>
        public readonly int? LastModifiedTime;
        /// <summary>
        /// Pricing rule modifier percentage
        /// </summary>
        public readonly double? ModifierPercentage;
        /// <summary>
        /// Pricing rule name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// One of MARKUP or DISCOUNT that describes the direction of the rate that is applied to a pricing plan.
        /// </summary>
        public readonly Pulumi.AwsNative.BillingConductor.PricingRuleType? Type;

        [OutputConstructor]
        private GetPricingRuleResult(
            string? arn,

            int? associatedPricingPlanCount,

            int? creationTime,

            string? description,

            int? lastModifiedTime,

            double? modifierPercentage,

            string? name,

            Pulumi.AwsNative.BillingConductor.PricingRuleType? type)
        {
            Arn = arn;
            AssociatedPricingPlanCount = associatedPricingPlanCount;
            CreationTime = creationTime;
            Description = description;
            LastModifiedTime = lastModifiedTime;
            ModifierPercentage = modifierPercentage;
            Name = name;
            Type = type;
        }
    }
}