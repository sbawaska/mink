/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vaikas/postgressource/pkg/apis/sources/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PostgresSourceLister helps list PostgresSources.
type PostgresSourceLister interface {
	// List lists all PostgresSources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PostgresSource, err error)
	// PostgresSources returns an object that can list and get PostgresSources.
	PostgresSources(namespace string) PostgresSourceNamespaceLister
	PostgresSourceListerExpansion
}

// postgresSourceLister implements the PostgresSourceLister interface.
type postgresSourceLister struct {
	indexer cache.Indexer
}

// NewPostgresSourceLister returns a new PostgresSourceLister.
func NewPostgresSourceLister(indexer cache.Indexer) PostgresSourceLister {
	return &postgresSourceLister{indexer: indexer}
}

// List lists all PostgresSources in the indexer.
func (s *postgresSourceLister) List(selector labels.Selector) (ret []*v1alpha1.PostgresSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PostgresSource))
	})
	return ret, err
}

// PostgresSources returns an object that can list and get PostgresSources.
func (s *postgresSourceLister) PostgresSources(namespace string) PostgresSourceNamespaceLister {
	return postgresSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PostgresSourceNamespaceLister helps list and get PostgresSources.
type PostgresSourceNamespaceLister interface {
	// List lists all PostgresSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PostgresSource, err error)
	// Get retrieves the PostgresSource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PostgresSource, error)
	PostgresSourceNamespaceListerExpansion
}

// postgresSourceNamespaceLister implements the PostgresSourceNamespaceLister
// interface.
type postgresSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PostgresSources in the indexer for a given namespace.
func (s postgresSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PostgresSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PostgresSource))
	})
	return ret, err
}

// Get retrieves the PostgresSource from the indexer for a given namespace and name.
func (s postgresSourceNamespaceLister) Get(name string) (*v1alpha1.PostgresSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("postgressource"), name)
	}
	return obj.(*v1alpha1.PostgresSource), nil
}
