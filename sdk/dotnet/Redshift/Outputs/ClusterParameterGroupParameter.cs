// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift.Outputs
{

    [OutputType]
    public sealed class ClusterParameterGroupParameter
    {
        public readonly string ParameterName;
        public readonly string ParameterValue;

        [OutputConstructor]
        private ClusterParameterGroupParameter(
            string parameterName,

            string parameterValue)
        {
            ParameterName = parameterName;
            ParameterValue = parameterValue;
        }
    }
}