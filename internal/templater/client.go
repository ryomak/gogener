package templater

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"gopkg.in/yaml.v2"
)

type option func(*remoteClient)

type remoteClient struct {
	BaseURL        *url.URL
	SettingYmlName string
	Client         *http.Client
}

type RemoteAppSetting struct {
	Name      string   `yaml:"name"`
	BgPath    string   `yaml:"bg-file-path"`
	Templates []string `yaml:"templates"`
}

func NewRemoteClient(urlStr string, ops ...option) (*remoteClient, error) {
	dir, filename := path.Split(urlStr)
	baseURL, err := url.Parse(dir)
	if err != nil {
		return nil, err
	}
	c := &remoteClient{
		BaseURL:        baseURL,
		SettingYmlName: filename,
		Client:         http.DefaultClient,
	}
	for _, op := range ops {
		op(c)
	}
	return c, nil
}

func WithHttpClient(client *http.Client) option {
	return func(c *remoteClient) {
		c.Client = client
	}
}

func (rc *remoteClient) fetchSettingFromYaml() (*RemoteAppSetting, error) {
	resp, err := rc.Client.Get(fmt.Sprintf("%s/%s", rc.BaseURL.String(), rc.SettingYmlName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ras := &RemoteAppSetting{}
	err = yaml.Unmarshal(data, &ras)
	if err != nil {
		return nil, err
	}
	return ras, nil
}

func (rc *remoteClient) RemoteSettingToAppTemplate() (*AppTemplate, error) {
	setting, err := rc.fetchSettingFromYaml()
	if err != nil {
		return nil, err
	}
	resp, err := rc.Client.Get(fmt.Sprintf("%s/%s", rc.BaseURL.String(), setting.BgPath))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	appTemplate := new(AppTemplate)
	appTemplate.Name = setting.Name
	appTemplate.Bg = string(data)
	templates := make([]*Templater, len(setting.Templates))
	for i, v := range setting.Templates {
		rt := remoteTmpl{
			Client:  rc.Client,
			Path:    v,
			BaseURL: rc.BaseURL,
		}
		dir, filename := path.Split(v)
		templates[i] = &Templater{
			rt,
			dir,
			filename,
		}
	}
	appTemplate.Tmpls = templates
	return appTemplate, nil
}

type remoteTmpl struct {
	Client  *http.Client
	Path    string
	BaseURL *url.URL
}

func (rt remoteTmpl) String() (string, error) {
	resp, err := rt.Client.Get(fmt.Sprintf("%s/%s", rt.BaseURL.String(), rt.Path))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("cannot get this file.")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
