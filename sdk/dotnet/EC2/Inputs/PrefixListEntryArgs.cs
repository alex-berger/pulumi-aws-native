// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class PrefixListEntryArgs : Pulumi.ResourceArgs
    {
        [Input("cidr", required: true)]
        public Input<string> Cidr { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        public PrefixListEntryArgs()
        {
        }
    }
}