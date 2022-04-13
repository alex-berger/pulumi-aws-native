// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getPublicRepository";
export * from "./getPullThroughCacheRule";
export * from "./getRegistryPolicy";
export * from "./getReplicationConfiguration";
export * from "./getRepository";
export * from "./publicRepository";
export * from "./pullThroughCacheRule";
export * from "./registryPolicy";
export * from "./replicationConfiguration";
export * from "./repository";

// Export enums:
export * from "../types/enums/ecr";

// Import resources to register:
import { PublicRepository } from "./publicRepository";
import { PullThroughCacheRule } from "./pullThroughCacheRule";
import { RegistryPolicy } from "./registryPolicy";
import { ReplicationConfiguration } from "./replicationConfiguration";
import { Repository } from "./repository";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:ecr:PublicRepository":
                return new PublicRepository(name, <any>undefined, { urn })
            case "aws-native:ecr:PullThroughCacheRule":
                return new PullThroughCacheRule(name, <any>undefined, { urn })
            case "aws-native:ecr:RegistryPolicy":
                return new RegistryPolicy(name, <any>undefined, { urn })
            case "aws-native:ecr:ReplicationConfiguration":
                return new ReplicationConfiguration(name, <any>undefined, { urn })
            case "aws-native:ecr:Repository":
                return new Repository(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "ecr", _module)
