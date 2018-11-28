package api

/*
TYPE: Activity
Represents a user activity, such as a MapReduce job, a Hive query, an Oozie workflow, etc.
*/
type Activity struct {
	Name              string         `json:"name"`              // Activity name.
	Type              ActivityType   `json:"type"`              // Activity type. Whether it's an MR job, a Pig job, a Hive query, etc.
	Parent            string         `json:"parent"`            // The name of the parent activity.
	StartTime         string         `json:"startTime"`         // The start time of this activity.
	FinishTime        string         `json:"finishTime"`        // The finish time of this activity.
	Id                string         `json:"id"`                // Activity id, which is unique within a MapReduce service.
	Status            ActivityStatus `json:"status"`            // Activity status.
	User              string         `json:"user"`              // The user who submitted this activity.
	Group             string         `json:"group"`             // The user-group of this activity.
	InputDir          string         `json:"inputDir"`          // The input data directory of the activity. An HDFS url.
	OutputDir         string         `json:"outputDir"`         // The output result directory of the activity. An HDFS url.
	Mapper            string         `json:"mapper"`            // The mapper class.
	Combiner          string         `json:"combiner"`          // The combiner class.
	Reducer           string         `json:"reducer"`           // The reducer class.
	QueueName         string         `json:"queueName"`         // The scheduler queue this activity is in.
	SchedulerPriority string         `json:"schedulerPriority"` // The scheduler priority of this activity.
}

/*
TYPE: ActivityList
A generic list.
*/
type ActivityList struct {
	Items []*Activity `json:"items"` // Array of ApiActivity
}

/*
TYPE: ActivityType
*/
type ActivityType struct {
	Value string
}

/*
TYPE: ActivityStatus
*/
type ActivityStatus struct {
	Value string
}

/*
TYPE: Audit
Models audit events from both CM and CM managed services like HDFS, HBase and
Hive. Audits for CM managed services are retrieved from Cloudera Navigator
server.
*/
type Audit struct {
	Timestamp     string `json:"timestamp"`     // When the audit event was captured.
	Service       string `json:"service"`       // Service name associated with this audit.
	Username      string `json:"username"`      // The user who performed this operation.
	Impersonator  string `json:"impersonator"`  // The impersonating user (or the proxy user) who submitted this operation. This is usually applicable when using services like Oozie or Hue, who can be configured to impersonate other users and submit jobs.
	IpAddress     string `json:"ipAddress"`     // The IP address that the client connected from.
	Command       string `json:"command"`       // The command/operation that was requested.
	Resource      string `json:"resource"`      // The resource that the operation was performed on.
	OperationText string `json:"operationText"` // The full text of the requested operation. E.g. the full Hive query.  Available since API v5.
	Allowed       bool   `json:"allowed"`       // Whether the operation was allowed or denied by the authorization system.
}

/*
TYPE: AuditList
A generic list.
*/
type AuditList struct {
	Items []*Audit `json:"items"` // array of ApiAudit
}

/*
TYPE: AuthRole
This is the model for user role scope in the API since v18. This is used to support granular permissions.
*/
type AuthRole struct {
	DisplayName          string                    `json:"displayName"`          //
	Clusters             []*ClusterRef             `json:"clusters"`             // array of ApiClusterRef
	Users                []*User2Ref               `json:"users"`                // array of ApiUser2Ref
	ExternalUserMappings []*ExternalUserMappingRef `json:"externalUserMappings"` // array of ApiExternalUserMappingRef
	BaseRole             *AuthRoleRef              `json:"baseRole"`             // A role this user possesses. In Cloudera Enterprise Datahub Edition, possible values are: ROLE_ADMIN ROLE_USER ROLE_LIMITED: Added in Cloudera Manager 5.0 ROLE_OPERATOR: Added in Cloudera Manager 5.1 ROLE_CONFIGURATOR: Added in Cloudera Manager 5.1 ROLE_CLUSTER_ADMIN: Added in Cloudera Manager 5.2 ROLE_BDR_ADMIN: Added in Cloudera Manager 5.2 ROLE_NAVIGATOR_ADMIN: Added in Cloudera Manager 5.2 ROLE_USER_ADMIN: Added in Cloudera Manager 5.2 ROLE_KEY_ADMIN: Added in Cloudera Manager 5.5 An empty role implies ROLE_USER.
	Uuid                 string                    `json:"uuid"`                 // Readonly. The UUID of the authRole.
	IsCustom             bool                      `json:"isCustom"`             //
}

/*
TYPE: AuthRoleAuthority
This represents an authority with a name and description.
*/
type AuthRoleAuthority struct {
	Name        string `json:"name"`        // The name of the authority.
	Description string `json:"description"` // The description of the authority.
}

/*
TYPE: AuthRoleList
A list of auth roles.
*/
type AuthRoleList struct {
	Items []*AuthRole `json:"items"` // array of ApiAuthRole
}

/*
TYPE: AuthRoleMetadata
This is the model for auth role metadata
*/
type AuthRoleMetadata struct {
	DisplayName   string               `json:"displayName"`   //
	Uuid          string               `json:"uuid"`          //
	Role          string               `json:"role"`          //
	Authorities   []*AuthRoleAuthority `json:"authorities"`   // array of ApiAuthRoleAuthority
	AllowedScopes []string             `json:"allowedScopes"` // array of string
}

/*
TYPE: AuthRoleMetadataList
A list of auth roles metadata.
*/
type AuthRoleMetadataList struct {
	Items []*AuthRoleMetadata `json:"items"` // array of ApiAuthRoleMetadata

}

/*
TYPE: AuthRoleRef
An authRoleRef to operate on ApiAuthRole object
*/
type AuthRoleRef struct {
	DisplayName string `json:"displayName"` // The name of the authRole.
	Uuid        string `json:"uuid"`        // The uuid of the authRole, which uniquely identifies it in a CM installation.
}

/*
TYPE: BatchRequest
A batch request, comprised of one or more request elements.
*/
type BatchRequest struct {
	Items []*BatchRequestElement `json:"items"` // array of ApiBatchRequestElement
}

/*
TYPE: BatchRequestElement
A single element of a batch request, often part of a list with other elements.
*/
type BatchRequestElement struct {
	Method      *HTTPMethod `json:"method"`      // The type of request (e.g. POST, GET, etc.).
	Url         string      `json:"url"`         // The URL of the request. Must not have a scheme, host, or port. The path should be prefixed with "/api/", and should include path and query parameters.
	Body        string      `json:"body"`        // Optional body of the request. Must be serialized in accordance with #getContentType(). For application/json, use com.cloudera.api.ApiObjectMapper.
	ContentType string      `json:"contentType"` // Content-Type header of the request element. If unset, the element will be treated as if the wildcard type had been specified unless it has a body, in which case it will fall back to application/json.
	AcceptType  string      `json:"acceptType"`  // Accept header of the request element. The response body (if it exists) will be in this representation. If unset, the element will be treated as if the wildcard type had been requested.
}

/*
TYPE: BatchResponse
A batch response, comprised of one or more response elements.
*/
type BatchResponse struct {
	Items   []*BatchResponseElement `json:"items"`   // array of ApiBatchResponseElement
	Success bool                    `json:"success"` // Read-only. True if every response element's ApiBatchResponseElement#getStatusCode() is in the range [200, 300), false otherwise.
}

/*
TYPE: BatchResponseElement
A single element of a batch response, often part of a list with other elements.
*/
type BatchResponseElement struct {
	StatusCode int8   `json:"statusCode"` // Read-only. The HTTP status code of the response.
	Response   string `json:"response"`   // Read-only. The (optional) serialized body of the response, in the representation produced by the corresponding API endpoint, such as application/json.
}

/*
TYPE: BulkCommandList
A list of commands.  This list is returned whenever commands are issued in bulk, and contains a second list with information about errors issuing specific commands.
*/
type BulkCommandList struct {
	Errors []string   `json:"errors"` // Array of string Errors that occurred when issuing individual commands.
	Items  []*Command `json:"items"`  // Array of ApiCommand
}

