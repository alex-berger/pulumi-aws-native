// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS
{
    public static class GetEndpoint
    {
        /// <summary>
        /// Resource Type definition for AWS::DMS::Endpoint
        /// </summary>
        public static Task<GetEndpointResult> InvokeAsync(GetEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEndpointResult>("aws-native:dms:getEndpoint", args ?? new GetEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DMS::Endpoint
        /// </summary>
        public static Output<GetEndpointResult> Invoke(GetEndpointInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEndpointResult>("aws-native:dms:getEndpoint", args ?? new GetEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEndpointArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetEndpointArgs()
        {
        }
    }

    public sealed class GetEndpointInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetEndpointInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEndpointResult
    {
        public readonly string? CertificateArn;
        public readonly string? DatabaseName;
        public readonly Outputs.EndpointDocDbSettings? DocDbSettings;
        public readonly Outputs.EndpointDynamoDbSettings? DynamoDbSettings;
        public readonly Outputs.EndpointElasticsearchSettings? ElasticsearchSettings;
        public readonly string? EndpointIdentifier;
        public readonly string? EndpointType;
        public readonly string? EngineName;
        public readonly string? ExternalId;
        public readonly string? ExtraConnectionAttributes;
        public readonly Outputs.EndpointGcpMySQLSettings? GcpMySQLSettings;
        public readonly Outputs.EndpointIbmDb2Settings? IbmDb2Settings;
        public readonly string? Id;
        public readonly Outputs.EndpointKafkaSettings? KafkaSettings;
        public readonly Outputs.EndpointKinesisSettings? KinesisSettings;
        public readonly Outputs.EndpointMicrosoftSqlServerSettings? MicrosoftSqlServerSettings;
        public readonly Outputs.EndpointMongoDbSettings? MongoDbSettings;
        public readonly Outputs.EndpointMySqlSettings? MySqlSettings;
        public readonly Outputs.EndpointNeptuneSettings? NeptuneSettings;
        public readonly Outputs.EndpointOracleSettings? OracleSettings;
        public readonly string? Password;
        public readonly int? Port;
        public readonly Outputs.EndpointPostgreSqlSettings? PostgreSqlSettings;
        public readonly Outputs.EndpointRedisSettings? RedisSettings;
        public readonly Outputs.EndpointRedshiftSettings? RedshiftSettings;
        public readonly Outputs.EndpointS3Settings? S3Settings;
        public readonly string? ServerName;
        public readonly string? SslMode;
        public readonly Outputs.EndpointSybaseSettings? SybaseSettings;
        public readonly ImmutableArray<Outputs.EndpointTag> Tags;
        public readonly string? Username;

        [OutputConstructor]
        private GetEndpointResult(
            string? certificateArn,

            string? databaseName,

            Outputs.EndpointDocDbSettings? docDbSettings,

            Outputs.EndpointDynamoDbSettings? dynamoDbSettings,

            Outputs.EndpointElasticsearchSettings? elasticsearchSettings,

            string? endpointIdentifier,

            string? endpointType,

            string? engineName,

            string? externalId,

            string? extraConnectionAttributes,

            Outputs.EndpointGcpMySQLSettings? gcpMySQLSettings,

            Outputs.EndpointIbmDb2Settings? ibmDb2Settings,

            string? id,

            Outputs.EndpointKafkaSettings? kafkaSettings,

            Outputs.EndpointKinesisSettings? kinesisSettings,

            Outputs.EndpointMicrosoftSqlServerSettings? microsoftSqlServerSettings,

            Outputs.EndpointMongoDbSettings? mongoDbSettings,

            Outputs.EndpointMySqlSettings? mySqlSettings,

            Outputs.EndpointNeptuneSettings? neptuneSettings,

            Outputs.EndpointOracleSettings? oracleSettings,

            string? password,

            int? port,

            Outputs.EndpointPostgreSqlSettings? postgreSqlSettings,

            Outputs.EndpointRedisSettings? redisSettings,

            Outputs.EndpointRedshiftSettings? redshiftSettings,

            Outputs.EndpointS3Settings? s3Settings,

            string? serverName,

            string? sslMode,

            Outputs.EndpointSybaseSettings? sybaseSettings,

            ImmutableArray<Outputs.EndpointTag> tags,

            string? username)
        {
            CertificateArn = certificateArn;
            DatabaseName = databaseName;
            DocDbSettings = docDbSettings;
            DynamoDbSettings = dynamoDbSettings;
            ElasticsearchSettings = elasticsearchSettings;
            EndpointIdentifier = endpointIdentifier;
            EndpointType = endpointType;
            EngineName = engineName;
            ExternalId = externalId;
            ExtraConnectionAttributes = extraConnectionAttributes;
            GcpMySQLSettings = gcpMySQLSettings;
            IbmDb2Settings = ibmDb2Settings;
            Id = id;
            KafkaSettings = kafkaSettings;
            KinesisSettings = kinesisSettings;
            MicrosoftSqlServerSettings = microsoftSqlServerSettings;
            MongoDbSettings = mongoDbSettings;
            MySqlSettings = mySqlSettings;
            NeptuneSettings = neptuneSettings;
            OracleSettings = oracleSettings;
            Password = password;
            Port = port;
            PostgreSqlSettings = postgreSqlSettings;
            RedisSettings = redisSettings;
            RedshiftSettings = redshiftSettings;
            S3Settings = s3Settings;
            ServerName = serverName;
            SslMode = sslMode;
            SybaseSettings = sybaseSettings;
            Tags = tags;
            Username = username;
        }
    }
}
