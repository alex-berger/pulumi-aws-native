// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alias";
export * from "./getAlias";
export * from "./getKey";
export * from "./getReplicaKey";
export * from "./key";
export * from "./replicaKey";

// Export enums:
export * from "../types/enums/kms";

// Import resources to register:
import { Alias } from "./alias";
import { Key } from "./key";
import { ReplicaKey } from "./replicaKey";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:kms:Alias":
                return new Alias(name, <any>undefined, { urn })
            case "aws-native:kms:Key":
                return new Key(name, <any>undefined, { urn })
            case "aws-native:kms:ReplicaKey":
                return new ReplicaKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "kms", _module)