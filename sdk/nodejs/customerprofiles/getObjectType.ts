// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An ObjectType resource of Amazon Connect Customer Profiles
 */
export function getObjectType(args: GetObjectTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetObjectTypeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:customerprofiles:getObjectType", {
        "domainName": args.domainName,
        "objectTypeName": args.objectTypeName,
    }, opts);
}

export interface GetObjectTypeArgs {
    /**
     * The unique name of the domain.
     */
    domainName: string;
    /**
     * The name of the profile object type.
     */
    objectTypeName: string;
}

export interface GetObjectTypeResult {
    /**
     * Indicates whether a profile should be created when data is received.
     */
    readonly allowProfileCreation?: boolean;
    /**
     * The time of this integration got created.
     */
    readonly createdAt?: string;
    /**
     * Description of the profile object type.
     */
    readonly description?: string;
    /**
     * The default encryption key
     */
    readonly encryptionKey?: string;
    /**
     * The default number of days until the data within the domain expires.
     */
    readonly expirationDays?: number;
    /**
     * A list of the name and ObjectType field.
     */
    readonly fields?: outputs.customerprofiles.ObjectTypeFieldMap[];
    /**
     * A list of unique keys that can be used to map data to the profile.
     */
    readonly keys?: outputs.customerprofiles.ObjectTypeKeyMap[];
    /**
     * The time of this integration got last updated at.
     */
    readonly lastUpdatedAt?: string;
    /**
     * The tags (keys and values) associated with the integration.
     */
    readonly tags?: outputs.customerprofiles.ObjectTypeTag[];
    /**
     * A unique identifier for the object template.
     */
    readonly templateId?: string;
}

export function getObjectTypeOutput(args: GetObjectTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetObjectTypeResult> {
    return pulumi.output(args).apply(a => getObjectType(a, opts))
}

export interface GetObjectTypeOutputArgs {
    /**
     * The unique name of the domain.
     */
    domainName: pulumi.Input<string>;
    /**
     * The name of the profile object type.
     */
    objectTypeName: pulumi.Input<string>;
}