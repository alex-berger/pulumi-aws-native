// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-stageoverrides.html
    /// </summary>
    [OutputType]
    public sealed class ApiGatewayManagedOverridesStageOverrides
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-stageoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-stageoverrides-accesslogsettings
        /// </summary>
        public readonly Outputs.ApiGatewayManagedOverridesAccessLogSettings? AccessLogSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-stageoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-stageoverrides-autodeploy
        /// </summary>
        public readonly bool? AutoDeploy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-stageoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-stageoverrides-defaultroutesettings
        /// </summary>
        public readonly Outputs.ApiGatewayManagedOverridesRouteSettings? DefaultRouteSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-stageoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-stageoverrides-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-stageoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-stageoverrides-routesettings
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? RouteSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-apigatewaymanagedoverrides-stageoverrides.html#cfn-apigatewayv2-apigatewaymanagedoverrides-stageoverrides-stagevariables
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? StageVariables;

        [OutputConstructor]
        private ApiGatewayManagedOverridesStageOverrides(
            Outputs.ApiGatewayManagedOverridesAccessLogSettings? accessLogSettings,

            bool? autoDeploy,

            Outputs.ApiGatewayManagedOverridesRouteSettings? defaultRouteSettings,

            string? description,

            Union<System.Text.Json.JsonElement, string>? routeSettings,

            Union<System.Text.Json.JsonElement, string>? stageVariables)
        {
            AccessLogSettings = accessLogSettings;
            AutoDeploy = autoDeploy;
            DefaultRouteSettings = defaultRouteSettings;
            Description = description;
            RouteSettings = routeSettings;
            StageVariables = stageVariables;
        }
    }
}