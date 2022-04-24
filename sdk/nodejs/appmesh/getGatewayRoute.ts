// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppMesh::GatewayRoute
 */
export function getGatewayRoute(args: GetGatewayRouteArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayRouteResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appmesh:getGatewayRoute", {
        "id": args.id,
    }, opts);
}

export interface GetGatewayRouteArgs {
    id: string;
}

export interface GetGatewayRouteResult {
    readonly arn?: string;
    readonly id?: string;
    readonly resourceOwner?: string;
    readonly spec?: outputs.appmesh.GatewayRouteSpec;
    readonly tags?: outputs.appmesh.GatewayRouteTag[];
    readonly uid?: string;
}

export function getGatewayRouteOutput(args: GetGatewayRouteOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayRouteResult> {
    return pulumi.output(args).apply(a => getGatewayRoute(a, opts))
}

export interface GetGatewayRouteOutputArgs {
    id: pulumi.Input<string>;
}
