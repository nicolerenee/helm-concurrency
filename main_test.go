package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/release"
)

var (
	statuses = []release.Status_Code{
		release.Status_UNKNOWN,
		release.Status_DEPLOYED,
		release.Status_DELETED,
		release.Status_DELETING,
		release.Status_FAILED,
		release.Status_PENDING_INSTALL,
		release.Status_PENDING_UPGRADE,
		release.Status_PENDING_ROLLBACK,
	}
	tiller     = "localhost:44134"
	ns         = "nicolerenee"
	helmClient = helm.NewClient(helm.Host(tiller))
	chartDir   = "./chart"
	wg         sync.WaitGroup
)

func releaseExists(t *testing.T, rls string) {
	l, err := helmClient.ListReleases(
		helm.ReleaseListNamespace(ns),
		helm.ReleaseListFilter(rls),
		helm.ReleaseListStatuses(statuses),
	)
	assert.Nil(t, err, "Search failed for %s", rls)
	assert.NotEqual(t, 0, len(l.Releases), "We should find at least 1 release for: %s", rls)

	var r *release.Release
	for _, i := range l.Releases {
		if i.GetName() == rls {
			r = i
		}
	}

	assert.NotNil(t, r, "We should have found release: %s in: %s", rls, l.Releases)
	if r != nil {
		assert.Equal(t, rls, r.Name)
	}
}

func TestAllReleasesExist(t *testing.T) {
	for i := 1; i <= 20; i++ {
		rls := fmt.Sprintf("test%d", i)
		releaseExists(t, rls)
	}
}

func TestAllReleasesExistInGoRoutine(t *testing.T) {
	wg.Add(20)
	for i := 1; i <= 20; i++ {
		rls := fmt.Sprintf("test%d", i)
		go func() {
			defer wg.Done()
			releaseExists(t, rls)
		}()
	}
	wg.Wait()
}
