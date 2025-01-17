// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Outputs
{

    [OutputType]
    public sealed class FileSystemLustreConfiguration
    {
        public readonly string? AutoImportPolicy;
        public readonly int? AutomaticBackupRetentionDays;
        public readonly bool? CopyTagsToBackups;
        public readonly string? DailyAutomaticBackupStartTime;
        public readonly string? DataCompressionType;
        public readonly string? DeploymentType;
        public readonly string? DriveCacheType;
        public readonly string? ExportPath;
        public readonly string? ImportPath;
        public readonly int? ImportedFileChunkSize;
        public readonly int? PerUnitStorageThroughput;
        public readonly string? WeeklyMaintenanceStartTime;

        [OutputConstructor]
        private FileSystemLustreConfiguration(
            string? autoImportPolicy,

            int? automaticBackupRetentionDays,

            bool? copyTagsToBackups,

            string? dailyAutomaticBackupStartTime,

            string? dataCompressionType,

            string? deploymentType,

            string? driveCacheType,

            string? exportPath,

            string? importPath,

            int? importedFileChunkSize,

            int? perUnitStorageThroughput,

            string? weeklyMaintenanceStartTime)
        {
            AutoImportPolicy = autoImportPolicy;
            AutomaticBackupRetentionDays = automaticBackupRetentionDays;
            CopyTagsToBackups = copyTagsToBackups;
            DailyAutomaticBackupStartTime = dailyAutomaticBackupStartTime;
            DataCompressionType = dataCompressionType;
            DeploymentType = deploymentType;
            DriveCacheType = driveCacheType;
            ExportPath = exportPath;
            ImportPath = importPath;
            ImportedFileChunkSize = importedFileChunkSize;
            PerUnitStorageThroughput = perUnitStorageThroughput;
            WeeklyMaintenanceStartTime = weeklyMaintenanceStartTime;
        }
    }
}
