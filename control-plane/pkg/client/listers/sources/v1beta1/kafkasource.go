/*
 * Copyright 2021 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1beta1 "knative.dev/eventing-kafka-broker/control-plane/pkg/apis/sources/v1beta1"
)

// KafkaSourceLister helps list KafkaSources.
// All objects returned here must be treated as read-only.
type KafkaSourceLister interface {
	// List lists all KafkaSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.KafkaSource, err error)
	// KafkaSources returns an object that can list and get KafkaSources.
	KafkaSources(namespace string) KafkaSourceNamespaceLister
	KafkaSourceListerExpansion
}

// kafkaSourceLister implements the KafkaSourceLister interface.
type kafkaSourceLister struct {
	listers.ResourceIndexer[*v1beta1.KafkaSource]
}

// NewKafkaSourceLister returns a new KafkaSourceLister.
func NewKafkaSourceLister(indexer cache.Indexer) KafkaSourceLister {
	return &kafkaSourceLister{listers.New[*v1beta1.KafkaSource](indexer, v1beta1.Resource("kafkasource"))}
}

// KafkaSources returns an object that can list and get KafkaSources.
func (s *kafkaSourceLister) KafkaSources(namespace string) KafkaSourceNamespaceLister {
	return kafkaSourceNamespaceLister{listers.NewNamespaced[*v1beta1.KafkaSource](s.ResourceIndexer, namespace)}
}

// KafkaSourceNamespaceLister helps list and get KafkaSources.
// All objects returned here must be treated as read-only.
type KafkaSourceNamespaceLister interface {
	// List lists all KafkaSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.KafkaSource, err error)
	// Get retrieves the KafkaSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.KafkaSource, error)
	KafkaSourceNamespaceListerExpansion
}

// kafkaSourceNamespaceLister implements the KafkaSourceNamespaceLister
// interface.
type kafkaSourceNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.KafkaSource]
}
