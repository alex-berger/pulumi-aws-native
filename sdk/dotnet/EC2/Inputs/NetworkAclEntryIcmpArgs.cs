// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class NetworkAclEntryIcmpArgs : Pulumi.ResourceArgs
    {
        [Input("code")]
        public Input<int>? Code { get; set; }

        [Input("type")]
        public Input<int>? Type { get; set; }

        public NetworkAclEntryIcmpArgs()
        {
        }
    }
}
