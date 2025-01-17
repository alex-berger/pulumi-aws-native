// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS.Inputs
{

    public sealed class EndpointOracleSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("accessAlternateDirectly")]
        public Input<bool>? AccessAlternateDirectly { get; set; }

        [Input("addSupplementalLogging")]
        public Input<bool>? AddSupplementalLogging { get; set; }

        [Input("additionalArchivedLogDestId")]
        public Input<int>? AdditionalArchivedLogDestId { get; set; }

        [Input("allowSelectNestedTables")]
        public Input<bool>? AllowSelectNestedTables { get; set; }

        [Input("archivedLogDestId")]
        public Input<int>? ArchivedLogDestId { get; set; }

        [Input("archivedLogsOnly")]
        public Input<bool>? ArchivedLogsOnly { get; set; }

        [Input("asmPassword")]
        public Input<string>? AsmPassword { get; set; }

        [Input("asmServer")]
        public Input<string>? AsmServer { get; set; }

        [Input("asmUser")]
        public Input<string>? AsmUser { get; set; }

        [Input("charLengthSemantics")]
        public Input<string>? CharLengthSemantics { get; set; }

        [Input("directPathNoLog")]
        public Input<bool>? DirectPathNoLog { get; set; }

        [Input("directPathParallelLoad")]
        public Input<bool>? DirectPathParallelLoad { get; set; }

        [Input("enableHomogenousTablespace")]
        public Input<bool>? EnableHomogenousTablespace { get; set; }

        [Input("extraArchivedLogDestIds")]
        private InputList<int>? _extraArchivedLogDestIds;
        public InputList<int> ExtraArchivedLogDestIds
        {
            get => _extraArchivedLogDestIds ?? (_extraArchivedLogDestIds = new InputList<int>());
            set => _extraArchivedLogDestIds = value;
        }

        [Input("failTasksOnLobTruncation")]
        public Input<bool>? FailTasksOnLobTruncation { get; set; }

        [Input("numberDatatypeScale")]
        public Input<int>? NumberDatatypeScale { get; set; }

        [Input("oraclePathPrefix")]
        public Input<string>? OraclePathPrefix { get; set; }

        [Input("parallelAsmReadThreads")]
        public Input<int>? ParallelAsmReadThreads { get; set; }

        [Input("readAheadBlocks")]
        public Input<int>? ReadAheadBlocks { get; set; }

        [Input("readTableSpaceName")]
        public Input<bool>? ReadTableSpaceName { get; set; }

        [Input("replacePathPrefix")]
        public Input<bool>? ReplacePathPrefix { get; set; }

        [Input("retryInterval")]
        public Input<int>? RetryInterval { get; set; }

        [Input("secretsManagerAccessRoleArn")]
        public Input<string>? SecretsManagerAccessRoleArn { get; set; }

        [Input("secretsManagerOracleAsmAccessRoleArn")]
        public Input<string>? SecretsManagerOracleAsmAccessRoleArn { get; set; }

        [Input("secretsManagerOracleAsmSecretId")]
        public Input<string>? SecretsManagerOracleAsmSecretId { get; set; }

        [Input("secretsManagerSecretId")]
        public Input<string>? SecretsManagerSecretId { get; set; }

        [Input("securityDbEncryption")]
        public Input<string>? SecurityDbEncryption { get; set; }

        [Input("securityDbEncryptionName")]
        public Input<string>? SecurityDbEncryptionName { get; set; }

        [Input("spatialDataOptionToGeoJsonFunctionName")]
        public Input<string>? SpatialDataOptionToGeoJsonFunctionName { get; set; }

        [Input("standbyDelayTime")]
        public Input<int>? StandbyDelayTime { get; set; }

        [Input("useAlternateFolderForOnline")]
        public Input<bool>? UseAlternateFolderForOnline { get; set; }

        [Input("useBFile")]
        public Input<bool>? UseBFile { get; set; }

        [Input("useDirectPathFullLoad")]
        public Input<bool>? UseDirectPathFullLoad { get; set; }

        [Input("useLogminerReader")]
        public Input<bool>? UseLogminerReader { get; set; }

        [Input("usePathPrefix")]
        public Input<string>? UsePathPrefix { get; set; }

        public EndpointOracleSettingsArgs()
        {
        }
    }
}
