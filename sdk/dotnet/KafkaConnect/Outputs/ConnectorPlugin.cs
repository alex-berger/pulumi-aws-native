// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect.Outputs
{

    /// <summary>
    /// Details about a Kafka Connect plugin which will be used with the connector.
    /// </summary>
    [OutputType]
    public sealed class ConnectorPlugin
    {
        public readonly Outputs.ConnectorCustomPlugin CustomPlugin;

        [OutputConstructor]
        private ConnectorPlugin(Outputs.ConnectorCustomPlugin customPlugin)
        {
            CustomPlugin = customPlugin;
        }
    }
}
