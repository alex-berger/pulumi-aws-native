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
    /// Details of the client authentication used by the Kafka cluster.
    /// </summary>
    [OutputType]
    public sealed class ConnectorKafkaClusterClientAuthentication
    {
        public readonly Pulumi.AwsNative.KafkaConnect.ConnectorKafkaClusterClientAuthenticationType AuthenticationType;

        [OutputConstructor]
        private ConnectorKafkaClusterClientAuthentication(Pulumi.AwsNative.KafkaConnect.ConnectorKafkaClusterClientAuthenticationType authenticationType)
        {
            AuthenticationType = authenticationType;
        }
    }
}
