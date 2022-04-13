// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./destination";
export * from "./deviceProfile";
export * from "./fuotaTask";
export * from "./getDestination";
export * from "./getDeviceProfile";
export * from "./getFuotaTask";
export * from "./getMulticastGroup";
export * from "./getPartnerAccount";
export * from "./getServiceProfile";
export * from "./getTaskDefinition";
export * from "./getWirelessDevice";
export * from "./getWirelessGateway";
export * from "./multicastGroup";
export * from "./partnerAccount";
export * from "./serviceProfile";
export * from "./taskDefinition";
export * from "./wirelessDevice";
export * from "./wirelessGateway";

// Export enums:
export * from "../types/enums/iotwireless";

// Import resources to register:
import { Destination } from "./destination";
import { DeviceProfile } from "./deviceProfile";
import { FuotaTask } from "./fuotaTask";
import { MulticastGroup } from "./multicastGroup";
import { PartnerAccount } from "./partnerAccount";
import { ServiceProfile } from "./serviceProfile";
import { TaskDefinition } from "./taskDefinition";
import { WirelessDevice } from "./wirelessDevice";
import { WirelessGateway } from "./wirelessGateway";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:iotwireless:Destination":
                return new Destination(name, <any>undefined, { urn })
            case "aws-native:iotwireless:DeviceProfile":
                return new DeviceProfile(name, <any>undefined, { urn })
            case "aws-native:iotwireless:FuotaTask":
                return new FuotaTask(name, <any>undefined, { urn })
            case "aws-native:iotwireless:MulticastGroup":
                return new MulticastGroup(name, <any>undefined, { urn })
            case "aws-native:iotwireless:PartnerAccount":
                return new PartnerAccount(name, <any>undefined, { urn })
            case "aws-native:iotwireless:ServiceProfile":
                return new ServiceProfile(name, <any>undefined, { urn })
            case "aws-native:iotwireless:TaskDefinition":
                return new TaskDefinition(name, <any>undefined, { urn })
            case "aws-native:iotwireless:WirelessDevice":
                return new WirelessDevice(name, <any>undefined, { urn })
            case "aws-native:iotwireless:WirelessGateway":
                return new WirelessGateway(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "iotwireless", _module)
