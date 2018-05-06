package ewok

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testJSON = `
{
  "id": 57451060,
  "name": "hraftd",
  "full_name": "euforia/hraftd",
  "owner": {
    "login": "euforia",
    "id": 788272,
    "avatar_url": "https://avatars2.githubusercontent.com/u/788272?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/euforia",
    "html_url": "https://github.com/euforia",
    "followers_url": "https://api.github.com/users/euforia/followers",
    "following_url": "https://api.github.com/users/euforia/following{/other_user}",
    "gists_url": "https://api.github.com/users/euforia/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/euforia/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/euforia/subscriptions",
    "organizations_url": "https://api.github.com/users/euforia/orgs",
    "repos_url": "https://api.github.com/users/euforia/repos",
    "events_url": "https://api.github.com/users/euforia/events{/privacy}",
    "received_events_url": "https://api.github.com/users/euforia/received_events",
    "type": "User",
    "site_admin": false
  },
  "private": false,
  "html_url": "https://github.com/euforia/hraftd",
  "description": "A reference use of Hashicorp's Raft implementation",
  "fork": true,
  "url": "https://api.github.com/repos/euforia/hraftd",
  "forks_url": "https://api.github.com/repos/euforia/hraftd/forks",
  "keys_url": "https://api.github.com/repos/euforia/hraftd/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/euforia/hraftd/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/euforia/hraftd/teams",
  "hooks_url": "https://api.github.com/repos/euforia/hraftd/hooks",
  "issue_events_url": "https://api.github.com/repos/euforia/hraftd/issues/events{/number}",
  "events_url": "https://api.github.com/repos/euforia/hraftd/events",
  "assignees_url": "https://api.github.com/repos/euforia/hraftd/assignees{/user}",
  "branches_url": "https://api.github.com/repos/euforia/hraftd/branches{/branch}",
  "tags_url": "https://api.github.com/repos/euforia/hraftd/tags",
  "blobs_url": "https://api.github.com/repos/euforia/hraftd/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/euforia/hraftd/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/euforia/hraftd/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/euforia/hraftd/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/euforia/hraftd/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/euforia/hraftd/languages",
  "stargazers_url": "https://api.github.com/repos/euforia/hraftd/stargazers",
  "contributors_url": "https://api.github.com/repos/euforia/hraftd/contributors",
  "subscribers_url": "https://api.github.com/repos/euforia/hraftd/subscribers",
  "subscription_url": "https://api.github.com/repos/euforia/hraftd/subscription",
  "commits_url": "https://api.github.com/repos/euforia/hraftd/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/euforia/hraftd/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/euforia/hraftd/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/euforia/hraftd/issues/comments{/number}",
  "contents_url": "https://api.github.com/repos/euforia/hraftd/contents/{+path}",
  "compare_url": "https://api.github.com/repos/euforia/hraftd/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/euforia/hraftd/merges",
  "archive_url": "https://api.github.com/repos/euforia/hraftd/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/euforia/hraftd/downloads",
  "issues_url": "https://api.github.com/repos/euforia/hraftd/issues{/number}",
  "pulls_url": "https://api.github.com/repos/euforia/hraftd/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/euforia/hraftd/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/euforia/hraftd/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/euforia/hraftd/labels{/name}",
  "releases_url": "https://api.github.com/repos/euforia/hraftd/releases{/id}",
  "deployments_url": "https://api.github.com/repos/euforia/hraftd/deployments",
  "created_at": "2016-04-30T16:17:46Z",
  "updated_at": "2016-06-18T09:37:48Z",
  "pushed_at": "2016-06-22T18:11:16Z",
  "git_url": "git://github.com/euforia/hraftd.git",
  "ssh_url": "git@github.com:euforia/hraftd.git",
  "clone_url": "https://github.com/euforia/hraftd.git",
  "svn_url": "https://github.com/euforia/hraftd",
  "homepage": "http://www.philipotoole.com/building-a-distributed-key-value-store-using-raft/",
  "size": 68,
  "stargazers_count": 1,
  "watchers_count": 1,
  "language": "Go",
  "has_issues": false,
  "has_projects": true,
  "has_downloads": true,
  "has_wiki": true,
  "has_pages": false,
  "forks_count": 0,
  "mirror_url": null,
  "archived": false,
  "open_issues_count": 0,
  "license": {
    "key": "mit",
    "name": "MIT License",
    "spdx_id": "MIT",
    "url": "https://api.github.com/licenses/mit"
  },
  "forks": 0,
  "open_issues": 0,
  "watchers": 1,
  "default_branch": "master"
}`

type testStruct2 struct {
	Labels []string
}

type testStruct1 struct {
	Name string
}

type testStruct struct {
	Map         map[int64]string
	NestedMap   map[string]interface{}
	Desc        *string
	NullPointer *testStruct1
	Pointer     *testStruct1 `pseudo:"pointer"`
	Struct      testStruct2
	Iface       interface{}
}

func testData() *testStruct {
	ii := 10
	data := &testStruct{
		Map: map[int64]string{
			1: "v1",
			2: "v2",
		},
		NestedMap: map[string]interface{}{
			"foo": map[int]interface{}{
				3: testStruct1{},
				4: &testStruct2{[]string{"label"}},
			},
			"bar": map[string]interface{}{
				"five": &testStruct2{},
				"six":  &ii,
			},
		},
		Pointer: &testStruct1{Name: "pointer"},
		Iface:   &testStruct2{},
	}

	return data
}

func Test_Ewok(t *testing.T) {
	d1 := testData()
	c1 := Config{}
	w1 := New(c1)
	w1.Index(d1)
	val1, ok := w1.Get(".Pointer.Name")
	assert.True(t, ok)
	assert.Equal(t, "pointer", val1.Interface().(string))
	c := 0
	w1.Iter(func(key string, value reflect.Value) bool {
		c++
		return false
	})
	assert.Equal(t, 1, c)
}

func Test_Ewok_TrimRoot(t *testing.T) {
	d2 := testData()
	c2 := Config{TrimRoot: true}
	w2 := New(c2)
	w2.Index(d2)
	val2, ok := w2.Get("Map.1")
	assert.True(t, ok)
	assert.Equal(t, "v1", val2.Interface().(string))
}

func Test_Ewok_Delimiter_FieldTag(t *testing.T) {
	d3 := testData()
	c3 := Config{Delimiter: "/", FieldTag: "pseudo"}
	w3 := New(c3)
	w3.Index(d3)
	last := c3.Delimiter
	w3.Iter(func(key string, value reflect.Value) bool {
		assert.Equal(t, uint8('/'), key[0])
		if last > key {
			t.Fatal("Order", last, key)
		}
		last = key
		return true
	})
	val3, ok := w3.Get("/pointer/Name")
	assert.True(t, ok)
	assert.Equal(t, "pointer", val3.Interface().(string))
}

func Test_Ewok_JSON_ScalarsOnly(t *testing.T) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(testJSON), &m)
	assert.Nil(t, err)

	conf := Config{TrimRoot: true, ScalarsOnly: true}
	w := New(conf)
	w.Index(m)

	_, ok := w.Get("owner")
	assert.False(t, ok)
	_, ok = w.Get("license")
	assert.False(t, ok)

	_, ok = w.Get("owner.id")
	assert.True(t, ok)
}
