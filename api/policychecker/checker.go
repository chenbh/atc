package policychecker

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sigs.k8s.io/yaml"
	"encoding/json"

	"github.com/chenbh/concourse/atc/api/accessor"
	"github.com/chenbh/concourse/atc/policy"
)

//go:generate counterfeiter . PolicyChecker

type PolicyChecker interface {
	Check(string, accessor.Access, *http.Request) (bool, error)
}

type checker struct {
	policyChecker *policy.Checker
}

func NewApiPolicyChecker(policyChecker *policy.Checker) PolicyChecker {
	if policyChecker == nil {
		return nil
	}
	return &checker{policyChecker: policyChecker}
}

func (c *checker) Check(action string, acc accessor.Access, req *http.Request) (bool, error) {
	// Ignore self invoked API calls.
	if acc.IsSystem() {
		return true, nil
	}

	// Actions in black will not go through policy check.
	if c.policyChecker.ShouldSkipAction(action) {
		return true, nil
	}

	// Only actions with specified http method will go through policy check.
	// But actions in white list will always go through policy check.
	if !c.policyChecker.ShouldCheckHttpMethod(req.Method) &&
		!c.policyChecker.ShouldCheckAction(action) {
		return true, nil
	}

	input := policy.PolicyCheckInput{
		HttpMethod:     req.Method,
		Action:         action,
		User:           acc.Claims().UserName,
		Team:           req.FormValue(":team_name"),
		Pipeline:       req.FormValue(":pipeline_name"),
	}

	switch ct := req.Header.Get("Content-type"); ct {
	case "application/json", "text/vnd.yaml", "text/yaml", "text/x-yaml", "application/x-yaml":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return false, err
		} else if body != nil && len(body) > 0 {
			if ct == "application/json" {
				err = json.Unmarshal(body, &input.Data)
			} else {
				err = yaml.Unmarshal(body, &input.Data)
			}
			if err != nil {
				return false, err
			}

			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}

	return c.policyChecker.Check(input)
}
