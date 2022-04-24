// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RefactorSpaces.Outputs
{

    [OutputType]
    public sealed class ServiceUrlEndpointInput
    {
        public readonly string? HealthUrl;
        public readonly string Url;

        [OutputConstructor]
        private ServiceUrlEndpointInput(
            string? healthUrl,

            string url)
        {
            HealthUrl = healthUrl;
            Url = url;
        }
    }
}
