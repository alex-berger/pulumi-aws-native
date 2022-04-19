// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class JobDatabaseTableOutputOptions
    {
        public readonly string TableName;
        public readonly Outputs.JobS3Location? TempDirectory;

        [OutputConstructor]
        private JobDatabaseTableOutputOptions(
            string tableName,

            Outputs.JobS3Location? tempDirectory)
        {
            TableName = tableName;
            TempDirectory = tempDirectory;
        }
    }
}
