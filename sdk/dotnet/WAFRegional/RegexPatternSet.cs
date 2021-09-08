// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFRegional
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html
    /// </summary>
    [AwsNativeResourceType("aws-native:wafregional:RegexPatternSet")]
    public partial class RegexPatternSet : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-regexpatternstrings
        /// </summary>
        [Output("regexPatternStrings")]
        public Output<ImmutableArray<string>> RegexPatternStrings { get; private set; } = null!;


        /// <summary>
        /// Create a RegexPatternSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegexPatternSet(string name, RegexPatternSetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:wafregional:RegexPatternSet", name, args ?? new RegexPatternSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegexPatternSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:wafregional:RegexPatternSet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RegexPatternSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegexPatternSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegexPatternSet(name, id, options);
        }
    }

    public sealed class RegexPatternSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("regexPatternStrings", required: true)]
        private InputList<string>? _regexPatternStrings;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-regexpatternstrings
        /// </summary>
        public InputList<string> RegexPatternStrings
        {
            get => _regexPatternStrings ?? (_regexPatternStrings = new InputList<string>());
            set => _regexPatternStrings = value;
        }

        public RegexPatternSetArgs()
        {
        }
    }
}
