// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/console/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConsoleCLIDownloadLister helps list ConsoleCLIDownloads.
type ConsoleCLIDownloadLister interface {
	// List lists all ConsoleCLIDownloads in the indexer.
	List(selector labels.Selector) (ret []*v1.ConsoleCLIDownload, err error)
	// Get retrieves the ConsoleCLIDownload from the index for a given name.
	Get(name string) (*v1.ConsoleCLIDownload, error)
	ConsoleCLIDownloadListerExpansion
}

// consoleCLIDownloadLister implements the ConsoleCLIDownloadLister interface.
type consoleCLIDownloadLister struct {
	indexer cache.Indexer
}

// NewConsoleCLIDownloadLister returns a new ConsoleCLIDownloadLister.
func NewConsoleCLIDownloadLister(indexer cache.Indexer) ConsoleCLIDownloadLister {
	return &consoleCLIDownloadLister{indexer: indexer}
}

// List lists all ConsoleCLIDownloads in the indexer.
func (s *consoleCLIDownloadLister) List(selector labels.Selector) (ret []*v1.ConsoleCLIDownload, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ConsoleCLIDownload))
	})
	return ret, err
}

// Get retrieves the ConsoleCLIDownload from the index for a given name.
func (s *consoleCLIDownloadLister) Get(name string) (*v1.ConsoleCLIDownload, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("consoleclidownload"), name)
	}
	return obj.(*v1.ConsoleCLIDownload), nil
}
