// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    public sealed class RuleRedshiftDataParametersArgs : Pulumi.ResourceArgs
    {
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("dbUser")]
        public Input<string>? DbUser { get; set; }

        [Input("secretManagerArn")]
        public Input<string>? SecretManagerArn { get; set; }

        [Input("sql", required: true)]
        public Input<string> Sql { get; set; } = null!;

        [Input("statementName")]
        public Input<string>? StatementName { get; set; }

        [Input("withEvent")]
        public Input<bool>? WithEvent { get; set; }

        public RuleRedshiftDataParametersArgs()
        {
        }
    }
}
