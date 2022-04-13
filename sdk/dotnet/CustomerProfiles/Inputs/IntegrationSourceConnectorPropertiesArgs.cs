// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Inputs
{

    public sealed class IntegrationSourceConnectorPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("marketo")]
        public Input<Inputs.IntegrationMarketoSourcePropertiesArgs>? Marketo { get; set; }

        [Input("s3")]
        public Input<Inputs.IntegrationS3SourcePropertiesArgs>? S3 { get; set; }

        [Input("salesforce")]
        public Input<Inputs.IntegrationSalesforceSourcePropertiesArgs>? Salesforce { get; set; }

        [Input("serviceNow")]
        public Input<Inputs.IntegrationServiceNowSourcePropertiesArgs>? ServiceNow { get; set; }

        [Input("zendesk")]
        public Input<Inputs.IntegrationZendeskSourcePropertiesArgs>? Zendesk { get; set; }

        public IntegrationSourceConnectorPropertiesArgs()
        {
        }
    }
}
