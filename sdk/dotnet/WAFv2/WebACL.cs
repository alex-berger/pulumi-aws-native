// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2
{
    /// <summary>
    /// Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
    /// </summary>
    [AwsNativeResourceType("aws-native:wafv2:WebACL")]
    public partial class WebACL : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("capacity")]
        public Output<int> Capacity { get; private set; } = null!;

        [Output("captchaConfig")]
        public Output<Outputs.WebACLCaptchaConfig?> CaptchaConfig { get; private set; } = null!;

        [Output("customResponseBodies")]
        public Output<Outputs.WebACLCustomResponseBodies?> CustomResponseBodies { get; private set; } = null!;

        [Output("defaultAction")]
        public Output<Outputs.WebACLDefaultAction> DefaultAction { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("labelNamespace")]
        public Output<string> LabelNamespace { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Collection of Rules.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.WebACLRule>> Rules { get; private set; } = null!;

        [Output("scope")]
        public Output<Pulumi.AwsNative.WAFv2.WebACLScope> Scope { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.WebACLTag>> Tags { get; private set; } = null!;

        [Output("visibilityConfig")]
        public Output<Outputs.WebACLVisibilityConfig> VisibilityConfig { get; private set; } = null!;


        /// <summary>
        /// Create a WebACL resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebACL(string name, WebACLArgs args, CustomResourceOptions? options = null)
            : base("aws-native:wafv2:WebACL", name, args ?? new WebACLArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebACL(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:wafv2:WebACL", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WebACL resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebACL Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebACL(name, id, options);
        }
    }

    public sealed class WebACLArgs : Pulumi.ResourceArgs
    {
        [Input("captchaConfig")]
        public Input<Inputs.WebACLCaptchaConfigArgs>? CaptchaConfig { get; set; }

        [Input("customResponseBodies")]
        public Input<Inputs.WebACLCustomResponseBodiesArgs>? CustomResponseBodies { get; set; }

        [Input("defaultAction", required: true)]
        public Input<Inputs.WebACLDefaultActionArgs> DefaultAction { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rules")]
        private InputList<Inputs.WebACLRuleArgs>? _rules;

        /// <summary>
        /// Collection of Rules.
        /// </summary>
        public InputList<Inputs.WebACLRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.WebACLRuleArgs>());
            set => _rules = value;
        }

        [Input("scope", required: true)]
        public Input<Pulumi.AwsNative.WAFv2.WebACLScope> Scope { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.WebACLTagArgs>? _tags;
        public InputList<Inputs.WebACLTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.WebACLTagArgs>());
            set => _tags = value;
        }

        [Input("visibilityConfig", required: true)]
        public Input<Inputs.WebACLVisibilityConfigArgs> VisibilityConfig { get; set; } = null!;

        public WebACLArgs()
        {
        }
    }
}