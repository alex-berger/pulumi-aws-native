// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGatewayV2::Integration
 */
export function getIntegration(args: GetIntegrationArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigatewayv2:getIntegration", {
        "id": args.id,
    }, opts);
}

export interface GetIntegrationArgs {
    id: string;
}

export interface GetIntegrationResult {
    readonly connectionId?: string;
    readonly connectionType?: string;
    readonly contentHandlingStrategy?: string;
    readonly credentialsArn?: string;
    readonly description?: string;
    readonly id?: string;
    readonly integrationMethod?: string;
    readonly integrationSubtype?: string;
    readonly integrationType?: string;
    readonly integrationUri?: string;
    readonly passthroughBehavior?: string;
    readonly payloadFormatVersion?: string;
    readonly requestParameters?: any;
    readonly requestTemplates?: any;
    readonly responseParameters?: any;
    readonly templateSelectionExpression?: string;
    readonly timeoutInMillis?: number;
    readonly tlsConfig?: outputs.apigatewayv2.IntegrationTlsConfig;
}

export function getIntegrationOutput(args: GetIntegrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIntegrationResult> {
    return pulumi.output(args).apply(a => getIntegration(a, opts))
}

export interface GetIntegrationOutputArgs {
    id: pulumi.Input<string>;
}