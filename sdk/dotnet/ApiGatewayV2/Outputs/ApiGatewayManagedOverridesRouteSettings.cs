// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Outputs
{

    [OutputType]
    public sealed class ApiGatewayManagedOverridesRouteSettings
    {
        public readonly bool? DataTraceEnabled;
        public readonly bool? DetailedMetricsEnabled;
        public readonly string? LoggingLevel;
        public readonly int? ThrottlingBurstLimit;
        public readonly double? ThrottlingRateLimit;

        [OutputConstructor]
        private ApiGatewayManagedOverridesRouteSettings(
            bool? dataTraceEnabled,

            bool? detailedMetricsEnabled,

            string? loggingLevel,

            int? throttlingBurstLimit,

            double? throttlingRateLimit)
        {
            DataTraceEnabled = dataTraceEnabled;
            DetailedMetricsEnabled = detailedMetricsEnabled;
            LoggingLevel = loggingLevel;
            ThrottlingBurstLimit = throttlingBurstLimit;
            ThrottlingRateLimit = throttlingRateLimit;
        }
    }
}