/*
TYPE: CdhUpgradeArgs
Arguments used for the CDH Upgrade command.
*/
type CdhUpgradeArgs struct {
	CdhParcelVersion   string                     `json:"cdhParcelVersion"`   // If using parcels, the full version of an already distributed parcel for the next major CDH version. Default is null, which indicates this is a package upgrade. Example versions are: '5.0.0-1.cdh5.0.0.p0.11' or '5.0.2-1.cdh5.0.2.p0.32'
	CdhPackageVersion  string                     `json:"cdhPackageVersion"`  // If using packages, the full version of the CDH packages being upgraded to, such as "5.1.2". These packages must already be installed on the cluster before running the upgrade command. For backwards compatibility, if "5.0.0" is specified here, then the upgrade command will relax validation of installed packages to match v6 behavior, only checking major version.  Introduced in v9. Has no effect in older API versions, which assume "5.0.0"
	RollingRestartArgs *RollingUpgradeClusterArgs `json:"rollingRestartArgs"` // If provided and rolling restart is available, will perform rolling restart with the requested arguments. If provided and rolling restart is not available, errors. If omitted, will do a regular restart.  Introduced in v9. Has no effect in older API versions, which must always do a hard restart.
}

/*
TYPE: Cluster
A cluster represents a set of interdependent services running on a set of hosts. All services on a given cluster are of the same software version (e.g. CDH4 or CDH5).
*/
type Cluster struct {
	Name              string        `json:"name"`              // The name of the cluster.  Immutable since API v6.  Prior to API v6, will contain the display name of the cluster.
	DisplayName       string        `json:"displayName"`       // The display name of the cluster that is shown in the UI.  Available since API v6.
	FullVersion       string        `json:"fullVersion"`       // The full CDH version of the cluster. The expected format is three dot separated version ints, e.g. "4.2.1" or "5.0.0". The full version takes precedence over the version field during cluster creation.  Available since API v6.
	MaintenanceMode   bool          `json:"maintenanceMode"`   // Readonly. Whether the cluster is in maintenance mode. Available since API v2.
	MaintenanceOwners []*EntityType `json:"maintenanceOwners"` // EntityType array of ApiEntityType Readonly. The list of objects that trigger this cluster to be in maintenance mode. Available since API v2.
	Services          []*Service    `json:"services"`          // Service array of ApiService Optional. Used during import/export of settings.
	Parcels           []*Parcel     `json:"parcels"`           // Parcel array of ApiParcel Optional. Used during import/export of settings. Available since API v4.
	ClusterUrl        string        `json:"clusterUrl"`        // Readonly. Link into the Cloudera Manager web UI for this specific cluster.  Available since API v10.
	HostsUrl          string        `json:"hostsUrl"`          // Readonly. Link into the Cloudera Manager web UI for host table for this cluster.  Available since API v11.
	EntityStatus      *EntityStatus `json:"entityStatus"`      // Readonly. The entity status for this cluster. Available since API v11.
	Uuid              string        `json:"uuid"`              // Readonly. The UUID of the cluster.  Available since API v15.
}

/*
TYPE: ClusterList
A list of clusters.
*/
type ClusterList struct {
	Items []*Cluster `json:"items"` // array of ApiCluster
}

/*
TYPE: ClusterRef
A clusterRef references a cluster. To operate on the cluster object, use the cluster API with the clusterName as the parameter.
*/
type ClusterRef struct {
	ClusterName string `json:"clusterName"` // The name of the cluster, which uniquely identifies it in a CM installation.
	DisplayName string `json:"displayName"` // The display name of the cluster. This is available from v30.
}

/*
TYPE: ClusterTemplate
Details of cluster template
*/
type ClusterTemplate struct {
	cdhVersion    string
	Products      []*ProductVersion              `json:"products"`      // array of ApiProductVersion
	Services      []*ClusterTemplateService      `json:"services"`      // array of ApiClusterTemplateService
	HostTemplates []*ClusterTemplateHostTemplate `json:"hostTemplates"` // array of ApiClusterTemplateHostTemplate
	DisplayName   string                         `json:"displayName"`   //
	CmVersion     string                         `json:"cmVersion"`     //
	Instantiator  *ClusterTemplateInstantiator   `json:"instantiator"`  // ApiClusterTemplateInstantiator
	Repositories  []string                       `json:"repositories"`  // array of string
}

/*
TYPE: ClusterTemplateConfig
Config Details: The config can either have a value or ref or variable.
*/
type ClusterTemplateConfig struct {
	Name       string `json:"name"`       //
	Value      string `json:"value"`      //
	Ref        string `json:"ref"`        //
	Variable   string `json:"variable"`   //
	AutoConfig bool   `json:"autoConfig"` //
}

/*
TYPE: ClusterTemplateHostInfo
This contains information about the host or host range on which provided host template will be applied.
*/
type ClusterTemplateHostInfo struct {
	HostName            string   `json:"hostName"`            //
	HostNameRange       string   `json:"hostNameRange"`       //
	RackId              string   `json:"rackId"`              //
	HostTemplateRefName string   `json:"hostTemplateRefName"` //
	RoleRefNames        []string `json:"roleRefNames"`        // of string
}

/*
TYPE: ClusterTemplateHostTemplate
Host templates will contain information about the role config groups that should be applied to a host. This basically means a host will have a role corresponding to each config group.
*/
type ClusterTemplateHostTemplate struct {
	RefName                  string   `json:"refName"`                  //
	RoleConfigGroupsRefNames []string `json:"roleConfigGroupsRefNames"` // of string
	Cardinality              int8     `json:"cardinality"`              //
}

/*
TYPE: ClusterTemplateInstantiator
Details of cluster template
*/
type ClusterTemplateInstantiator struct {
	ClusterName      string                                `json:"clusterName"`      //
	Hosts            []*ClusterTemplateHostInfo            `json:"hosts"`            // array of ApiClusterTemplateHostInfo
	Variables        []*ClusterTemplateVariable            `json:"variables"`        // array of ApiClusterTemplateVariable
	RoleConfigGroups []*ClusterTemplateRoleConfigGroupInfo `json:"roleConfigGroups"` // array of ApiClusterTemplateRoleConfigGroupInfo
}

/*
TYPE: ClusterTemplateRole
Role info: This will contain information related to a role referred by some configuration. During import type this role must be materizalized.
*/
type ClusterTemplateRole struct {
	RefName  string `json:"refName"`  //
	RoleType string `json:"roleType"` //
}

/*
TYPE: ClusterTemplateRoleConfigGroup
Role config group info.
*/
type ClusterTemplateRoleConfigGroup struct {
	RefName     string                   `json:"refName"`     //
	RoleType    string                   `json:"roleType"`    //
	Base        bool                     `json:"base"`        //
	DisplayName string                   `json:"displayName"` //
	Configs     []*ClusterTemplateConfig `json:"configs"`     // array of ApiClusterTemplateConfig
}

/*
TYPE: ClusterTemplateRoleConfigGroupInfo
During import time information related to all the non-base config groups must be provided.
*/
type ClusterTemplateRoleConfigGroupInfo struct {
	RcgRefName string `json:"rcgRefName"` //
	Name       string `json:"name"`       //
}

/*
TYPE: ClusterTemplateService
Service information
*/
type ClusterTemplateService struct {
	RefName          string                            `json:"refName"`          //
	ServiceType      string                            `json:"serviceType"`      //
	ServiceConfigs   []*ClusterTemplateConfig          `json:"serviceConfigs"`   // of ApiClusterTemplateConfig
	RoleConfigGroups []*ClusterTemplateRoleConfigGroup `json:"roleConfigGroups"` // of ApiClusterTemplateRoleConfigGroup
	Roles            []*ClusterTemplateRole            `json:"roles"`            // of ApiClusterTemplateRole
	DisplayName      string                            `json:"displayName"`      //
}

/*
TYPE: ClusterTemplateVariable
Variable that is referred in cluster template.
*/
type ClusterTemplateVariable struct {
	Name  string `json:"name"`  //
	Value string `json:"value"` //
}

