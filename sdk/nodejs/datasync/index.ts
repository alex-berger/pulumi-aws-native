// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./agent";
export * from "./getAgent";
export * from "./getLocationEFS";
export * from "./getLocationFSxLustre";
export * from "./getLocationFSxOpenZFS";
export * from "./getLocationFSxWindows";
export * from "./getLocationHDFS";
export * from "./getLocationNFS";
export * from "./getLocationObjectStorage";
export * from "./getLocationS3";
export * from "./getLocationSMB";
export * from "./getTask";
export * from "./locationEFS";
export * from "./locationFSxLustre";
export * from "./locationFSxOpenZFS";
export * from "./locationFSxWindows";
export * from "./locationHDFS";
export * from "./locationNFS";
export * from "./locationObjectStorage";
export * from "./locationS3";
export * from "./locationSMB";
export * from "./task";

// Export enums:
export * from "../types/enums/datasync";

// Import resources to register:
import { Agent } from "./agent";
import { LocationEFS } from "./locationEFS";
import { LocationFSxLustre } from "./locationFSxLustre";
import { LocationFSxOpenZFS } from "./locationFSxOpenZFS";
import { LocationFSxWindows } from "./locationFSxWindows";
import { LocationHDFS } from "./locationHDFS";
import { LocationNFS } from "./locationNFS";
import { LocationObjectStorage } from "./locationObjectStorage";
import { LocationS3 } from "./locationS3";
import { LocationSMB } from "./locationSMB";
import { Task } from "./task";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:datasync:Agent":
                return new Agent(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationEFS":
                return new LocationEFS(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationFSxLustre":
                return new LocationFSxLustre(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationFSxOpenZFS":
                return new LocationFSxOpenZFS(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationFSxWindows":
                return new LocationFSxWindows(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationHDFS":
                return new LocationHDFS(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationNFS":
                return new LocationNFS(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationObjectStorage":
                return new LocationObjectStorage(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationS3":
                return new LocationS3(name, <any>undefined, { urn })
            case "aws-native:datasync:LocationSMB":
                return new LocationSMB(name, <any>undefined, { urn })
            case "aws-native:datasync:Task":
                return new Task(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "datasync", _module)
