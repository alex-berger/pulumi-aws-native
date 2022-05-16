// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    /// <summary>
    /// Resource Type definition for AWS::SageMaker::ModelQualityJobDefinition
    /// </summary>
    [AwsNativeResourceType("aws-native:sagemaker:ModelQualityJobDefinition")]
    public partial class ModelQualityJobDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// The time at which the job definition was created.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of job definition.
        /// </summary>
        [Output("jobDefinitionArn")]
        public Output<string> JobDefinitionArn { get; private set; } = null!;

        [Output("jobDefinitionName")]
        public Output<string?> JobDefinitionName { get; private set; } = null!;

        [Output("jobResources")]
        public Output<Outputs.ModelQualityJobDefinitionMonitoringResources> JobResources { get; private set; } = null!;

        [Output("modelQualityAppSpecification")]
        public Output<Outputs.ModelQualityJobDefinitionModelQualityAppSpecification> ModelQualityAppSpecification { get; private set; } = null!;

        [Output("modelQualityBaselineConfig")]
        public Output<Outputs.ModelQualityJobDefinitionModelQualityBaselineConfig?> ModelQualityBaselineConfig { get; private set; } = null!;

        [Output("modelQualityJobInput")]
        public Output<Outputs.ModelQualityJobDefinitionModelQualityJobInput> ModelQualityJobInput { get; private set; } = null!;

        [Output("modelQualityJobOutputConfig")]
        public Output<Outputs.ModelQualityJobDefinitionMonitoringOutputConfig> ModelQualityJobOutputConfig { get; private set; } = null!;

        [Output("networkConfig")]
        public Output<Outputs.ModelQualityJobDefinitionNetworkConfig?> NetworkConfig { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        [Output("stoppingCondition")]
        public Output<Outputs.ModelQualityJobDefinitionStoppingCondition?> StoppingCondition { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ModelQualityJobDefinitionTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ModelQualityJobDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModelQualityJobDefinition(string name, ModelQualityJobDefinitionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:ModelQualityJobDefinition", name, args ?? new ModelQualityJobDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModelQualityJobDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:ModelQualityJobDefinition", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ModelQualityJobDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModelQualityJobDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ModelQualityJobDefinition(name, id, options);
        }
    }

    public sealed class ModelQualityJobDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("jobDefinitionName")]
        public Input<string>? JobDefinitionName { get; set; }

        [Input("jobResources", required: true)]
        public Input<Inputs.ModelQualityJobDefinitionMonitoringResourcesArgs> JobResources { get; set; } = null!;

        [Input("modelQualityAppSpecification", required: true)]
        public Input<Inputs.ModelQualityJobDefinitionModelQualityAppSpecificationArgs> ModelQualityAppSpecification { get; set; } = null!;

        [Input("modelQualityBaselineConfig")]
        public Input<Inputs.ModelQualityJobDefinitionModelQualityBaselineConfigArgs>? ModelQualityBaselineConfig { get; set; }

        [Input("modelQualityJobInput", required: true)]
        public Input<Inputs.ModelQualityJobDefinitionModelQualityJobInputArgs> ModelQualityJobInput { get; set; } = null!;

        [Input("modelQualityJobOutputConfig", required: true)]
        public Input<Inputs.ModelQualityJobDefinitionMonitoringOutputConfigArgs> ModelQualityJobOutputConfig { get; set; } = null!;

        [Input("networkConfig")]
        public Input<Inputs.ModelQualityJobDefinitionNetworkConfigArgs>? NetworkConfig { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("stoppingCondition")]
        public Input<Inputs.ModelQualityJobDefinitionStoppingConditionArgs>? StoppingCondition { get; set; }

        [Input("tags")]
        private InputList<Inputs.ModelQualityJobDefinitionTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.ModelQualityJobDefinitionTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ModelQualityJobDefinitionTagArgs>());
            set => _tags = value;
        }

        public ModelQualityJobDefinitionArgs()
        {
        }
    }
}