/*
TYPE: ClusterUtilization
Utilization report information of a Cluster.
*/
type ClusterUtilization struct {
	TotalCpuCores                   int8                   `json:"totalCpuCores"`                   // Average int8 of CPU cores available in the cluster during the report window.
	AvgCpuUtilization               int8                   `json:"avgCpuUtilization"`               // Average CPU consumption for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster.
	MaxCpuUtilization               int8                   `json:"maxCpuUtilization"`               // Maximum CPU consumption for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster.
	AvgCpuDailyPeak                 int8                   `json:"avgCpuDailyPeak"`                 // Average daily peak CPU consumption for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster.
	AvgWorkloadCpu                  int8                   `json:"avgWorkloadCpu"`                  // Average CPU consumption by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala.
	MaxWorkloadCpu                  int8                   `json:"maxWorkloadCpu"`                  // Maximum CPU consumption by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala.
	AvgWorkloadCpuDailyPeak         int8                   `json:"avgWorkloadCpuDailyPeak"`         // Average daily peak CPU consumption by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala.
	TotalMemory                     int8                   `json:"totalMemory"`                     // Average physical memory (in bytes) available in the cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster.
	AvgMemoryUtilization            int8                   `json:"avgMemoryUtilization"`            // Average memory consumption (as percentage of total memory) for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster.
	MaxMemoryUtilization            int8                   `json:"maxMemoryUtilization"`            // Maximum memory consumption (as percentage of total memory) for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster.
	AvgMemoryDailyPeak              int8                   `json:"avgMemoryDailyPeak"`              // Average daily peak memory consumption (as percentage of total memory) for the entire cluster during the report window. This includes consumption by user workloads in YARN and Impala, as well as consumption by all services running in the cluster.
	AvgWorkloadMemory               int8                   `json:"avgWorkloadMemory"`               // Average memory consumption (as percentage of total memory) by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala.
	MaxWorkloadMemory               int8                   `json:"maxWorkloadMemory"`               // Maximum memory consumption (as percentage of total memory) by workloads that ran on the cluster. This includes consumption by user workloads in YARN and Impala
	AvgWorkloadMemoryDailyPeak      int8                   `json:"avgWorkloadMemoryDailyPeak"`      // Average daily peak memory consumption (as percentage of total memory) by workloads that ran on the cluster during the report window. This includes consumption by user workloads in YARN and Impala.
	TenantUtilizations              *TenantUtilizationList `json:"tenantUtilizations"`              // A list of tenant utilization reports.
	MaxCpuUtilizationTimestampMs    int8                   `json:"maxCpuUtilizationTimestampMs"`    // Timestamp corresponding to maximum CPU utilization for the entire cluster during the report window.
	MaxMemoryUtilizationTimestampMs int8                   `json:"maxMemoryUtilizationTimestampMs"` // Timestamp corresponding to maximum memory utilization for the entire cluster during the report window.
	MaxWorkloadCpuTimestampMs       int8                   `json:"maxWorkloadCpuTimestampMs"`       // Timestamp corresponds to maximum CPU consumption by workloads that ran on the cluster during the report window.
	MaxWorkloadMemoryTimestampMs    int8                   `json:"maxWorkloadMemoryTimestampMs"`    // Timestamp corresponds to maximum memory resource consumption by workloads that ran on the cluster during the report window.
	ErrorMessage                    string                 `json:"errorMessage"`                    // Error message while generating utilization report.
}

/*
TYPE: CmPeer
Information about a Cloudera Manager peer instance.  The requirement and usage of username and password properties are dependent on the clouderaManagerCreatedUser flag.  When creating peers, if 'clouderaManagerCreatedUser' is true, the username/password should be the credentials of a user with administrator privileges on the remote Cloudera Manager. These credentials are not stored, they are used to connect to the peer and create a user in that peer. The newly created user is stored and used for communication with that peer. If 'clouderaManagerCreatedUser' is false, which is not applicable to REPLICATION peer type, the username/password to the remote Cloudera Manager are directly stored and used for all communications with that peer.  When updating peers, if 'clouderaManagerCreatedUser' is true and username/password are set, a new remote user will be created. If 'clouderaManagerCreatedUser' is false and username/password are set, the stored username/password will be updated.
*/
type CmPeer struct {
	Name                       string      `json:"name"`                       // The name of the remote CM instance. Immutable during update.
	Type                       *CmPeerType `json:"type"`                       // The type of the remote CM instance. Immutable during update. Available since API v11.
	Url                        string      `json:"url"`                        // The URL of the remote CM instance. Mutable during update.
	Username                   string      `json:"username"`                   // When creating peers, if 'clouderaManagerCreatedUser' is true, this should be the remote admin username for creating a user in remote Cloudera Manager. The created remote user will then be stored in the local Cloudera Manager DB and used in later communication. If 'clouderaManagerCreatedUser' is false, which is not applicable to REPLICATION peer type, Cloudera Manager will store this username in the local DB directly and use it together with 'password' for communication. Mutable during update. When set during update, if 'clouderaManagerCreatedUser' is true, a new user in remote Cloudera Manager is created, the newly created remote user will be stored in the local DB. An attempt to delete the previously created remote user will be made; If 'clouderaManagerCreatedUser' is false, the username/password in the local DB will be updated.
	Password                   string      `json:"password"`                   // When creating peers, if 'clouderaManagerCreatedUser' is true, this should be the remote admin password for creating a user in remote Cloudera Manager. The created remote user will then be stored in the local Cloudera Manager DB and used in later communication. If 'clouderaManagerCreatedUser' is false, which is not applicable to REPLICATION peer type, Cloudera Manager will store this password in the local DB directly and use it together with 'username' for communication. Mutable during update. When set during update, if 'clouderaManagerCreatedUser' is true, a new user in remote Cloudera Manager is created, the newly created remote user will be stored in the local DB. An attempt to delete the previously created remote user will be made; If 'clouderaManagerCreatedUser' is false, the username/password in the local DB will be updated.
	ClouderaManagerCreatedUser bool        `json:"clouderaManagerCreatedUser"` // If true, Cloudera Manager creates a remote user using the given username/password and stores the created user in local DB for use in later communication. Cloudera Manager will also try to delete the created remote user when deleting such peers. If false, Cloudera Manager will store the provided username/password in the local DB and use them in later communication. 'false' value on this field is not applicable to REPLICATION peer type. Available since API v11. Immutable during update. Should not be set when updating peers.
}

/*
TYPE: CmPeerList
A list of Cloudera Manager peers.
*/
type CmPeerList struct {
	Items []*CmPeer `json:"items"` // array of ApiCmPeer
}

/*
TYPE: CmPeerType
Enum for CM peer types.
*/
type CmPeerType struct {
	// ENUM
}

/*
TYPE: CollectDiagnosticDataArguments
Arguments used for the collectDiagnosticData command.
*/
type CollectDiagnosticDataArguments struct {
	BundleSizeBytes                int8     `json:"bundleSizeBytes"`                // The maximum approximate bundle size of the output file
	StartTime                      string   `json:"startTime"`                      // This parameter is ignored between CM 4.5 and CM 5.7 versions. For versions from CM 4.5 to CM 5.7, use endTime and bundleSizeBytes instead. For CM 5.7+ versions, startTime is an optional parameter that is with endTime and bundleSizeBytes. This was introduced to perform diagnostic data estimation and collection of global diagnostics data for a certain time range. The start time (in ISO 8601 format) of the period to collection statistics for.
	EndTime                        string   `json:"endTime"`                        // The end time (in ISO 8601 format) of the period to collection statistics for.
	IncludeInfoLog                 bool     `json:"includeInfoLog"`                 // This parameter is ignored as of CM 4.5. INFO logs are always collected. Whether to include INFO level logs. WARN, ERROR, and FATAL level logs are always included.
	Ticketint                      string   `json:"ticketint"`                      // The support ticket int8 to attach to this data collection.
	Comments                       string   `json:"comments"`                       // Comments to include with this data collection.
	ClusterName                    string   `json:"clusterName"`                    // Name of the cluster to collect. If null, collects from all clusters.
	EnableMonitorMetricsCollection bool     `json:"enableMonitorMetricsCollection"` // Flag to enable collection of metrics for chart display.
	Roles                          []string `json:"roles"`                          // array of string List of roles for which to get logs and metrics. If set, this restricts the roles for log and metrics collection to the list specified. If empty, the default is to get logs for all roles (in the selected cluster, if one is selected). Introduced in API v10 of the API.
}

