/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	certv1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/vmware-tanzu/octant/pkg/plugin"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/store"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

var pluginName = "certificates.certmanager.k8s.io"

// This is a sample plugin showing the features of Octant's plugin API.
func main() {
	// Remove the prefix from the go logger since Octant will print logs with timestamps.
	log.SetPrefix("")

	// This plugin is interested in certificates
	certGVK := schema.GroupVersionKind{Group: "cert-manager.io", Version: "v1", Kind: "Certificate"}

	// Tell Octant to call this plugin when printing configuration for certificates
	capabilities := &plugin.Capabilities{
		SupportsPrinterConfig: []schema.GroupVersionKind{certGVK},
		IsModule:              false,
	}

	// Set up what should happen when Octant calls this plugin.
	options := []service.PluginOption{
		service.WithPrinter(handlePrint),
	}

	// Use the plugin service helper to register this plugin.
	p, err := service.Register(pluginName, "extra summary details for certificates", capabilities, options...)
	if err != nil {
		log.Fatal(err)
	}

	// The plugin can log and the log messages will show up in Octant.
	log.Printf("%s is starting", pluginName)
	p.Serve()
}

// handlePrint is called when Octant wants to print an object.
func handlePrint(request *service.PrintRequest) (plugin.PrintResponse, error) {
	if request.Object == nil {
		return plugin.PrintResponse{}, errors.Errorf("object is nil")
	}

	// Octant has a helper function to generate a key from an object. The key
	// is used to find the object in the cluster.
	key, err := store.KeyFromObject(request.Object)
	if err != nil {
		return plugin.PrintResponse{}, err
	}
	u, err := request.DashboardClient.Get(request.Context(), key)
	if err != nil {
		return plugin.PrintResponse{}, err
	}

	// The plugin can check if the object it requested exists.
	if u == nil {
		return plugin.PrintResponse{}, errors.New("object doesn't exist")
	}

	var cert certv1.Certificate
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &cert)
	if err != nil {
		return plugin.PrintResponse{}, err
	}

	config := []component.SummarySection{}
	dnsNameList := []component.Component{}

	var dnsNameComp component.Component
	for _, dnsname := range cert.Spec.DNSNames {
		dnsNameComp = component.NewText(dnsname)
		dnsNameList = append(dnsNameList, dnsNameComp)
	}
	config = append(config, component.SummarySection{Header: "DNS Names", Content: component.NewList(component.TitleFromString(""), dnsNameList)})
	// When printing an object, you can create multiple types of content. In this
	// example, the plugin is:
	//
	// * adding a field to the configuration section for this object.
	// * adding a field to the status section for this object.
	// * create a new piece of content that will be embedded in the
	//   summary section for the component.
	return plugin.PrintResponse{
		Config: config,
		// Status: status,
	}, nil
}
