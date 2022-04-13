// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RefactorSpaces.Inputs
{

    public sealed class RouteUriPathRouteInputArgs : Pulumi.ResourceArgs
    {
        [Input("activationState", required: true)]
        public Input<Pulumi.AwsNative.RefactorSpaces.RouteActivationState> ActivationState { get; set; } = null!;

        [Input("includeChildPaths")]
        public Input<bool>? IncludeChildPaths { get; set; }

        [Input("methods")]
        private InputList<Pulumi.AwsNative.RefactorSpaces.RouteMethod>? _methods;
        public InputList<Pulumi.AwsNative.RefactorSpaces.RouteMethod> Methods
        {
            get => _methods ?? (_methods = new InputList<Pulumi.AwsNative.RefactorSpaces.RouteMethod>());
            set => _methods = value;
        }

        [Input("sourcePath")]
        public Input<string>? SourcePath { get; set; }

        public RouteUriPathRouteInputArgs()
        {
        }
    }
}