/*
TYPE: Command
Provides detailed information about a submitted command.  There are two types of commands: synchronous and asynchronous. Synchronous commands complete immediately, and their results are passed back in the returned command object after the execution of an API call. Outside of that returned object, there is no way to check the result of a synchronous command.  Asynchronous commands have unique non-negative IDs. They may still be running when the API call returns. Clients can check the status of such commands using the API.
*/
type Command struct {
	Id            int8         `json:"id"`            // The command ID.
	Name          string       `json:"name"`          // The command name.
	StartTime     string       `json:"startTime"`     // The start time.
	EndTime       string       `json:"endTime"`       // The end time, if the command is finished.
	Active        bool         `json:"active"`        // Whether the command is currently active.
	Success       bool         `json:"success"`       // If the command is finished, whether it was successful.
	ResultMessage string       `json:"resultMessage"` // If the command is finished, the result message.
	ResultDataUrl string       `json:"resultDataUrl"` // URL to the command's downloadable result data, if any exists.
	ClusterRef    *ClusterRef  `json:"clusterRef"`    // Reference to the cluster (for cluster commands only).
	ServiceRef    *ServiceRef  `json:"serviceRef"`    // Reference to the service (for service commands only).
	RoleRef       *RoleRef     `json:"roleRef"`       // Reference to the role (for role commands only).
	HostRef       *HostRef     `json:"hostRef"`       // Reference to the host (for host commands only).
	Parent        *Command     `json:"parent"`        // Reference to the parent command, if any.
	Children      *CommandList `json:"children"`      // List of child commands. Only available in the full view.  The list contains only the summary view of the children.
	CanRetry      bool         `json:"canRetry"`      // If the command can be retried. Available since V11
}

/*
TYPE: CommandList
A list of commands.
*/
type CommandList struct {
	Items []*Command `json:"items"` //Command  array of ApiCommand
}

/*
TYPE: CommandMetadata
Provides metadata information about a command.
*/
type CommandMetadata struct {
	Name      string `json:"name"`      // The name of of the command.
	ArgSchema string `json:"argSchema"` // The command arguments schema. This is in the form of json schema and describes the structure of the command arguments. If null, the command does not take arguments.
}

/*
TYPE: CommandMetadataList
A list of command metadata.
*/
type CommandMetadataList struct {
	Items []*CommandMetadata `json:"items"` //CommandMetadata  array of ApiCommandMetadata The list of command metadata objects.
}

/*
TYPE: CommissionState
Represents the Commission state of an entity.
*/
type CommissionState struct {
	// ENUM
}

/*
TYPE: Config
Model for a configuration parameter. When an entry's value property is not available, it means the entry is not configured. This means that the default value for the entry, if any, will be used. Setting a value to null also can be used to unset any previously set value for the parameter, reverting to the default value (if any).
*/
type Config struct {
	Name                         string           `json:"name"`                         // Readonly. The canonical name that identifies this configuration parameter.
	Value                        string           `json:"value"`                        // The user-defined value. When absent, the default value (if any) will be used. Can also be absent, when enumerating allowed configs.
	Required                     bool             `json:"required"`                     // Readonly. Requires "full" view. Whether this configuration is required for the object. If any required configuration is not set, operations on the object may not work.
	Default                      string           `json:"default"`                      // Readonly. Requires "full" view. The default value.
	DisplayName                  string           `json:"displayName"`                  // Readonly. Requires "full" view. A user-friendly name of the parameters, as would have been shown in the web UI.
	Description                  string           `json:"description"`                  // Readonly. Requires "full" view. A textual description of the parameter.
	RelatedName                  string           `json:"relatedName"`                  // Readonly. Requires "full" view. If applicable, contains the related configuration variable used by the source project.
	Sensitive                    bool             `json:"sensitive"`                    // Readonly. Whether this configuration is sensitive, i.e. contains information such as passwords, which might affect how the value of this configuration might be shared by the caller. Available since v14.
	ValidationState              *ValidationState `json:"validationState"`              // Readonly. Requires "full" view. State of the configuration parameter after validation.
	ValidationMessage            string           `json:"validationMessage"`            // Readonly. Requires "full" view. A message explaining the parameter's validation state.
	ValidationWarningsSuppressed bool             `json:"validationWarningsSuppressed"` // Readonly. Requires "full" view. Whether validation warnings associated with this parameter are suppressed. In general, suppressed validation warnings are hidden in the Cloudera Manager UI. Configurations that do not produce warnings will not contain this field.
}

/*
TYPE: ConfigList
A list of configuration data.
*/
type ConfigList struct {
	Items []*Config `json:"items"` // array of ApiConfig
}

/*
TYPE: ConfigStalenessStatus
Represents the configuration staleness status of an entity.
*/
type ConfigStalenessStatus struct {
	// ENUM
}

/*
TYPE: ConfigureForKerberosArguments
Arguments used to configure a cluster for Kerberos.
*/
type ConfigureForKerberosArguments struct {
	DatanodeTransceiverPort int8 `json:"datanodeTransceiverPort"` // The HDFS DataNode transceiver port to use. This will be applied to all DataNode role configuration groups. If not specified, this will default to 1004.
	DatanodeWebPort         int8 `json:"datanodeWebPort"`         // The HDFS DataNode web port to use. This will be applied to all DataNode role configuration groups. If not specified, this will default to 1006.
}

/*
TYPE: Dashboard
A dashboard definition. Dashboards are composed of tsquery-based charts.
*/
type Dashboard struct {
	Name string `json:"name"` // Returns the dashboard name.
	Json string `json:"json"` // Returns the json structure for the dashboard. This should be treated as an opaque blob.
}

/*
TYPE: DashboardList
A list of dashboard definitions.
*/
type DashboardList struct {
	Items []*Dashboard `json:"items"` // array of ApiDashboard
}

/*
TYPE: Deployment
This objects represents a deployment including all clusters, hosts, services, roles, etc in the system. It can be used to save and restore all settings.
*/
type Deployment struct {
	Timestamp         string            `json:"timestamp"`         // Readonly. This timestamp is provided when you request a deployment and is not required (or even read) when creating a deployment. This timestamp is useful if you have multiple deployments saved and want to determine which one to use as a restore point.
	Clusters          []*Cluster        `json:"clusters"`          // array of ApiCluster List of clusters in the system including their services, roles and complete config values.
	Hosts             []*Host           `json:"hosts"`             // array of ApiHost List of hosts in the system
	Users             []*User           `json:"users"`             // array of ApiUser List of all users in the system
	VersionInfo       *VersionInfo      `json:"versionInfo"`       // Full version information about the running Cloudera Manager instance
	ManagementService *Service          `json:"managementService"` // The full configuration of the Cloudera Manager management service including all the management roles and their config values
	ManagerSettings   *ConfigList       `json:"managerSettings"`   // The full configuration of Cloudera Manager itself including licensing info
	AllHostsConfig    *ConfigList       `json:"allHostsConfig"`    // Configuration parameters that apply to all hosts, unless overridden at the host level. Available since API v3.
	Peers             []*CmPeer         `json:"peers"`             // of ApiCmPeer The list of peers configured in Cloudera Manager. Available since API v3.
	HostTemplates     *HostTemplateList `json:"hostTemplates"`     // The list of all host templates in Cloudera Manager.
}

/*
TYPE: Deployment2
This objects represents a deployment including all clusters, hosts, services, roles, etc in the system. It can be used to save and restore all settings. This model will be used v18 and beyond since users will be represented by ApiUser2 v18 and beyond.
*/
type Deployment2 struct {
	Timestamp         string           `json:"timestamp"`         // Readonly. This timestamp is provided when you request a deployment and is not required (or even read) when creating a deployment. This timestamp is useful if you have multiple deployments saved and want to determine which one to use as a restore point.
	Clusters          []*Cluster       `json:"clusters"`          // array of ApiCluster List of clusters in the system including their services, roles and complete config values.
	Hosts             []*Host          `json:"hosts"`             // array of ApiHost List of hosts in the system
	Users             []*User2         `json:"users"`             // array of ApiUser2 List of all users in the system
	VersionInfo       VersionInfo      `json:"versionInfo"`       // Full version information about the running Cloudera Manager instance
	ManagementService Service          `json:"managementService"` // The full configuration of the Cloudera Manager management service including all the management roles and their config values
	ManagerSettings   ConfigList       `json:"managerSettings"`   // The full configuration of Cloudera Manager itself including licensing info
	AllHostsConfig    ConfigList       `json:"allHostsConfig"`    // Configuration parameters that apply to all hosts, unless overridden at the host level. Available since API v3.
	Peers             []*CmPeer        `json:"peers"`             // array of ApiCmPeer The list of peers configured in Cloudera Manager. Available since API v3.
	HostTemplates     HostTemplateList `json:"hostTemplates"`     // The list of all host templates in Cloudera Manager.
}

