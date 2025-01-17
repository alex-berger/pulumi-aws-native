// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./discoverer";
export * from "./getDiscoverer";
export * from "./getRegistry";
export * from "./getRegistryPolicy";
export * from "./getSchema";
export * from "./registry";
export * from "./registryPolicy";
export * from "./schema";

// Import resources to register:
import { Discoverer } from "./discoverer";
import { Registry } from "./registry";
import { RegistryPolicy } from "./registryPolicy";
import { Schema } from "./schema";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:eventschemas:Discoverer":
                return new Discoverer(name, <any>undefined, { urn })
            case "aws-native:eventschemas:Registry":
                return new Registry(name, <any>undefined, { urn })
            case "aws-native:eventschemas:RegistryPolicy":
                return new RegistryPolicy(name, <any>undefined, { urn })
            case "aws-native:eventschemas:Schema":
                return new Schema(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "eventschemas", _module)
