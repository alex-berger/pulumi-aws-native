// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.
 */
export function getKeySigningKey(args: GetKeySigningKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetKeySigningKeyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53:getKeySigningKey", {
        "hostedZoneId": args.hostedZoneId,
        "name": args.name,
    }, opts);
}

export interface GetKeySigningKeyArgs {
    /**
     * The unique string (ID) used to identify a hosted zone.
     */
    hostedZoneId: string;
    /**
     * An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
     */
    name: string;
}

export interface GetKeySigningKeyResult {
    /**
     * A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
     */
    readonly status?: enums.route53.KeySigningKeyStatus;
}

export function getKeySigningKeyOutput(args: GetKeySigningKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKeySigningKeyResult> {
    return pulumi.output(args).apply(a => getKeySigningKey(a, opts))
}

export interface GetKeySigningKeyOutputArgs {
    /**
     * The unique string (ID) used to identify a hosted zone.
     */
    hostedZoneId: pulumi.Input<string>;
    /**
     * An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
     */
    name: pulumi.Input<string>;
}
