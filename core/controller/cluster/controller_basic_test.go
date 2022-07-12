package cluster

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"testing"
	"time"

	"g.hz.netease.com/horizon/core/common"
	"g.hz.netease.com/horizon/lib/q"
	mockcd "g.hz.netease.com/horizon/mock/pkg/cluster/cd"
	appmodels "g.hz.netease.com/horizon/pkg/application/models"
	applicationservice "g.hz.netease.com/horizon/pkg/application/service"
	userauth "g.hz.netease.com/horizon/pkg/authentication/user"
	clustermodels "g.hz.netease.com/horizon/pkg/cluster/models"
	envmodels "g.hz.netease.com/horizon/pkg/environmentregion/models"
	groupmodels "g.hz.netease.com/horizon/pkg/group/models"
	groupservice "g.hz.netease.com/horizon/pkg/group/service"
	harbordao "g.hz.netease.com/horizon/pkg/harbor/dao"
	harbormodels "g.hz.netease.com/horizon/pkg/harbor/models"
	membermodels "g.hz.netease.com/horizon/pkg/member/models"
	regionmodels "g.hz.netease.com/horizon/pkg/region/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListClusterByNameFuzzily(t *testing.T) {
	// init data
	var groups []*groupmodels.Group
	for i := 0; i < 5; i++ {
		name := "groupForClusterFuzzily" + strconv.Itoa(i)
		group, err := manager.GroupManager.Create(ctx, &groupmodels.Group{
			Name:     name,
			Path:     name,
			ParentID: 0,
		})
		assert.Nil(t, err)
		assert.NotNil(t, group)
		groups = append(groups, group)
	}

	var applications []*appmodels.Application
	for i := 0; i < 5; i++ {
		group := groups[i]
		name := "appForClusterFuzzily" + strconv.Itoa(i)
		application, err := manager.ApplicationManager.Create(ctx, &appmodels.Application{
			GroupID:         group.ID,
			Name:            name,
			Priority:        "P3",
			GitURL:          "ssh://git.com",
			GitSubfolder:    "/test",
			GitRef:          "master",
			Template:        "javaapp",
			TemplateRelease: "v1.0.0",
		}, nil)
		assert.Nil(t, err)
		assert.NotNil(t, application)
		applications = append(applications, application)
	}

	region, err := manager.RegionMgr.Create(ctx, &regionmodels.Region{
		Name:        "hzFuzzily",
		DisplayName: "HZFuzzily",
	})
	assert.Nil(t, err)
	assert.NotNil(t, region)

	for i := 0; i < 5; i++ {
		application := applications[i]
		name := "fuzzilyCluster" + strconv.Itoa(i)
		cluster, err := manager.ClusterMgr.Create(ctx, &clustermodels.Cluster{
			ApplicationID:   application.ID,
			Name:            name,
			EnvironmentName: "testFuzzily",
			RegionName:      "hzFuzzily",
		}, nil, nil)
		assert.Nil(t, err)
		assert.NotNil(t, cluster)
	}

	c = &controller{
		clusterMgr:     manager.ClusterMgr,
		applicationMgr: manager.ApplicationManager,
		applicationSvc: applicationservice.NewService(groupservice.NewService(manager), manager),
		groupManager:   manager.GroupManager,
		memberManager:  manager.MemberManager,
	}

	count, resps, err := c.ListClusterByNameFuzzily(ctx, "", "fuzzilyCluster", nil)
	assert.Nil(t, err)
	assert.Equal(t, 5, count)
	assert.Equal(t, "fuzzilyCluster4", resps[0].Name)
	assert.Equal(t, "fuzzilyCluster3", resps[1].Name)
	assert.Equal(t, "fuzzilyCluster0", resps[4].Name)
	for _, resp := range resps {
		b, _ := json.Marshal(resp)
		t.Logf("%v", string(b))
	}
}

