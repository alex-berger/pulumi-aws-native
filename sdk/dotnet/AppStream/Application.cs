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
    /// Resource Type definition for AWS::AppStream::Application
    /// </summary>
    [AwsNativeResourceType("aws-native:appstream:Application")]
    public partial class Application : Pulumi.CustomResource
    {
        [Output("appBlockArn")]
        public Output<string> AppBlockArn { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("attributesToDelete")]
        public Output<ImmutableArray<string>> AttributesToDelete { get; private set; } = null!;

        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        [Output("iconS3Location")]
        public Output<Outputs.ApplicationS3Location> IconS3Location { get; private set; } = null!;

        [Output("instanceFamilies")]
        public Output<ImmutableArray<string>> InstanceFamilies { get; private set; } = null!;

        [Output("launchParameters")]
        public Output<string?> LaunchParameters { get; private set; } = null!;

        [Output("launchPath")]
        public Output<string> LaunchPath { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("platforms")]
        public Output<ImmutableArray<string>> Platforms { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.ApplicationTag>> Tags { get; private set; } = null!;

        [Output("workingDirectory")]
        public Output<string?> WorkingDirectory { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appstream:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appstream:Application", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : Pulumi.ResourceArgs
    {
        [Input("appBlockArn", required: true)]
        public Input<string> AppBlockArn { get; set; } = null!;

        [Input("attributesToDelete")]
        private InputList<string>? _attributesToDelete;
        public InputList<string> AttributesToDelete
        {
            get => _attributesToDelete ?? (_attributesToDelete = new InputList<string>());
            set => _attributesToDelete = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("iconS3Location", required: true)]
        public Input<Inputs.ApplicationS3LocationArgs> IconS3Location { get; set; } = null!;

        [Input("instanceFamilies", required: true)]
        private InputList<string>? _instanceFamilies;
        public InputList<string> InstanceFamilies
        {
            get => _instanceFamilies ?? (_instanceFamilies = new InputList<string>());
            set => _instanceFamilies = value;
        }

        [Input("launchParameters")]
        public Input<string>? LaunchParameters { get; set; }

        [Input("launchPath", required: true)]
        public Input<string> LaunchPath { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("platforms", required: true)]
        private InputList<string>? _platforms;
        public InputList<string> Platforms
        {
            get => _platforms ?? (_platforms = new InputList<string>());
            set => _platforms = value;
        }

        [Input("tags")]
        private InputList<Inputs.ApplicationTagArgs>? _tags;
        public InputList<Inputs.ApplicationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ApplicationTagArgs>());
            set => _tags = value;
        }

        [Input("workingDirectory")]
        public Input<string>? WorkingDirectory { get; set; }

        public ApplicationArgs()
        {
        }
    }
}
