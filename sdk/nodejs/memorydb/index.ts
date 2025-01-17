// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./acl";
export * from "./cluster";
export * from "./getACL";
export * from "./getCluster";
export * from "./getParameterGroup";
export * from "./getSubnetGroup";
export * from "./getUser";
export * from "./parameterGroup";
export * from "./subnetGroup";
export * from "./user";

// Export enums:
export * from "../types/enums/memorydb";

// Import resources to register:
import { ACL } from "./acl";
import { Cluster } from "./cluster";
import { ParameterGroup } from "./parameterGroup";
import { SubnetGroup } from "./subnetGroup";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:memorydb:ACL":
                return new ACL(name, <any>undefined, { urn })
            case "aws-native:memorydb:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "aws-native:memorydb:ParameterGroup":
                return new ParameterGroup(name, <any>undefined, { urn })
            case "aws-native:memorydb:SubnetGroup":
                return new SubnetGroup(name, <any>undefined, { urn })
            case "aws-native:memorydb:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "memorydb", _module)
