// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Macie.Inputs
{

    public sealed class FindingsFilterFindingCriteriaArgs : Pulumi.ResourceArgs
    {
        [Input("criterion")]
        public Input<Inputs.FindingsFilterCriterionArgs>? Criterion { get; set; }

        public FindingsFilterFindingCriteriaArgs()
        {
        }
    }
}