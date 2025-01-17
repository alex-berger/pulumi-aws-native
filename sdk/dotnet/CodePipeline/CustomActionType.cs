// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline
{
    /// <summary>
    /// Resource Type definition for AWS::CodePipeline::CustomActionType
    /// </summary>
    [Obsolete(@"CustomActionType is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:codepipeline:CustomActionType")]
    public partial class CustomActionType : Pulumi.CustomResource
    {
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        [Output("configurationProperties")]
        public Output<ImmutableArray<Outputs.CustomActionTypeConfigurationProperties>> ConfigurationProperties { get; private set; } = null!;

        [Output("inputArtifactDetails")]
        public Output<Outputs.CustomActionTypeArtifactDetails> InputArtifactDetails { get; private set; } = null!;

        [Output("outputArtifactDetails")]
        public Output<Outputs.CustomActionTypeArtifactDetails> OutputArtifactDetails { get; private set; } = null!;

        [Output("provider")]
        public Output<string> Provider { get; private set; } = null!;

        [Output("settings")]
        public Output<Outputs.CustomActionTypeSettings?> Settings { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.CustomActionTypeTag>> Tags { get; private set; } = null!;

        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a CustomActionType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomActionType(string name, CustomActionTypeArgs args, CustomResourceOptions? options = null)
            : base("aws-native:codepipeline:CustomActionType", name, args ?? new CustomActionTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomActionType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:codepipeline:CustomActionType", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing CustomActionType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomActionType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomActionType(name, id, options);
        }
    }

    public sealed class CustomActionTypeArgs : Pulumi.ResourceArgs
    {
        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        [Input("configurationProperties")]
        private InputList<Inputs.CustomActionTypeConfigurationPropertiesArgs>? _configurationProperties;
        public InputList<Inputs.CustomActionTypeConfigurationPropertiesArgs> ConfigurationProperties
        {
            get => _configurationProperties ?? (_configurationProperties = new InputList<Inputs.CustomActionTypeConfigurationPropertiesArgs>());
            set => _configurationProperties = value;
        }

        [Input("inputArtifactDetails", required: true)]
        public Input<Inputs.CustomActionTypeArtifactDetailsArgs> InputArtifactDetails { get; set; } = null!;

        [Input("outputArtifactDetails", required: true)]
        public Input<Inputs.CustomActionTypeArtifactDetailsArgs> OutputArtifactDetails { get; set; } = null!;

        [Input("provider", required: true)]
        public Input<string> Provider { get; set; } = null!;

        [Input("settings")]
        public Input<Inputs.CustomActionTypeSettingsArgs>? Settings { get; set; }

        [Input("tags")]
        private InputList<Inputs.CustomActionTypeTagArgs>? _tags;
        public InputList<Inputs.CustomActionTypeTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CustomActionTypeTagArgs>());
            set => _tags = value;
        }

        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public CustomActionTypeArgs()
        {
        }
    }
}
