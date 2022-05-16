// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Outputs
{

    [OutputType]
    public sealed class IntegrationSourceFlowConfig
    {
        public readonly string? ConnectorProfileName;
        public readonly Pulumi.AwsNative.CustomerProfiles.IntegrationConnectorType ConnectorType;
        public readonly Outputs.IntegrationIncrementalPullConfig? IncrementalPullConfig;
        public readonly Outputs.IntegrationSourceConnectorProperties SourceConnectorProperties;

        [OutputConstructor]
        private IntegrationSourceFlowConfig(
            string? connectorProfileName,

            Pulumi.AwsNative.CustomerProfiles.IntegrationConnectorType connectorType,

            Outputs.IntegrationIncrementalPullConfig? incrementalPullConfig,

            Outputs.IntegrationSourceConnectorProperties sourceConnectorProperties)
        {
            ConnectorProfileName = connectorProfileName;
            ConnectorType = connectorType;
            IncrementalPullConfig = incrementalPullConfig;
            SourceConnectorProperties = sourceConnectorProperties;
        }
    }
}
