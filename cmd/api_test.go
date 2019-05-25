package cmd

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMD5(t *testing.T) {
	a := HpArray{
		user:     "user",
		password: "password",
	}

	expected := "2b7cc318da9ba9d03912592c2f34a1ec"

	if generateMD5(a.user, a.password) != expected {
		t.Log("MD5 sum was incorrect")
		t.Fail()
	}
}

func TestGetSessionKey(t *testing.T) {
	expectedKey := "GSESS0005889fdb54118f"
	f, err := ioutil.ReadFile("testdata/sessionKeyResponse.xml")
	if err != nil {
		t.Log("No session Key response testdata file")
	}
	testServer :=
		httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(f)
			}),
		)
	defer testServer.Close()

	a := HpArray{
		URL:      testServer.URL,
		user:     "user",
		password: "password",
		Client:   testServer.Client(),
	}

	a.getSessionKey()
	if a.sessionKey != expectedKey {
		t.Errorf("Session Key incorrectly parsed")
	}
}

func TestShowHosts(t *testing.T) {
	expected := []string{"TestHost1", "TestHost2"}
	f, err := ioutil.ReadFile("testdata/hostsResponse.xml")
	if err != nil {
		t.Log("No hosts response testdata file")
	}
	testServer :=
		httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(f)
			}),
		)
	defer testServer.Close()

	a := HpArray{
		URL:      testServer.URL,
		user:     "user",
		password: "password",
		Client:   testServer.Client(),
	}

	result, _ := a.ShowHosts()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Host list incorrectly parsed")
	}
}
func TestShowAllVolumes(t *testing.T) {
	expected := []string{"test_volume_a", "test_volume_b", "test_snapshot_volume_c"}
	f, err := ioutil.ReadFile("testdata/volumesResponse.xml")
	if err != nil {
		t.Log("No volumes response testdata file")
	}
	testServer :=
		httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(f)
			}),
		)
	defer testServer.Close()

	a := HpArray{
		URL:      testServer.URL,
		user:     "user",
		password: "password",
		Client:   testServer.Client(),
	}

	result, _ := a.ShowVolumes("all")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Volumes list incorrectly parsed, got %v expected %v", result, expected)
	}
}

func TestShowAllVolumesSnapshots(t *testing.T) {
	expected := []string{"test_snapshot_volume_c"}
	f, err := ioutil.ReadFile("testdata/volumesResponse.xml")
	if err != nil {
		t.Log("No volumes response testdata file")
	}
	testServer :=
		httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(f)
			}),
		)
	defer testServer.Close()

	a := HpArray{
		URL:      testServer.URL,
		user:     "user",
		password: "password",
		Client:   testServer.Client(),
	}

	result, _ := a.ShowVolumes("snapshot")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Spanshot list invalid, got %v expected %v", result, expected)
	}
}
func TestShowAllVolumesVolumes(t *testing.T) {
	expected := []string{"test_volume_a", "test_volume_b"}
	f, err := ioutil.ReadFile("testdata/volumesResponse.xml")
	if err != nil {
		t.Log("No volumes response testdata file")
	}
	testServer :=
		httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write(f)
			}),
		)
	defer testServer.Close()

	a := HpArray{
		URL:      testServer.URL,
		user:     "user",
		password: "password",
		Client:   testServer.Client(),
	}

	result, _ := a.ShowVolumes("volume")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Volume list invalid, got %v expected %v", result, expected)
	}
}
