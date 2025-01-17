// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Definition of AWS::RefactorSpaces::Route Resource Type
 */
export function getRoute(args: GetRouteArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:refactorspaces:getRoute", {
        "applicationIdentifier": args.applicationIdentifier,
        "environmentIdentifier": args.environmentIdentifier,
        "routeIdentifier": args.routeIdentifier,
    }, opts);
}

export interface GetRouteArgs {
    applicationIdentifier: string;
    environmentIdentifier: string;
    routeIdentifier: string;
}

export interface GetRouteResult {
    readonly arn?: string;
    readonly pathResourceToId?: string;
    readonly routeIdentifier?: string;
    /**
     * Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.
     */
    readonly tags?: outputs.refactorspaces.RouteTag[];
}

export function getRouteOutput(args: GetRouteOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteResult> {
    return pulumi.output(args).apply(a => getRoute(a, opts))
}

export interface GetRouteOutputArgs {
    applicationIdentifier: pulumi.Input<string>;
    environmentIdentifier: pulumi.Input<string>;
    routeIdentifier: pulumi.Input<string>;
}
