// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    public sealed class JobStatisticOverrideArgs : Pulumi.ResourceArgs
    {
        [Input("parameters", required: true)]
        public Input<Inputs.JobParameterMapArgs> Parameters { get; set; } = null!;

        [Input("statistic", required: true)]
        public Input<string> Statistic { get; set; } = null!;

        public JobStatisticOverrideArgs()
        {
        }
    }
}
