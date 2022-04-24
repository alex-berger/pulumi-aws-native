// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Inputs
{

    public sealed class OrganizationConfigRuleOrganizationCustomRuleMetadataArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputParameters")]
        public Input<string>? InputParameters { get; set; }

        [Input("lambdaFunctionArn", required: true)]
        public Input<string> LambdaFunctionArn { get; set; } = null!;

        [Input("maximumExecutionFrequency")]
        public Input<string>? MaximumExecutionFrequency { get; set; }

        [Input("organizationConfigRuleTriggerTypes", required: true)]
        private InputList<string>? _organizationConfigRuleTriggerTypes;
        public InputList<string> OrganizationConfigRuleTriggerTypes
        {
            get => _organizationConfigRuleTriggerTypes ?? (_organizationConfigRuleTriggerTypes = new InputList<string>());
            set => _organizationConfigRuleTriggerTypes = value;
        }

        [Input("resourceIdScope")]
        public Input<string>? ResourceIdScope { get; set; }

        [Input("resourceTypesScope")]
        private InputList<string>? _resourceTypesScope;
        public InputList<string> ResourceTypesScope
        {
            get => _resourceTypesScope ?? (_resourceTypesScope = new InputList<string>());
            set => _resourceTypesScope = value;
        }

        [Input("tagKeyScope")]
        public Input<string>? TagKeyScope { get; set; }

        [Input("tagValueScope")]
        public Input<string>? TagValueScope { get; set; }

        public OrganizationConfigRuleOrganizationCustomRuleMetadataArgs()
        {
        }
    }
}
