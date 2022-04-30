// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail
{
    public static class GetDatabase
    {
        /// <summary>
        /// Resource Type definition for AWS::Lightsail::Database
        /// </summary>
        public static Task<GetDatabaseResult> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseResult>("aws-native:lightsail:getDatabase", args ?? new GetDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Lightsail::Database
        /// </summary>
        public static Output<GetDatabaseResult> Invoke(GetDatabaseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatabaseResult>("aws-native:lightsail:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name to use for your new Lightsail database resource.
        /// </summary>
        [Input("relationalDatabaseName", required: true)]
        public string RelationalDatabaseName { get; set; } = null!;

        public GetDatabaseArgs()
        {
        }
    }

    public sealed class GetDatabaseInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name to use for your new Lightsail database resource.
        /// </summary>
        [Input("relationalDatabaseName", required: true)]
        public Input<string> RelationalDatabaseName { get; set; } = null!;

        public GetDatabaseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseResult
    {
        /// <summary>
        /// When true, enables automated backup retention for your database. Updates are applied during the next maintenance window because this can result in an outage.
        /// </summary>
        public readonly bool? BackupRetention;
        /// <summary>
        /// Indicates the certificate that needs to be associated with the database.
        /// </summary>
        public readonly string? CaCertificateIdentifier;
        public readonly string? DatabaseArn;
        /// <summary>
        /// The daily time range during which automated backups are created for your new database if automated backups are enabled.
        /// </summary>
        public readonly string? PreferredBackupWindow;
        /// <summary>
        /// The weekly time range during which system maintenance can occur on your new database.
        /// </summary>
        public readonly string? PreferredMaintenanceWindow;
        /// <summary>
        /// Specifies the accessibility options for your new database. A value of true specifies a database that is available to resources outside of your Lightsail account. A value of false specifies a database that is available only to your Lightsail resources in the same region as your database.
        /// </summary>
        public readonly bool? PubliclyAccessible;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseTag> Tags;

        [OutputConstructor]
        private GetDatabaseResult(
            bool? backupRetention,

            string? caCertificateIdentifier,

            string? databaseArn,

            string? preferredBackupWindow,

            string? preferredMaintenanceWindow,

            bool? publiclyAccessible,

            ImmutableArray<Outputs.DatabaseTag> tags)
        {
            BackupRetention = backupRetention;
            CaCertificateIdentifier = caCertificateIdentifier;
            DatabaseArn = databaseArn;
            PreferredBackupWindow = preferredBackupWindow;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            PubliclyAccessible = publiclyAccessible;
            Tags = tags;
        }
    }
}