func TestListUserClustersByNameFuzzily(t *testing.T) {
	// init data
	region, err := manager.RegionMgr.Create(ctx, &regionmodels.Region{
		Name:        "hzUserClustersFuzzily",
		DisplayName: "HZUserClusters",
	})
	assert.Nil(t, err)
	assert.NotNil(t, region)

	er, err := manager.EnvironmentRegionMgr.CreateEnvironmentRegion(ctx, &envmodels.EnvironmentRegion{
		EnvironmentName: "testUserClustersFuzzily",
		RegionName:      "hzUserClustersFuzzily",
	})
	assert.Nil(t, err)

	var groups []*groupmodels.Group
	for i := 0; i < 5; i++ {
		name := "groupForUserClusterFuzzily" + strconv.Itoa(i)
		group, err := manager.GroupManager.Create(ctx, &groupmodels.Group{
			Name:     name,
			Path:     name,
			ParentID: 0,
		})
		assert.Nil(t, err)
		assert.NotNil(t, group)
		groups = append(groups, group)
	}

	var applications []*appmodels.Application
	for i := 0; i < 5; i++ {
		group := groups[i]
		name := "appForUserClusterFuzzily" + strconv.Itoa(i)
		application, err := manager.ApplicationManager.Create(ctx, &appmodels.Application{
			GroupID:         group.ID,
			Name:            name,
			Priority:        "P3",
			GitURL:          "ssh://git.com",
			GitSubfolder:    "/test",
			GitRef:          "master",
			Template:        "javaapp",
			TemplateRelease: "v1.0.0",
		}, nil)
		assert.Nil(t, err)
		assert.NotNil(t, application)
		applications = append(applications, application)
	}

	var clusters []*clustermodels.Cluster
	for i := 0; i < 5; i++ {
		application := applications[i]
		name := "userClusterFuzzily" + strconv.Itoa(i)
		cluster, err := manager.ClusterMgr.Create(ctx, &clustermodels.Cluster{
			ApplicationID:   application.ID,
			Name:            name,
			EnvironmentName: "testUserClustersFuzzily",
			RegionName:      "hzUserClustersFuzzily",
			GitURL:          "ssh://git@g.hz.netease.com:22222/music-cloud-native/horizon/horizon.git",
		}, nil, nil)
		assert.Nil(t, err)
		assert.NotNil(t, cluster)
		clusters = append(clusters, cluster)
	}

	// nolint
	ctx = context.WithValue(ctx, common.UserContextKey(), &userauth.DefaultInfo{
		Name: "Matt",
		ID:   uint(2),
	})
	_, err = manager.MemberManager.Create(ctx, &membermodels.Member{
		ResourceType: membermodels.TypeGroup,
		ResourceID:   groups[0].ID,
		Role:         "owner",
		MemberType:   membermodels.MemberUser,
		MemberNameID: 2,
	})
	assert.Nil(t, err)

	_, err = manager.MemberManager.Create(ctx, &membermodels.Member{
		ResourceType: membermodels.TypeApplication,
		ResourceID:   applications[1].ID,
		Role:         "owner",
		MemberType:   membermodels.MemberUser,
		MemberNameID: 2,
	})
	assert.Nil(t, err)

	_, err = manager.MemberManager.Create(ctx, &membermodels.Member{
		ResourceType: membermodels.TypeApplicationCluster,
		ResourceID:   clusters[3].ID,
		Role:         "owner",
		MemberType:   membermodels.MemberUser,
		MemberNameID: 2,
	})
	assert.Nil(t, err)

	c = &controller{
		clusterMgr:     manager.ClusterMgr,
		applicationMgr: manager.ApplicationManager,
		applicationSvc: applicationservice.NewService(groupservice.NewService(manager), manager),
		groupManager:   manager.GroupManager,
		memberManager:  manager.MemberManager,
	}

	count, resps, err := c.ListUserClusterByNameFuzzily(ctx, er.EnvironmentName, "cluster", nil)
	assert.Nil(t, err)
	assert.Equal(t, 3, count)
	assert.Equal(t, "userClusterFuzzily3", resps[0].Name)
	assert.Equal(t, "userClusterFuzzily1", resps[1].Name)
	assert.Equal(t, "userClusterFuzzily0", resps[2].Name)
	for _, resp := range resps {
		b, _ := json.Marshal(resp)
		t.Logf("%v", string(b))
	}

	count, resps, err = c.ListUserClusterByNameFuzzily(ctx, er.EnvironmentName, "userCluster", &q.Query{
		PageSize: 2,
	})
	assert.Nil(t, err)
	assert.Equal(t, 3, count)
	assert.Equal(t, "userClusterFuzzily3", resps[0].Name)
	assert.Equal(t, "userClusterFuzzily1", resps[1].Name)
	for _, resp := range resps {
		b, _ := json.Marshal(resp)
		t.Logf("%v", string(b))
	}
}

