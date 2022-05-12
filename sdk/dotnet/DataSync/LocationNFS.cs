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
    /// Resource schema for AWS::DataSync::LocationNFS
    /// </summary>
    [AwsNativeResourceType("aws-native:datasync:LocationNFS")]
    public partial class LocationNFS : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the NFS location.
        /// </summary>
        [Output("locationArn")]
        public Output<string> LocationArn { get; private set; } = null!;

        /// <summary>
        /// The URL of the NFS location that was described.
        /// </summary>
        [Output("locationUri")]
        public Output<string> LocationUri { get; private set; } = null!;

        [Output("mountOptions")]
        public Output<Outputs.LocationNFSMountOptions?> MountOptions { get; private set; } = null!;

        [Output("onPremConfig")]
        public Output<Outputs.LocationNFSOnPremConfig> OnPremConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the NFS server. This value is the IP address or DNS name of the NFS server.
        /// </summary>
        [Output("serverHostname")]
        public Output<string> ServerHostname { get; private set; } = null!;

        /// <summary>
        /// The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.
        /// </summary>
        [Output("subdirectory")]
        public Output<string> Subdirectory { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.LocationNFSTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a LocationNFS resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocationNFS(string name, LocationNFSArgs args, CustomResourceOptions? options = null)
            : base("aws-native:datasync:LocationNFS", name, args ?? new LocationNFSArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocationNFS(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:datasync:LocationNFS", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LocationNFS resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocationNFS Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LocationNFS(name, id, options);
        }
    }

    public sealed class LocationNFSArgs : Pulumi.ResourceArgs
    {
        [Input("mountOptions")]
        public Input<Inputs.LocationNFSMountOptionsArgs>? MountOptions { get; set; }

        [Input("onPremConfig", required: true)]
        public Input<Inputs.LocationNFSOnPremConfigArgs> OnPremConfig { get; set; } = null!;

        /// <summary>
        /// The name of the NFS server. This value is the IP address or DNS name of the NFS server.
        /// </summary>
        [Input("serverHostname", required: true)]
        public Input<string> ServerHostname { get; set; } = null!;

        /// <summary>
        /// The subdirectory in the NFS file system that is used to read data from the NFS source location or write data to the NFS destination.
        /// </summary>
        [Input("subdirectory", required: true)]
        public Input<string> Subdirectory { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.LocationNFSTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.LocationNFSTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LocationNFSTagArgs>());
            set => _tags = value;
        }

        public LocationNFSArgs()
        {
        }
    }
}
