// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Neptune
{
    public static class GetDBCluster
    {
        /// <summary>
        /// Resource Type definition for AWS::Neptune::DBCluster
        /// </summary>
        public static Task<GetDBClusterResult> InvokeAsync(GetDBClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDBClusterResult>("aws-native:neptune:getDBCluster", args ?? new GetDBClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Neptune::DBCluster
        /// </summary>
        public static Output<GetDBClusterResult> Invoke(GetDBClusterInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDBClusterResult>("aws-native:neptune:getDBCluster", args ?? new GetDBClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBClusterArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDBClusterArgs()
        {
        }
    }

    public sealed class GetDBClusterInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDBClusterInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDBClusterResult
    {
        public readonly ImmutableArray<Outputs.DBClusterRole> AssociatedRoles;
        public readonly int? BackupRetentionPeriod;
        public readonly string? ClusterResourceId;
        public readonly string? DBClusterParameterGroupName;
        public readonly bool? DeletionProtection;
        public readonly ImmutableArray<string> EnableCloudwatchLogsExports;
        public readonly string? Endpoint;
        public readonly bool? IamAuthEnabled;
        public readonly string? Id;
        public readonly int? Port;
        public readonly string? PreferredBackupWindow;
        public readonly string? PreferredMaintenanceWindow;
        public readonly string? ReadEndpoint;
        public readonly ImmutableArray<Outputs.DBClusterTag> Tags;
        public readonly ImmutableArray<string> VpcSecurityGroupIds;

        [OutputConstructor]
        private GetDBClusterResult(
            ImmutableArray<Outputs.DBClusterRole> associatedRoles,

            int? backupRetentionPeriod,

            string? clusterResourceId,

            string? dBClusterParameterGroupName,

            bool? deletionProtection,

            ImmutableArray<string> enableCloudwatchLogsExports,

            string? endpoint,

            bool? iamAuthEnabled,

            string? id,

            int? port,

            string? preferredBackupWindow,

            string? preferredMaintenanceWindow,

            string? readEndpoint,

            ImmutableArray<Outputs.DBClusterTag> tags,

            ImmutableArray<string> vpcSecurityGroupIds)
        {
            AssociatedRoles = associatedRoles;
            BackupRetentionPeriod = backupRetentionPeriod;
            ClusterResourceId = clusterResourceId;
            DBClusterParameterGroupName = dBClusterParameterGroupName;
            DeletionProtection = deletionProtection;
            EnableCloudwatchLogsExports = enableCloudwatchLogsExports;
            Endpoint = endpoint;
            IamAuthEnabled = iamAuthEnabled;
            Id = id;
            Port = port;
            PreferredBackupWindow = preferredBackupWindow;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            ReadEndpoint = readEndpoint;
            Tags = tags;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
        }
    }
}
