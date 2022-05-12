// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    /// <summary>
    /// A dimension can be used to limit the scope of a metric used in a security profile for AWS IoT Device Defender.
    /// </summary>
    [AwsNativeResourceType("aws-native:iot:Dimension")]
    public partial class Dimension : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN (Amazon resource name) of the created dimension.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the dimension.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the value or list of values for the dimension.
        /// </summary>
        [Output("stringValues")]
        public Output<ImmutableArray<string>> StringValues { get; private set; } = null!;

        /// <summary>
        /// Metadata that can be used to manage the dimension.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DimensionTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of the dimension.
        /// </summary>
        [Output("type")]
        public Output<Pulumi.AwsNative.IoT.DimensionType> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Dimension resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dimension(string name, DimensionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:iot:Dimension", name, args ?? new DimensionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dimension(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iot:Dimension", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Dimension resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dimension Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Dimension(name, id, options);
        }
    }

    public sealed class DimensionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for the dimension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("stringValues", required: true)]
        private InputList<string>? _stringValues;

        /// <summary>
        /// Specifies the value or list of values for the dimension.
        /// </summary>
        public InputList<string> StringValues
        {
            get => _stringValues ?? (_stringValues = new InputList<string>());
            set => _stringValues = value;
        }

        [Input("tags")]
        private InputList<Inputs.DimensionTagArgs>? _tags;

        /// <summary>
        /// Metadata that can be used to manage the dimension.
        /// </summary>
        public InputList<Inputs.DimensionTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DimensionTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Specifies the type of the dimension.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.IoT.DimensionType> Type { get; set; } = null!;

        public DimensionArgs()
        {
        }
    }
}
