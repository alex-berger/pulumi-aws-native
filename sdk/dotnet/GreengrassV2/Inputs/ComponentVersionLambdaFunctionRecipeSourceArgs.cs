// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Inputs
{

    public sealed class ComponentVersionLambdaFunctionRecipeSourceArgs : Pulumi.ResourceArgs
    {
        [Input("componentDependencies")]
        public Input<object>? ComponentDependencies { get; set; }

        [Input("componentLambdaParameters")]
        public Input<Inputs.ComponentVersionLambdaExecutionParametersArgs>? ComponentLambdaParameters { get; set; }

        [Input("componentName")]
        public Input<string>? ComponentName { get; set; }

        [Input("componentPlatforms")]
        private InputList<Inputs.ComponentVersionComponentPlatformArgs>? _componentPlatforms;
        public InputList<Inputs.ComponentVersionComponentPlatformArgs> ComponentPlatforms
        {
            get => _componentPlatforms ?? (_componentPlatforms = new InputList<Inputs.ComponentVersionComponentPlatformArgs>());
            set => _componentPlatforms = value;
        }

        [Input("componentVersion")]
        public Input<string>? ComponentVersion { get; set; }

        [Input("lambdaArn")]
        public Input<string>? LambdaArn { get; set; }

        public ComponentVersionLambdaFunctionRecipeSourceArgs()
        {
        }
    }
}
