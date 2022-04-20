// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowErrorHandlingConfigArgs : Pulumi.ResourceArgs
    {
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        [Input("failOnFirstError")]
        public Input<bool>? FailOnFirstError { get; set; }

        public FlowErrorHandlingConfigArgs()
        {
        }
    }
}
