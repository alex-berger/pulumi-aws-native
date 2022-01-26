// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Forecast
{
    /// <summary>
    /// Represents a dataset group that holds a collection of related datasets
    /// </summary>
    [AwsNativeResourceType("aws-native:forecast:DatasetGroup")]
    public partial class DatasetGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.
        /// </summary>
        [Output("datasetArns")]
        public Output<ImmutableArray<string>> DatasetArns { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the dataset group to delete.
        /// </summary>
        [Output("datasetGroupArn")]
        public Output<string> DatasetGroupArn { get; private set; } = null!;

        /// <summary>
        /// A name for the dataset group.
        /// </summary>
        [Output("datasetGroupName")]
        public Output<string> DatasetGroupName { get; private set; } = null!;

        /// <summary>
        /// The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.
        /// </summary>
        [Output("domain")]
        public Output<Pulumi.AwsNative.Forecast.DatasetGroupDomain> Domain { get; private set; } = null!;

        /// <summary>
        /// The tags of Application Insights application.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DatasetGroupTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetGroup(string name, DatasetGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:forecast:DatasetGroup", name, args ?? new DatasetGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:forecast:DatasetGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DatasetGroup(name, id, options);
        }
    }

    public sealed class DatasetGroupArgs : Pulumi.ResourceArgs
    {
        [Input("datasetArns")]
        private InputList<string>? _datasetArns;

        /// <summary>
        /// An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the dataset group.
        /// </summary>
        public InputList<string> DatasetArns
        {
            get => _datasetArns ?? (_datasetArns = new InputList<string>());
            set => _datasetArns = value;
        }

        /// <summary>
        /// A name for the dataset group.
        /// </summary>
        [Input("datasetGroupName")]
        public Input<string>? DatasetGroupName { get; set; }

        /// <summary>
        /// The domain associated with the dataset group. When you add a dataset to a dataset group, this value and the value specified for the Domain parameter of the CreateDataset operation must match.
        /// </summary>
        [Input("domain", required: true)]
        public Input<Pulumi.AwsNative.Forecast.DatasetGroupDomain> Domain { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.DatasetGroupTagArgs>? _tags;

        /// <summary>
        /// The tags of Application Insights application.
        /// </summary>
        public InputList<Inputs.DatasetGroupTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DatasetGroupTagArgs>());
            set => _tags = value;
        }

        public DatasetGroupArgs()
        {
        }
    }
}