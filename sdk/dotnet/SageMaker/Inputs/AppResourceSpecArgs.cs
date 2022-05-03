// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class AppResourceSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The instance type that the image version runs on.
        /// </summary>
        [Input("instanceType")]
        public Input<Pulumi.AwsNative.SageMaker.AppResourceSpecInstanceType>? InstanceType { get; set; }

        /// <summary>
        /// The ARN of the SageMaker image that the image version belongs to.
        /// </summary>
        [Input("sageMakerImageArn")]
        public Input<string>? SageMakerImageArn { get; set; }

        /// <summary>
        /// The ARN of the image version created on the instance.
        /// </summary>
        [Input("sageMakerImageVersionArn")]
        public Input<string>? SageMakerImageVersionArn { get; set; }

        public AppResourceSpecArgs()
        {
        }
    }
}
