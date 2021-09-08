// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html
    /// </summary>
    [OutputType]
    public sealed class ChannelOutputGroup
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-outputgroupsettings
        /// </summary>
        public readonly Outputs.ChannelOutputGroupSettings? OutputGroupSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-outputs
        /// </summary>
        public readonly ImmutableArray<Outputs.ChannelOutput> Outputs;

        [OutputConstructor]
        private ChannelOutputGroup(
            string? name,

            Outputs.ChannelOutputGroupSettings? outputGroupSettings,

            ImmutableArray<Outputs.ChannelOutput> outputs)
        {
            Name = name;
            OutputGroupSettings = outputGroupSettings;
            Outputs = outputs;
        }
    }
}