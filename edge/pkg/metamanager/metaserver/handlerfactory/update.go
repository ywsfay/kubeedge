package handlerfactory

/*

import (
	utiltrace "k8s.io/utils/trace"
)

// UpdateResource returns a function that will handle a resource update
func UpdateResource(r rest.Updater, scope *handlers.RequestScope, admit admission.Interface) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// For performance tracking purposes.
		trace := utiltrace.New("Update", utiltrace.Field{Key: "url", Value: req.URL.Path}, utiltrace.Field{Key: "user-agent", Value: &lazyTruncatedUserAgent{req}}, utiltrace.Field{Key: "client", Value: &lazyClientIP{req}})
		defer trace.LogIfLong(500 * time.Millisecond)

		if isDryRun(req.URL) && !utilfeature.DefaultFeatureGate.Enabled(features.DryRun) {
			responsewriters.ErrorNegotiated(errors.NewBadRequest("the dryRun feature is disabled"), scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}

		// TODO: we either want to remove timeout or document it (if we document, move timeout out of this function and declare it in api_installer)
		timeout := parseTimeout(req.URL.Query().Get("timeout"))

		namespace, name, err := scope.Namer.Name(req)
		if err != nil {
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}
		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()
		ctx = request.WithNamespace(ctx, namespace)

		outputMediaType, _, err := negotiation.NegotiateOutputMediaType(req, scope.Serializer, scope)
		if err != nil {
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}

		body, err := limitedReadBody(req, scope.MaxRequestBodyBytes)
		if err != nil {
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}

		options := &metav1.UpdateOptions{}
		if err := metainternalversionscheme.ParameterCodec.DecodeParameters(req.URL.Query(), scope.MetaGroupVersion, options); err != nil {
			err = errors.NewBadRequest(err.Error())
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}
		if errs := validation.ValidateUpdateOptions(options); len(errs) > 0 {
			err := errors.NewInvalid(schema.GroupKind{Group: metav1.GroupName, Kind: "UpdateOptions"}, "", errs)
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}
		options.TypeMeta.SetGroupVersionKind(metav1.SchemeGroupVersion.WithKind("UpdateOptions"))

		s, err := negotiation.NegotiateInputSerializer(req, false, scope.Serializer)
		if err != nil {
			return
		}
		defaultGVK := scope.Kind
		original := r.New()

		trace.Step("About to convert to expected version")
		decoder := scope.Serializer.DecoderToVersion(s.Serializer, scope.HubGroupVersion)
		obj, gvk, err := decoder.Decode(body, &defaultGVK, original)
		if err != nil {
			err = transformDecodeError(scope.Typer, err, original, gvk, body)
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}
		if gvk.GroupVersion() != defaultGVK.GroupVersion() {
			err = errors.NewBadRequest(fmt.Sprintf("the API version in the data (%s) does not match the expected API version (%s)", gvk.GroupVersion(), defaultGVK.GroupVersion()))
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}
		trace.Step("Conversion done")

		ae := request.AuditEventFrom(ctx)
		audit.LogRequestObject(ae, obj, scope.Resource, scope.Subresource, scope.Serializer)
		admit = admission.WithAudit(admit, ae)

		if err := checkName(obj, name, namespace, scope.Namer); err != nil {
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}

		transformers := []rest.TransformFunc{}

		trace.Step("About to store object in database")
		wasCreated := false
		requestFunc := func() (runtime.Object, error) {
			obj, created, err := r.Update(
				ctx,
				name,
				rest.DefaultUpdatedObjectInfo(obj, transformers...),
				rest.ValidateAllObjectFunc,
				rest.ValidateAllObjectUpdateFunc,
				false,
				options,
			)
			wasCreated = created
			return obj, err
		}
		result, err := finishRequest(timeout, func() (runtime.Object, error) {
			result, err := requestFunc()
			return result, err
		})
		if err != nil {
			responsewriters.ErrorNegotiated(err, scope.Serializer, scope.Kind.GroupVersion(), w, req)
			return
		}
		trace.Step("Object stored in database")

		status := http.StatusOK
		if wasCreated {
			status = http.StatusCreated
		}

		transformResponseObject(ctx, scope, trace, req, w, status, outputMediaType, result)
	}
}
*/