/*
TYPE: DisableJtHaArguments
Arguments used for disable JT HA command.
*/
type DisableJtHaArguments struct {
	ActiveName string `json:"activeName"` // Name of the JobTracker that will be active after HA is disabled.
}

/*
TYPE: DisableLlamaHaArguments
Arguments used for disable Llama HA command.
*/
type DisableLlamaHaArguments struct {
	ActiveName string `json:"activeName"` // Name of the Llama role that will be active after HA is disabled.
}

/*
TYPE: DisableNnHaArguments
Arguments used for Disable NameNode High Availability command.
*/
type DisableNnHaArguments struct {
	ActiveNnName         string   `json:"activeNnName"`         // Name of the NamdeNode role that is going to be active after High Availability is disabled.
	SnnHostId            string   `json:"snnHostId"`            // Id of the host where the new SecondaryNameNode will be created.
	SnnCheckpointDirList []string `json:"snnCheckpointDirList"` // of string List of directories used for checkpointing by the new SecondaryNameNode.
	SnnName              string   `json:"snnName"`              // Name of the new SecondaryNameNode role (Optional).
}

/*
TYPE: DisableOozieHaArguments
ApiDisableRmHaArguments Arguments used for Disable RM HA command.
*/
type DisableOozieHaArguments struct {
	ActiveName string `json:"activeName"` // Name of the Oozie Server that will be active after HA is disabled.
}

/*
TYPE: DisableSentryHaArgs
Arguments used for disable Sentry HA API call.
*/
type DisableSentryHaArgs struct {
	ActiveName string `json:"activeName"` // Name of the single role that will remain active after HA is disabled.
}

/*
TYPE: Echo
The echoMessage carries a message to be echoed back from the API service.
*/
type Echo struct {
}

/*
TYPE: EnableJtHaArguments
Arguments used for enable JT HA command.
*/
type EnableJtHaArguments struct {
}

/*
TYPE: EnableLlamaHaArguments
Arguments used for enable Llama HA command.
*/
type EnableLlamaHaArguments struct {
}

/*
TYPE: EnableLlamaRmArguments
Arguments used for enable Llama RM command.
*/
type EnableLlamaRmArguments struct {
}

/*
TYPE: EnableNnHaArguments
Arguments used for Enable NameNode High Availability command.
*/
type EnableNnHaArguments struct {
}

/*
TYPE: EnableOozieHaArguments

ApiEnableRmHaArguments Arguments used for enable RM HA command.
*/
type EnableOozieHaArguments struct {
	message string
}

/*
TYPE: EnableSentryHaArgs
Arguments used for enable Sentry HA command.
*/
type EnableSentryHaArgs struct {
	NewSentryHostId   string                           `json:"newSentryHostId"`   //	Id of host on which new Sentry Server role will be added.
	NewSentryRoleName string                           `json:"newSentryRoleName"` //	Name of the new Sentry Server role to be created. This is an optional argument.
	ZkServiceName     string                           `json:"zkServiceName"`     //	Name of the ZooKeeper service that will be used for Sentry HA. This is an optional parameter if the Sentry to ZooKeeper dependency is already set in CM.
	RrcArgs           *SimpleRollingRestartClusterArgs `json:"rrcArgs"`           //
}

/*
TYPE: EntityStatus
The single value used by the Cloudera Manager UI to represent the status of the entity. It is computed from a variety of other entity-specific states, not all values apply to all entities. For example, STARTING/STOPPING do not apply to a host.
*/
type EntityStatus struct {
	// ENUM
}

/*
TYPE: EntityType
Represents the types of entities.
*/
type EntityType struct {
	//
}

/*
TYPE: Event
Events model noteworthy incidents in Cloudera Manager or the managed Hadoop cluster. An event carries its event category, severity, and a string content. They also have generic attributes, which are free-form key value pairs. Important events may be promoted into alerts.
*/
type Event struct {
	Id           string            `json:"id"`           //	A unique ID for this event.
	Content      string            `json:"content"`      //	The content payload of this event.
	TimeOccurred string            `json:"timeOccurred"` //	When the event was generated.
	TimeReceived string            `json:"timeReceived"` //	When the event was stored by Cloudera Manager. Events do not arrive in the order that they are generated. If you are writing an event poller, this is a useful field to query.
	Category     string            `json:"category"`     //	The category of this event -- whether it is a health event, an audit event, an activity event, etc.
	Severity     *EventSeverity    `json:"severity"`     //	The severity of the event.
	Alert        bool              `json:"alert"`        //	Whether the event is promoted to an alert according to configuration.
	Attributes   []*EventAttribute `json:"attributes"`   //	A list of key-value attribute pairs.
}

/*
TYPE: EventAttribute

EventCategory
*/
type EventAttribute struct {
	Name   string    `json:"name"`   //
	Values []*string `json:"values"` //string	array of string
}

/*
TYPE: EventQueryResult
A generic list.
*/
type EventQueryResult struct {
	TotalResults int32    `json:"totalResults"` //	The total number of matched results. Some are possibly not shown due to pagination.
	Items        []*Event `json:"items"`        // //
}

/*
TYPE: EventSeverity

ExternalAccount Represents an instantiation of an external account type, referencing a supported external account type, via the typeName field, along with suitable configuration to access an external resource of the provided type. The typeName field must match the name of an external account type.
*/
type EventSeverity struct {
	// ENUM
}

/*
TYPE: ExternalAccountCategory
Type representing an external account category.
*/
type ExternalAccountCategory struct {
	Name             string      `json:"name"`             //	Represents the intial name of the account; used to uniquely identify this account.
	DisplayName      string      `json:"displayName"`      //	Represents a modifiable label to identify this account for user-visible purposes.
	CreatedTime      string      `json:"createdTime"`      //	Represents the time of creation for this account.
	LastModifiedTime string      `json:"lastModifiedTime"` //	Represents the last modification time for this account.
	TypeName         string      `json:"typeName"`         //	Represents the Type ID of a supported external account type. The type represented by this field dictates which configuration options must be defined for this account.
	AccountConfigs   *ConfigList `json:"accountConfigs"`   //	Represents the account configuration for this account. When an account is retrieved from the server, the configs returned must match allowed configuration for the type of this account. When specified for creation of a new account or for the update of an existing account, this field must include every required configuration parameter specified in the type's definition, with the account configuration's value field specified to represent the specific configuration desired for this account.
}

/*
TYPE: ExternalAccountCategoryList
Represents a list of external account categories.
*/
type ExternalAccountCategoryList struct {
	Name        string `json:"name"`        //	Represents an identifier for a category.
	DisplayName string `json:"displayName"` //	Represents a localized display name for a category.
	Description string `json:"description"` //	Represents a localized description for a category.
}

/*
TYPE: ExternalAccountList
Represents a list of external accounts.
*/
type ExternalAccountList struct {
	Items []*ExternalAccount `json:"items"` //
}

type ExternalAccount struct {}
/*
TYPE: ExternalAccountType
A supported external account type. An external account type represents an external authentication source that is used by Cloudera Manager in its APIs to take suitable actions that require authentication to an external service. An external account type is uniquely identified by a server-generated ID and identifies with a category identifier: e.g. The "AWS" category has an account type "AWS_Access_Key_Authorization"
*/
type ExternalAccountType struct {
	Name                  string      `json:"name"`                  //	Represents the immutable name for this account.
	CategoryName          string      `json:"categoryName"`          //	Represents the category of this account.
	Type                  string      `json:"type"`                  //	Represents the type for this account.
	DisplayName           string      `json:"displayName"`           //	Represents the localized display name for this account.
	Description           string      `json:"description"`           //	Represents the localized description for this account type.
	AllowedAccountConfigs *ConfigList `json:"allowedAccountConfigs"` //	Represents the list of allowed account configs.
}

/*
TYPE: ExternalAccountTypeList
Represents a list of external account types.
*/
type ExternalAccountTypeList struct {
	Items []*ExternalAccountType `json:"items"` //
}

