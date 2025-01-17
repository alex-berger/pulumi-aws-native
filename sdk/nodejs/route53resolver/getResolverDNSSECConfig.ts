// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.
 */
export function getResolverDNSSECConfig(args: GetResolverDNSSECConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetResolverDNSSECConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53resolver:getResolverDNSSECConfig", {
        "id": args.id,
    }, opts);
}

export interface GetResolverDNSSECConfigArgs {
    /**
     * Id
     */
    id: string;
}

export interface GetResolverDNSSECConfigResult {
    /**
     * Id
     */
    readonly id?: string;
    /**
     * AccountId
     */
    readonly ownerId?: string;
    /**
     * ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
     */
    readonly validationStatus?: enums.route53resolver.ResolverDNSSECConfigValidationStatus;
}

export function getResolverDNSSECConfigOutput(args: GetResolverDNSSECConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResolverDNSSECConfigResult> {
    return pulumi.output(args).apply(a => getResolverDNSSECConfig(a, opts))
}

export interface GetResolverDNSSECConfigOutputArgs {
    /**
     * Id
     */
    id: pulumi.Input<string>;
}
