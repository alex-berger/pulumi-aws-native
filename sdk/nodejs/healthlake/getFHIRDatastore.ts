// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * HealthLake FHIR Datastore
 */
export function getFHIRDatastore(args: GetFHIRDatastoreArgs, opts?: pulumi.InvokeOptions): Promise<GetFHIRDatastoreResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:healthlake:getFHIRDatastore", {
        "datastoreId": args.datastoreId,
    }, opts);
}

export interface GetFHIRDatastoreArgs {
    datastoreId: string;
}

export interface GetFHIRDatastoreResult {
    readonly createdAt?: outputs.healthlake.FHIRDatastoreCreatedAt;
    readonly datastoreArn?: string;
    readonly datastoreEndpoint?: string;
    readonly datastoreId?: string;
    readonly datastoreStatus?: enums.healthlake.FHIRDatastoreDatastoreStatus;
    readonly tags?: outputs.healthlake.FHIRDatastoreTag[];
}

export function getFHIRDatastoreOutput(args: GetFHIRDatastoreOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFHIRDatastoreResult> {
    return pulumi.output(args).apply(a => getFHIRDatastore(a, opts))
}

export interface GetFHIRDatastoreOutputArgs {
    datastoreId: pulumi.Input<string>;
}