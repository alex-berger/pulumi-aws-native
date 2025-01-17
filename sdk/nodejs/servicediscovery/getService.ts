// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceDiscovery::Service
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:servicediscovery:getService", {
        "id": args.id,
    }, opts);
}

export interface GetServiceArgs {
    id: string;
}

export interface GetServiceResult {
    readonly arn?: string;
    readonly description?: string;
    readonly dnsConfig?: outputs.servicediscovery.ServiceDnsConfig;
    readonly healthCheckConfig?: outputs.servicediscovery.ServiceHealthCheckConfig;
    readonly id?: string;
    readonly tags?: outputs.servicediscovery.ServiceTag[];
}

export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply(a => getService(a, opts))
}

export interface GetServiceOutputArgs {
    id: pulumi.Input<string>;
}
