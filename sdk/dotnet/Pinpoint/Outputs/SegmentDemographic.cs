// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentdimensions-demographic.html
    /// </summary>
    [OutputType]
    public sealed class SegmentDemographic
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentdimensions-demographic.html#cfn-pinpoint-segment-segmentdimensions-demographic-appversion
        /// </summary>
        public readonly Outputs.SegmentSetDimension? AppVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentdimensions-demographic.html#cfn-pinpoint-segment-segmentdimensions-demographic-channel
        /// </summary>
        public readonly Outputs.SegmentSetDimension? Channel;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentdimensions-demographic.html#cfn-pinpoint-segment-segmentdimensions-demographic-devicetype
        /// </summary>
        public readonly Outputs.SegmentSetDimension? DeviceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentdimensions-demographic.html#cfn-pinpoint-segment-segmentdimensions-demographic-make
        /// </summary>
        public readonly Outputs.SegmentSetDimension? Make;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentdimensions-demographic.html#cfn-pinpoint-segment-segmentdimensions-demographic-model
        /// </summary>
        public readonly Outputs.SegmentSetDimension? Model;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-segmentdimensions-demographic.html#cfn-pinpoint-segment-segmentdimensions-demographic-platform
        /// </summary>
        public readonly Outputs.SegmentSetDimension? Platform;

        [OutputConstructor]
        private SegmentDemographic(
            Outputs.SegmentSetDimension? appVersion,

            Outputs.SegmentSetDimension? channel,

            Outputs.SegmentSetDimension? deviceType,

            Outputs.SegmentSetDimension? make,

            Outputs.SegmentSetDimension? model,

            Outputs.SegmentSetDimension? platform)
        {
            AppVersion = appVersion;
            Channel = channel;
            DeviceType = deviceType;
            Make = make;
            Model = model;
            Platform = platform;
        }
    }
}
