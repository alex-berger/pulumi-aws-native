// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    /// <summary>
    /// Resource schema for AWS::DataSync::LocationEFS.
    /// </summary>
    [AwsNativeResourceType("aws-native:datasync:LocationEFS")]
    public partial class LocationEFS : Pulumi.CustomResource
    {
        [Output("ec2Config")]
        public Output<Outputs.LocationEFSEc2Config> Ec2Config { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the Amazon EFS file system.
        /// </summary>
        [Output("efsFilesystemArn")]
        public Output<string> EfsFilesystemArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.
        /// </summary>
        [Output("locationArn")]
        public Output<string> LocationArn { get; private set; } = null!;

        /// <summary>
        /// The URL of the EFS location that was described.
        /// </summary>
        [Output("locationUri")]
        public Output<string> LocationUri { get; private set; } = null!;

        /// <summary>
        /// A subdirectory in the location's path. This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination.
        /// </summary>
        [Output("subdirectory")]
        public Output<string?> Subdirectory { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.LocationEFSTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a LocationEFS resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocationEFS(string name, LocationEFSArgs args, CustomResourceOptions? options = null)
            : base("aws-native:datasync:LocationEFS", name, args ?? new LocationEFSArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocationEFS(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:datasync:LocationEFS", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LocationEFS resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocationEFS Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LocationEFS(name, id, options);
        }
    }

    public sealed class LocationEFSArgs : Pulumi.ResourceArgs
    {
        [Input("ec2Config", required: true)]
        public Input<Inputs.LocationEFSEc2ConfigArgs> Ec2Config { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the Amazon EFS file system.
        /// </summary>
        [Input("efsFilesystemArn", required: true)]
        public Input<string> EfsFilesystemArn { get; set; } = null!;

        /// <summary>
        /// A subdirectory in the location's path. This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination.
        /// </summary>
        [Input("subdirectory")]
        public Input<string>? Subdirectory { get; set; }

        [Input("tags")]
        private InputList<Inputs.LocationEFSTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.LocationEFSTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LocationEFSTagArgs>());
            set => _tags = value;
        }

        public LocationEFSArgs()
        {
        }
    }
}
