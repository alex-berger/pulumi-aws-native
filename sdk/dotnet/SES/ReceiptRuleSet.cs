// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES
{
    /// <summary>
    /// Resource Type definition for AWS::SES::ReceiptRuleSet
    /// </summary>
    [Obsolete(@"ReceiptRuleSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ses:ReceiptRuleSet")]
    public partial class ReceiptRuleSet : Pulumi.CustomResource
    {
        [Output("ruleSetName")]
        public Output<string?> RuleSetName { get; private set; } = null!;


        /// <summary>
        /// Create a ReceiptRuleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReceiptRuleSet(string name, ReceiptRuleSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ses:ReceiptRuleSet", name, args ?? new ReceiptRuleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReceiptRuleSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ses:ReceiptRuleSet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ReceiptRuleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReceiptRuleSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ReceiptRuleSet(name, id, options);
        }
    }

    public sealed class ReceiptRuleSetArgs : Pulumi.ResourceArgs
    {
        [Input("ruleSetName")]
        public Input<string>? RuleSetName { get; set; }

        public ReceiptRuleSetArgs()
        {
        }
    }
}
