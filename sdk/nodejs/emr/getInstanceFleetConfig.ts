// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EMR::InstanceFleetConfig
 */
export function getInstanceFleetConfig(args: GetInstanceFleetConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceFleetConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:emr:getInstanceFleetConfig", {
        "id": args.id,
    }, opts);
}

export interface GetInstanceFleetConfigArgs {
    id: string;
}

export interface GetInstanceFleetConfigResult {
    readonly id?: string;
    readonly targetOnDemandCapacity?: number;
    readonly targetSpotCapacity?: number;
}

export function getInstanceFleetConfigOutput(args: GetInstanceFleetConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceFleetConfigResult> {
    return pulumi.output(args).apply(a => getInstanceFleetConfig(a, opts))
}

export interface GetInstanceFleetConfigOutputArgs {
    id: pulumi.Input<string>;
}
