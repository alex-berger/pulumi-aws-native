// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-serviceprofile.html
    /// </summary>
    [AwsNativeResourceType("aws-native:iotwireless:ServiceProfile")]
    public partial class ServiceProfile : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-serviceprofile.html#cfn-iotwireless-serviceprofile-lorawan
        /// </summary>
        [Output("loRaWAN")]
        public Output<Outputs.ServiceProfileLoRaWANServiceProfile?> LoRaWAN { get; private set; } = null!;

        [Output("loRaWANChannelMask")]
        public Output<string> LoRaWANChannelMask { get; private set; } = null!;

        [Output("loRaWANDevStatusReqFreq")]
        public Output<int> LoRaWANDevStatusReqFreq { get; private set; } = null!;

        [Output("loRaWANDlBucketSize")]
        public Output<int> LoRaWANDlBucketSize { get; private set; } = null!;

        [Output("loRaWANDlRate")]
        public Output<int> LoRaWANDlRate { get; private set; } = null!;

        [Output("loRaWANDlRatePolicy")]
        public Output<string> LoRaWANDlRatePolicy { get; private set; } = null!;

        [Output("loRaWANDrMax")]
        public Output<int> LoRaWANDrMax { get; private set; } = null!;

        [Output("loRaWANDrMin")]
        public Output<int> LoRaWANDrMin { get; private set; } = null!;

        [Output("loRaWANHrAllowed")]
        public Output<bool> LoRaWANHrAllowed { get; private set; } = null!;

        [Output("loRaWANMinGwDiversity")]
        public Output<int> LoRaWANMinGwDiversity { get; private set; } = null!;

        [Output("loRaWANNwkGeoLoc")]
        public Output<bool> LoRaWANNwkGeoLoc { get; private set; } = null!;

        [Output("loRaWANPrAllowed")]
        public Output<bool> LoRaWANPrAllowed { get; private set; } = null!;

        [Output("loRaWANRaAllowed")]
        public Output<bool> LoRaWANRaAllowed { get; private set; } = null!;

        [Output("loRaWANReportDevStatusBattery")]
        public Output<bool> LoRaWANReportDevStatusBattery { get; private set; } = null!;

        [Output("loRaWANReportDevStatusMargin")]
        public Output<bool> LoRaWANReportDevStatusMargin { get; private set; } = null!;

        [Output("loRaWANTargetPer")]
        public Output<int> LoRaWANTargetPer { get; private set; } = null!;

        [Output("loRaWANUlBucketSize")]
        public Output<int> LoRaWANUlBucketSize { get; private set; } = null!;

        [Output("loRaWANUlRate")]
        public Output<int> LoRaWANUlRate { get; private set; } = null!;

        [Output("loRaWANUlRatePolicy")]
        public Output<string> LoRaWANUlRatePolicy { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-serviceprofile.html#cfn-iotwireless-serviceprofile-name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-serviceprofile.html#cfn-iotwireless-serviceprofile-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceProfile(string name, ServiceProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:iotwireless:ServiceProfile", name, args ?? new ServiceProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iotwireless:ServiceProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceProfile(name, id, options);
        }
    }

    public sealed class ServiceProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-serviceprofile.html#cfn-iotwireless-serviceprofile-lorawan
        /// </summary>
        [Input("loRaWAN")]
        public Input<Inputs.ServiceProfileLoRaWANServiceProfileArgs>? LoRaWAN { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-serviceprofile.html#cfn-iotwireless-serviceprofile-name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-serviceprofile.html#cfn-iotwireless-serviceprofile-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ServiceProfileArgs()
        {
        }
    }
}