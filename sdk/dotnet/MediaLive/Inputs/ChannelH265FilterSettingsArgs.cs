// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265filtersettings.html
    /// </summary>
    public sealed class ChannelH265FilterSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h265filtersettings.html#cfn-medialive-channel-h265filtersettings-temporalfiltersettings
        /// </summary>
        [Input("temporalFilterSettings")]
        public Input<Inputs.ChannelTemporalFilterSettingsArgs>? TemporalFilterSettings { get; set; }

        public ChannelH265FilterSettingsArgs()
        {
        }
    }
}
