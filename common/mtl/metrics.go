// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mtl

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Registry 注册中心，用于注册指标
var Registry *prometheus.Registry

// InitMetric 初始化指标
func InitMetric(serviceName string, metricsPort string, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	// 注册 go 运行时相关指标
	Registry.MustRegister(collectors.NewGoCollector())
	// 注册进程相关指标
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	// Prometheus 注册到 Consul 注册中心
	r, _ := consul.NewConsulRegister(registryAddr)

	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	// 构建注册信息
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registryInfo)
	// 注册关闭 Hook 函数，在服务关闭时注销注册信息
	server.RegisterShutdownHook(func() {
		err := r.Deregister(registryInfo)
		if err != nil {
			klog.Fatal(err)
			return
		}
	})
	// 启动 HTTP 服务，暴露指标
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	// 异步启动 Server
	go http.ListenAndServe(metricsPort, nil) //nolint:errcheck
	return r, registryInfo
}
