// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class EndpointConfigServerlessConfigArgs : Pulumi.ResourceArgs
    {
        [Input("maxConcurrency", required: true)]
        public Input<int> MaxConcurrency { get; set; } = null!;

        [Input("memorySizeInMB", required: true)]
        public Input<int> MemorySizeInMB { get; set; } = null!;

        public EndpointConfigServerlessConfigArgs()
        {
        }
    }
}