/*
TYPE: ExternalUserMapping
This is the model for external user mapping information in the API, v19 and beyond. These can be of 4 types : LDAP group, SAML, SAML attribute and External Script.
*/
type ExternalUserMapping struct {
}

/*
TYPE: ExternalUserMappingList
A list of external user mappings.
*/
type ExternalUserMappingList struct {
}

/*
TYPE: ExternalUserMappingRef
An externalUserMappingRef references an externalUserMapping.
*/
type ExternalUserMappingRef struct {
}

/*
TYPE: ExternalUserMappingType
Enum for external user mapping types
*/
type ExternalUserMappingType struct {
}

/*
TYPE: GenerateHostCertsArguments
Arguments to install certificates on a host
*/
type GenerateHostCertsArguments struct {
}

/*
TYPE: HBaseSnapshot
An HBase snapshot descriptor.
*/
type HBaseSnapshot struct {
}

/*
TYPE: HBaseSnapshotError
A HBase snapshot operation error.
*/
type HBaseSnapshotError struct {
}

/*
TYPE: HBaseSnapshotPolicyArguments
HBase specific snapshot policy arguments.
*/
type HBaseSnapshotPolicyArguments struct {
}

/*
TYPE: HBaseSnapshotResult
Detailed information about an HBase snapshot command.
*/
type HBaseSnapshotResult struct {
}

/*
TYPE: HdfsCloudReplicationArguments
Replication arguments for HDFS.
*/
type HdfsCloudReplicationArguments struct {
}

/*
TYPE: HdfsDisableHaArguments
Arguments used for the HDFS disable HA command.
*/
type HdfsDisableHaArguments struct {
}

/*
TYPE: HdfsFailoverArguments
Arguments used when enabling HDFS automatic failover.
*/
type HdfsFailoverArguments struct {
}

/*
TYPE: HdfsHaArguments
Arguments used for HDFS HA commands.
*/
type HdfsHaArguments struct {
}

/*
TYPE: HdfsReplicationArguments
Replication arguments for HDFS.
*/
type HdfsReplicationArguments struct {
}

/*
TYPE: HdfsReplicationCounter
A counter in an HDFS replication job.
*/
type HdfsReplicationCounter struct {
}

/*
TYPE: HdfsReplicationResult
Detailed information about an HDFS replication job.
*/
type HdfsReplicationResult struct {
}

/*
TYPE: HdfsSnapshot
An HDFS snapshot descriptor.
*/
type HdfsSnapshot struct {
}

/*
TYPE: HdfsSnapshotError
An HDFS snapshot operation error.
*/
type HdfsSnapshotError struct {
}

/*
TYPE: HdfsSnapshotPolicyArguments
HDFS specific snapshot policy arguments.
*/
type HdfsSnapshotPolicyArguments struct {
}

/*
TYPE: HdfsSnapshotResult
Detailed information about an HDFS snapshot command.
*/
type HdfsSnapshotResult struct {
}

/*
TYPE: HdfsUsageReport
A generic list.
*/
type HdfsUsageReport struct {
}

/*
TYPE: HdfsUsageReportRow

ApiHealthCheck Represents a result from a health test performed by Cloudera Manager for an entity.
*/
type HdfsUsageReportRow struct {
}

/*
TYPE: HealthSummary
Represents of the high-level health status of a subject in the cluster.
*/
type HealthSummary struct {
}

/*
TYPE: HiveCloudReplicationArguments
Replication arguments for Hive services.
*/
type HiveCloudReplicationArguments struct {
}

/*
TYPE: HiveReplicationArguments
Replication arguments for Hive services.
*/
type HiveReplicationArguments struct {
}

/*
TYPE: HiveReplicationError
A Hive replication error.
*/
type HiveReplicationError struct {
}

/*
TYPE: HiveReplicationResult
Detailed information about a Hive replication job.
*/
type HiveReplicationResult struct {
}

/*
TYPE: HiveTable
A Hive table identifier.
*/
type HiveTable struct {
}

/*
TYPE: HiveUDF
An hive UDF identifier.
*/
type HiveUDF struct {
}

/*
TYPE: Host
This is the model for a host in the system.
*/
type Host struct {
}

/*
TYPE: HostInstallArguments
Arguments to perform installation on one or more hosts
*/
type HostInstallArguments struct {
}

/*
TYPE: HostList
A list of ApiHost objects
*/
type HostList struct {
}

/*
TYPE: HostNameList
A list of host names.
*/
type HostNameList struct {
}

/*
TYPE: HostRef
A reference to a host.
*/
type HostRef struct {
}

/*
TYPE: HostRefList
A list of host references.
*/
type HostRefList struct {
}

/*
TYPE: HostTemplate
A host template belongs to a cluster and contains a set of role config groups for slave roles (such as DataNodes and TaskTrackers) from services in the cluster. At most one role config group per role type can be present in a host template. Host templates can be applied to fresh hosts (those with no roles on them) in order to create a role for each of the role groups on each host.
*/
type HostTemplate struct {
}

/*
TYPE: HostTemplateList
A list of host templates.
*/
type HostTemplateList struct {
}

/*
TYPE: ImpalaCancelResponse
The response from an Impala cancel query response.
*/
type ImpalaCancelResponse struct {
}

/*
TYPE: ImpalaQuery
Represents an Impala Query.
*/
type ImpalaQuery struct {
}

/*
TYPE: ImpalaQueryAttribute
Metadata about an Impala query attribute.
*/
type ImpalaQueryAttribute struct {
}

/*
TYPE: ImpalaQueryAttributeList
The list of all the attributes that are applicable to Impala queries.
*/
type ImpalaQueryAttributeList struct {
}

/*
TYPE: ImpalaQueryDetailsResponse
A query details response.
*/
type ImpalaQueryDetailsResponse struct {
}

/*
TYPE: ImpalaQueryResponse
The response contains a list of queries and warnings.
*/
type ImpalaQueryResponse struct {
}

/*
TYPE: ImpalaTenantUtilization
Utilization report information of a tenant of Impala application.
*/
type ImpalaTenantUtilization struct {
}

/*
TYPE: ImpalaTenantUtilizationList
A list of impala tenant utilization reports.
*/
type ImpalaTenantUtilizationList struct {
}

/*
TYPE: ImpalaUDF
An impala UDF identifier.
*/
type ImpalaUDF struct {
}

/*
TYPE: ImpalaUtilization
Utilization report information of a Impala application service.
*/
type ImpalaUtilization struct {
}

/*
TYPE: ImpalaUtilizationHistogram
Histogram of Impala utilization.
*/
type ImpalaUtilizationHistogram struct {
}

/*
TYPE: ImpalaUtilizationHistogramBin
Histogram bin of Impala utilization.
*/
type ImpalaUtilizationHistogramBin struct {
}

/*
TYPE: ImpalaUtilizationHistogramBinList
A generic list.
*/
type ImpalaUtilizationHistogramBinList struct {
}

/*
TYPE: JournalNodeArguments
Arguments used as part of ApiEnableNnHaArguments to specify JournalNodes.
*/
type JournalNodeArguments struct {
}

/*
TYPE: KerberosInfo
Kerberos information of a Cluster or Cloudera Manager.
*/
type KerberosInfo struct {
}

/*
TYPE: License
Information about the Cloudera Manager license.
*/
type License struct {
}

/*
TYPE: LicensedFeatureUsage
Information about the int8 of nodes using which product features.  Usage information is provided for individual clusters, as well as totals across all clusters.
*/
type LicensedFeatureUsage struct {
}

/*
TYPE: ListBase
A generic list.
*/
type ListBase struct {
}

/*
TYPE: MapEntry
Models a map entry, with a key and a value. By forming a list of these entries you can have the equivalent of Map<String, String> (since JAX-B doesn't support maps).
*/
type MapEntry struct {
}

/*
TYPE: Metric
A metric represents a specific metric monitored by the Cloudera Management Services, and a list of values matching a user query.  These fields are available only in the "full" view: displayName description
*/
type Metric struct {
}

/*
TYPE: MetricData
A single data point of metric data.
*/
type MetricData struct {
}

/*
TYPE: MetricList
A list of ApiMetric objects
*/
type MetricList struct {
}

/*
TYPE: MetricSchema
A metric schema represents the schema for a specific metric monitored by the Cloudera Management Services.
*/
type MetricSchema struct {
}

