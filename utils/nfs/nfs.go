package nfs

import (
	"Glue-API/model"
	"Glue-API/utils"
	"encoding/json"
	"errors"
	"os/exec"
)

func NfsClusterCreate(yaml_file string) (output string, err error) {
	var stdCreate []byte
	cmd := exec.Command("ceph", "orch", "apply", "-i", yaml_file)
	stdCreate, err = cmd.CombinedOutput()
	if err != nil {
		err = errors.New(string(stdCreate))
		utils.FancyHandleError(err)
		output = "Fail"
		return
	} else {
		output = "Success"
	}
	return

}
func NfsClusterDelete(cluster_id string) (output string, err error) {
	var stdDelete []byte
	cluster_rm_cmd := exec.Command("ceph", "nfs", "cluster", "rm", cluster_id)
	stdDelete, err = cluster_rm_cmd.CombinedOutput()
	if err != nil {
		err = errors.New(string(stdDelete))
		utils.FancyHandleError(err)
		output = "Fail"
		return
	} else {
		output = "Success"
	}
	return
}
func NfsExportCreateOrUpdate(cluster_id string, json_file string) (output string, err error) {
	var stdCreate []byte
	cmd := exec.Command("ceph", "nfs", "export", "apply", cluster_id, "-i", json_file)
	stdCreate, err = cmd.CombinedOutput()
	if err != nil {
		err = errors.New(string(stdCreate))
		utils.FancyHandleError(err)
		output = "Fail"
		return
	} else {
		output = "Success"
	}
	return
}
func NfsExportDelete(cluster_id string, pseudo string) (output string, err error) {
	var stdCreate []byte

	cmd := exec.Command("ceph", "nfs", "export", "rm", cluster_id, pseudo)
	stdCreate, err = cmd.CombinedOutput()
	if err != nil {
		err = errors.New(string(stdCreate))
		utils.FancyHandleError(err)
		output = "Fail"
		return
	} else {
		output = "Success"
	}
	return
}
func NfsClusterList(cluster_id string) (dat model.NfsClusterList, err error) {
	var stdout []byte
	if cluster_id == "" {
		cmd := exec.Command("ceph", "nfs", "cluster", "info")
		stdout, err = cmd.CombinedOutput()
		if err != nil {
			return
		}

		if err = json.Unmarshal(stdout, &dat); err != nil {
			return
		}
	} else {
		cmd := exec.Command("ceph", "nfs", "cluster", "info", cluster_id)
		stdout, err = cmd.CombinedOutput()
		if err != nil {
			return
		}

		if err = json.Unmarshal(stdout, &dat); err != nil {
			return
		}
	}
	return
}
func NfsExportDetailed(cluster_id string) (dat model.NfsExportDetailed, err error) {
	var stdout []byte

	cmd := exec.Command("ceph", "nfs", "export", "ls", cluster_id, "--detailed")
	stdout, err = cmd.CombinedOutput()
	if err != nil {
		return
	}

	if err = json.Unmarshal(stdout, &dat); err != nil {
		return
	}
	return
}
func NfsClusterLs() (dat model.NfsClusterInfoList, err error) {
	var stdout []byte

	cmd := exec.Command("ceph", "nfs", "cluster", "ls")
	stdout, err = cmd.CombinedOutput()
	if err != nil {
		return
	}

	if err = json.Unmarshal(stdout, &dat); err != nil {
		return
	}
	return
}
