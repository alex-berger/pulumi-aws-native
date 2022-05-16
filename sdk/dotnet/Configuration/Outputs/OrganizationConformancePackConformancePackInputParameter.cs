// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    /// <summary>
    /// Input parameters in the form of key-value pairs for the conformance pack.
    /// </summary>
    [OutputType]
    public sealed class OrganizationConformancePackConformancePackInputParameter
    {
        public readonly string ParameterName;
        public readonly string ParameterValue;

        [OutputConstructor]
        private OrganizationConformancePackConformancePackInputParameter(
            string parameterName,

            string parameterValue)
        {
            ParameterName = parameterName;
            ParameterValue = parameterValue;
        }
    }
}