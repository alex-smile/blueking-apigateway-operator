debug: true

operator:
  withKube: true
  withLeader: true
  agentMode: true

  defaultGateway: "bk-default"
  defaultStage: "default"

dashboard:
  etcd:
    endpoints: "bk-apigateway-etcd:2379"
    keyPrefix: "/bk-gateway-apigw/default"
    username: "root"
    password: "blueking"

apisix:
  etcd:
    endpoints: "bk-apigateway-etcd:2379"
    keyPrefix: "/bk-gateway-apisix"
    username: "root"
    password: "blueking"
  resourceStoreMode: "etcd"
  virtualStage:
    operatorExternalHost: "bk-apigateway-operator"
    operatorExternalHealthProbePort: 6004
    extraApisixResources: "/data/config/extra-resources.yaml"

httpServer:
  bindAddress: "0.0.0.0"
  bindAddressV6: "[::]"
  bindPort: 6004

logger:
  default:
    level: info
    writer: os
    settings: {name: stdout}
  controller:
    level: info
    writer: os
    settings: {name: stdout}
    # writer: file
    # settings: {name: operator_system.log, size: 100, backups: 10, age: 7, path: /}