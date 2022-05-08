// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall
{
    /// <summary>
    /// Resource type definition for AWS::NetworkFirewall::RuleGroup
    /// </summary>
    [AwsNativeResourceType("aws-native:networkfirewall:RuleGroup")]
    public partial class RuleGroup : Pulumi.CustomResource
    {
        [Output("capacity")]
        public Output<int> Capacity { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("ruleGroup")]
        public Output<Outputs.RuleGroup?> RuleGroupValue { get; private set; } = null!;

        [Output("ruleGroupArn")]
        public Output<string> RuleGroupArn { get; private set; } = null!;

        [Output("ruleGroupId")]
        public Output<string> RuleGroupId { get; private set; } = null!;

        [Output("ruleGroupName")]
        public Output<string> RuleGroupName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.RuleGroupTag>> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<Pulumi.AwsNative.NetworkFirewall.RuleGroupTypeEnum> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RuleGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RuleGroup(string name, RuleGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:networkfirewall:RuleGroup", name, args ?? new RuleGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RuleGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:networkfirewall:RuleGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RuleGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RuleGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RuleGroup(name, id, options);
        }
    }

    public sealed class RuleGroupArgs : Pulumi.ResourceArgs
    {
        [Input("capacity", required: true)]
        public Input<int> Capacity { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ruleGroup")]
        public Input<Inputs.RuleGroupArgs>? RuleGroupValue { get; set; }

        [Input("ruleGroupName")]
        public Input<string>? RuleGroupName { get; set; }

        [Input("tags")]
        private InputList<Inputs.RuleGroupTagArgs>? _tags;
        public InputList<Inputs.RuleGroupTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.RuleGroupTagArgs>());
            set => _tags = value;
        }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.NetworkFirewall.RuleGroupTypeEnum> Type { get; set; } = null!;

        public RuleGroupArgs()
        {
        }
    }
}