func TestController_FreeOrDeleteClusterFailed(t *testing.T) {
	mockCtl := gomock.NewController(t)
	cd := mockcd.NewMockCD(mockCtl)
	cd.EXPECT().DeleteCluster(gomock.Any(), gomock.Any()).Return(errors.New("test")).AnyTimes()

	c = &controller{
		cd:             cd,
		clusterMgr:     manager.ClusterMgr,
		applicationMgr: manager.ApplicationManager,
		applicationSvc: applicationservice.NewService(groupservice.NewService(manager), manager),
		groupManager:   manager.GroupManager,
		envMgr:         manager.EnvMgr,
		regionMgr:      manager.RegionMgr,
	}

	id, err := harbordao.NewDAO(db).Create(ctx, &harbormodels.Harbor{
		Server: "http://127.0.0.1",
	})
	assert.Nil(t, err)
	region, err := manager.RegionMgr.Create(ctx, &regionmodels.Region{
		Name:        "TestController_FreeOrDeleteClusterFailed",
		DisplayName: "TestController_FreeOrDeleteClusterFailed",
		HarborID:    id,
	})
	assert.Nil(t, err)
	assert.NotNil(t, region)

	group, err := manager.GroupManager.Create(ctx, &groupmodels.Group{
		Name:     "TestController_FreeOrDeleteClusterFailed",
		Path:     "/TestController_FreeOrDeleteClusterFailed",
		ParentID: 0,
	})
	assert.Nil(t, err)
	assert.NotNil(t, group)

	application, err := manager.ApplicationManager.Create(ctx, &appmodels.Application{
		GroupID:         group.ID,
		Name:            "TestController_FreeOrDeleteClusterFailed",
		Priority:        "P3",
		GitURL:          "ssh://git.com",
		GitSubfolder:    "/test",
		GitRef:          "master",
		Template:        "javaapp",
		TemplateRelease: "v1.0.0",
	}, nil)
	assert.Nil(t, err)
	assert.NotNil(t, application)

	cluster, err := manager.ClusterMgr.Create(ctx, &clustermodels.Cluster{
		ApplicationID:   application.ID,
		Name:            "TestController_FreeOrDeleteClusterFailed",
		EnvironmentName: "TestController_FreeOrDeleteClusterFailed",
		RegionName:      region.Name,
		GitURL:          "",
	}, nil, nil)
	assert.Nil(t, err)
	assert.NotNil(t, cluster)

	// if failed to free, status should be set to empty
	err = c.FreeCluster(ctx, cluster.ID)
	assert.Nil(t, err)
	time.Sleep(time.Second)
	cluster, err = manager.ClusterMgr.GetByID(ctx, cluster.ID)
	assert.Nil(t, err)
	assert.Equal(t, "", cluster.Status)

	// if failed to delete, status should be set to empty
	err = c.DeleteCluster(ctx, cluster.ID)
	assert.Nil(t, err)
	time.Sleep(time.Second)
	cluster, err = manager.ClusterMgr.GetByID(ctx, cluster.ID)
	assert.Nil(t, err)
	assert.Equal(t, "", cluster.Status)
}
