// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceAclConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("allowedGroupsColumnName", required: true)]
        public Input<string> AllowedGroupsColumnName { get; set; } = null!;

        public DataSourceAclConfigurationArgs()
        {
        }
    }
}
