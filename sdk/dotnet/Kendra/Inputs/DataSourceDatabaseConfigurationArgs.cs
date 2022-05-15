// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    public sealed class DataSourceDatabaseConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("aclConfiguration")]
        public Input<Inputs.DataSourceAclConfigurationArgs>? AclConfiguration { get; set; }

        [Input("columnConfiguration", required: true)]
        public Input<Inputs.DataSourceColumnConfigurationArgs> ColumnConfiguration { get; set; } = null!;

        [Input("connectionConfiguration", required: true)]
        public Input<Inputs.DataSourceConnectionConfigurationArgs> ConnectionConfiguration { get; set; } = null!;

        [Input("databaseEngineType", required: true)]
        public Input<Pulumi.AwsNative.Kendra.DataSourceDatabaseEngineType> DatabaseEngineType { get; set; } = null!;

        [Input("sqlConfiguration")]
        public Input<Inputs.DataSourceSqlConfigurationArgs>? SqlConfiguration { get; set; }

        [Input("vpcConfiguration")]
        public Input<Inputs.DataSourceVpcConfigurationArgs>? VpcConfiguration { get; set; }

        public DataSourceDatabaseConfigurationArgs()
        {
        }
    }
}
