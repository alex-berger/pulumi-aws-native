// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PinpointEmail
{
    /// <summary>
    /// Resource Type definition for AWS::PinpointEmail::ConfigurationSetEventDestination
    /// </summary>
    [Obsolete(@"ConfigurationSetEventDestination is not yet supported by AWS Cloud Control API, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:pinpointemail:ConfigurationSetEventDestination")]
    public partial class ConfigurationSetEventDestination : Pulumi.CustomResource
    {
        [Output("configurationSetName")]
        public Output<string> ConfigurationSetName { get; private set; } = null!;

        [Output("eventDestination")]
        public Output<Outputs.ConfigurationSetEventDestinationEventDestination?> EventDestination { get; private set; } = null!;

        [Output("eventDestinationName")]
        public Output<string> EventDestinationName { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationSetEventDestination resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationSetEventDestination(string name, ConfigurationSetEventDestinationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:pinpointemail:ConfigurationSetEventDestination", name, args ?? new ConfigurationSetEventDestinationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationSetEventDestination(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:pinpointemail:ConfigurationSetEventDestination", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationSetEventDestination resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationSetEventDestination Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigurationSetEventDestination(name, id, options);
        }
    }

    public sealed class ConfigurationSetEventDestinationArgs : Pulumi.ResourceArgs
    {
        [Input("configurationSetName", required: true)]
        public Input<string> ConfigurationSetName { get; set; } = null!;

        [Input("eventDestination")]
        public Input<Inputs.ConfigurationSetEventDestinationEventDestinationArgs>? EventDestination { get; set; }

        [Input("eventDestinationName", required: true)]
        public Input<string> EventDestinationName { get; set; } = null!;

        public ConfigurationSetEventDestinationArgs()
        {
        }
    }
}