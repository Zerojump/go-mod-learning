package rocket_demo

import (
	"fmt"
	"go13-learning/src/commons"
	"gopkg.in/couchbase/gocb.v1"
	"os"
	"testing"
)

var bucket *gocb.Bucket
var cluster *gocb.Cluster

func TestMain(m *testing.M) {
	fmt.Println("begin")
	clusterTmp, err := gocb.Connect("couchbase://192.168.204.128")
	commons.FailOnError(err, "connect couchbase server err")
	cluster = clusterTmp
	err = cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "123456",
	})
	commons.FailOnError(err, "auth err")

	tmp, err := cluster.OpenBucket("bk_test_001", "")
	commons.FailOnError(err, "couchbase cluster open tmp err")

	manager := tmp.Manager("", "")
	err = manager.CreatePrimaryIndex("", true, false)
	commons.FailOnError(err, "create primary index err")

	bucket = tmp
	exitCode := m.Run()
	fmt.Println("end")
	// 退出
	os.Exit(exitCode)
}

func TestHello(t *testing.T) {
	upsert, err := bucket.Upsert("u:kingarthur", User{
		Id:        "kingarthur",
		Email:     "kingarthur@couchbase.com",
		Interests: []string{"Holy Grail", "African Swalllows"},
	}, 0)
	commons.FailOnError(err, "upsert err")
	t.Log("upsert: ", upsert)

	var inUser User
	get, err := bucket.Get("u:kingarthur", &inUser)
	commons.FailOnError(err, "get err")
	t.Log("User:", inUser, "get: ", get)

	rows, err := bucket.ExecuteN1qlQuery(gocb.NewN1qlQuery("select * from bk_test_001 where $1 in interests"), []interface{}{"African Swallows"})
	commons.FailOnError(err, "exec n1ql err")
	var row interface{}
	for rows.Next(&row) {
		t.Log("Row: ", row)

	}
}

func TestCounter(t *testing.T) {
	initVal, cas, err := bucket.Counter("counter_id", 20, 100, 0)
	commons.FailOnError(err, "counter err")
	t.Log("initVal: ", initVal, " cas: ", cas)

	initVal, cas, err = bucket.Counter("counter_id", 1, 100, 0)
	commons.FailOnError(err, "counter err")
	t.Log("initVal: ", initVal, " cas: ", cas)

	initVal, cas, err = bucket.Counter("counter_id", -50, 100, 0)
	commons.FailOnError(err, "counter err")
	t.Log("initVal: ", initVal, " cas: ", cas)
}

func TestIndexConsistency(t *testing.T) {
	_, err := bucket.Insert("user:bob", User{
		Id:        "bob",
		Email:     "bob@couchbase.com",
		Interests: []string{"Holy Grail", "African Swalllows"},
	}, 180)
	commons.FailOnError(err, "upsert err")

	//query := gocb.NewN1qlQuery("SELECT uid, email FROM bk_test_001 WHERE uid = $1")
	//rows, err := cluster.ExecuteN1qlQuery(query,[]interface{}{"alice"})
	query := gocb.NewN1qlQuery("SELECT uid, email FROM bk_test_001") //.Consistency(gocb.RequestPlus)
	rows, err := cluster.ExecuteN1qlQuery(query, nil)
	commons.FailOnError(err, "exec n1ql err")
	var row interface{}
	for rows.Next(&row) {
		t.Log("Row: ", row)

	}
}

func TestAnalytics(t *testing.T) {
	//query := gocb.NewAnalyticsQuery("select experience, level from `gamesim-sample` where name = $name")
	//params := make(map[string]interface{})
	//params["name"] = "Aaron0"
	//rows, err := bucket.ExecuteAnalyticsQuery(query, params)
	//commons.FailOnError(err, "exec n1ql err")
	//var row interface{}
	//for rows.Next(&row) {
	//	t.Log("Row: ", row)
	//
	//}
	query := gocb.NewAnalyticsQuery("select email, interests from `bk_test_001` where uid = $uid")
	params := make(map[string]interface{})
	params["uid"] = "kingarthur"
	rows, err := bucket.ExecuteAnalyticsQuery(query, params)
	commons.FailOnError(err, "exec n1ql err")
	var row interface{}
	for rows.Next(&row) {
		t.Log("Row: ", row)

	}
}
