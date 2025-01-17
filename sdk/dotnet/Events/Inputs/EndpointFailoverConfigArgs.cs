// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    public sealed class EndpointFailoverConfigArgs : Pulumi.ResourceArgs
    {
        [Input("primary", required: true)]
        public Input<Inputs.EndpointPrimaryArgs> Primary { get; set; } = null!;

        [Input("secondary", required: true)]
        public Input<Inputs.EndpointSecondaryArgs> Secondary { get; set; } = null!;

        public EndpointFailoverConfigArgs()
        {
        }
    }
}
