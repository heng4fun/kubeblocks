cluster: {
	metadata: {
		namespace: string
		name:      string
	}
}
component: {
	clusterDefName: string
	clusterType:    string
	type:           string
	name:           string
	service: {
		ports: [...]
		type: string
	}
	podSpec: containers: [...]
	volumeClaimTemplates: [...]
}

service: {
	"apiVersion": "v1"
	"kind":       "Service"
	"metadata": {
		namespace: cluster.metadata.namespace
		name:      "\(cluster.metadata.name)-\(component.name)"
		labels: {
			"app.kubernetes.io/name":     "\(component.clusterType)-\(component.clusterDefName)"
			"app.kubernetes.io/instance": cluster.metadata.name
			// "app.kubernetes.io/version" : # TODO
			"app.kubernetes.io/component-name": "\(component.name)"
		}
	}
	"spec": {
		"selector": {
			"app.kubernetes.io/instance":       "\(cluster.metadata.name)"
			"app.kubernetes.io/component-name": "\(component.name)"
		}
		ports: component.service.ports
		if component.service.type != _|_ {
			type: component.service.type
		}
	}
}