/*
TYPE: MetricSchemaList
A list of ApiMetricSchema objects
*/
type MetricSchemaList struct {
}

/*
TYPE: MigrateRolesArguments

ApiMr2AppInformation Represents MapReduce2 information for a YARN application.
*/
type MigrateRolesArguments struct {
}

/*
TYPE: MrUsageReport
A generic list.
*/
type MrUsageReport struct {
}

/*
TYPE: MrUsageReportRow

ApiNameservice Provides information about an HDFS nameservice.  Nameservices can be either a stand-alone NameNode, a NameNode paired with a SecondaryNameNode, or a high-availability pair formed by an active and a stand-by NameNode.  The following fields are only available in the object's full view: healthSummary healthChecks
*/
type MrUsageReportRow struct {
}

/*
TYPE: NameserviceList
A list of HDFS nameservices.
*/
type NameserviceList struct {
}

/*
TYPE: Parcel
A Parcel encapsulate a specific product and version. For example, (CDH 4.1). A parcel is downloaded, distributed to all the machines of a cluster and then allowed to be activated.  > The available parcels are determined by which cluster they will be running on. For example, a SLES parcel won't show up for a RHEL cluster.
*/
type Parcel struct {
}

/*
TYPE: ParcelList
A list of ApiParcel.
*/
type ParcelList struct {
}

/*
TYPE: ParcelRef
A parcelRef references a parcel. Each parcel is identified by its "parcelName" and "parcelVersion", and the "clusterName" of the cluster that is using it. To operate on the parcel object, use the API with the those fields as parameters.
*/
type ParcelRef struct {
}

/*
TYPE: ParcelState
The ApiParcelState encapsulates the state of a parcel while it is in transition and reports any errors that may have occurred..  The complete progress of a parcel is broken up into two different reporting indicators - progress and count. Progress is the primary indicator that reports the global state of transitions. For example, when downloading, progress and totalProgress will show the current int8 of bytes downloaded and the total int8 of bytes needed to be downloaded respectively.  The count and totalCount indicator is used when a state transition affects multiple hosts. The count and totalCount show the current int8 of hosts completed and the total int8 of hosts respectively. For example, during distribution, the progress and totalProgress will show how many bytes have been transferred to each host and the count will indicate how many hosts of of totalCount have had parcels unpacked.  Along with the two progress indicators, the ApiParcelState shows both errors and warnings that may have turned up during a state transition.
*/
type ParcelState struct {
}

/*
TYPE: ParcelUsage
This object provides a complete view of the usage of parcels in a given cluster - particularly which parcels are in use for which roles.
*/
type ParcelUsage struct {
}

/*
TYPE: ParcelUsageHost
This object is used to represent a host within an ApiParcelUsage.
*/
type ParcelUsageHost struct {
}

/*
TYPE: ParcelUsageParcel
This object is used to represent a parcel within an ApiParcelUsage.
*/
type ParcelUsageParcel struct {
}

/*
TYPE: ParcelUsageRack
This object is used to represent a rack within an ApiParcelUsage.
*/
type ParcelUsageRack struct {
}

/*
TYPE: ParcelUsageRole
This object is used to represent a role within an ApiParcelUsage.
*/
type ParcelUsageRole struct {
}

/*
TYPE: PrincipalList
A list of kerberos principals.
*/
type PrincipalList struct {
}

/*
TYPE: Process
A process represents a unix process to be managed by the Cloudera Manager agents. A process can be a daemon, e.g. if it is associated with a running role. It can also be a one-off process which is expected to start, run and finish.
*/
type Process struct {
}

/*
TYPE: ProductVersion

ApiReplicationCommand Information about a replication command.  This object holds all the information a regular ApiCommand object provides, and adds specific information about the results of a replication command.  Depending on the type of the service where the replication was run, a different result property will be populated.
*/
type ProductVersion struct {
}

/*
TYPE: ReplicationCommandList
A list of replication commands.
*/
type ReplicationCommandList struct {
}

/*
TYPE: ReplicationDiagnosticsCollectionArgs
Optional arguments for diagnostics collection.
*/
type ReplicationDiagnosticsCollectionArgs struct {
}

/*
TYPE: ReplicationSchedule
A replication job schedule.  Replication jobs have service-specific arguments. This object has methods to retrieve arguments for all supported types of replication, but only one argument type is allowed to be set; the backend will check that the provided argument matches the service type where the replication is being scheduled.  The replication job's arguments should match the underlying service. Refer to each property's documentation to find out which properties correspond to which services.
*/
type ReplicationSchedule struct {
}

/*
TYPE: ReplicationScheduleList
A list of replication schedules.
*/
type ReplicationScheduleList struct {
}

/*
TYPE: ReplicationState
The state of Hive/HDFS Replication.
*/
type ReplicationState struct {
}

/*
TYPE: RestartClusterArgs
Arguments used for Cluster Restart command. Since V11: If both restartOnlyStaleServices and restartServiceNames are specified, a service must be specified in restartServiceNames and also be stale, in order to be restarted.
*/
type RestartClusterArgs struct {
}

/*
TYPE: Role
A role represents a specific entity that participate in a service. Examples are JobTrackers, DataNodes, HBase Masters. Each role is assigned a host where it runs on.
*/
type Role struct {
}

/*
TYPE: RoleConfigGroup
A role config group contains roles of the same role type sharing the same configuration. While each role has to belong to a group, a role config group may be empty. There exists a default role config group for each role type. Default groups cannot be removed nor created. The name of a role config group is unique and cannot be changed. The configuration of individual roles may be overridden on role level.
*/
type RoleConfigGroup struct {
}

/*
TYPE: RoleConfigGroupList
A list of role config groups.
*/
type RoleConfigGroupList struct {
}

/*
TYPE: RoleConfigGroupRef

ApiRoleList A list of roles.
*/
type RoleConfigGroupRef struct {
}

/*
TYPE: RoleNameList
A list of role names.
*/
type RoleNameList struct {
}

/*
TYPE: RoleRef
A roleRef references a role. Each role is identified by its "roleName", the "serviceName" for the service it belongs to, and the "clusterName" in which the service resides. To operate on the role object, use the API with the those fields as parameters.
*/
type RoleRef struct {
}

/*
TYPE: RoleState
Represents the configured run state of a role.
*/
type RoleState struct {
}

/*
TYPE: RoleTypeConfig
Role type configuration information.
*/
type RoleTypeConfig struct {
}

/*
TYPE: RoleTypeList
A list of roles types that exists for a given service.
*/
type RoleTypeList struct {
}

/*
TYPE: RolesToInclude
Roles to include during a cluster rolling restart.
*/
type RolesToInclude struct {
}

/*
TYPE: RollEditsArgs
Arguments used for the Roll Edits command.
*/
type RollEditsArgs struct {
}

/*
TYPE: RollingRestartArgs
Arguments used for Rolling Restart commands.
*/
type RollingRestartArgs struct {
}

/*
TYPE: RollingRestartClusterArgs
Arguments used for Rolling Restart Cluster command.
*/
type RollingRestartClusterArgs struct {
}

/*
TYPE: RollingUpgradeClusterArgs
Rolling upgrade arguments used in the CDH Upgrade Command. Part of ApiCdhUpgradeArgs.
*/
type RollingUpgradeClusterArgs struct {
}

/*
TYPE: RollingUpgradeServicesArgs
Arguments used for Rolling Upgrade command.
*/
type RollingUpgradeServicesArgs struct {
}

/*
TYPE: Schedule
Base class for commands that can be scheduled in Cloudera Manager.  Note that schedule IDs are not preserved upon import.
*/
type Schedule struct {
}

/*
TYPE: ScheduleInterval
Represents the unit for the repeat interval for schedules.
*/
type ScheduleInterval struct {
}

/*
TYPE: ScmDbInfo
Cloudera Manager server's database information
*/
type ScmDbInfo struct {
}

/*
TYPE: Service
A service (such as HDFS, MapReduce, HBase) runs in a cluster. It has roles, which are the actual entities (NameNode, DataNodes, etc.) that perform the service's functions.  HDFS services and health checks In CDH4, HDFS services may not present any health checks. This will happen if the service has more than one nameservice configured. In those cases, the health information will be available by fetching information about the nameservices instead.  The health summary is still available, and reflects a service-wide summary.
*/
type Service struct {
}

