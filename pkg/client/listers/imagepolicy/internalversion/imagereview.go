/*
Copyright The Kubernetes Authors.

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

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	imagepolicy "k8s.io/kubernetes/pkg/apis/imagepolicy"
)

// ImageReviewLister helps list ImageReviews.
type ImageReviewLister interface {
	// List lists all ImageReviews in the indexer.
	List(selector labels.Selector) (ret []*imagepolicy.ImageReview, err error)
	// Get retrieves the ImageReview from the index for a given name.
	Get(name string) (*imagepolicy.ImageReview, error)
	ImageReviewListerExpansion
}

// imageReviewLister implements the ImageReviewLister interface.
type imageReviewLister struct {
	indexer cache.Indexer
}

// NewImageReviewLister returns a new ImageReviewLister.
func NewImageReviewLister(indexer cache.Indexer) ImageReviewLister {
	return &imageReviewLister{indexer: indexer}
}

// List lists all ImageReviews in the indexer.
func (s *imageReviewLister) List(selector labels.Selector) (ret []*imagepolicy.ImageReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*imagepolicy.ImageReview))
	})
	return ret, err
}

// Get retrieves the ImageReview from the index for a given name.
func (s *imageReviewLister) Get(name string) (*imagepolicy.ImageReview, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(imagepolicy.Resource("imagereview"), name)
	}
	return obj.(*imagepolicy.ImageReview), nil
}
