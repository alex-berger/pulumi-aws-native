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
    /// Data privacy setting of the Bot.
    /// </summary>
    public sealed class DataPrivacyPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("childDirected", required: true)]
        public Input<bool> ChildDirected { get; set; } = null!;

        public DataPrivacyPropertiesArgs()
        {
        }
    }
}
