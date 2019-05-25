package cmd

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// HpArray holds array credentials and URL
type HpArray struct {
	URL        string
	user       string
	password   string
	sessionKey string
	Client     *http.Client
}

// getSessionKey sends GET request towards Array API in order to retrieve session key
func (a *HpArray) getSessionKey() {
	var lr LoginResponse

	hash := generateMD5(a.user, a.password)

	url := fmt.Sprintf("%s/api/login/%s", a.URL, hash)
	resp, err := a.Client.Get(url)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Invalid return code from array")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = xml.Unmarshal(bodyBytes, &lr)
	if err != nil {
		log.Fatal("Error:", err)
	}

	for _, name := range lr.OBJECT.PROPERTY {
		if name.Name == "response" {
			a.sessionKey = name.Text
		}
	}
}

// closeSession sends request closing the API session
func (a *HpArray) closeSession() error {
	url := fmt.Sprintf("%s/api/exit", a.URL)
	resp, err := a.Client.Get(url)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Could not close the session")
	}
	return nil

}

// ShowHosts prints out list of hosts
func (a *HpArray) ShowHosts() ([]string, error) {
	var hr HostsResponse
	hosts := []string{}
	url := fmt.Sprintf("%s/api/show/hosts", a.URL)
	req, _ := http.NewRequest("GET", url, nil)
	a.getSessionKey()
	req.Header.Set("sessionKey", a.sessionKey)
	req.Header.Set("dataType", "ipa")
	resp, err := a.Client.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	defer a.closeSession()

	if resp.StatusCode != 200 {
		log.Fatal("Could not get host list")
		return nil, nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = xml.Unmarshal(bodyBytes, &hr)
	if err != nil {
		log.Fatal("Error:", err)
	}

	for _, o := range hr.OBJECT {
		if o.Basetype == "hosts" {
			for _, p := range o.PROPERTY {
				if p.Name == "host-name" {
					hosts = append(hosts, p.Text)
				}
			}
		}
	}
	return hosts, nil
}

// ShowVolumes prints list of volumes to stdout
func (a *HpArray) ShowVolumes(volumeType string) ([]string, error) {
	var vr VolumesResponse
	volumes := []string{}
	url := fmt.Sprintf("%s/api/show/volumes", a.URL)
	req, _ := http.NewRequest("GET", url, nil)
	a.getSessionKey()
	req.Header.Set("sessionKey", a.sessionKey)
	req.Header.Set("dataType", "ipa")
	resp, err := a.Client.Do(req)
	if err != nil {
		log.Fatal("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	defer a.closeSession()

	if resp.StatusCode != 200 {
		log.Fatal("Could not get volumes list")
		return nil, nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = xml.Unmarshal(bodyBytes, &vr)
	if err != nil {
		log.Fatal("Error:", err)
	}

	for _, o := range vr.OBJECT {
		if o.Basetype == "volumes" {

			n := ""
			for _, p := range o.PROPERTY {
				if p.Name == "volume-name" {
					n = p.Text
				}
				if p.Name == "volume-type" && p.Text == "snapshot" {
					if Contains([]string{"snapshot", "all"}, volumeType) {
						volumes = append(volumes, n)
					}
				} else if p.Name == "volume-type" && Contains([]string{"standard", "master volume"}, p.Text) {
					if Contains([]string{"volume", "all"}, volumeType) {
						volumes = append(volumes, n)
					}
				}
			}
		}
	}
	return volumes, nil
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// generateMD5 create hash for array API authentication
func generateMD5(a string, b string) string {
	hs := fmt.Sprintf("%s_%s", a, b)
	h := md5.New()
	io.WriteString(h, hs)
	return fmt.Sprintf("%x", h.Sum(nil))
}
