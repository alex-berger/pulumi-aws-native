// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream
{
    /// <summary>
    /// Resource Type definition for AWS::AppStream::AppBlock
    /// </summary>
    [AwsNativeResourceType("aws-native:appstream:AppBlock")]
    public partial class AppBlock : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("setupScriptDetails")]
        public Output<Outputs.AppBlockScriptDetails> SetupScriptDetails { get; private set; } = null!;

        [Output("sourceS3Location")]
        public Output<Outputs.AppBlockS3Location> SourceS3Location { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.AppBlockTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AppBlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppBlock(string name, AppBlockArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appstream:AppBlock", name, args ?? new AppBlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppBlock(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appstream:AppBlock", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AppBlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppBlock Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppBlock(name, id, options);
        }
    }

    public sealed class AppBlockArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("setupScriptDetails", required: true)]
        public Input<Inputs.AppBlockScriptDetailsArgs> SetupScriptDetails { get; set; } = null!;

        [Input("sourceS3Location", required: true)]
        public Input<Inputs.AppBlockS3LocationArgs> SourceS3Location { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.AppBlockTagArgs>? _tags;
        public InputList<Inputs.AppBlockTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AppBlockTagArgs>());
            set => _tags = value;
        }

        public AppBlockArgs()
        {
        }
    }
}
