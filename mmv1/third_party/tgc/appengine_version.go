package google

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AppEngineVersionAssetType string = "appengine.googleapis.com/Version"

func resourceAppEngineVersion() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AppEngineVersionAssetType,
		Convert:   GetAppEngineVersionCaiObject,
	}
}

func GetAppEngineVersionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//appengine.googleapis.com/projects/{{project}}/zones/{{zone}}/version/{{group}}") //nao tenho ctz
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAppEngineVersionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AppEngineVersionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://appengine.googleapis.com/$discovery/rest?version=v1",
				DiscoveryName:        "Version",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAppEngineVersionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandAppEngineVersionName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	automaticScalingProp, err := expandAppEngineVersionAutomaticScaling(d.Get("automaticScaling"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("automaticScaling"); !tpgresource.IsEmptyValue(reflect.ValueOf(automaticScalingProp)) && (ok || !reflect.DeepEqual(v, automaticScalingProp)) {
		obj["automaticScaling"] = automaticScalingProp
	}

	networkProp, err := expandAppEngineVersionNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}

	basicScalingProp, err := expandAppEngineVersionBasicScaling(d.Get("basicScaling"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("basicScaling"); !tpgresource.IsEmptyValue(reflect.ValueOf(basicScalingProp)) && (ok || !reflect.DeepEqual(v, basicScalingProp)) {
		obj["basicScaling"] = basicScalingProp
	}

	manualScalingProp, err := expandAppEngineVersionManualScaling(d.Get("manualScaling"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("manualScaling"); !tpgresource.IsEmptyValue(reflect.ValueOf(manualScalingProp)) && (ok || !reflect.DeepEqual(v, manualScalingProp)) {
		obj["manualScaling"] = manualScalingProp
	}

	instanceClassProp, err := expandAppEngineVersionInstanceClass(d.Get("instanceClass"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instanceClass"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceClassProp)) && (ok || !reflect.DeepEqual(v, instanceClassProp)) {
		obj["instanceClass"] = instanceClassProp
	}

	zonesProp, err := expandAppEngineVersionZones(d.Get("zones"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zones"); !tpgresource.IsEmptyValue(reflect.ValueOf(zonesProp)) && (ok || !reflect.DeepEqual(v, zonesProp)) {
		obj["zones"] = zonesProp
	}

	resourcesProp, err := expandAppEngineVersionResources(d.Get("resources"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourcesProp)) && (ok || !reflect.DeepEqual(v, resourcesProp)) {
		obj["resources"] = resourcesProp
	}

	runtimeProp, err := expandAppEngineVersionRuntime(d.Get("runtime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtime"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeProp)) && (ok || !reflect.DeepEqual(v, runtimeProp)) {
		obj["runtime"] = runtimeProp
	}

	runtimeChannelProp, err := expandAppEngineVersionRuntimeChannel(d.Get("runtimeChannel"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtimeChannel"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeChannelProp)) && (ok || !reflect.DeepEqual(v, runtimeChannelProp)) {
		obj["runtimeChannel"] = runtimeChannelProp
	}

	threadsafeProp, err := expandAppEngineVersionThreadsafe(d.Get("threadsafe"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("threadsafe"); !tpgresource.IsEmptyValue(reflect.ValueOf(threadsafeProp)) && (ok || !reflect.DeepEqual(v, threadsafeProp)) {
		obj["threadsafe"] = threadsafeProp
	}

	vmProp, err := expandAppEngineVersionVm(d.Get("vm"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vm"); !tpgresource.IsEmptyValue(reflect.ValueOf(vmProp)) && (ok || !reflect.DeepEqual(v, vmProp)) {
		obj["vm"] = vmProp
	}

	flexibleRuntimeSettingsProp, err := expandAppEngineVersionFlexibleRuntimeSettings(d.Get("flexibleRuntimeSettings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("flexibleRuntimeSettings"); !tpgresource.IsEmptyValue(reflect.ValueOf(flexibleRuntimeSettingsProp)) && (ok || !reflect.DeepEqual(v, flexibleRuntimeSettingsProp)) {
		obj["flexibleRuntimeSettings"] = flexibleRuntimeSettingsProp
	}

	appEngineApisProp, err := expandAppEngineVersionAppEngineApis(d.Get("appEngineApis"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("appEngineApis"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineApisProp)) && (ok || !reflect.DeepEqual(v, appEngineApisProp)) {
		obj["appEngineApis"] = appEngineApisProp
	}

	envProp, err := expandAppEngineVersionEnv(d.Get("env"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("env"); !tpgresource.IsEmptyValue(reflect.ValueOf(envProp)) && (ok || !reflect.DeepEqual(v, envProp)) {
		obj["env"] = envProp
	}

	createdByProp, err := expandAppEngineVersionCreatedBy(d.Get("createdBy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("createdBy"); !tpgresource.IsEmptyValue(reflect.ValueOf(createdByProp)) && (ok || !reflect.DeepEqual(v, createdByProp)) {
		obj["createdBy"] = createdByProp
	}

	createTimeProp, err := expandAppEngineVersionCreateTime(d.Get("createTime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("createTime"); !tpgresource.IsEmptyValue(reflect.ValueOf(createTimeProp)) && (ok || !reflect.DeepEqual(v, createTimeProp)) {
		obj["createTime"] = createTimeProp
	}

	diskUsageBytesProp, err := expandAppEngineVersionDiskUsageBytes(d.Get("diskUsageBytes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("diskUsageBytes"); !tpgresource.IsEmptyValue(reflect.ValueOf(diskUsageBytesProp)) && (ok || !reflect.DeepEqual(v, diskUsageBytesProp)) {
		obj["diskUsageBytes"] = diskUsageBytesProp
	}

	runtimeApiVersionProp, err := expandAppEngineVersionRuntimeApiVersion(d.Get("runtimeApiVersion"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtimeApiVersion"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeApiVersionProp)) && (ok || !reflect.DeepEqual(v, runtimeApiVersionProp)) {
		obj["runtimeApiVersion"] = runtimeApiVersionProp
	}

	runtimeMainExecutablePathProp, err := expandAppEngineVersionRuntimeMainExecutablePath(d.Get("runtimeMainExecutablePath"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtimeMainExecutablePath"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeMainExecutablePathProp)) && (ok || !reflect.DeepEqual(v, runtimeMainExecutablePathProp)) {
		obj["runtimeMainExecutablePath"] = runtimeMainExecutablePathProp
	}

	serviceAccountProp, err := expandAppEngineVersionServiceAccount(d.Get("serviceAccount"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("serviceAccount"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}

	handlersProp, err := expandAppEngineVersionHandlers(d.Get("handlers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("handlers"); !tpgresource.IsEmptyValue(reflect.ValueOf(handlersProp)) && (ok || !reflect.DeepEqual(v, handlersProp)) {
		obj["handlers"] = handlersProp
	}

	errorHandlersProp, err := expandAppEngineVersionErrorHandlers(d.Get("errorHandlers"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("errorHandlers"); !tpgresource.IsEmptyValue(reflect.ValueOf(errorHandlersProp)) && (ok || !reflect.DeepEqual(v, errorHandlersProp)) {
		obj["errorHandlers"] = errorHandlersProp
	}

	librariesProp, err := expandAppEngineVersionLibraries(d.Get("libraries"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("libraries"); !tpgresource.IsEmptyValue(reflect.ValueOf(librariesProp)) && (ok || !reflect.DeepEqual(v, librariesProp)) {
		obj["libraries"] = librariesProp
	}

	apiConfigProp, err := expandAppEngineVersionApiConfig(d.Get("apiConfig"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("apiConfig"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiConfigProp)) && (ok || !reflect.DeepEqual(v, apiConfigProp)) {
		obj["apiConfig"] = apiConfigProp
	}

	deploymentProp, err := expandAppEngineVersionDeployment(d.Get("deployment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deployment"); !tpgresource.IsEmptyValue(reflect.ValueOf(deploymentProp)) && (ok || !reflect.DeepEqual(v, deploymentProp)) {
		obj["deployment"] = deploymentProp
	}

	return obj, nil
}

func expandAppEngineVersionName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionAutomaticScaling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	coolDownPeriodProp, err := expandAppEngineVersionCoolDownPeriod(d.Get("coolDownPeriod"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("coolDownPeriod"); !tpgresource.IsEmptyValue(reflect.ValueOf(coolDownPeriodProp)) && (ok || !reflect.DeepEqual(v, coolDownPeriodProp)) {
		obj["coolDownPeriod"] = coolDownPeriodProp
	}

	cpuUtilizationProp, err := expandAppEngineVersionCpuUtilization(d.Get("cpuUtilization"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cpuUtilization"); !tpgresource.IsEmptyValue(reflect.ValueOf(cpuUtilizationProp)) && (ok || !reflect.DeepEqual(v, cpuUtilizationProp)) {
		obj["cpuUtilization"] = cpuUtilizationProp
	}

	maxConcurrentRequestsProp, err := expandAppEngineVersionMaxConcurrentRequests(d.Get("maxConcurrentRequests"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maxConcurrentRequests"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxConcurrentRequestsProp)) && (ok || !reflect.DeepEqual(v, maxConcurrentRequestsProp)) {
		obj["maxConcurrentRequests"] = maxConcurrentRequestsProp
	}

	maxIdleInstancesProp, err := expandAppEngineVersionMaxIdleInstances(d.Get("maxIdleInstances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maxIdleInstances"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxIdleInstancesProp)) && (ok || !reflect.DeepEqual(v, maxIdleInstancesProp)) {
		obj["maxIdleInstances"] = maxIdleInstancesProp
	}

	maxTotalInstancesProp, err := expandAppEngineVersionMaxTotalInstances(d.Get("maxTotalInstances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maxTotalInstances"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxTotalInstancesProp)) && (ok || !reflect.DeepEqual(v, maxTotalInstancesProp)) {
		obj["maxTotalInstances"] = maxTotalInstancesProp
	}

	maxPendingLatencyProp, err := expandAppEngineVersionMaxPendingLatency(d.Get("maxPendingLatency"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maxPendingLatency"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxPendingLatencyProp)) && (ok || !reflect.DeepEqual(v, maxPendingLatencyProp)) {
		obj["maxPendingLatency"] = maxPendingLatencyProp
	}

	minIdleInstancesProp, err := expandAppEngineVersionMinIdleInstances(d.Get("minIdleInstances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("minIdleInstances"); !tpgresource.IsEmptyValue(reflect.ValueOf(minIdleInstancesProp)) && (ok || !reflect.DeepEqual(v, minIdleInstancesProp)) {
		obj["minIdleInstances"] = minIdleInstancesProp
	}

	minTotalInstancesProp, err := expandAppEngineVersionMinTotalInstances(d.Get("minTotalInstances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("minTotalInstances"); !tpgresource.IsEmptyValue(reflect.ValueOf(minTotalInstancesProp)) && (ok || !reflect.DeepEqual(v, minTotalInstancesProp)) {
		obj["minTotalInstances"] = minTotalInstancesProp
	}

	minPendingLatencyProp, err := expandAppEngineVersionMinPendingLatency(d.Get("minPendingLatency"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("minPendingLatency"); !tpgresource.IsEmptyValue(reflect.ValueOf(minPendingLatencyProp)) && (ok || !reflect.DeepEqual(v, minPendingLatencyProp)) {
		obj["minPendingLatency"] = minPendingLatencyProp
	}

	requestUtilizationProp, err := expandAppEngineVersionRequestUtilization(d.Get("requestUtilization"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("requestUtilization"); !tpgresource.IsEmptyValue(reflect.ValueOf(requestUtilizationProp)) && (ok || !reflect.DeepEqual(v, requestUtilizationProp)) {
		obj["requestUtilization"] = requestUtilizationProp
	}

	diskUtilizationProp, err := expandAppEngineVersionDiskUtilization(d.Get("diskUtilization"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("diskUtilization"); !tpgresource.IsEmptyValue(reflect.ValueOf(diskUtilizationProp)) && (ok || !reflect.DeepEqual(v, diskUtilizationProp)) {
		obj["diskUtilization"] = diskUtilizationProp
	}

	networkUtilizationProp, err := expandAppEngineVersionNetworkUtilization(d.Get("networkUtilization"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("networkUtilization"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkUtilizationProp)) && (ok || !reflect.DeepEqual(v, networkUtilizationProp)) {
		obj["networkUtilization"] = networkUtilizationProp
	}

	standardSchedulerSettingsProp, err := expandAppEngineVersionStandardSchedulerSettings(d.Get("standardSchedulerSettings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("standardSchedulerSettings"); !tpgresource.IsEmptyValue(reflect.ValueOf(standardSchedulerSettingsProp)) && (ok || !reflect.DeepEqual(v, standardSchedulerSettingsProp)) {
		obj["standardSchedulerSettings"] = standardSchedulerSettingsProp
	}

	return obj, nil
}

func expandAppEngineVersionNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	forwardedPortsProp, err := expandAppEngineVersionForwardedPorts(d.Get("forwardedPorts"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("forwardedPorts"); !tpgresource.IsEmptyValue(reflect.ValueOf(forwardedPortsProp)) && (ok || !reflect.DeepEqual(v, forwardedPortsProp)) {
		obj["forwardedPorts"] = forwardedPortsProp
	}

	instanceTagProp, err := expandAppEngineVersionInstanceTag(d.Get("instanceTag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instanceTag"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceTagProp)) && (ok || !reflect.DeepEqual(v, instanceTagProp)) {
		obj["instanceTag"] = instanceTagProp
	}

	nameProp, err := expandAppEngineVersionName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	subnetworkNameProp, err := expandAppEngineVersionSubnetworkName(d.Get("subnetworkName"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("subnetworkName"); !tpgresource.IsEmptyValue(reflect.ValueOf(subnetworkNameProp)) && (ok || !reflect.DeepEqual(v, subnetworkNameProp)) {
		obj["subnetworkName"] = subnetworkNameProp
	}

	sessionAffinityProp, err := expandAppEngineVersionSessionAffinity(d.Get("sessionAffinity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sessionAffinity"); !tpgresource.IsEmptyValue(reflect.ValueOf(sessionAffinityProp)) && (ok || !reflect.DeepEqual(v, sessionAffinityProp)) {
		obj["sessionAffinity"] = sessionAffinityProp
	}

	return obj, nil

}

func expandAppEngineVersionForwardedPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionInstanceTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionSubnetworkName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionSessionAffinity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionCoolDownPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionCpuUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	aggregationWindowLengthProp, err := expandAppEngineVersionAggregationWindowLength(d.Get("aggregationWindowLength"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("aggregationWindowLength"); !tpgresource.IsEmptyValue(reflect.ValueOf(aggregationWindowLengthProp)) && (ok || !reflect.DeepEqual(v, aggregationWindowLengthProp)) {
		obj["aggregationWindowLength"] = aggregationWindowLengthProp
	}

	targetUtilizationProp, err := expandAppEngineVersionTargetUtilization(d.Get("targetUtilization"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetUtilization"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetUtilizationProp)) && (ok || !reflect.DeepEqual(v, targetUtilizationProp)) {
		obj["targetUtilization"] = targetUtilizationProp
	}

	return obj, nil
}

func expandAppEngineVersionAggregationWindowLength(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMaxConcurrentRequests(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMaxIdleInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMaxTotalInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMaxPendingLatency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMinIdleInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMinTotalInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMinPendingLatency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionRequestUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	targetRequestCountPerSecondProp, err := expandAppEngineVersionTargetRequestCountPerSecond(d.Get("targetRequestCountPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetRequestCountPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetRequestCountPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetRequestCountPerSecondProp)) {
		obj["targetRequestCountPerSecond"] = targetRequestCountPerSecondProp
	}

	targetConcurrentRequestsProp, err := expandAppEngineVersionTargetConcurrentRequests(d.Get("targetConcurrentRequests"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetConcurrentRequests"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetConcurrentRequestsProp)) && (ok || !reflect.DeepEqual(v, targetConcurrentRequestsProp)) {
		obj["targetConcurrentRequests"] = targetConcurrentRequestsProp
	}

	return obj, nil
}

func expandAppEngineVersionTargetRequestCountPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetConcurrentRequests(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionDiskUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	targetWriteBytesPerSecondProp, err := expandAppEngineVersionTargetWriteBytesPerSecond(d.Get("targetWriteBytesPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetWriteBytesPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetWriteBytesPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetWriteBytesPerSecondProp)) {
		obj["targetWriteBytesPerSecond"] = targetWriteBytesPerSecondProp
	}

	targetWriteOpsPerSecondProp, err := expandAppEngineVersionTargetWriteOpsPerSecond(d.Get("targetWriteOpsPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetWriteOpsPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetWriteOpsPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetWriteOpsPerSecondProp)) {
		obj["targetWriteOpsPerSecond"] = targetWriteOpsPerSecondProp
	}

	targetReadBytesPerSecondProp, err := expandAppEngineVersionTargetReadBytesPerSecond(d.Get("targetReadBytesPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetReadBytesPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetReadBytesPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetReadBytesPerSecondProp)) {
		obj["targetReadBytesPerSecond"] = targetReadBytesPerSecondProp
	}

	targetReadOpsPerSecondProp, err := expandAppEngineVersionTargetReadOpsPerSecond(d.Get("targetReadOpsPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetReadOpsPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetReadOpsPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetReadOpsPerSecondProp)) {
		obj["targetReadOpsPerSecond"] = targetReadOpsPerSecondProp
	}

	return obj, nil
}

func expandAppEngineVersionTargetWriteOpsPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetReadBytesPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetWriteBytesPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetReadOpsPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionNetworkUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	targetSentBytesPerSecondProp, err := expandAppEngineVersionTargetSentBytesPerSecond(d.Get("targetSentBytesPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetSentBytesPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetSentBytesPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetSentBytesPerSecondProp)) {
		obj["targetSentBytesPerSecond"] = targetSentBytesPerSecondProp
	}

	targetSentPacketsPerSecondProp, err := expandAppEngineVersionTargetSentPacketsPerSecond(d.Get("targetSentPacketsPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetSentPacketsPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetSentPacketsPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetSentPacketsPerSecondProp)) {
		obj["targetSentPacketsPerSecond"] = targetSentPacketsPerSecondProp
	}

	targetReceivedBytesPerSecondProp, err := expandAppEngineVersionTargetReceivedBytesPerSecond(d.Get("targetReceivedBytesPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetReceivedBytesPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetReceivedBytesPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetReceivedBytesPerSecondProp)) {
		obj["targetReceivedBytesPerSecond"] = targetReceivedBytesPerSecondProp
	}

	targetReceivedPacketsPerSecondProp, err := expandAppEngineVersionTargetReceivedPacketsPerSecond(d.Get("targetReceivedPacketsPerSecond"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetReceivedPacketsPerSecond"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetReceivedPacketsPerSecondProp)) && (ok || !reflect.DeepEqual(v, targetReceivedPacketsPerSecondProp)) {
		obj["targetReceivedPacketsPerSecond"] = targetReceivedPacketsPerSecondProp
	}

	return obj, nil
}

func expandAppEngineVersionTargetSentBytesPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetSentPacketsPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetReceivedBytesPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetReceivedPacketsPerSecond(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionStandardSchedulerSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	targetCpuUtilizationProp, err := expandAppEngineVersionTargetCpuUtilization(d.Get("targetCpuUtilization"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetCpuUtilization"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetCpuUtilizationProp)) && (ok || !reflect.DeepEqual(v, targetCpuUtilizationProp)) {
		obj["targetCpuUtilization"] = targetCpuUtilizationProp
	}

	targetThroughputUtilizationProp, err := expandAppEngineVersionTargetThroughputUtilization(d.Get("targetThroughputUtilization"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("targetThroughputUtilization"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetThroughputUtilizationProp)) && (ok || !reflect.DeepEqual(v, targetThroughputUtilizationProp)) {
		obj["targetThroughputUtilization"] = targetThroughputUtilizationProp
	}

	minInstancesProp, err := expandAppEngineVersionMinInstances(d.Get("minInstances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("minInstances"); !tpgresource.IsEmptyValue(reflect.ValueOf(minInstancesProp)) && (ok || !reflect.DeepEqual(v, minInstancesProp)) {
		obj["minInstances"] = minInstancesProp
	}

	maxInstancesProp, err := expandAppEngineVersionMaxInstances(d.Get("maxInstances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maxInstances"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxInstancesProp)) && (ok || !reflect.DeepEqual(v, maxInstancesProp)) {
		obj["maxInstances"] = maxInstancesProp
	}

	return obj, nil
}

func expandAppEngineVersionTargetCpuUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionTargetThroughputUtilization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMinInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMaxInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionBasicScaling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	idleTimeoutProp, err := expandAppEngineVersionIdleTimeout(d.Get("idleTimeout"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("idleTimeout"); !tpgresource.IsEmptyValue(reflect.ValueOf(idleTimeoutProp)) && (ok || !reflect.DeepEqual(v, idleTimeoutProp)) {
		obj["idleTimeout"] = idleTimeoutProp
	}

	maxInstancesProp, err := expandAppEngineVersionMaxInstances(d.Get("maxInstances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maxInstances"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxInstancesProp)) && (ok || !reflect.DeepEqual(v, maxInstancesProp)) {
		obj["maxInstances"] = maxInstancesProp
	}

	return obj, nil
}

func expandAppEngineVersionIdleTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionManualScaling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	instancesProp, err := expandAppEngineVersionInstances(d.Get("instances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instances"); !tpgresource.IsEmptyValue(reflect.ValueOf(instancesProp)) && (ok || !reflect.DeepEqual(v, instancesProp)) {
		obj["instances"] = instancesProp
	}

	return obj, nil
}

func expandAppEngineVersionInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionInstanceClass(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionZones(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	cpuProp, err := expandAppEngineVersionCpu(d.Get("cpu"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cpu"); !tpgresource.IsEmptyValue(reflect.ValueOf(cpuProp)) && (ok || !reflect.DeepEqual(v, cpuProp)) {
		obj["cpu"] = cpuProp
	}

	diskGbProp, err := expandAppEngineVersionDiskGb(d.Get("diskGb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("diskGb"); !tpgresource.IsEmptyValue(reflect.ValueOf(diskGbProp)) && (ok || !reflect.DeepEqual(v, diskGbProp)) {
		obj["diskGb"] = diskGbProp
	}

	memoryGbProp, err := expandAppEngineVersionMemoryGb(d.Get("memoryGb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("memoryGb"); !tpgresource.IsEmptyValue(reflect.ValueOf(memoryGbProp)) && (ok || !reflect.DeepEqual(v, memoryGbProp)) {
		obj["memoryGb"] = memoryGbProp
	}

	volumesProp, err := expandAppEngineVersionVolumes(d.Get("volumes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("volumes"); !tpgresource.IsEmptyValue(reflect.ValueOf(volumesProp)) && (ok || !reflect.DeepEqual(v, volumesProp)) {
		obj["volumes"] = volumesProp
	}

	kmsKeyReferenceProp, err := expandAppEngineVersionKmsKeyReference(d.Get("kmsKeyReference"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kmsKeyReference"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyReferenceProp)) && (ok || !reflect.DeepEqual(v, kmsKeyReferenceProp)) {
		obj["kmsKeyReference"] = kmsKeyReferenceProp
	}

	return obj, nil
}

func expandAppEngineVersionCpu(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionDiskGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMemoryGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionVolumes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	nameProp, err := expandAppEngineVersionName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	volumeTypeProp, err := expandAppEngineVersionVolumeType(d.Get("volumeType"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("volumeType"); !tpgresource.IsEmptyValue(reflect.ValueOf(volumeTypeProp)) && (ok || !reflect.DeepEqual(v, volumeTypeProp)) {
		obj["volumeType"] = volumeTypeProp
	}

	sizeGbProp, err := expandAppEngineVersionSizeGb(d.Get("sizeGb"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sizeGb"); !tpgresource.IsEmptyValue(reflect.ValueOf(sizeGbProp)) && (ok || !reflect.DeepEqual(v, sizeGbProp)) {
		obj["sizeGb"] = sizeGbProp
	}

	return obj, nil
}

func expandAppEngineVersionKmsKeyReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionVolumeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionRuntime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionRuntimeChannel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionThreadsafe(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionVm(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionFlexibleRuntimeSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	operatingSystemProp, err := expandAppEngineVersionOperatingSystem(d.Get("operatingSystem"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("operatingSystem"); !tpgresource.IsEmptyValue(reflect.ValueOf(operatingSystemProp)) && (ok || !reflect.DeepEqual(v, operatingSystemProp)) {
		obj["operatingSystem"] = operatingSystemProp
	}

	runtimeVersionProp, err := expandAppEngineVersionrRuntimeVersion(d.Get("runtimeVersion"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("runtimeVersion"); !tpgresource.IsEmptyValue(reflect.ValueOf(runtimeVersionProp)) && (ok || !reflect.DeepEqual(v, runtimeVersionProp)) {
		obj["runtimeVersion"] = runtimeVersionProp
	}

	return obj, nil
}

func expandAppEngineVersionOperatingSystem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionrRuntimeVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionAppEngineApis(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionEnv(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionCreateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionDiskUsageBytes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionRuntimeApiVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionRuntimeMainExecutablePath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionCreatedBy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionStaticFiles(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	pathProp, err := expandAppEngineVersionPath(d.Get("path"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("path"); !tpgresource.IsEmptyValue(reflect.ValueOf(pathProp)) && (ok || !reflect.DeepEqual(v, pathProp)) {
		obj["path"] = pathProp
	}

	uploadPathRegexProp, err := expandAppEngineVersionUploadPathRegex(d.Get("uploadPathRegex"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("uploadPathRegex"); !tpgresource.IsEmptyValue(reflect.ValueOf(uploadPathRegexProp)) && (ok || !reflect.DeepEqual(v, uploadPathRegexProp)) {
		obj["uploadPathRegex"] = uploadPathRegexProp
	}

	//OBJETO DE STRING, NAO SEI SE TA CERTO
	//https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StaticFilesHandler
	httpHeadersProp, err := expandAppEngineVersionHttpHeaders(d.Get("httpHeaders"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("httpHeaders"); !tpgresource.IsEmptyValue(reflect.ValueOf(httpHeadersProp)) && (ok || !reflect.DeepEqual(v, httpHeadersProp)) {
		obj["httpHeaders"] = httpHeadersProp
	}

	mimeTypeProp, err := expandAppEngineVersionMimeType(d.Get("mimeType"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mimeType"); !tpgresource.IsEmptyValue(reflect.ValueOf(mimeTypeProp)) && (ok || !reflect.DeepEqual(v, mimeTypeProp)) {
		obj["mimeType"] = mimeTypeProp
	}

	expirationProp, err := expandAppEngineVersionExpiration(d.Get("expiration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("expiration"); !tpgresource.IsEmptyValue(reflect.ValueOf(expirationProp)) && (ok || !reflect.DeepEqual(v, expirationProp)) {
		obj["expiration"] = expirationProp
	}

	requireMatchingFileProp, err := expandAppEngineVersionRequireMatchingFile(d.Get("requireMatchingFile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("requireMatchingFile"); !tpgresource.IsEmptyValue(reflect.ValueOf(requireMatchingFileProp)) && (ok || !reflect.DeepEqual(v, requireMatchingFileProp)) {
		obj["requireMatchingFile"] = requireMatchingFileProp
	}

	applicationReadableProp, err := expandAppEngineVersionApplicationReadable(d.Get("applicationReadable"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("applicationReadable"); !tpgresource.IsEmptyValue(reflect.ValueOf(applicationReadableProp)) && (ok || !reflect.DeepEqual(v, applicationReadableProp)) {
		obj["applicationReadable"] = applicationReadableProp
	}

	return obj, nil
}

func expandAppEngineVersionPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionUploadPathRegex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionHttpHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMimeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionExpiration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionRequireMatchingFile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionApplicationReadable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionUrlRegex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionHandlers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	urlRegexProp, err := expandAppEngineVersionUrlRegex(d.Get("urlRegex"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("urlRegex"); !tpgresource.IsEmptyValue(reflect.ValueOf(urlRegexProp)) && (ok || !reflect.DeepEqual(v, urlRegexProp)) {
		obj["urlRegex"] = urlRegexProp
	}

	staticFilesProp, err := expandAppEngineVersionStaticFiles(d.Get("staticFiles"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("staticFiles"); !tpgresource.IsEmptyValue(reflect.ValueOf(staticFilesProp)) && (ok || !reflect.DeepEqual(v, staticFilesProp)) {
		obj["staticFiles"] = staticFilesProp
	}

	scriptProp, err := expandAppEngineVersionScript(d.Get("script"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("script"); !tpgresource.IsEmptyValue(reflect.ValueOf(scriptProp)) && (ok || !reflect.DeepEqual(v, scriptProp)) {
		obj["script"] = scriptProp
	}

	apiEndpointProp, err := expandAppEngineVersionApiEndpoint(d.Get("apiEndpoint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("apiEndpoint"); !tpgresource.IsEmptyValue(reflect.ValueOf(apiEndpointProp)) && (ok || !reflect.DeepEqual(v, apiEndpointProp)) {
		obj["apiEndpoint"] = apiEndpointProp
	}

	return obj, nil
}

func expandAppEngineVersionScript(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	scriptPathProp, err := expandAppEngineVersionScriptPath(d.Get("scriptPath"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("scriptPath"); !tpgresource.IsEmptyValue(reflect.ValueOf(scriptPathProp)) && (ok || !reflect.DeepEqual(v, scriptPathProp)) {
		obj["scriptPath"] = scriptPathProp
	}

	return obj, nil
}

func expandAppEngineVersionApiEndpoint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	scriptPathProp, err := expandAppEngineVersionScriptPath(d.Get("scriptPath"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("scriptPath"); !tpgresource.IsEmptyValue(reflect.ValueOf(scriptPathProp)) && (ok || !reflect.DeepEqual(v, scriptPathProp)) {
		obj["scriptPath"] = scriptPathProp
	}

	return obj, nil
}

func expandAppEngineVersionScriptPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionErrorHandlers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	staticFileProp, err := expandAppEngineVersionStaticFile(d.Get("staticFile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("staticFile"); !tpgresource.IsEmptyValue(reflect.ValueOf(staticFileProp)) && (ok || !reflect.DeepEqual(v, staticFileProp)) {
		obj["staticFile"] = staticFileProp
	}

	mimeTypeFileProp, err := expandAppEngineVersionMimeTypeFile(d.Get("mimeTypeFile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mimeTypeFile"); !tpgresource.IsEmptyValue(reflect.ValueOf(mimeTypeFileProp)) && (ok || !reflect.DeepEqual(v, mimeTypeFileProp)) {
		obj["mimeTypeFile"] = mimeTypeFileProp
	}

	return obj, nil
}

func expandAppEngineVersionStaticFile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionMimeTypeFile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionLibraries(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	nameProp, err := expandAppEngineVersionName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	versionProp, err := expandAppEngineVersionVersion(d.Get("version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}

	return obj, nil
}

func expandAppEngineVersionVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionApiConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	scriptProp, err := expandAppEngineVersionScript(d.Get("script"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("script"); !tpgresource.IsEmptyValue(reflect.ValueOf(scriptProp)) && (ok || !reflect.DeepEqual(v, scriptProp)) {
		obj["script"] = scriptProp
	}

	urlProp, err := expandAppEngineUrl(d.Get("url"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("url"); !tpgresource.IsEmptyValue(reflect.ValueOf(urlProp)) && (ok || !reflect.DeepEqual(v, urlProp)) {
		obj["url"] = urlProp
	}

	return obj, nil
}

func expandAppEngineUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineVersionDeployment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	zipProp, err := expandAppEngineVersionZip(d.Get("zip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zip"); !tpgresource.IsEmptyValue(reflect.ValueOf(zipProp)) && (ok || !reflect.DeepEqual(v, zipProp)) {
		obj["zip"] = zipProp
	}

	filesProp, err := expandAppEngineFiles(d.Get("files"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("files"); !tpgresource.IsEmptyValue(reflect.ValueOf(filesProp)) && (ok || !reflect.DeepEqual(v, filesProp)) {
		obj["files"] = filesProp
	}

	return obj, nil
}

func expandAppEngineVersionZip(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	sourceUrlProp, err := expandAppEngineVersionSourceUrl(d.Get("sourceUrl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sourceUrl"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceUrlProp)) && (ok || !reflect.DeepEqual(v, sourceUrlProp)) {
		obj["sourceUrl"] = sourceUrlProp
	}

	filesCountProp, err := expandAppEngineFilesCount(d.Get("filesCount"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filesCount"); !tpgresource.IsEmptyValue(reflect.ValueOf(filesCountProp)) && (ok || !reflect.DeepEqual(v, filesCountProp)) {
		obj["filesCount"] = filesCountProp
	}

	return obj, nil
}

func expandAppEngineVersionSourceUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFilesCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFiles(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	obj := make(map[string]interface{})

	sourceUrlProp, err := expandAppEngineVersionSourceUrl(d.Get("sourceUrl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sourceUrl"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceUrlProp)) && (ok || !reflect.DeepEqual(v, sourceUrlProp)) {
		obj["sourceUrl"] = sourceUrlProp
	}

	filesCountProp, err := expandAppEngineFilesCount(d.Get("filesCount"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filesCount"); !tpgresource.IsEmptyValue(reflect.ValueOf(filesCountProp)) && (ok || !reflect.DeepEqual(v, filesCountProp)) {
		obj["filesCount"] = filesCountProp
	}

	return obj, nil
}