/*
TYPE: ServiceConfig
Service and role type configuration.
*/
type ServiceConfig struct {
}

/*
TYPE: ServiceList
A list of services.
*/
type ServiceList struct {
}

/*
TYPE: ServiceRef
A serviceRef references a service. It is identified by the "serviceName", "clusterName" (name of the cluster which the service belongs to) and an optional "peerName" (to reference a remote service i.e. services managed by other CM instances). To operate on the service object, use the API with those fields as parameters.
*/
type ServiceRef struct {
}

/*
TYPE: ServiceState
Represents the configured run state of a service.
*/
type ServiceState struct {
}

/*
TYPE: ServiceTypeList
A list of service types that exists for a given cluster.
*/
type ServiceTypeList struct {
}

/*
TYPE: ShutdownReadiness
Cloudera Manager server's shutdown readiness
*/
type ShutdownReadiness struct {
}

/*
TYPE: SimpleRollingRestartClusterArgs
Basic arguments used for Rolling Restart Cluster commands.
*/
type SimpleRollingRestartClusterArgs struct {
}

/*
TYPE: SnapshotCommand
Information about snapshot commands.  This object holds all the information a regular ApiCommand object provides, and adds specific information about the results of a snapshot command.  Depending on the type of the service where the snapshot command was run, a different result property will be populated.
*/
type SnapshotCommand struct {
}

/*
TYPE: SnapshotCommandList
A list of snapshot commands.
*/
type SnapshotCommandList struct {
}

/*
TYPE: SnapshotPolicy
A snapshot policy.  Snapshot policies have service specific arguments. This object has methods to retrieve arguments for all supported types of snapshots, but only one argument type is allowed to be set; the backend will check that the provided argument matches the type of the service with which the snapshot policy is associated.
*/
type SnapshotPolicy struct {
}

/*
TYPE: SnapshotPolicyList
A list of snapshot policies.
*/
type SnapshotPolicyList struct {
}

/*
TYPE: TenantUtilization
Utilization report information of a tenant.
*/
type TenantUtilization struct {
}

/*
TYPE: TenantUtilizationList
A list of tenant utilization reports.
*/
type TenantUtilizationList struct {
}

/*
TYPE: TimeSeries
A time series represents a stream of data points. Each data point contains a time and a value. Time series are returned by executing a tsquery.
*/
type TimeSeries struct {
}

/*
TYPE: TimeSeriesAggregateStatistics
Statistics related to one time series aggregate data point. It is available from v6 for data points containing aggregate data. It includes further statistics about the data point. An aggregate can be across entities (e.g., fd_open_across_datanodes), over time (e.g., a daily point for the fd_open metric for a specific DataNode), or both (e.g., a daily point for the fd_open_across_datanodes metric). If the data point is for non-aggregate date this will return null.
*/
type TimeSeriesAggregateStatistics struct {
}

/*
TYPE: TimeSeriesCrossEntityMetadata
A class holding additional metadata to the ApiTimeSeriesAggregateStatistics class that applies specifically to cross-entity aggregate metrics.
*/
type TimeSeriesCrossEntityMetadata struct {
}

/*
TYPE: TimeSeriesData
A single data point of time series data.
*/
type TimeSeriesData struct {
}

/*
TYPE: TimeSeriesEntityAttribute
A time series entity attribute represents a possible attribute of a time series entity type monitored by the Cloudera Management Services.  Available since API v11.
*/
type TimeSeriesEntityAttribute struct {
}

/*
TYPE: TimeSeriesEntityAttributeList
A list of ApiTimeSeriesEntityAttribute objects
*/
type TimeSeriesEntityAttributeList struct {
}

/*
TYPE: TimeSeriesEntityType
Describe a time series entity type and attributes associated with this entity type.  Available since API v11.
*/
type TimeSeriesEntityType struct {
}

/*
TYPE: TimeSeriesEntityTypeList
A list of ApiTimeSeriesEntityType objects
*/
type TimeSeriesEntityTypeList struct {
}

/*
TYPE: TimeSeriesMetadata
Metadata for a time series.
*/
type TimeSeriesMetadata struct {
}

/*
TYPE: TimeSeriesRequest
Request object containing information needed for querying timeseries data. Available since API v11.
*/
type TimeSeriesRequest struct {
}

/*
TYPE: TimeSeriesResponse
The time series response for a time series query.
*/
type TimeSeriesResponse struct {
}

/*
TYPE: TimeSeriesResponseList
A generic list.
*/
type TimeSeriesResponseList struct {
}

/*
TYPE: User
This is the model for user information in the API prior to v18. Post v18, please refer to ApiUser2.java.  Note that any method that returns user information will not contain any password information. The password property is only used when creating or updating users.
*/
type User struct {
}

/*
TYPE: User2
This is the model for user information in the API, v18 and beyond.  Note that any method that returns user information will not contain any password information. The password property is only used when creating or updating users.
*/
type User2 struct {
}

/*
TYPE: User2List
A list of users.
*/
type User2List struct {
}

/*
TYPE: User2Ref
A userRef references a user.
*/
type User2Ref struct {
}

/*
TYPE: UserList
A list of users.
*/
type UserList struct {
}

/*
TYPE: UserSession
This is the model for interactive user session information in the API.  A user may have more than one active session. Each such session will have its own session object.
*/
type UserSession struct {
}

/*
TYPE: UserSessionList
A list of user sessions.
*/
type UserSessionList struct {
}

/*
TYPE: VersionInfo
Version information of Cloudera Manager itself.
*/
type VersionInfo struct {
	Version        string `json:"version"`        // Version.
	Snapshot       bool   `json:"snapshot"`       // Whether this build is a development snapshot.
	BuildUser      string `json:"buildUser"`      // The user performing the build.
	BuildTimestamp string `json:"buildTimestamp"` // Build timestamp.
	GitHash        string `json:"gitHash"`        // Source control management hash.
}

/*
TYPE: WatchedDir

ApiWatchedDirList A list of watched directories.
*/
type WatchedDir struct {
}

/*
TYPE: YarnApplication
Represents a Yarn application
*/
type YarnApplication struct {
}

/*
TYPE: YarnApplicationAttribute
Metadata about a YARN application attribute.
*/
type YarnApplicationAttribute struct {
}

/*
TYPE: YarnApplicationAttributeList
The list of all attributes that are applicable to YARN applications.
*/
type YarnApplicationAttributeList struct {
}

/*
TYPE: YarnApplicationDiagnosticsCollectionArgs
Arguments used for collecting diagnostics data for Yarn applications
*/
type YarnApplicationDiagnosticsCollectionArgs struct {
}

/*
TYPE: YarnApplicationResponse
The response contains a list of applications and warnings.
*/
type YarnApplicationResponse struct {
}

/*
TYPE: YarnKillResponse
The response from an Yarn kill application response.
*/
type YarnKillResponse struct {
}

/*
TYPE: YarnTenantUtilization
Utilization report information of a tenant of Yarn application.
*/
type YarnTenantUtilization struct {
}

/*
TYPE: YarnTenantUtilizationList
A list of yarn tenant utilization reports.
*/
type YarnTenantUtilizationList struct {
}

/*
TYPE: YarnUtilization
Utilization report information of a Yarn application service.
*/
type YarnUtilization struct {
}

/*
TYPE: HTTPMethod

HaStatus
*/
type HTTPMethod struct {
}

/*
TYPE: ReplicationOption
This will decide how cloud replication will take place
*/
type ReplicationOption struct {
}

/*
TYPE: ReplicationStrategy
The strategy for distributing the file replication tasks among the mappers of the MR job associated with a replication.
*/
type ReplicationStrategy struct {
}

/*
TYPE: ScmDbType
Enum for Cloudera Manager DB type. Note that DERBY and SQLITE3 are not supported DBs
*/
type ScmDbType struct {
}

/*
TYPE: ShutdownReadinessState
Enum for Cloudera Manager shutdown readiness state.
*/
type ShutdownReadinessState struct {
}

/*
TYPE: Storage

ValidationState
*/
type ValidationState struct {
}

/*
TYPE: Storage
*/
type Storage struct {
}

/*
TYPE: ZooKeeperServerMode
The state of the Zookeeper server.piActivity Represents a user activity, such as a MapReduce job, a Hive query, an Oozie workflow, etc.
*/
type ZooKeeperServerMode struct {
}
