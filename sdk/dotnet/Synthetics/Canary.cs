// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics
{
    /// <summary>
    /// Resource Type definition for AWS::Synthetics::Canary
    /// </summary>
    [AwsNativeResourceType("aws-native:synthetics:Canary")]
    public partial class Canary : Pulumi.CustomResource
    {
        /// <summary>
        /// Provide artifact configuration
        /// </summary>
        [Output("artifactConfig")]
        public Output<Outputs.CanaryArtifactConfig?> ArtifactConfig { get; private set; } = null!;

        /// <summary>
        /// Provide the s3 bucket output location for test results
        /// </summary>
        [Output("artifactS3Location")]
        public Output<string> ArtifactS3Location { get; private set; } = null!;

        /// <summary>
        /// Provide the canary script source
        /// </summary>
        [Output("code")]
        public Output<Outputs.CanaryCode> Code { get; private set; } = null!;

        /// <summary>
        /// Deletes associated lambda resources created by Synthetics if set to True. Default is False
        /// </summary>
        [Output("deleteLambdaResourcesOnCanaryDeletion")]
        public Output<bool?> DeleteLambdaResourcesOnCanaryDeletion { get; private set; } = null!;

        /// <summary>
        /// Lambda Execution role used to run your canaries
        /// </summary>
        [Output("executionRoleArn")]
        public Output<string> ExecutionRoleArn { get; private set; } = null!;

        /// <summary>
        /// Retention period of failed canary runs represented in number of days
        /// </summary>
        [Output("failureRetentionPeriod")]
        public Output<int?> FailureRetentionPeriod { get; private set; } = null!;

        /// <summary>
        /// Name of the canary.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provide canary run configuration
        /// </summary>
        [Output("runConfig")]
        public Output<Outputs.CanaryRunConfig?> RunConfig { get; private set; } = null!;

        /// <summary>
        /// Runtime version of Synthetics Library
        /// </summary>
        [Output("runtimeVersion")]
        public Output<string> RuntimeVersion { get; private set; } = null!;

        /// <summary>
        /// Frequency to run your canaries
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.CanarySchedule> Schedule { get; private set; } = null!;

        /// <summary>
        /// Runs canary if set to True. Default is False
        /// </summary>
        [Output("startCanaryAfterCreation")]
        public Output<bool> StartCanaryAfterCreation { get; private set; } = null!;

        /// <summary>
        /// State of the canary
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Retention period of successful canary runs represented in number of days
        /// </summary>
        [Output("successRetentionPeriod")]
        public Output<int?> SuccessRetentionPeriod { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.CanaryTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Provide VPC Configuration if enabled.
        /// </summary>
        [Output("vPCConfig")]
        public Output<Outputs.CanaryVPCConfig?> VPCConfig { get; private set; } = null!;

        /// <summary>
        /// Visual reference configuration for visual testing
        /// </summary>
        [Output("visualReference")]
        public Output<Outputs.CanaryVisualReference?> VisualReference { get; private set; } = null!;


        /// <summary>
        /// Create a Canary resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Canary(string name, CanaryArgs args, CustomResourceOptions? options = null)
            : base("aws-native:synthetics:Canary", name, args ?? new CanaryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Canary(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:synthetics:Canary", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Canary resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Canary Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Canary(name, id, options);
        }
    }

    public sealed class CanaryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Provide artifact configuration
        /// </summary>
        [Input("artifactConfig")]
        public Input<Inputs.CanaryArtifactConfigArgs>? ArtifactConfig { get; set; }

        /// <summary>
        /// Provide the s3 bucket output location for test results
        /// </summary>
        [Input("artifactS3Location", required: true)]
        public Input<string> ArtifactS3Location { get; set; } = null!;

        /// <summary>
        /// Provide the canary script source
        /// </summary>
        [Input("code", required: true)]
        public Input<Inputs.CanaryCodeArgs> Code { get; set; } = null!;

        /// <summary>
        /// Deletes associated lambda resources created by Synthetics if set to True. Default is False
        /// </summary>
        [Input("deleteLambdaResourcesOnCanaryDeletion")]
        public Input<bool>? DeleteLambdaResourcesOnCanaryDeletion { get; set; }

        /// <summary>
        /// Lambda Execution role used to run your canaries
        /// </summary>
        [Input("executionRoleArn", required: true)]
        public Input<string> ExecutionRoleArn { get; set; } = null!;

        /// <summary>
        /// Retention period of failed canary runs represented in number of days
        /// </summary>
        [Input("failureRetentionPeriod")]
        public Input<int>? FailureRetentionPeriod { get; set; }

        /// <summary>
        /// Name of the canary.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Provide canary run configuration
        /// </summary>
        [Input("runConfig")]
        public Input<Inputs.CanaryRunConfigArgs>? RunConfig { get; set; }

        /// <summary>
        /// Runtime version of Synthetics Library
        /// </summary>
        [Input("runtimeVersion", required: true)]
        public Input<string> RuntimeVersion { get; set; } = null!;

        /// <summary>
        /// Frequency to run your canaries
        /// </summary>
        [Input("schedule", required: true)]
        public Input<Inputs.CanaryScheduleArgs> Schedule { get; set; } = null!;

        /// <summary>
        /// Runs canary if set to True. Default is False
        /// </summary>
        [Input("startCanaryAfterCreation", required: true)]
        public Input<bool> StartCanaryAfterCreation { get; set; } = null!;

        /// <summary>
        /// Retention period of successful canary runs represented in number of days
        /// </summary>
        [Input("successRetentionPeriod")]
        public Input<int>? SuccessRetentionPeriod { get; set; }

        [Input("tags")]
        private InputList<Inputs.CanaryTagArgs>? _tags;
        public InputList<Inputs.CanaryTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CanaryTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Provide VPC Configuration if enabled.
        /// </summary>
        [Input("vPCConfig")]
        public Input<Inputs.CanaryVPCConfigArgs>? VPCConfig { get; set; }

        /// <summary>
        /// Visual reference configuration for visual testing
        /// </summary>
        [Input("visualReference")]
        public Input<Inputs.CanaryVisualReferenceArgs>? VisualReference { get; set; }

        public CanaryArgs()
        {
        }
    }
}
