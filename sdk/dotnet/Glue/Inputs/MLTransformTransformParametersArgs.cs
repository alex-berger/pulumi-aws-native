// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    public sealed class MLTransformTransformParametersArgs : Pulumi.ResourceArgs
    {
        [Input("findMatchesParameters")]
        public Input<Inputs.MLTransformFindMatchesParametersArgs>? FindMatchesParameters { get; set; }

        [Input("transformType", required: true)]
        public Input<string> TransformType { get; set; } = null!;

        public MLTransformTransformParametersArgs()
        {
        }
    }
}
