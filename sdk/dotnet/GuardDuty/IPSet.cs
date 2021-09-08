// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GuardDuty
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html
    /// </summary>
    [AwsNativeResourceType("aws-native:guardduty:IPSet")]
    public partial class IPSet : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-activate
        /// </summary>
        [Output("activate")]
        public Output<bool> Activate { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-detectorid
        /// </summary>
        [Output("detectorId")]
        public Output<string> DetectorId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-format
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;


        /// <summary>
        /// Create a IPSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IPSet(string name, IPSetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:guardduty:IPSet", name, args ?? new IPSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IPSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:guardduty:IPSet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing IPSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IPSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IPSet(name, id, options);
        }
    }

    public sealed class IPSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-activate
        /// </summary>
        [Input("activate", required: true)]
        public Input<bool> Activate { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-detectorid
        /// </summary>
        [Input("detectorId", required: true)]
        public Input<string> DetectorId { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-format
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IPSetArgs()
        {
        }
    }
}
