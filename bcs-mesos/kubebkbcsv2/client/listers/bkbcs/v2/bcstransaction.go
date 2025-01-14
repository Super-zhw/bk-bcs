/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BcsTransactionLister helps list BcsTransactions.
type BcsTransactionLister interface {
	// List lists all BcsTransactions in the indexer.
	List(selector labels.Selector) (ret []*v2.BcsTransaction, err error)
	// BcsTransactions returns an object that can list and get BcsTransactions.
	BcsTransactions(namespace string) BcsTransactionNamespaceLister
	BcsTransactionListerExpansion
}

// bcsTransactionLister implements the BcsTransactionLister interface.
type bcsTransactionLister struct {
	indexer cache.Indexer
}

// NewBcsTransactionLister returns a new BcsTransactionLister.
func NewBcsTransactionLister(indexer cache.Indexer) BcsTransactionLister {
	return &bcsTransactionLister{indexer: indexer}
}

// List lists all BcsTransactions in the indexer.
func (s *bcsTransactionLister) List(selector labels.Selector) (ret []*v2.BcsTransaction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsTransaction))
	})
	return ret, err
}

// BcsTransactions returns an object that can list and get BcsTransactions.
func (s *bcsTransactionLister) BcsTransactions(namespace string) BcsTransactionNamespaceLister {
	return bcsTransactionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BcsTransactionNamespaceLister helps list and get BcsTransactions.
type BcsTransactionNamespaceLister interface {
	// List lists all BcsTransactions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.BcsTransaction, err error)
	// Get retrieves the BcsTransaction from the indexer for a given namespace and name.
	Get(name string) (*v2.BcsTransaction, error)
	BcsTransactionNamespaceListerExpansion
}

// bcsTransactionNamespaceLister implements the BcsTransactionNamespaceLister
// interface.
type bcsTransactionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BcsTransactions in the indexer for a given namespace.
func (s bcsTransactionNamespaceLister) List(selector labels.Selector) (ret []*v2.BcsTransaction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsTransaction))
	})
	return ret, err
}

// Get retrieves the BcsTransaction from the indexer for a given namespace and name.
func (s bcsTransactionNamespaceLister) Get(name string) (*v2.BcsTransaction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("bcstransaction"), name)
	}
	return obj.(*v2.BcsTransaction), nil
}
