// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew
{
    /// <summary>
    /// Resource schema for AWS::DataBrew::Dataset.
    /// </summary>
    [AwsNativeResourceType("aws-native:databrew:Dataset")]
    public partial class Dataset : Pulumi.CustomResource
    {
        /// <summary>
        /// Dataset format
        /// </summary>
        [Output("format")]
        public Output<Pulumi.AwsNative.DataBrew.DatasetFormat?> Format { get; private set; } = null!;

        /// <summary>
        /// Format options for dataset
        /// </summary>
        [Output("formatOptions")]
        public Output<Outputs.DatasetFormatOptions?> FormatOptions { get; private set; } = null!;

        /// <summary>
        /// Input
        /// </summary>
        [Output("input")]
        public Output<Outputs.DatasetInput> Input { get; private set; } = null!;

        /// <summary>
        /// Dataset name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// PathOptions
        /// </summary>
        [Output("pathOptions")]
        public Output<Outputs.DatasetPathOptions?> PathOptions { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.DatasetTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Dataset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dataset(string name, DatasetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:databrew:Dataset", name, args ?? new DatasetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dataset(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:databrew:Dataset", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Dataset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dataset Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Dataset(name, id, options);
        }
    }

    public sealed class DatasetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dataset format
        /// </summary>
        [Input("format")]
        public Input<Pulumi.AwsNative.DataBrew.DatasetFormat>? Format { get; set; }

        /// <summary>
        /// Format options for dataset
        /// </summary>
        [Input("formatOptions")]
        public Input<Inputs.DatasetFormatOptionsArgs>? FormatOptions { get; set; }

        /// <summary>
        /// Input
        /// </summary>
        [Input("input", required: true)]
        public Input<Inputs.DatasetInputArgs> Input { get; set; } = null!;

        /// <summary>
        /// Dataset name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// PathOptions
        /// </summary>
        [Input("pathOptions")]
        public Input<Inputs.DatasetPathOptionsArgs>? PathOptions { get; set; }

        [Input("tags")]
        private InputList<Inputs.DatasetTagArgs>? _tags;
        public InputList<Inputs.DatasetTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DatasetTagArgs>());
            set => _tags = value;
        }

        public DatasetArgs()
        {
        }
    }
}
