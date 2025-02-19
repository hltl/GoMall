package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	// 服务注册的TTL时间
	serviceRegisterTTL = 15
)

// ServiceInstance 表示一个服务实例
type ServiceInstance struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Address  string            `json:"address"`
	Port     int               `json:"port"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ServiceRegistry 服务注册接口
type ServiceRegistry interface {
	// Register 注册服务实例
	Register(ctx context.Context, service *ServiceInstance) error
	// Deregister 注销服务实例
	Deregister(ctx context.Context, service *ServiceInstance) error
	// Close 关闭注册中心连接
	Close() error
}

// EtcdRegistry etcd实现的服务注册中心
type EtcdRegistry struct {
	client        *clientv3.Client
	leasesID      clientv3.LeaseID
	keepAliveCh   <-chan *clientv3.LeaseKeepAliveResponse
	keepAliveDone chan struct{}
}

// NewEtcdRegistry 创建etcd注册中心实例
func NewEtcdRegistry(endpoints []string) (*EtcdRegistry, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("create etcd client failed: %v", err)
	}

	return &EtcdRegistry{
		client:        client,
		keepAliveDone: make(chan struct{}),
	}, nil
}

// Register 注册服务
func (r *EtcdRegistry) Register(ctx context.Context, service *ServiceInstance) error {
	// 创建租约
	resp, err := r.client.Grant(ctx, serviceRegisterTTL)
	if err != nil {
		logrus.WithError(err).Error("创建etcd租约失败")
		return fmt.Errorf("create lease failed: %v", err)
	}
	r.leasesID = resp.ID
	logrus.WithField("lease_id", resp.ID).Debug("成功创建etcd租约")

	// 设置续租
	keepAliveCh, err := r.client.KeepAlive(ctx, resp.ID)
	if err != nil {
		logrus.WithError(err).Error("设置etcd续租失败")
		return fmt.Errorf("set keep alive failed: %v", err)
	}
	r.keepAliveCh = keepAliveCh
	logrus.Debug("成功设置etcd续租")

	// 监控续租情况
	go func() {
		for {
			select {
			case resp := <-r.keepAliveCh:
				if resp == nil {
					logrus.Warn("服务续租已停止")
					return
				}
				logrus.WithFields(logrus.Fields{
					"lease_id": resp.ID,
					"ttl":      resp.TTL,
				}).Debug("成功续租服务实例")
			case <-r.keepAliveDone:
				logrus.Info("服务注册监控已关闭")
				return
			}
		}
	}()

	// 注册服务实例
	key := fmt.Sprintf("/services/%s/%s", service.Name, service.ID)
	value, _ := json.Marshal(service)

	logrus.WithFields(logrus.Fields{
		"key":   key,
		"value": string(value),
	}).Debug("正在注册服务到etcd")

	_, err = r.client.Put(ctx, key, string(value), clientv3.WithLease(r.leasesID))
	if err != nil {
		logrus.WithError(err).Error("注册服务到etcd失败")
		return fmt.Errorf("put service instance failed: %v", err)
	}

	logrus.WithFields(logrus.Fields{
		"service_id":   service.ID,
		"service_name": service.Name,
		"address":      service.Address,
		"port":         service.Port,
	}).Info("服务实例注册成功")

	return nil
}

// Deregister 注销服务
func (r *EtcdRegistry) Deregister(ctx context.Context, service *ServiceInstance) error {
	close(r.keepAliveDone)

	key := fmt.Sprintf("/services/%s/%s", service.Name, service.ID)
	_, err := r.client.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("deregister service instance failed: %v", err)
	}

	logrus.WithFields(logrus.Fields{
		"service_id":   service.ID,
		"service_name": service.Name,
	}).Info("服务实例注销成功")

	return nil
}

// Close 关闭注册中心连接
func (r *EtcdRegistry) Close() error {
	close(r.keepAliveDone)
	return r.client.Close()
}
