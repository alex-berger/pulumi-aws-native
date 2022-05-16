// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalytics.Outputs
{

    [OutputType]
    public sealed class ApplicationOutputResourceOutput
    {
        public readonly Outputs.ApplicationOutputResourceDestinationSchema DestinationSchema;
        public readonly Outputs.ApplicationOutputResourceKinesisFirehoseOutput? KinesisFirehoseOutput;
        public readonly Outputs.ApplicationOutputResourceKinesisStreamsOutput? KinesisStreamsOutput;
        public readonly Outputs.ApplicationOutputResourceLambdaOutput? LambdaOutput;
        public readonly string? Name;

        [OutputConstructor]
        private ApplicationOutputResourceOutput(
            Outputs.ApplicationOutputResourceDestinationSchema destinationSchema,

            Outputs.ApplicationOutputResourceKinesisFirehoseOutput? kinesisFirehoseOutput,

            Outputs.ApplicationOutputResourceKinesisStreamsOutput? kinesisStreamsOutput,

            Outputs.ApplicationOutputResourceLambdaOutput? lambdaOutput,

            string? name)
        {
            DestinationSchema = destinationSchema;
            KinesisFirehoseOutput = kinesisFirehoseOutput;
            KinesisStreamsOutput = kinesisStreamsOutput;
            LambdaOutput = lambdaOutput;
            Name = name;
        }
    }
}