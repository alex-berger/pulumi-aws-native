// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.BillingConductor
{
    /// <summary>
    /// Pricing Plan enables you to customize your billing details consistent with the usage that accrues in each of your billing groups.
    /// </summary>
    [Obsolete(@"PricingPlan is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:billingconductor:PricingPlan")]
    public partial class PricingPlan : Pulumi.CustomResource
    {
        /// <summary>
        /// Pricing Plan ARN
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in UNIX epoch time format
        /// </summary>
        [Output("creationTime")]
        public Output<int> CreationTime { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Latest modified timestamp in UNIX epoch time format
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<int> LastModifiedTime { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("pricingRuleArns")]
        public Output<ImmutableArray<string>> PricingRuleArns { get; private set; } = null!;

        /// <summary>
        /// Number of associated pricing rules
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;


        /// <summary>
        /// Create a PricingPlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PricingPlan(string name, PricingPlanArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:billingconductor:PricingPlan", name, args ?? new PricingPlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PricingPlan(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:billingconductor:PricingPlan", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PricingPlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PricingPlan Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PricingPlan(name, id, options);
        }
    }

    public sealed class PricingPlanArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pricingRuleArns")]
        private InputList<string>? _pricingRuleArns;
        public InputList<string> PricingRuleArns
        {
            get => _pricingRuleArns ?? (_pricingRuleArns = new InputList<string>());
            set => _pricingRuleArns = value;
        }

        public PricingPlanArgs()
        {
        }
    }
}