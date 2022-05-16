// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    public sealed class JobDatabaseTableOutputOptionsArgs : Pulumi.ResourceArgs
    {
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        [Input("tempDirectory")]
        public Input<Inputs.JobS3LocationArgs>? TempDirectory { get; set; }

        public JobDatabaseTableOutputOptionsArgs()
        {
        }
    }
}
