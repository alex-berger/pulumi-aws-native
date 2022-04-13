// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Inputs
{

    /// <summary>
    /// Settings for logging audio of conversations between Amazon Lex and a user. You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.
    /// </summary>
    public sealed class BotAudioLogSettingArgs : Pulumi.ResourceArgs
    {
        [Input("destination", required: true)]
        public Input<Inputs.BotAudioLogDestinationArgs> Destination { get; set; } = null!;

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public BotAudioLogSettingArgs()
        {
        }
    }
}
