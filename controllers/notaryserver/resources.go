package notaryserver

import (
	"context"

	"github.com/pkg/errors"

	goharborv1alpha2 "github.com/goharbor/harbor-operator/api/v1alpha2"
)

func (r *Reconciler) InitResources() error {
	return errors.Wrap(r.InitConfigMaps(), "configmaps")
}

func (r *Reconciler) AddResources(ctx context.Context, notaryserver *goharborv1alpha2.NotaryServer) error {
	service, err := r.GetService(ctx, notaryserver)
	if err != nil {
		return errors.Wrap(err, "cannot get service")
	}

	_, err = r.Controller.AddBasicObjectToManage(ctx, service)
	if err != nil {
		return errors.Wrapf(err, "cannot add service %s", service.GetName())
	}

	configMap, err := r.GetConfigMap(ctx, notaryserver)
	if err != nil {
		return errors.Wrap(err, "cannot get configMap")
	}

	configMapResource, err := r.Controller.AddInstantResourceToManage(ctx, configMap)
	if err != nil {
		return errors.Wrapf(err, "cannot add configMap %s", configMap.GetName())
	}

	deployment, err := r.GetDeployment(ctx, notaryserver)
	if err != nil {
		return errors.Wrap(err, "cannot get deployment")
	}

	_, err = r.Controller.AddDeploymentToManage(ctx, deployment, configMapResource)
	if err != nil {
		return errors.Wrapf(err, "cannot add deployment %s", deployment.GetName())
	}

	return nil
}