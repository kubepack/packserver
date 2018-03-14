/*
Copyright 2018 The Kubepack Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	tamal "github.com/kubepack/packserver/apis/tamal"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PackLister helps list Packs.
type PackLister interface {
	// List lists all Packs in the indexer.
	List(selector labels.Selector) (ret []*tamal.Pack, err error)
	// Packs returns an object that can list and get Packs.
	Packs(namespace string) PackNamespaceLister
	PackListerExpansion
}

// packLister implements the PackLister interface.
type packLister struct {
	indexer cache.Indexer
}

// NewPackLister returns a new PackLister.
func NewPackLister(indexer cache.Indexer) PackLister {
	return &packLister{indexer: indexer}
}

// List lists all Packs in the indexer.
func (s *packLister) List(selector labels.Selector) (ret []*tamal.Pack, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*tamal.Pack))
	})
	return ret, err
}

// Packs returns an object that can list and get Packs.
func (s *packLister) Packs(namespace string) PackNamespaceLister {
	return packNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PackNamespaceLister helps list and get Packs.
type PackNamespaceLister interface {
	// List lists all Packs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*tamal.Pack, err error)
	// Get retrieves the Pack from the indexer for a given namespace and name.
	Get(name string) (*tamal.Pack, error)
	PackNamespaceListerExpansion
}

// packNamespaceLister implements the PackNamespaceLister
// interface.
type packNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Packs in the indexer for a given namespace.
func (s packNamespaceLister) List(selector labels.Selector) (ret []*tamal.Pack, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*tamal.Pack))
	})
	return ret, err
}

// Get retrieves the Pack from the indexer for a given namespace and name.
func (s packNamespaceLister) Get(name string) (*tamal.Pack, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(tamal.Resource("pack"), name)
	}
	return obj.(*tamal.Pack), nil
}
