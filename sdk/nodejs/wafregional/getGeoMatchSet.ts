// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAFRegional::GeoMatchSet
 */
export function getGeoMatchSet(args: GetGeoMatchSetArgs, opts?: pulumi.InvokeOptions): Promise<GetGeoMatchSetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:wafregional:getGeoMatchSet", {
        "id": args.id,
    }, opts);
}

export interface GetGeoMatchSetArgs {
    id: string;
}

export interface GetGeoMatchSetResult {
    readonly geoMatchConstraints?: outputs.wafregional.GeoMatchSetGeoMatchConstraint[];
    readonly id?: string;
}

export function getGeoMatchSetOutput(args: GetGeoMatchSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGeoMatchSetResult> {
    return pulumi.output(args).apply(a => getGeoMatchSet(a, opts))
}

export interface GetGeoMatchSetOutputArgs {
    id: pulumi.Input<string>;
}
