// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2022 Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumIdentityLister helps list CiliumIdentities.
// All objects returned here must be treated as read-only.
type CiliumIdentityLister interface {
	// List lists all CiliumIdentities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumIdentity, err error)
	// Get retrieves the CiliumIdentity from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2.CiliumIdentity, error)
	CiliumIdentityListerExpansion
}

// ciliumIdentityLister implements the CiliumIdentityLister interface.
type ciliumIdentityLister struct {
	indexer cache.Indexer
}

// NewCiliumIdentityLister returns a new CiliumIdentityLister.
func NewCiliumIdentityLister(indexer cache.Indexer) CiliumIdentityLister {
	return &ciliumIdentityLister{indexer: indexer}
}

// List lists all CiliumIdentities in the indexer.
func (s *ciliumIdentityLister) List(selector labels.Selector) (ret []*v2.CiliumIdentity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumIdentity))
	})
	return ret, err
}

// Get retrieves the CiliumIdentity from the index for a given name.
func (s *ciliumIdentityLister) Get(name string) (*v2.CiliumIdentity, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("ciliumidentity"), name)
	}
	return obj.(*v2.CiliumIdentity), nil
}
