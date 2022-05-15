// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    public sealed class ApplicationOutputResourceOutputArgs : Pulumi.ResourceArgs
    {
        [Input("destinationSchema", required: true)]
        public Input<Inputs.ApplicationOutputResourceDestinationSchemaArgs> DestinationSchema { get; set; } = null!;

        [Input("kinesisFirehoseOutput")]
        public Input<Inputs.ApplicationOutputResourceKinesisFirehoseOutputArgs>? KinesisFirehoseOutput { get; set; }

        [Input("kinesisStreamsOutput")]
        public Input<Inputs.ApplicationOutputResourceKinesisStreamsOutputArgs>? KinesisStreamsOutput { get; set; }

        [Input("lambdaOutput")]
        public Input<Inputs.ApplicationOutputResourceLambdaOutputArgs>? LambdaOutput { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApplicationOutputResourceOutputArgs()
        {
        }
    }
}
