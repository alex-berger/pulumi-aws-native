// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch
{
    /// <summary>
    /// Resource Type definition for AWS::CloudWatch::AnomalyDetector
    /// </summary>
    [Obsolete(@"AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:cloudwatch:AnomalyDetector")]
    public partial class AnomalyDetector : Pulumi.CustomResource
    {
        [Output("configuration")]
        public Output<Outputs.AnomalyDetectorConfiguration?> Configuration { get; private set; } = null!;

        [Output("dimensions")]
        public Output<ImmutableArray<Outputs.AnomalyDetectorDimension>> Dimensions { get; private set; } = null!;

        [Output("metricMathAnomalyDetector")]
        public Output<Outputs.AnomalyDetectorMetricMathAnomalyDetector?> MetricMathAnomalyDetector { get; private set; } = null!;

        [Output("metricName")]
        public Output<string?> MetricName { get; private set; } = null!;

        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        [Output("singleMetricAnomalyDetector")]
        public Output<Outputs.AnomalyDetectorSingleMetricAnomalyDetector?> SingleMetricAnomalyDetector { get; private set; } = null!;

        [Output("stat")]
        public Output<string?> Stat { get; private set; } = null!;


        /// <summary>
        /// Create a AnomalyDetector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AnomalyDetector(string name, AnomalyDetectorArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:cloudwatch:AnomalyDetector", name, args ?? new AnomalyDetectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AnomalyDetector(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudwatch:AnomalyDetector", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AnomalyDetector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AnomalyDetector Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AnomalyDetector(name, id, options);
        }
    }

    public sealed class AnomalyDetectorArgs : Pulumi.ResourceArgs
    {
        [Input("configuration")]
        public Input<Inputs.AnomalyDetectorConfigurationArgs>? Configuration { get; set; }

        [Input("dimensions")]
        private InputList<Inputs.AnomalyDetectorDimensionArgs>? _dimensions;
        public InputList<Inputs.AnomalyDetectorDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.AnomalyDetectorDimensionArgs>());
            set => _dimensions = value;
        }

        [Input("metricMathAnomalyDetector")]
        public Input<Inputs.AnomalyDetectorMetricMathAnomalyDetectorArgs>? MetricMathAnomalyDetector { get; set; }

        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("singleMetricAnomalyDetector")]
        public Input<Inputs.AnomalyDetectorSingleMetricAnomalyDetectorArgs>? SingleMetricAnomalyDetector { get; set; }

        [Input("stat")]
        public Input<string>? Stat { get; set; }

        public AnomalyDetectorArgs()
        {
        }
    }
}
