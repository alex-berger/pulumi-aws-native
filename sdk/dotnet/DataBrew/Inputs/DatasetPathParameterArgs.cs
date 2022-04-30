// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    /// <summary>
    /// A key-value pair to associate dataset parameter name with its definition.
    /// </summary>
    public sealed class DatasetPathParameterArgs : Pulumi.ResourceArgs
    {
        [Input("datasetParameter", required: true)]
        public Input<Inputs.DatasetParameterArgs> DatasetParameter { get; set; } = null!;

        [Input("pathParameterName", required: true)]
        public Input<string> PathParameterName { get; set; } = null!;

        public DatasetPathParameterArgs()
        {
        }
    }
}
